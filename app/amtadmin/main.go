package main

import (
	"encoding/json"
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/jesand/crowds/amt"
	xsdt "github.com/metaleap/go-xsd/types"
	"os"
	"reflect"
	"sort"
)

const (
	USAGE = `hiclusterd - Web service hosting for hicluster

Usage:
  amtadmin balance --amt=<path> [--sandbox]
  amtadmin hit --amt=<path> --id=<id> [--sandbox]
  amtadmin -h | --help
  amtadmin --version

Options:
  balance       Get the account balance
  hits          List HIT information
  --amt=<path>  The path to a file containing AMT credentials
  --sandbox     Address the AMT sandbox instead of the production site
  --id=<id>     The ID of the object you want to view
`
)

type AmtCred struct {
	AccessKey, SecretKey string
}

func main() {

	// Parse the command line
	args, _ := docopt.Parse(USAGE, nil, true, "1.0", false)

	// Initialize the AMT client
	var (
		credPath = args["--amt"].(string)
		sandbox  = args["--sandbox"].(bool)
		amtCred  AmtCred
		client   *amt.AmtClient
	)
	if f, err := os.Open(credPath); err != nil {
		fmt.Printf("Error: Could not open %s - %v", credPath, err)
		return
	} else if err = json.NewDecoder(f).Decode(&amtCred); err != nil {
		fmt.Printf("Error: Could not parse %s - %v", credPath, err)
		return
	} else {
		client = amt.NewClient(amtCred.AccessKey, amtCred.SecretKey, sandbox)
	}

	switch {
	case args["balance"].(bool):
		RunBalance(client)
	case args["hit"].(bool):
		hitId, _ := args["--id"].(string)
		RunHit(client, hitId)
	}
}

func getObjectFields(object interface{}, vals map[string]string) {
	v := reflect.Indirect(reflect.ValueOf(object))
	if !v.IsValid() {
		return
	}
	t := v.Type()
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			getObjectFields(v.Index(i).Interface(), vals)
		}
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			switch field.Type.Kind() {
			case reflect.Struct, reflect.Ptr, reflect.Slice:
				getObjectFields(v.Field(i).Interface(), vals)
			default:
				if field.Type == reflect.TypeOf(xsdt.Int(0)) {
					vals[field.Name] = fmt.Sprintf("%d", v.Field(i).Interface())
				} else if field.Type == reflect.TypeOf(xsdt.Long(0)) {
					vals[field.Name] = fmt.Sprintf("%d", v.Field(i).Interface())
				} else {
					vals[field.Name] = v.Field(i).String()
				}
			}
		}
	}
}

func printObject(object interface{}) {
	var (
		fields   []string
		vals     = make(map[string]string)
		fieldLen int
	)
	getObjectFields(object, vals)
	for name, _ := range vals {
		fields = append(fields, name)
		if len(name) > fieldLen {
			fieldLen = len(name)
		}
	}
	sort.Strings(fields)
	format := fmt.Sprintf("%%%ds: %%s\n", fieldLen)
	for _, name := range fields {
		fmt.Printf(format, name, vals[name])
	}
}

func RunBalance(client *amt.AmtClient) {
	balance, err := client.GetAccountBalance()
	if err != nil {
		fmt.Printf("Error: The AMT request failed: %v\n", err)
		return
	}
	printObject(balance)
}

func RunHit(client *amt.AmtClient, hitId string) {
	if hit, err := client.GetHIT(hitId); err != nil {
		fmt.Printf("Error: The AMT request failed: %v\n", err)
		return
	} else if len(hit.Hits) > 0 && hit.Hits[0].Request != nil &&
		hit.Hits[0].Request.Errors != nil {

		printObject(hit.Hits[0].Request)
	} else {
		printObject(hit)
	}
}
