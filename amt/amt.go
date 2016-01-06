// Package amt implements an API for crowdsourcing questions through
// Amazon Mechanical Turk. This package implement the low-level API calls
// exposed by AMT.
//
// The operations exposed here are documented by Amazon:
// http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_OperationsArticle.html
//
// Known issues:
// - The QuestionForm is not smart enough to marshal elements in the correct
//   order. This is a drawback of encoding/xml.
// - QuestionFormAnswers Marshals several empty fields.
package amt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// AmtClient is an initialized structure for interacting with AMT.
type AmtClient struct {

	// The access key for your AMT account
	AWSAccessKeyId string

	// The secret key (password) for your AMT account
	SecretKey string

	// The root URL to which requests should be sent
	UrlRoot string
}

// Initialize a new client for AMT.
func NewClient(accessKeyId, secretKey string, sandbox bool) *AmtClient {
	urlRoot := URL_PROD
	if sandbox {
		urlRoot = URL_SANDBOX
	}
	return &AmtClient{
		AWSAccessKeyId: accessKeyId,
		SecretKey:      secretKey,
		UrlRoot:        urlRoot,
	}
}

// amtRequest wraps a Request type from amtgen with an operation name. It can
// safely be marshalled into a REST request.
type amtRequest struct {
	AWSAccessKeyId, Signature string
	Service, Version          string
	Operation, Timestamp      string
	Request                   interface{}
}

// Formats the current time in the format required by AMT
func FormatNow() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}

// Formats a timestamp in the format required by AMT (2005-01-31T23:59:59Z)
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

// Sets default fields and cryptographically signs the request.
func (client AmtClient) signRequest(operation string, request interface{}) (amtRequest, error) {
	t := reflect.TypeOf(request)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return amtRequest{}, errors.New("signRequest() requires a struct ptr as its second arg")
	} else if reflect.ValueOf(request).IsNil() {
		return amtRequest{}, errors.New("signRequest() requires a non-nil struct ptr as its second arg")
	}

	req := amtRequest{
		AWSAccessKeyId: client.AWSAccessKeyId,
		Operation:      operation,
		Request:        request,
		Service:        AMT_SERVICE,
		Timestamp:      FormatNow(),
		Version:        API_VERSION,
	}
	req.Signature = client.signatureFor(req.Service, req.Operation, req.Timestamp)
	return req, nil
}

func (client AmtClient) signatureFor(service, operation, timestamp string) string {
	mac := hmac.New(sha1.New, []byte(client.SecretKey))
	io.WriteString(mac, service)
	io.WriteString(mac, operation)
	io.WriteString(mac, timestamp)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func isEmptyValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Ptr, reflect.Slice:
		if v.IsNil() {
			return true
		}

	case reflect.Struct:
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			if !isEmptyValue(v.Field(i)) {
				return false
			}
		}
		return true

	default:
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
	return false
}

func packField(n string, v reflect.Value, justIndexed bool) map[string]string {
	m := make(map[string]string)
	t := v.Type()

	switch t.Kind() {
	case reflect.Ptr:
		if !justIndexed {
			return packField(n+".1", v.Elem(), true)
		} else {
			return packField(n, v.Elem(), false)
		}

	case reflect.Slice:
		st := t.Elem()
		if st.Kind() == reflect.Ptr {
			st = st.Elem()
		}
		if st.Kind() == reflect.Struct {
			for i := 0; i < v.Len(); i++ {
				ni := fmt.Sprintf("%s.%d", n, i+1)
				for k, v := range packField(ni, v.Index(i), true) {
					if v != "" {
						m[k] = v
					}
				}
			}
		} else {

			// Potentially error-prone noun singularization
			if strings.HasSuffix(n, "ses") {
				// For: AssignmentStatuses
				n = n[:len(n)-2]
			} else if strings.HasSuffix(n, "s") {
				// For: WorkerIds, PolicyLevels
				n = n[:len(n)-1]
			}

			var vals []string
			for i := 0; i < v.Len(); i++ {
				for _, v := range packField(n, v.Index(i), true) {
					if v != "" {
						vals = append(vals, v)
					}
				}
			}
			m[n] = strings.Join(vals, ",")
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := t.Field(i)
			var ni string
			if f.Anonymous {
				ni = n
				justIndexed = false
			} else {
				ni = fmt.Sprintf("%s.%s", n, t.Field(i).Name)
				justIndexed = true
			}
			for k, v := range packField(ni, v.Field(i), justIndexed) {
				if v != "" {
					m[k] = v
				}
			}
		}

	default:
		m[n] = fmt.Sprint(v.Interface())
	}

	return m
}

// Send a request and decode the response into the given struct.
func (client AmtClient) sendRequest(request amtRequest, response interface{}) error {
	req, err := http.NewRequest("GET", client.UrlRoot, nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()
	query.Add("AWSAccessKeyId", request.AWSAccessKeyId)
	query.Add("Operation", request.Operation)
	query.Add("Service", request.Service)
	query.Add("Signature", request.Signature)
	query.Add("Timestamp", request.Timestamp)
	query.Add("Version", request.Version)
	if request.Request != nil {
		args := reflect.ValueOf(request.Request).Elem().FieldByName("Requests").Index(0).Elem()
		argType := args.Type()
		for i := 0; i < args.NumField(); i++ {
			fName := argType.FieldByIndex([]int{i, 0}).Name
			fValue := args.FieldByIndex([]int{i, 0})
			if !isEmptyValue(fValue) {
				for key, value := range packField(fName, fValue, false) {
					query.Add(key, value)
				}
			}
		}
	}
	req.URL.RawQuery = query.Encode()

	if resp, err := http.DefaultClient.Do(req); err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with HTTP status %d: %s", resp.StatusCode, resp.Status)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		resp.Body.Close()
		contents = []byte(strings.TrimSpace(string(contents)))
		if len(contents) == 0 {
			return fmt.Errorf("%s returned an empty response.", request.Operation)
		}

		wrapped := []byte(fmt.Sprintf(
			"%s<response xmlns='http://requester.mturk.amazonaws.com/doc/%s'>%s</response>",
			xml.Header, API_VERSION, string(contents)))
		err = xml.Unmarshal(wrapped, response)
		if err == nil && len(contents) > 0 &&
			isEmptyValue(reflect.ValueOf(response).Elem().Field(1)) {

			return fmt.Errorf("%s returned an empty response struct. Parse error? Response was: %s",
				request.Operation, string(contents))
		}
		return err
	}
}
