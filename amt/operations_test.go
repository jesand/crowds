package amt

import (
	"fmt"
	amtgen "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurk/2014-08-15/AWSMechanicalTurkRequester.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
)

const (
	FAKE_ACCESS_KEY = "FakeAccessKey"
	FAKE_SECRET_KEY = "FakeSecretKey"

	ASSIGNMENT_ID  = "FakeAssignment"
	FEEDBACK       = "Fake Feedback"
	HIT_ID         = "FakeHit"
	HIT_LAYOUT_ID  = "FakeHitLayout"
	HIT_TYPE_ID    = "FakeHitType"
	QUAL_ID        = "FakeQualification"
	QUAL_VALUE     = 7
	QUAL_VALUE_STR = "7"
	REQUEST_ID     = "ece2785b-6292-4b12-a60e-4c34847a7916"
	WORKER_ID      = "FakeWorker"
)

var (
	srv             *httptest.Server
	srvResponse     string
	srvUrlArgs      []string
	srvUrlTimestamp string
)

func newTestClient() *amtClient {
	openSrv()
	client := NewClient(FAKE_ACCESS_KEY, FAKE_SECRET_KEY, false).(*amtClient)
	client.UrlRoot = srv.URL
	return client
}

func openSrv() {
	srvResponse = ""
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		srvUrlTimestamp = r.Form.Get("Timestamp")
		var urlKeys []string
		for key, _ := range r.Form {
			urlKeys = append(urlKeys, key)
		}
		sort.Strings(urlKeys)
		srvUrlArgs = nil
		for _, key := range urlKeys {
			srvUrlArgs = append(srvUrlArgs,
				fmt.Sprintf("%s=%s", key, strings.Join(r.Form[key], "&")))
		}
		fmt.Fprint(w, srvResponse)
	}))
}

func closeSrv() {
	srv.Close()
	srv = nil
}

func TestApproveAssignment(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call ApproveAssignment", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<ApproveAssignmentResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<ApproveAssignmentResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</ApproveAssignmentResult>
				</ApproveAssignmentResponse>
				`
			result, err := client.ApproveAssignment(ASSIGNMENT_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TApproveAssignmentResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.ApproveAssignmentResults, ShouldHaveLength, 1)
				So(result.ApproveAssignmentResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"Operation=ApproveAssignment",
					"RequesterFeedback=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "ApproveAssignment", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestApproveRejectedAssignment(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call ApproveRejectedAssignment", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<ApproveRejectedAssignmentResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<ApproveRejectedAssignmentResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</ApproveRejectedAssignmentResult>
				</ApproveRejectedAssignmentResponse>`
			result, err := client.ApproveRejectedAssignment(ASSIGNMENT_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TApproveRejectedAssignmentResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.ApproveRejectedAssignmentResults, ShouldHaveLength, 1)
				So(result.ApproveRejectedAssignmentResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"Operation=ApproveRejectedAssignment",
					"RequesterFeedback=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "ApproveRejectedAssignment", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestAssignQualification(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call AssignQualification", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<AssignQualificationResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<AssignQualificationResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</AssignQualificationResult>
				</AssignQualificationResponse>`
			result, err := client.AssignQualification(QUAL_ID, WORKER_ID, QUAL_VALUE, false)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TAssignQualificationResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.AssignQualificationResults, ShouldHaveLength, 1)
				So(result.AssignQualificationResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"IntegerValue=" + QUAL_VALUE_STR,
					"Operation=AssignQualification",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "AssignQualification", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
					"WorkerId=" + WORKER_ID,
				})
			})
		})
	})
}

func TestBlockWorker(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call BlockWorker", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<BlockWorkerResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<BlockWorkerResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</BlockWorkerResult>
				</BlockWorkerResponse>`
			result, err := client.BlockWorker(WORKER_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TBlockWorkerResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.BlockWorkerResults, ShouldHaveLength, 1)
				So(result.BlockWorkerResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=BlockWorker",
					"Reason=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "BlockWorker", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
					"WorkerId=" + WORKER_ID,
				})
			})
		})
	})
}

func TestChangeHITTypeOfHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call ChangeHITTypeOfHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<ChangeHITTypeOfHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<ChangeHITTypeOfHITResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</ChangeHITTypeOfHITResult>
				</ChangeHITTypeOfHITResponse>`
			result, err := client.ChangeHITTypeOfHIT(HIT_ID, HIT_TYPE_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TChangeHITTypeOfHITResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.ChangeHITTypeOfHITResults, ShouldHaveLength, 1)
				So(result.ChangeHITTypeOfHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"HITTypeId=" + HIT_TYPE_ID,
					"Operation=ChangeHITTypeOfHIT",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "ChangeHITTypeOfHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestCreateHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call CreateHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<CreateHITResponse>
					<OperationRequest>
						<RequestId>` + REQUEST_ID + `</RequestId>
					</OperationRequest>
					<HIT>
						<Request>
							<IsValid>True</IsValid>
						</Request>
						<HITId>` + HIT_ID + `</HITId>
					</HIT>
				</CreateHITResponse>`
			result, err := client.CreateHIT("title", "description", "question",
				HIT_LAYOUT_ID, map[string]string{"name1": "val1", "name2": "val2"},
				0.5, 10, 20, 30, 40, []string{"key1", "key2", "key3"},
				nil, nil, nil, "requesterAnnotation", "uniqueRequestToken")
			Convey("Then the correct result was returned", func() {
				var (
					exp amtgen.TxsdCreateHITResponse
					op  amtgen.TxsdOperationRequest
					res amtgen.Thit
				)
				exp.OperationRequest = &op
				op.RequestId = REQUEST_ID
				exp.Hits = append(exp.Hits, &res)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.HITId = HIT_ID
				So(result, ShouldResemble, exp)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentDurationInSeconds=10",
					"AutoApprovalDelayInSeconds=40",
					"Description=description",
					"HITLayoutId=" + HIT_LAYOUT_ID,
					"HITLayoutParameters.1.Name=name1",
					"HITLayoutParameters.1.Value=val1",
					"HITLayoutParameters.2.Name=name2",
					"HITLayoutParameters.2.Value=val2",
					"Keywords=key1,key2,key3",
					"LifetimeInSeconds=20",
					"MaxAssignments=30",
					"Operation=CreateHIT",
					"Question=question",
					"RequesterAnnotation=requesterAnnotation",
					"Reward.1.Amount=0.5",
					"Reward.1.CurrencyCode=USD",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "CreateHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Title=title",
					"UniqueRequestToken=uniqueRequestToken",
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestCreateHITFromHITTypeId(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call CreateHITFromHITTypeId", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<CreateHITResponse>
					<OperationRequest>
						<RequestId>` + REQUEST_ID + `</RequestId>
					</OperationRequest>
					<HIT>
						<Request>
							<IsValid>True</IsValid>
						</Request>
						<HITId>` + HIT_ID + `</HITId>
						<HITTypeId>` + HIT_TYPE_ID + `</HITTypeId>
					</HIT>
				</CreateHITResponse>`
			result, err := client.CreateHITFromHITTypeId(HIT_TYPE_ID, "question",
				HIT_LAYOUT_ID, map[string]string{"name1": "val1", "name2": "val2"},
				10, 20, nil, nil, "requesterAnnotation", "uniqueRequestToken")
			Convey("Then the correct result was returned", func() {
				var (
					exp amtgen.TxsdCreateHITResponse
					op  amtgen.TxsdOperationRequest
					res amtgen.Thit
				)
				exp.OperationRequest = &op
				op.RequestId = REQUEST_ID
				exp.Hits = append(exp.Hits, &res)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.HITId = HIT_ID
				res.HITTypeId = HIT_TYPE_ID
				So(result, ShouldResemble, exp)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITLayoutId=" + HIT_LAYOUT_ID,
					"HITLayoutParameters.1.Name=name1",
					"HITLayoutParameters.1.Value=val1",
					"HITLayoutParameters.2.Name=name2",
					"HITLayoutParameters.2.Value=val2",
					"HITTypeId=" + HIT_TYPE_ID,
					"LifetimeInSeconds=10",
					"MaxAssignments=20",
					"Operation=CreateHIT",
					"Question=question",
					"RequesterAnnotation=requesterAnnotation",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "CreateHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"UniqueRequestToken=uniqueRequestToken",
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestCreateQualificationType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call CreateQualificationType", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<CreateQualificationTypeResponse>
					<OperationRequest>
						<RequestId>` + REQUEST_ID + `</RequestId>
					</OperationRequest>
					<QualificationType>
						<Request>
							<IsValid>True</IsValid>
						</Request>
						<QualificationTypeId>` + QUAL_ID + `</QualificationTypeId>
						<CreationTime>2009-07-13T17:26:33Z</CreationTime>
						<Name>name</Name>
						<Description>description</Description>
						<QualificationTypeStatus>Active</QualificationTypeStatus>
						<AutoGranted>True</AutoGranted>
					</QualificationType>
				</CreateQualificationTypeResponse>`
			result, err := client.CreateQualificationType("name", "description",
				[]string{"key1", "key2", "key3"}, 10,
				"qualificationTypeStatus", "test", "answerKey",
				20, true, 30)
			Convey("Then the correct result was returned", func() {
				var (
					exp amtgen.TxsdCreateQualificationTypeResponse
					op  amtgen.TxsdOperationRequest
					res amtgen.TQualificationType
				)
				exp.OperationRequest = &op
				op.RequestId = REQUEST_ID
				exp.QualificationTypes = append(exp.QualificationTypes, &res)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.QualificationTypeId = QUAL_ID
				res.CreationTime = "2009-07-13T17:26:33Z"
				res.Name = "name"
				res.Description = "description"
				res.QualificationTypeStatus = "Active"
				res.AutoGranted = xsdt.Boolean(true)
				So(result, ShouldResemble, exp)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AnswerKey=answerKey",
					"AutoGranted=true",
					"AutoGrantedValue=30",
					"Description=description",
					"Keywords=key1,key2,key3",
					"Name=name",
					"Operation=CreateQualificationType",
					"QualificationTypeStatus=qualificationTypeStatus",
					"RetryDelayInSeconds=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "CreateQualificationType", srvUrlTimestamp),
					"Test=test",
					"TestDurationInSeconds=20",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestDisableHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call DisableHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<DisableHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<DisableHITResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</DisableHITResult>
				</DisableHITResponse>`
			result, err := client.DisableHIT(HIT_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TDisableHITResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.DisableHITResults, ShouldHaveLength, 1)
				So(result.DisableHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"Operation=DisableHIT",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "DisableHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestDisposeHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call DisposeHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<DisposeHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<DisposeHITResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</DisposeHITResult>
				</DisposeHITResponse>`
			result, err := client.DisposeHIT(HIT_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TDisposeHITResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.DisposeHITResults, ShouldHaveLength, 1)
				So(result.DisposeHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"Operation=DisposeHIT",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "DisposeHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestDisposeQualificationType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call DisposeQualificationType", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<DisposeQualificationTypeResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<DisposeQualificationTypeResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</DisposeQualificationTypeResult>
				</DisposeQualificationTypeResponse>`
			result, err := client.DisposeQualificationType(QUAL_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TDisposeQualificationTypeResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.DisposeQualificationTypeResults, ShouldHaveLength, 1)
				So(result.DisposeQualificationTypeResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=DisposeQualificationType",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "DisposeQualificationType", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestExtendHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call ExtendHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<ExtendHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<ExtendHITResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</ExtendHITResult>
				</ExtendHITResponse>`
			result, err := client.ExtendHIT(HIT_ID, 10, 20, "uniqueRequestToken")
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TExtendHITResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.ExtendHITResults, ShouldHaveLength, 1)
				So(result.ExtendHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"ExpirationIncrementInSeconds=20",
					"HITId=" + HIT_ID,
					"MaxAssignmentsIncrement=10",
					"Operation=ExtendHIT",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "ExtendHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"UniqueRequestToken=uniqueRequestToken",
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestForceExpireHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call ForceExpireHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<ForceExpireHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<ForceExpireHITResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</ForceExpireHITResult>
				</ForceExpireHITResponse>`
			result, err := client.ForceExpireHIT(HIT_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TForceExpireHITResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.ForceExpireHITResults, ShouldHaveLength, 1)
				So(result.ForceExpireHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"Operation=ForceExpireHIT",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "ForceExpireHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetAccountBalance(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetAccountBalance", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetAccountBalanceResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>` +
				`<GetAccountBalanceResult>` +
				`<Request>` +
				`<IsValid>True</IsValid>` +
				`</Request>` +
				`<AvailableBalance>` +
				`<Amount>10000.000</Amount>` +
				`<CurrencyCode>USD</CurrencyCode>` +
				`<FormattedPrice>$10,000.00</FormattedPrice>` +
				`</AvailableBalance>` +
				`</GetAccountBalanceResult>` +
				`</GetAccountBalanceResponse>`
			result, err := client.GetAccountBalance()
			So(err, ShouldBeNil)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetAccountBalanceResult
					bal amtgen.TPrice
				)
				bal.Amount = xsdt.Decimal("10000.000")
				bal.CurrencyCode = "USD"
				bal.FormattedPrice = "$10,000.00"
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.AvailableBalance = &bal
				So(result.GetAccountBalanceResults, ShouldHaveLength, 1)
				So(result.GetAccountBalanceResults[0].AvailableBalance, ShouldResemble, &bal)
				So(result.GetAccountBalanceResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetAccountBalance",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetAccountBalance", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetAssignment(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetAssignment", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetAssignmentResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetAssignmentResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<Assignment>
						<AssignmentId>` + ASSIGNMENT_ID + `</AssignmentId>
						<WorkerId>` + WORKER_ID + `</WorkerId>
						<HITId>` + HIT_ID + `</HITId>
						<AssignmentStatus>Approved</AssignmentStatus>
						<AutoApprovalTime>2012-08-12T19:21:54Z</AutoApprovalTime>
						<AcceptTime>2012-07-13T19:21:40Z</AcceptTime>
						<SubmitTime>2012-07-13T19:21:54Z</SubmitTime>
						<ApprovalTime>2012-07-13T19:27:54Z</ApprovalTime>
						<Answer>
							&lt;QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd"&gt;
								[XML-formatted answer]
							&lt;/QuestionFormAnswers&gt;
						</Answer>
					</Assignment>
					<HIT>
						<HITId>` + HIT_ID + `</HITId>
						<HITTypeId>` + HIT_TYPE_ID + `</HITTypeId>
						<CreationTime>2012-07-07T00:56:40Z</CreationTime>
						<Title>Location</Title>
						<Description>Answer this Question</Description>
						<Question>
							&lt;QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd"&gt;
								[XML-formatted question]
							&lt;/QuestionForm&gt;
						</Question>
						<HITStatus>Assignable</HITStatus>
						<MaxAssignments>1</MaxAssignments>
						<Reward>
							<Amount>5.00</Amount>
							<CurrencyCode>USD</CurrencyCode>
							<FormattedPrice>$5.00</FormattedPrice>
						</Reward>
						<AutoApprovalDelayInSeconds>2592000</AutoApprovalDelayInSeconds>
						<Expiration>2012-07-14T00:56:40Z</Expiration>
						<AssignmentDurationInSeconds>30</AssignmentDurationInSeconds>
						<NumberOfSimilarHITs>1</NumberOfSimilarHITs>
						<HITReviewStatus>NotReviewed</HITReviewStatus>
					</HIT>
				</GetAssignmentResult>
				</GetAssignmentResponse>`
			result, err := client.GetAssignment(ASSIGNMENT_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetAssignmentResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.Assignment = &amtgen.TAssignment{}
				res.Assignment.AssignmentId = ASSIGNMENT_ID
				res.Assignment.WorkerId = WORKER_ID
				res.Assignment.HITId = HIT_ID
				res.Assignment.AssignmentStatus = "Approved"
				res.Assignment.AutoApprovalTime = "2012-08-12T19:21:54Z"
				res.Assignment.AcceptTime = "2012-07-13T19:21:40Z"
				res.Assignment.SubmitTime = "2012-07-13T19:21:54Z"
				res.Assignment.ApprovalTime = "2012-07-13T19:27:54Z"
				res.Assignment.Answer = `
							<QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd">
								[XML-formatted answer]
							</QuestionFormAnswers>
						`

				res.Hit = &amtgen.Thit{}
				res.Hit.HITId = HIT_ID
				res.Hit.HITTypeId = HIT_TYPE_ID
				res.Hit.CreationTime = "2012-07-07T00:56:40Z"
				res.Hit.Title = "Location"
				res.Hit.Description = "Answer this Question"
				res.Hit.Question = `
							<QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd">
								[XML-formatted question]
							</QuestionForm>
						`
				res.Hit.HITStatus = "Assignable"
				res.Hit.MaxAssignments = 1
				res.Hit.Reward = &amtgen.TPrice{}
				res.Hit.Reward.Amount = xsdt.Decimal("5.00")
				res.Hit.Reward.CurrencyCode = "USD"
				res.Hit.Reward.FormattedPrice = "$5.00"
				res.Hit.AutoApprovalDelayInSeconds = 2592000
				res.Hit.Expiration = "2012-07-14T00:56:40Z"
				res.Hit.AssignmentDurationInSeconds = 30
				res.Hit.HITReviewStatus = "NotReviewed"

				So(result.GetAssignmentResults, ShouldHaveLength, 1)
				So(result.GetAssignmentResults[0].Hit, ShouldResemble, res.Hit)
				So(result.GetAssignmentResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"Operation=GetAssignment",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetAssignment", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetAssignmentsForHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetAssignmentsForHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetAssignmentsForHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetAssignmentsForHITResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>1</NumResults>
					<TotalNumResults>1</TotalNumResults>
					<PageNumber>1</PageNumber>
					<Assignment>
						<AssignmentId>` + ASSIGNMENT_ID + `</AssignmentId>
						<WorkerId>` + WORKER_ID + `</WorkerId>
						<HITId>` + HIT_ID + `</HITId>
						<AssignmentStatus>Approved</AssignmentStatus>
						<AutoApprovalTime>2009-08-12T19:21:54Z</AutoApprovalTime>
						<AcceptTime>2009-07-13T19:21:40Z</AcceptTime>
						<SubmitTime>2009-07-13T19:21:54Z</SubmitTime>
						<ApprovalTime>2009-07-13T19:27:54Z</ApprovalTime>
						<Answer>
							&lt;QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd"&gt;
								[XML-formatted answer]
							&lt;/QuestionFormAnswers&gt;
						</Answer>
					</Assignment>
				</GetAssignmentsForHITResult>
				</GetAssignmentsForHITResponse>`
			result, err := client.GetAssignmentsForHIT(HIT_ID,
				[]string{"status1", "status2"}, "sortProperty", true, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetAssignmentsForHITResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 1
				res.TotalNumResults = 1
				res.PageNumber = 1
				res.Assignments = append(res.Assignments, &amtgen.TAssignment{})
				res.Assignments[0].AssignmentId = ASSIGNMENT_ID
				res.Assignments[0].WorkerId = WORKER_ID
				res.Assignments[0].HITId = HIT_ID
				res.Assignments[0].AssignmentStatus = "Approved"
				res.Assignments[0].AutoApprovalTime = "2009-08-12T19:21:54Z"
				res.Assignments[0].AcceptTime = "2009-07-13T19:21:40Z"
				res.Assignments[0].SubmitTime = "2009-07-13T19:21:54Z"
				res.Assignments[0].ApprovalTime = "2009-07-13T19:27:54Z"
				res.Assignments[0].Answer = `
							<QuestionFormAnswers xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionFormAnswers.xsd">
								[XML-formatted answer]
							</QuestionFormAnswers>
						`

				So(result.GetAssignmentsForHITResults, ShouldHaveLength, 1)
				So(result.GetAssignmentsForHITResults[0].Assignments, ShouldHaveLength, 1)
				So(result.GetAssignmentsForHITResults[0].Assignments[0], ShouldResemble, res.Assignments[0])
				So(result.GetAssignmentsForHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentStatus=status1,status2",
					"HITId=" + HIT_ID,
					"Operation=GetAssignmentsForHIT",
					"PageNumber=20",
					"PageSize=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetAssignmentsForHIT", srvUrlTimestamp),
					"SortDirection=Ascending",
					"SortProperty=sortProperty",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetBlockedWorkers(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetBlockedWorkers", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetBlockedWorkersResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetBlockedWorkersResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<PageNumber>1</PageNumber>
					<NumResults>2</NumResults>
					<TotalNumResults>2</TotalNumResults>
					<WorkerBlock>
						<WorkerId>A2QWESAMPLE1</WorkerId>
						<Reason>Poor Quality Work on Categorization</Reason>
					</WorkerBlock>
					<WorkerBlock>
						<WorkerId>A2QWESAMPLE2</WorkerId>
						<Reason>Poor Quality Work on Photo Moderation</Reason>
					</WorkerBlock>
				</GetBlockedWorkersResult>
				</GetBlockedWorkersResponse>`
			result, err := client.GetBlockedWorkers(10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetBlockedWorkersResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.PageNumber = 1
				res.NumResults = 2
				res.TotalNumResults = 2
				res.WorkerBlocks = append(res.WorkerBlocks,
					&amtgen.TWorkerBlock{}, &amtgen.TWorkerBlock{})
				res.WorkerBlocks[0].WorkerId = "A2QWESAMPLE1"
				res.WorkerBlocks[0].Reason = "Poor Quality Work on Categorization"
				res.WorkerBlocks[1].WorkerId = "A2QWESAMPLE2"
				res.WorkerBlocks[1].Reason = "Poor Quality Work on Photo Moderation"
				So(result.GetBlockedWorkersResults, ShouldHaveLength, 1)
				So(result.GetBlockedWorkersResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetBlockedWorkers",
					"PageNumber=20",
					"PageSize=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetBlockedWorkers", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetBonusPayments(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetBonusPayments", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetBonusPaymentsResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetBonusPaymentsResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>0</NumResults>
					<TotalNumResults>0</TotalNumResults>
					<PageNumber>1</PageNumber>
				</GetBonusPaymentsResult>
				</GetBonusPaymentsResponse>`
			result, err := client.GetBonusPayments(HIT_ID, ASSIGNMENT_ID, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetBonusPaymentsResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.PageNumber = 1
				res.NumResults = 0
				res.TotalNumResults = 0
				So(result.GetBonusPaymentsResults, ShouldHaveLength, 1)
				So(result.GetBonusPaymentsResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"HITId=" + HIT_ID,
					"Operation=GetBonusPayments",
					"PageNumber=20",
					"PageSize=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetBonusPayments", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetFileUploadURL(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetFileUploadURL", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetFileUploadURLResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetFileUploadURLResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<FileUploadURL>http://s3.amazonaws.com/myawsbucket/puppy.jpg</FileUploadURL>
				</GetFileUploadURLResult>
				</GetFileUploadURLResponse>`
			result, err := client.GetFileUploadURL(ASSIGNMENT_ID, "questionIdentifier")
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetFileUploadURLResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.FileUploadURL = "http://s3.amazonaws.com/myawsbucket/puppy.jpg"
				So(result.GetFileUploadURLResults, ShouldHaveLength, 1)
				So(result.GetFileUploadURLResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"Operation=GetFileUploadURL",
					"QuestionIdentifier=questionIdentifier",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetFileUploadURL", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<HIT>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<HITId>ZZRZPTY4ERDZWJ868JCZ</HITId>
					<HITTypeId>NYVZTQ1QVKJZXCYZCZVZ</HITTypeId>
					<CreationTime>2009-07-07T00:56:40Z</CreationTime>
					<Title>Location</Title>
					<Description>Select the image that best represents</Description>
					<Question>
						&lt;QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd"&gt;
							[XML-formatted question data]
						&lt;/QuestionForm&gt;
					</Question>
					<HITStatus>Assignable</HITStatus>
					<MaxAssignments>1</MaxAssignments>
					<Reward>
						<Amount>5.00</Amount>
						<CurrencyCode>USD</CurrencyCode>
						<FormattedPrice>$5.00</FormattedPrice>
					</Reward>
					<AutoApprovalDelayInSeconds>2592000</AutoApprovalDelayInSeconds>
					<Expiration>2009-07-14T00:56:40Z</Expiration>
					<AssignmentDurationInSeconds>30</AssignmentDurationInSeconds>
					<NumberOfSimilarHITs>1</NumberOfSimilarHITs>
					<HITReviewStatus>NotReviewed</HITReviewStatus>
				</HIT>
				</GetHITResponse>`
			result, err := client.GetHIT(HIT_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.Thit
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.HITId = "ZZRZPTY4ERDZWJ868JCZ"
				res.HITTypeId = "NYVZTQ1QVKJZXCYZCZVZ"
				res.CreationTime = "2009-07-07T00:56:40Z"
				res.Title = "Location"
				res.Description = "Select the image that best represents"
				res.Question = `
						<QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd">
							[XML-formatted question data]
						</QuestionForm>
					`
				res.HITStatus = "Assignable"
				res.MaxAssignments = 1
				res.Reward = &amtgen.TPrice{}
				res.Reward.Amount = xsdt.Decimal("5.00")
				res.Reward.CurrencyCode = "USD"
				res.Reward.FormattedPrice = "$5.00"
				res.AutoApprovalDelayInSeconds = 2592000
				res.Expiration = "2009-07-14T00:56:40Z"
				res.AssignmentDurationInSeconds = 30
				res.HITReviewStatus = "NotReviewed"
				So(result.Hits, ShouldHaveLength, 1)
				So(result.Hits[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"Operation=GetHIT",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetHITsForQualificationType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetHITsForQualificationType", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetHITsForQualificationTypeResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetHITsForQualificationTypeResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>1</NumResults>
					<TotalNumResults>1</TotalNumResults>
					<PageNumber>1</PageNumber>
					<HIT>
						<HITId>123RVWYBAZW00EXAMPLE</HITId>
						<HITTypeId>T100CN9P324W00EXAMPLE</HITTypeId>
						<CreationTime>2009-06-15T12:00:01</CreationTime>
						<HITStatus>Assignable</HITStatus>
						<MaxAssignments>5</MaxAssignments>
						<AutoApprovalDelayInSeconds>86400</AutoApprovalDelayInSeconds>
						<Expiration>2009-04-29T00:17:32Z</Expiration>
						<AssignmentDurationInSeconds>300</AssignmentDurationInSeconds>
						<Reward>
							<Amount>0.25</Amount>
							<CurrencyCode>USD</CurrencyCode>
							<FormattedPrice>$0.25</FormattedPrice>
						</Reward>
						<Title>Location and Photograph Identification</Title>
						<Description>Select the image that best represents...</Description>
						<Keywords>location, photograph, image, identification, opinion</Keywords>
						<Question>
							&lt;QuestionForm&gt;
							[XML-encoded Question data]
							&lt;/QuestionForm&gt;
						</Question>
						<QualificationRequirement>
							<QualificationTypeId>789RVWYBAZW00EXAMPLE</QualificationTypeId>
							<Comparator>GreaterThan</Comparator>
							<IntegerValue>18</IntegerValue>
						</QualificationRequirement>
						<HITReviewStatus>NotReviewed</HITReviewStatus>
					</HIT>
				</GetHITsForQualificationTypeResult>
				</GetHITsForQualificationTypeResponse>`
			result, err := client.GetHITsForQualificationType(QUAL_ID, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetHITsForQualificationTypeResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 1
				res.TotalNumResults = 1
				res.PageNumber = 1
				res.Hits = append(res.Hits, &amtgen.Thit{})
				res.Hits[0].HITId = "123RVWYBAZW00EXAMPLE"
				res.Hits[0].HITTypeId = "T100CN9P324W00EXAMPLE"
				res.Hits[0].CreationTime = "2009-06-15T12:00:01"
				res.Hits[0].HITStatus = "Assignable"
				res.Hits[0].MaxAssignments = 5
				res.Hits[0].AutoApprovalDelayInSeconds = 86400
				res.Hits[0].Expiration = "2009-04-29T00:17:32Z"
				res.Hits[0].AssignmentDurationInSeconds = 300
				res.Hits[0].Reward = &amtgen.TPrice{}
				res.Hits[0].Reward.Amount = xsdt.Decimal("0.25")
				res.Hits[0].Reward.CurrencyCode = "USD"
				res.Hits[0].Reward.FormattedPrice = "$0.25"
				res.Hits[0].Title = "Location and Photograph Identification"
				res.Hits[0].Description = "Select the image that best represents..."
				res.Hits[0].Keywords = "location, photograph, image, identification, opinion"
				res.Hits[0].Question = `
							<QuestionForm>
							[XML-encoded Question data]
							</QuestionForm>
						`
				res.Hits[0].QualificationRequirements = append(res.Hits[0].QualificationRequirements, &amtgen.TQualificationRequirement{})
				res.Hits[0].QualificationRequirements[0].QualificationTypeId = "789RVWYBAZW00EXAMPLE"
				res.Hits[0].QualificationRequirements[0].Comparator = "GreaterThan"
				res.Hits[0].QualificationRequirements[0].IntegerValues = []xsdt.Int{18}
				res.Hits[0].HITReviewStatus = "NotReviewed"
				So(result.GetHITsForQualificationTypeResults, ShouldHaveLength, 1)
				So(result.GetHITsForQualificationTypeResults[0].Hits, ShouldHaveLength, 1)
				So(result.GetHITsForQualificationTypeResults[0].Hits[0].Reward, ShouldResemble, res.Hits[0].Reward)
				So(result.GetHITsForQualificationTypeResults[0].Hits[0].QualificationRequirements, ShouldHaveLength, 1)
				So(result.GetHITsForQualificationTypeResults[0].Hits[0].QualificationRequirements[0], ShouldResemble, res.Hits[0].QualificationRequirements[0])
				So(result.GetHITsForQualificationTypeResults[0].Hits[0], ShouldResemble, res.Hits[0])
				So(result.GetHITsForQualificationTypeResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetHITsForQualificationType",
					"PageNumber=20",
					"PageSize=10",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetHITsForQualificationType", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetQualificationsForQualificationType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetQualificationsForQualificationType", func() {
			// Note: this is not the response from the API docs. The API docs
			// have a QualificationRequest instead of a Qualification, which
			// is counter to the .xsd and the expectation for this API call.
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetQualificationsForQualificationTypeResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetQualificationsForQualificationTypeResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>1</NumResults>
					<TotalNumResults>1</TotalNumResults>
					<PageNumber>1</PageNumber>
					<Qualification>
						<QualificationTypeId>789RVWYBAZW00EXAMPLE</QualificationTypeId>
						<SubjectId>AZ3456EXAMPLE</SubjectId>
						<GrantTime>2005-01-31T23:59:59Z</GrantTime>
						<IntegerValue>95</IntegerValue>
					</Qualification>
				</GetQualificationsForQualificationTypeResult>
				</GetQualificationsForQualificationTypeResponse>`
			result, err := client.GetQualificationsForQualificationType(
				QUAL_ID, true, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetQualificationsForQualificationTypeResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 1
				res.TotalNumResults = 1
				res.PageNumber = 1
				res.Qualifications = append(res.Qualifications, &amtgen.TQualification{})
				res.Qualifications[0].QualificationTypeId = "789RVWYBAZW00EXAMPLE"
				res.Qualifications[0].SubjectId = "AZ3456EXAMPLE"
				res.Qualifications[0].GrantTime = "2005-01-31T23:59:59Z"
				res.Qualifications[0].IntegerValue = 95
				So(result.GetQualificationsForQualificationTypeResults, ShouldHaveLength, 1)
				So(result.GetQualificationsForQualificationTypeResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetQualificationsForQualificationType",
					"PageNumber=20",
					"PageSize=10",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetQualificationsForQualificationType", srvUrlTimestamp),
					"Status=Granted",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetQualificationRequests(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetQualificationRequests", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetQualificationRequestsResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetQualificationRequestsResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>1</NumResults>
					<TotalNumResults>1</TotalNumResults>
					<PageNumber>1</PageNumber>
					<QualificationRequest>
						<QualificationRequestId>789RVWYBAZW00EXAMPLE951RVWYBAZW00EXAMPLE</QualificationRequestId>
						<QualificationTypeId>789RVWYBAZW00EXAMPLE</QualificationTypeId>
						<SubjectId>AZ3456EXAMPLE</SubjectId>
						<Test>
							&lt;QuestionForm&gt;
							[XML-encoded question data]
							&lt;/QuestionForm&gt;
						</Test>
						<Answer>
							&lt;QuestionFormAnswers&gt;
							[XML-encoded answer data]
							&lt;/QuestionFormAnswers&gt;
						</Answer>
						<SubmitTime>2005-12-01T23:59:59Z</SubmitTime>
					</QualificationRequest>
				</GetQualificationRequestsResult>
				</GetQualificationRequestsResponse>`
			result, err := client.GetQualificationRequests(QUAL_ID, "sortProperty", true, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetQualificationRequestsResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 1
				res.TotalNumResults = 1
				res.PageNumber = 1
				res.QualificationRequests = append(res.QualificationRequests, &amtgen.TQualificationRequest{})
				res.QualificationRequests[0].QualificationRequestId = "789RVWYBAZW00EXAMPLE951RVWYBAZW00EXAMPLE"
				res.QualificationRequests[0].QualificationTypeId = "789RVWYBAZW00EXAMPLE"
				res.QualificationRequests[0].SubjectId = "AZ3456EXAMPLE"
				res.QualificationRequests[0].Test = `
							<QuestionForm>
							[XML-encoded question data]
							</QuestionForm>
						`
				res.QualificationRequests[0].Answer = `
							<QuestionFormAnswers>
							[XML-encoded answer data]
							</QuestionFormAnswers>
						`
				res.QualificationRequests[0].SubmitTime = "2005-12-01T23:59:59Z"
				So(result.GetQualificationRequestsResults, ShouldHaveLength, 1)
				So(result.GetQualificationRequestsResults[0].QualificationRequests, ShouldHaveLength, 1)
				So(result.GetQualificationRequestsResults[0].QualificationRequests[0], ShouldResemble, res.QualificationRequests[0])
				So(result.GetQualificationRequestsResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetQualificationRequests",
					"PageNumber=20",
					"PageSize=10",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetQualificationRequests", srvUrlTimestamp),
					"SortDirection=Ascending",
					"SortProperty=sortProperty",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetQualificationScore(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetQualificationScore", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetQualificationScoreResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<Qualification>
					<QualificationTypeId>789RVWYBAZW00EXAMPLE</QualificationTypeId>
					<SubjectId>AZ3456EXAMPLE</SubjectId>
					<GrantTime>2005-01-31T23:59:59Z</GrantTime>
					<IntegerValue>95</IntegerValue>
				</Qualification>
				</GetQualificationScoreResponse>`
			result, err := client.GetQualificationScore(QUAL_ID, WORKER_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TQualification
				)
				res.QualificationTypeId = "789RVWYBAZW00EXAMPLE"
				res.SubjectId = "AZ3456EXAMPLE"
				res.GrantTime = "2005-01-31T23:59:59Z"
				res.IntegerValue = 95
				So(result.Qualifications, ShouldHaveLength, 1)
				So(result.Qualifications[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetQualificationScore",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetQualificationScore", srvUrlTimestamp),
					"SubjectId=" + WORKER_ID,
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetQualificationType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetQualificationType", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetQualificationTypeResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<QualificationType>
					<QualificationTypeId>789RVWYBAZW00EXAMPLE</QualificationTypeId>
					<CreationTime>2005-01-31T23:59:59Z</CreationTime>
					<Name>EnglishWritingAbility</Name>
					<Description>The ability to write and edit text...</Description>
					<Keywords>English, text, write, edit, language</Keywords>
					<QualificationTypeStatus>Active</QualificationTypeStatus>
					<RetryDelayInSeconds>86400</RetryDelayInSeconds>
					<IsRequestable>true</IsRequestable>
				</QualificationType>
				</GetQualificationTypeResponse>`
			result, err := client.GetQualificationType(QUAL_ID)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TQualificationType
				)
				res.QualificationTypeId = "789RVWYBAZW00EXAMPLE"
				res.CreationTime = "2005-01-31T23:59:59Z"
				res.Name = "EnglishWritingAbility"
				res.Description = "The ability to write and edit text..."
				res.Keywords = "English, text, write, edit, language"
				res.QualificationTypeStatus = "Active"
				res.RetryDelayInSeconds = 86400
				res.IsRequestable = true
				So(result.QualificationTypes, ShouldHaveLength, 1)
				So(result.QualificationTypes[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=GetQualificationType",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetQualificationType", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetRequesterStatistic(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetRequesterStatistic", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetStatisticResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetStatisticResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<Statistic>NumberAssignmentsApproved</Statistic>
					<TimePeriod>ThirtyDays</TimePeriod>
					<DataPoint>
						<Date>2011-09-05T07:00:00Z</Date>
						<DoubleValue>281</DoubleValue>
					</DataPoint>
				</GetStatisticResult>
				</GetStatisticResponse>`
			result, err := client.GetRequesterStatistic("statistic", "timePeriod", 10)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetStatisticResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.Statistic = "NumberAssignmentsApproved"
				res.TimePeriod = "ThirtyDays"
				res.DataPoints = append(res.DataPoints, &amtgen.TDataPoint{})
				res.DataPoints[0].Date = "2011-09-05T07:00:00Z"
				res.DataPoints[0].DoubleValue = 281
				So(result.GetStatisticResults, ShouldHaveLength, 1)
				So(result.GetStatisticResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Count=10",
					"Operation=GetRequesterStatistic",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetRequesterStatistic", srvUrlTimestamp),
					"Statistic=statistic",
					"TimePeriod=timePeriod",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetRequesterWorkerStatistic(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetRequesterWorkerStatistic", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetStatisticResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetStatisticResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<WorkerId>A1Z4X5D207ALZF</WorkerId>
					<Statistic>NumberAssignmentsApproved</Statistic>
					<TimePeriod>ThirtyDays</TimePeriod>
					<DataPoint>
						<Date>2011-09-05T07:00:00Z</Date>
						<DoubleValue>281</DoubleValue>
					</DataPoint>
				</GetStatisticResult>
				</GetStatisticResponse>`
			result, err := client.GetRequesterWorkerStatistic("statistic", WORKER_ID, "timePeriod", 10)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetStatisticResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.WorkerId = "A1Z4X5D207ALZF"
				res.Statistic = "NumberAssignmentsApproved"
				res.TimePeriod = "ThirtyDays"
				res.DataPoints = append(res.DataPoints, &amtgen.TDataPoint{})
				res.DataPoints[0].Date = "2011-09-05T07:00:00Z"
				res.DataPoints[0].DoubleValue = 281
				So(result.GetStatisticResults, ShouldHaveLength, 1)
				So(result.GetStatisticResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Count=10",
					"Operation=GetRequesterWorkerStatistic",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetRequesterWorkerStatistic", srvUrlTimestamp),
					"Statistic=statistic",
					"TimePeriod=timePeriod",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
					"WorkerId=" + WORKER_ID,
				})
			})
		})
	})
}

func TestGetReviewableHITs(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetReviewableHITs", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetReviewableHITsResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetReviewableHITsResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>1</NumResults>
					<TotalNumResults>1</TotalNumResults>
					<PageNumber>1</PageNumber>
					<HIT>
						<HITId>GBHZVQX3EHXZ2AYDY2T0</HITId>
					</HIT>
				</GetReviewableHITsResult>
				</GetReviewableHITsResponse>`
			result, err := client.GetReviewableHITs(HIT_TYPE_ID, "status", "sortProperty", true, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetReviewableHITsResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 1
				res.TotalNumResults = 1
				res.PageNumber = 1
				res.Hits = append(res.Hits, &amtgen.Thit{})
				res.Hits[0].HITId = "GBHZVQX3EHXZ2AYDY2T0"
				So(result.GetReviewableHITsResults, ShouldHaveLength, 1)
				So(result.GetReviewableHITsResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITTypeId=" + HIT_TYPE_ID,
					"Operation=GetReviewableHITs",
					"PageNumber=20",
					"PageSize=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetReviewableHITs", srvUrlTimestamp),
					"SortDirection=Ascending",
					"SortProperty=sortProperty",
					"Status=status",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGetReviewResultsForHIT(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GetReviewResultsForHIT", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GetReviewResultsForHITResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GetReviewResultsForHITResult>
					<HITId>1AAAAAAAAABBBBBBBBBBCCCCCCCCCC</HITId>
					<AssignmentReviewPolicy>
						<PolicyName>ScoreMyKnownAnswers/2011-09-01</PolicyName>
					</AssignmentReviewPolicy>
					<HITReviewPolicy>
						<PolicyName>SimplePlurality/2011-09-01</PolicyName>
					</HITReviewPolicy>
					<AssignmentReviewReport>
						<ReviewResult>
							<SubjectId>1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF</SubjectId>
							<SubjectType>Assignment</SubjectType>
							<QuestionId>Question_2</QuestionId>
							<Key>KnownAnswerCorrect</Key>
							<Value>1</Value>
						</ReviewResult>
						<ReviewResult>
							<SubjectId>1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF</SubjectId>
							<SubjectType>Assignment</SubjectType>
							<Key>KnownAnswerScore</Key>
							<Value>100</Value>
						</ReviewResult>
						<ReviewResult>
							<SubjectId>1GGGGGGGGGHHHHHHHHHHIIIIIIIIII</SubjectId>
							<SubjectType>Assignment</SubjectType>
							<QuestionId>Question_2</QuestionId>
							<Key>KnownAnswerCorrect</Key>
							<Value>0</Value>
						</ReviewResult>
						<ReviewResult>
							<SubjectId>1GGGGGGGGGHHHHHHHHHHIIIIIIIIII</SubjectId>
							<SubjectType>Assignment</SubjectType>
							<Key>KnownAnswerScore</Key>
							<Value>0</Value>
						</ReviewResult>
						<ReviewAction>
							<ActionName>review</ActionName>
							<ObjectId>1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF</ObjectId>
							<ObjectType>Assignment</ObjectType>
							<Status>SUCCEEDED</Status>
							<Result>Reviewed one known answer; 1/1 correct</Result>
						</ReviewAction>
						<ReviewAction>
							<ActionName>approve</ActionName>
							<ObjectId>1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF</ObjectId>
							<ObjectType>Assignment</ObjectType>
							<Status>SUCCEEDED</Status>
							<Result>Approved</Result>
						</ReviewAction>
						<ReviewAction>
							<ActionName>review</ActionName>
							<ObjectId>1GGGGGGGGGHHHHHHHHHHIIIIIIIIII</ObjectId>
							<ObjectType>Assignment</ObjectType>
							<Status>SUCCEEDED</Status>
							<Result>Reviewed one known answer; 0/1 correct</Result>
						</ReviewAction>
						<ReviewAction>
							<ActionName>reject</ActionName>
							<ObjectId>1GGGGGGGGGHHHHHHHHHHIIIIIIIIII</ObjectId>
							<ObjectType>Assignment</ObjectType>
							<Status>SUCCEEDED</Status>
							<Result>Rejected</Result>
						</ReviewAction>
					</AssignmentReviewReport>
					<HITReviewReport>
						<ReviewResult>
							<SubjectId>1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF</SubjectId>
							<SubjectType>Assignment</SubjectType>
							<QuestionId>Question_1</QuestionId>
							<Key>AgreedWithPlurality</Key>
							<Value>1</Value>
						</ReviewResult>
						<ReviewResult>
							<SubjectId>1GGGGGGGGGHHHHHHHHHHIIIIIIIIII</SubjectId>
							<SubjectType>Assignment</SubjectType>
							<QuestionId>Question_1</QuestionId>
							<Key>AgreedWithPlurality</Key>
							<Value>1</Value>
						</ReviewResult>
						<ReviewResult>
							<SubjectId>1AAAAAAAAABBBBBBBBBBCCCCCCCCCC</SubjectId>
							<SubjectType>HIT</SubjectType>
							<QuestionId>Question_1</QuestionId>
							<Key>PluralityAnswer</Key>
							<Value>true</Value>
						</ReviewResult>
						<ReviewResult>
							<SubjectId>1AAAAAAAAABBBBBBBBBBCCCCCCCCCC</SubjectId>
							<SubjectType>HIT</SubjectType>
							<QuestionId>Question_1</QuestionId>
							<Key>PluralityLevel</Key>
							<Value>100</Value>
						</ReviewResult>
						<ReviewAction>
							<ActionName>approve</ActionName>
							<ObjectId>1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF</ObjectId>
							<ObjectType>Assignment</ObjectType>
							<Status>SUCCEEDED</Status>
							<Result>Already approved</Result>
						</ReviewAction>
						<ReviewAction>
							<ActionName>approve</ActionName>
							<ObjectId>1GGGGGGGGGHHHHHHHHHHIIIIIIIIII</ObjectId>
							<ObjectType>Assignment</ObjectType>
							<Status>FAILED</Status>
							<Result>Assignment was in an invalid state for this operation.</Result>
							<ErrorCode>AWS.MechanicalTurk.InvalidAssignmentState</ErrorCode>
						</ReviewAction>
					</HITReviewReport>
				</GetReviewResultsForHITResult>
				</GetReviewResultsForHITResponse>`
			result, err := client.GetReviewResultsForHIT(HIT_ID, []string{"lvl1", "lvl2"}, true, false, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGetReviewResultsForHITResult
				)
				res.HITId = "1AAAAAAAAABBBBBBBBBBCCCCCCCCCC"
				res.AssignmentReviewPolicy = &amtgen.TReviewPolicy{}
				res.AssignmentReviewPolicy.PolicyName = "ScoreMyKnownAnswers/2011-09-01"
				res.HITReviewPolicy = &amtgen.TReviewPolicy{}
				res.HITReviewPolicy.PolicyName = "SimplePlurality/2011-09-01"
				res.AssignmentReviewReport = &amtgen.TReviewReport{}

				res.AssignmentReviewReport.ReviewResults = append(res.AssignmentReviewReport.ReviewResults,
					&amtgen.TReviewResultDetail{}, &amtgen.TReviewResultDetail{}, &amtgen.TReviewResultDetail{}, &amtgen.TReviewResultDetail{})
				res.AssignmentReviewReport.ReviewResults[0].SubjectId = "1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF"
				res.AssignmentReviewReport.ReviewResults[0].SubjectType = "Assignment"
				res.AssignmentReviewReport.ReviewResults[0].QuestionId = "Question_2"
				res.AssignmentReviewReport.ReviewResults[0].Key = "KnownAnswerCorrect"
				res.AssignmentReviewReport.ReviewResults[0].Value = "1"

				res.AssignmentReviewReport.ReviewResults[1].SubjectId = "1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF"
				res.AssignmentReviewReport.ReviewResults[1].SubjectType = "Assignment"
				res.AssignmentReviewReport.ReviewResults[1].Key = "KnownAnswerScore"
				res.AssignmentReviewReport.ReviewResults[1].Value = "100"

				res.AssignmentReviewReport.ReviewResults[2].SubjectId = "1GGGGGGGGGHHHHHHHHHHIIIIIIIIII"
				res.AssignmentReviewReport.ReviewResults[2].SubjectType = "Assignment"
				res.AssignmentReviewReport.ReviewResults[2].QuestionId = "Question_2"
				res.AssignmentReviewReport.ReviewResults[2].Key = "KnownAnswerCorrect"
				res.AssignmentReviewReport.ReviewResults[2].Value = "0"

				res.AssignmentReviewReport.ReviewResults[3].SubjectId = "1GGGGGGGGGHHHHHHHHHHIIIIIIIIII"
				res.AssignmentReviewReport.ReviewResults[3].SubjectType = "Assignment"
				res.AssignmentReviewReport.ReviewResults[3].Key = "KnownAnswerScore"
				res.AssignmentReviewReport.ReviewResults[3].Value = "0"

				res.AssignmentReviewReport.ReviewActions = append(res.AssignmentReviewReport.ReviewActions,
					&amtgen.TReviewActionDetail{}, &amtgen.TReviewActionDetail{}, &amtgen.TReviewActionDetail{}, &amtgen.TReviewActionDetail{})
				res.AssignmentReviewReport.ReviewActions[0].ActionName = "review"
				res.AssignmentReviewReport.ReviewActions[0].ObjectId = "1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF"
				res.AssignmentReviewReport.ReviewActions[0].ObjectType = "Assignment"
				res.AssignmentReviewReport.ReviewActions[0].Status = "SUCCEEDED"
				res.AssignmentReviewReport.ReviewActions[0].Result = "Reviewed one known answer; 1/1 correct"

				res.AssignmentReviewReport.ReviewActions[1].ActionName = "approve"
				res.AssignmentReviewReport.ReviewActions[1].ObjectId = "1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF"
				res.AssignmentReviewReport.ReviewActions[1].ObjectType = "Assignment"
				res.AssignmentReviewReport.ReviewActions[1].Status = "SUCCEEDED"
				res.AssignmentReviewReport.ReviewActions[1].Result = "Approved"

				res.AssignmentReviewReport.ReviewActions[2].ActionName = "review"
				res.AssignmentReviewReport.ReviewActions[2].ObjectId = "1GGGGGGGGGHHHHHHHHHHIIIIIIIIII"
				res.AssignmentReviewReport.ReviewActions[2].ObjectType = "Assignment"
				res.AssignmentReviewReport.ReviewActions[2].Status = "SUCCEEDED"
				res.AssignmentReviewReport.ReviewActions[2].Result = "Reviewed one known answer; 0/1 correct"

				res.AssignmentReviewReport.ReviewActions[3].ActionName = "reject"
				res.AssignmentReviewReport.ReviewActions[3].ObjectId = "1GGGGGGGGGHHHHHHHHHHIIIIIIIIII"
				res.AssignmentReviewReport.ReviewActions[3].ObjectType = "Assignment"
				res.AssignmentReviewReport.ReviewActions[3].Status = "SUCCEEDED"
				res.AssignmentReviewReport.ReviewActions[3].Result = "Rejected"

				res.HITReviewReport = &amtgen.TReviewReport{}
				res.HITReviewReport.ReviewResults = append(res.HITReviewReport.ReviewResults,
					&amtgen.TReviewResultDetail{}, &amtgen.TReviewResultDetail{}, &amtgen.TReviewResultDetail{}, &amtgen.TReviewResultDetail{})
				res.HITReviewReport.ReviewResults[0].SubjectId = "1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF"
				res.HITReviewReport.ReviewResults[0].SubjectType = "Assignment"
				res.HITReviewReport.ReviewResults[0].QuestionId = "Question_1"
				res.HITReviewReport.ReviewResults[0].Key = "AgreedWithPlurality"
				res.HITReviewReport.ReviewResults[0].Value = "1"

				res.HITReviewReport.ReviewResults[1].SubjectId = "1GGGGGGGGGHHHHHHHHHHIIIIIIIIII"
				res.HITReviewReport.ReviewResults[1].SubjectType = "Assignment"
				res.HITReviewReport.ReviewResults[1].QuestionId = "Question_1"
				res.HITReviewReport.ReviewResults[1].Key = "AgreedWithPlurality"
				res.HITReviewReport.ReviewResults[1].Value = "1"

				res.HITReviewReport.ReviewResults[2].SubjectId = "1AAAAAAAAABBBBBBBBBBCCCCCCCCCC"
				res.HITReviewReport.ReviewResults[2].SubjectType = "HIT"
				res.HITReviewReport.ReviewResults[2].QuestionId = "Question_1"
				res.HITReviewReport.ReviewResults[2].Key = "PluralityAnswer"
				res.HITReviewReport.ReviewResults[2].Value = "true"

				res.HITReviewReport.ReviewResults[3].SubjectId = "1AAAAAAAAABBBBBBBBBBCCCCCCCCCC"
				res.HITReviewReport.ReviewResults[3].SubjectType = "HIT"
				res.HITReviewReport.ReviewResults[3].QuestionId = "Question_1"
				res.HITReviewReport.ReviewResults[3].Key = "PluralityLevel"
				res.HITReviewReport.ReviewResults[3].Value = "100"

				res.HITReviewReport.ReviewActions = append(res.HITReviewReport.ReviewActions,
					&amtgen.TReviewActionDetail{}, &amtgen.TReviewActionDetail{})

				res.HITReviewReport.ReviewActions[0].ActionName = "approve"
				res.HITReviewReport.ReviewActions[0].ObjectId = "1DDDDDDDDDEEEEEEEEEEFFFFFFFFFF"
				res.HITReviewReport.ReviewActions[0].ObjectType = "Assignment"
				res.HITReviewReport.ReviewActions[0].Status = "SUCCEEDED"
				res.HITReviewReport.ReviewActions[0].Result = "Already approved"

				res.HITReviewReport.ReviewActions[1].ActionName = "approve"
				res.HITReviewReport.ReviewActions[1].ObjectId = "1GGGGGGGGGHHHHHHHHHHIIIIIIIIII"
				res.HITReviewReport.ReviewActions[1].ObjectType = "Assignment"
				res.HITReviewReport.ReviewActions[1].Status = "FAILED"
				res.HITReviewReport.ReviewActions[1].Result = "Assignment was in an invalid state for this operation."
				res.HITReviewReport.ReviewActions[1].ErrorCode = "AWS.MechanicalTurk.InvalidAssignmentState"

				So(result.GetReviewResultsForHITResults, ShouldHaveLength, 1)
				So(result.GetReviewResultsForHITResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"Operation=GetReviewResultsForHIT",
					"PageNumber=20",
					"PageSize=10",
					"PolicyLevel=lvl1,lvl2",
					"RetrieveActions=true",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GetReviewResultsForHIT", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestGrantBonus(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GrantBonus", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GrantBonusResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GrantBonusResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</GrantBonusResult>
				</GrantBonusResponse>`
			result, err := client.GrantBonus(WORKER_ID, ASSIGNMENT_ID, 1.5, FEEDBACK, "uniqueRequestToken")
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGrantBonusResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.GrantBonusResults, ShouldHaveLength, 1)
				So(result.GrantBonusResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"BonusAmount.1.Amount=1.5",
					"BonusAmount.1.CurrencyCode=USD",
					"Operation=GrantBonus",
					"Reason=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GrantBonus", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"UniqueRequestToken=uniqueRequestToken",
					"Version=2014-08-15",
					"WorkerId=" + WORKER_ID,
				})
			})
		})
	})
}

func TestGrantQualification(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call GrantQualification", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<GrantQualificationResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<GrantQualificationResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</GrantQualificationResult>
				</GrantQualificationResponse>`
			result, err := client.GrantQualification(QUAL_ID, QUAL_VALUE)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TGrantQualificationResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.GrantQualificationResults, ShouldHaveLength, 1)
				So(result.GrantQualificationResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"IntegerValue=" + QUAL_VALUE_STR,
					"Operation=GrantQualification",
					"QualificationRequestId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "GrantQualification", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestNotifyWorkers(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call NotifyWorkers", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<NotifyWorkersResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<NotifyWorkersResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</NotifyWorkersResult>
				</NotifyWorkersResponse>`
			result, err := client.NotifyWorkers("subject", "messageText", []string{WORKER_ID})
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TNotifyWorkersResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.NotifyWorkersResults, ShouldHaveLength, 1)
				So(result.NotifyWorkersResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"MessageText=messageText",
					"Operation=NotifyWorkers",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "NotifyWorkers", srvUrlTimestamp),
					"Subject=subject",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
					"WorkerId=" + WORKER_ID,
				})
			})
		})
	})
}

func TestRegisterHITType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call RegisterHITType", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<RegisterHITTypeResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<RegisterHITTypeResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<HITTypeId>KZ3GKTRXBWGYX8WXBW60</HITTypeId>
				</RegisterHITTypeResult>
				</RegisterHITTypeResponse>`
			result, err := client.RegisterHITType("title", "description",
				1.5, 10, 20, []string{"key1", "key2", "key3"}, nil)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TRegisterHITTypeResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.HITTypeId = "KZ3GKTRXBWGYX8WXBW60"
				So(result.RegisterHITTypeResults, ShouldHaveLength, 1)
				So(result.RegisterHITTypeResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentDurationInSeconds=10",
					"AutoApprovalDelayInSeconds=20",
					"Description=description",
					"Keywords=key1,key2,key3",
					"Operation=RegisterHITType",
					"Reward.1.Amount=1.5",
					"Reward.1.CurrencyCode=USD",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "RegisterHITType", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Title=title",
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestRejectAssignment(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call RejectAssignment", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<RejectAssignmentResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<RejectAssignmentResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</RejectAssignmentResult>
				</RejectAssignmentResponse>`
			result, err := client.RejectAssignment(ASSIGNMENT_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TRejectAssignmentResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.RejectAssignmentResults, ShouldHaveLength, 1)
				So(result.RejectAssignmentResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AssignmentId=" + ASSIGNMENT_ID,
					"Operation=RejectAssignment",
					"RequesterFeedback=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "RejectAssignment", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestRejectQualificationRequest(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call RejectQualificationRequest", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<RejectQualificationRequestResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<RejectQualificationRequestResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</RejectQualificationRequestResult>
				</RejectQualificationRequestResponse>`
			result, err := client.RejectQualificationRequest(QUAL_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TRejectQualificationRequestResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.RejectQualificationRequestResults, ShouldHaveLength, 1)
				So(result.RejectQualificationRequestResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=RejectQualificationRequest",
					"QualificationRequestId=" + QUAL_ID,
					"Reason=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "RejectQualificationRequest", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestRevokeQualification(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call RevokeQualification", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<RevokeQualificationResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<RevokeQualificationResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</RevokeQualificationResult>
				</RevokeQualificationResponse>`
			result, err := client.RevokeQualification(WORKER_ID, QUAL_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TRevokeQualificationResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.RevokeQualificationResults, ShouldHaveLength, 1)
				So(result.RevokeQualificationResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=RevokeQualification",
					"QualificationTypeId=" + QUAL_ID,
					"Reason=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "RevokeQualification", srvUrlTimestamp),
					"SubjectId=" + WORKER_ID,
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestSearchHITs(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call SearchHITs", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<SearchHITsResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<SearchHITsResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>2</NumResults>
					<TotalNumResults>2</TotalNumResults>
					<PageNumber>1</PageNumber>

					<HIT>
						<HITId>GBHZVQX3EHXZ2AYDY2T0</HITId>
						<HITTypeId>NYVZTQ1QVKJZXCYZCZVZ</HITTypeId>
						<CreationTime>2009-04-22T00:17:32Z</CreationTime>
						<Title>Location</Title>
						<Description>Select the image that best represents</Description>
						<HITStatus>Reviewable</HITStatus>
						<MaxAssignments>1</MaxAssignments>
						<Reward>
							<Amount>5.00</Amount>
							<CurrencyCode>USD</CurrencyCode>
							<FormattedPrice>$5.00</FormattedPrice>
						</Reward>
						<AutoApprovalDelayInSeconds>2592000</AutoApprovalDelayInSeconds>
						<Expiration>2009-04-29T00:17:32Z</Expiration>
						<AssignmentDurationInSeconds>30</AssignmentDurationInSeconds>
						<NumberOfAssignmentsPending>0</NumberOfAssignmentsPending>
						<NumberOfAssignmentsAvailable>0</NumberOfAssignmentsAvailable>
						<NumberOfAssignmentsCompleted>1</NumberOfAssignmentsCompleted>
					</HIT>

					<HIT>
						<HITId>ZZRZPTY4ERDZWJ868JCZ</HITId>
						<HITTypeId>NYVZTQ1QVKJZXCYZCZVZ</HITTypeId>
						<CreationTime>2009-07-07T00:56:40Z</CreationTime>
						<Title>Location</Title>
						<Description>Select the image that best represents</Description>
						<HITStatus>Assignable</HITStatus>
						<MaxAssignments>1</MaxAssignments>
						<Reward>
							<Amount>5.00</Amount>
							<CurrencyCode>USD</CurrencyCode>
							<FormattedPrice>$5.00</FormattedPrice>
						</Reward>
						<AutoApprovalDelayInSeconds>2592000</AutoApprovalDelayInSeconds>
						<Expiration>2009-07-14T00:56:40Z</Expiration>
						<AssignmentDurationInSeconds>30</AssignmentDurationInSeconds>
						<NumberOfAssignmentsPending>0</NumberOfAssignmentsPending>
						<NumberOfAssignmentsAvailable>1</NumberOfAssignmentsAvailable>
						<NumberOfAssignmentsCompleted>0</NumberOfAssignmentsCompleted>
					</HIT>
				</SearchHITsResult>
				</SearchHITsResponse>`
			result, err := client.SearchHITs("sortProperty", true, 10, 20)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TSearchHITsResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 2
				res.TotalNumResults = 2
				res.PageNumber = 1
				res.Hits = append(res.Hits, &amtgen.Thit{}, &amtgen.Thit{})

				res.Hits[0].HITId = "GBHZVQX3EHXZ2AYDY2T0"
				res.Hits[0].HITTypeId = "NYVZTQ1QVKJZXCYZCZVZ"
				res.Hits[0].CreationTime = "2009-04-22T00:17:32Z"
				res.Hits[0].Title = "Location"
				res.Hits[0].Description = "Select the image that best represents"
				res.Hits[0].HITStatus = "Reviewable"
				res.Hits[0].MaxAssignments = 1
				res.Hits[0].Reward = &amtgen.TPrice{}
				res.Hits[0].Reward.Amount = xsdt.Decimal("5.00")
				res.Hits[0].Reward.CurrencyCode = "USD"
				res.Hits[0].Reward.FormattedPrice = "$5.00"
				res.Hits[0].AutoApprovalDelayInSeconds = 2592000
				res.Hits[0].Expiration = "2009-04-29T00:17:32Z"
				res.Hits[0].AssignmentDurationInSeconds = 30
				res.Hits[0].NumberOfAssignmentsPending = 0
				res.Hits[0].NumberOfAssignmentsAvailable = 0
				res.Hits[0].NumberOfAssignmentsCompleted = 1

				res.Hits[1].HITId = "ZZRZPTY4ERDZWJ868JCZ"
				res.Hits[1].HITTypeId = "NYVZTQ1QVKJZXCYZCZVZ"
				res.Hits[1].CreationTime = "2009-07-07T00:56:40Z"
				res.Hits[1].Title = "Location"
				res.Hits[1].Description = "Select the image that best represents"
				res.Hits[1].HITStatus = "Assignable"
				res.Hits[1].MaxAssignments = 1
				res.Hits[1].Reward = &amtgen.TPrice{}
				res.Hits[1].Reward.Amount = xsdt.Decimal("5.00")
				res.Hits[1].Reward.CurrencyCode = "USD"
				res.Hits[1].Reward.FormattedPrice = "$5.00"
				res.Hits[1].AutoApprovalDelayInSeconds = 2592000
				res.Hits[1].Expiration = "2009-07-14T00:56:40Z"
				res.Hits[1].AssignmentDurationInSeconds = 30
				res.Hits[1].NumberOfAssignmentsPending = 0
				res.Hits[1].NumberOfAssignmentsAvailable = 1
				res.Hits[1].NumberOfAssignmentsCompleted = 0

				So(result.SearchHITsResults, ShouldHaveLength, 1)
				So(result.SearchHITsResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=SearchHITs",
					"PageNumber=20",
					"PageSize=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "SearchHITs", srvUrlTimestamp),
					"SortDirection=Ascending",
					"SortProperty=sortProperty",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestSearchQualificationTypes(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call SearchQualificationTypes", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<SearchQualificationTypesResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<SearchQualificationTypesResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
					<NumResults>10</NumResults>
					<TotalNumResults>5813</TotalNumResults>
					<QualificationType>
						<QualificationTypeId>WKAZMYZDCYCZP412TZEZ</QualificationTypeId>
						<CreationTime>2009-05-17T10:05:15Z</CreationTime>
						<Name>WebReviews Qualification Master Test</Name>
						<Description>
							This qualification will allow you to earn more on the WebReviews HITs.
						</Description>
						<Keywords>WebReviews, webreviews, web reviews</Keywords>
						<QualificationTypeStatus>Active</QualificationTypeStatus>
						<Test>
							&lt;QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd"&gt;
								[XML-formatted question data]
							&lt;/QuestionForm&gt;
						</Test>
						<TestDurationInSeconds>1200</TestDurationInSeconds>
					</QualificationType>
				</SearchQualificationTypesResult>
				</SearchQualificationTypesResponse>`
			result, err := client.SearchQualificationTypes("query", "sortProperty", true, 10, 20, false, true)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TSearchQualificationTypesResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				res.NumResults = 10
				res.TotalNumResults = 5813
				res.QualificationTypes = append(res.QualificationTypes, &amtgen.TQualificationType{})
				res.QualificationTypes[0].QualificationTypeId = "WKAZMYZDCYCZP412TZEZ"
				res.QualificationTypes[0].CreationTime = "2009-05-17T10:05:15Z"
				res.QualificationTypes[0].Name = "WebReviews Qualification Master Test"
				res.QualificationTypes[0].Description = "\n\t\t\t\t\t\t\tThis qualification will allow you to earn more on the WebReviews HITs.\n\t\t\t\t\t\t"
				res.QualificationTypes[0].Keywords = "WebReviews, webreviews, web reviews"
				res.QualificationTypes[0].QualificationTypeStatus = "Active"
				res.QualificationTypes[0].TestDurationInSeconds = 1200
				res.QualificationTypes[0].Test = `
							<QuestionForm xmlns="http://mechanicalturk.amazonaws.com/AWSMechanicalTurkDataSchemas/2005-10-01/QuestionForm.xsd">
								[XML-formatted question data]
							</QuestionForm>
						`
				So(result.SearchQualificationTypesResults, ShouldHaveLength, 1)
				So(result.SearchQualificationTypesResults[0].QualificationTypes, ShouldHaveLength, 1)
				So(result.SearchQualificationTypesResults[0].QualificationTypes[0], ShouldResemble, res.QualificationTypes[0])
				So(result.SearchQualificationTypesResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"MustBeOwnedByCaller=true",
					"Operation=SearchQualificationTypes",
					"PageNumber=20",
					"PageSize=10",
					"Query=query",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "SearchQualificationTypes", srvUrlTimestamp),
					"SortDirection=Ascending",
					"SortProperty=sortProperty",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestSendTestEventNotification(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call SendTestEventNotification", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<SendTestEventNotificationResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<SendTestEventNotificationResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</SendTestEventNotificationResult>
				</SendTestEventNotificationResponse>`
			result, err := client.SendTestEventNotification(nil, "testEventType")
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TSendTestEventNotificationResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.SendTestEventNotificationResults, ShouldHaveLength, 1)
				So(result.SendTestEventNotificationResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=SendTestEventNotification",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "SendTestEventNotification", srvUrlTimestamp),
					"TestEventType=testEventType",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestSetHITAsReviewing(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call SetHITAsReviewing", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<SetHITAsReviewingResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<SetHITAsReviewingResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</SetHITAsReviewingResult>
				</SetHITAsReviewingResponse>`
			result, err := client.SetHITAsReviewing(HIT_ID, true)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TSetHITAsReviewingResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.SetHITAsReviewingResults, ShouldHaveLength, 1)
				So(result.SetHITAsReviewingResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"HITId=" + HIT_ID,
					"Operation=SetHITAsReviewing",
					"Revert=true",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "SetHITAsReviewing", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestSetHITTypeNotification(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call SetHITTypeNotification", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<SetHITTypeNotificationResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<SetHITTypeNotificationResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</SetHITTypeNotificationResult>
				</SetHITTypeNotificationResponse>`
			result, err := client.SetHITTypeNotification(HIT_TYPE_ID, nil, true)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TSetHITTypeNotificationResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.SetHITTypeNotificationResults, ShouldHaveLength, 1)
				So(result.SetHITTypeNotificationResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Active=true",
					"HITTypeId=" + HIT_TYPE_ID,
					"Operation=SetHITTypeNotification",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "SetHITTypeNotification", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestUnblockWorker(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call UnblockWorker", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<UnblockWorkerResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<UnblockWorkerResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</UnblockWorkerResult>
				</UnblockWorkerResponse>`
			result, err := client.UnblockWorker(WORKER_ID, FEEDBACK)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TUnblockWorkerResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.UnblockWorkerResults, ShouldHaveLength, 1)
				So(result.UnblockWorkerResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"Operation=UnblockWorker",
					"Reason=" + FEEDBACK,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "UnblockWorker", srvUrlTimestamp),
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
					"WorkerId=" + WORKER_ID,
				})
			})
		})
	})
}

func TestUpdateQualificationScore(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call UpdateQualificationScore", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<UpdateQualificationScoreResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<UpdateQualificationScoreResult>
					<Request>
						<IsValid>True</IsValid>
					</Request>
				</UpdateQualificationScoreResult>
				</UpdateQualificationScoreResponse>`
			result, err := client.UpdateQualificationScore(QUAL_ID, WORKER_ID, QUAL_VALUE)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TUpdateQualificationScoreResult
				)
				res.Request = &amtgen.TxsdRequest{}
				res.Request.IsValid = "True"
				So(result.UpdateQualificationScoreResults, ShouldHaveLength, 1)
				So(result.UpdateQualificationScoreResults[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"IntegerValue=" + QUAL_VALUE_STR,
					"Operation=UpdateQualificationScore",
					"QualificationTypeId=" + QUAL_ID,
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "UpdateQualificationScore", srvUrlTimestamp),
					"SubjectId=" + WORKER_ID,
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}

func TestUpdateQualificationType(t *testing.T) {
	client := newTestClient()
	defer closeSrv()
	Convey("Given an initialized client", t, func() {
		Convey("When I call UpdateQualificationType", func() {
			srvResponse = `<?xml version="1.0"?>` + "\n" +
				`<UpdateQualificationTypeResponse>` +
				`<OperationRequest>` +
				`<RequestId>5d38ca17-77d6-4caa-b3f0-7cabc17c56ad</RequestId>` +
				`</OperationRequest>
				<QualificationType>
					<QualificationTypeId>789RVWYBAZW00EXAMPLE</QualificationTypeId>
					<CreationTime>2009-06-15T12:00:01Z</CreationTime>
					<Name>EnglishWritingAbility</Name>
					<Description>The ability to write and edit text...</Description>
					<Keywords>English, text, write, edit, language</Keywords>
					<QualificationTypeStatus>Active</QualificationTypeStatus>
					<RetryDelayInSeconds>86400</RetryDelayInSeconds>
					<IsRequestable>true</IsRequestable>
				</QualificationType>
				</UpdateQualificationTypeResponse>`
			result, err := client.UpdateQualificationType(QUAL_ID, 10,
				"qualificationTypeStatus", "description", "test",
				"answerKey", 20, true, 30)
			Convey("Then the correct result was returned", func() {
				var (
					res amtgen.TQualificationType
				)
				// res.Request = &amtgen.TxsdRequest{}
				// res.Request.IsValid = "True"
				res.QualificationTypeId = "789RVWYBAZW00EXAMPLE"
				res.CreationTime = "2009-06-15T12:00:01Z"
				res.Name = "EnglishWritingAbility"
				res.Description = "The ability to write and edit text..."
				res.Keywords = "English, text, write, edit, language"
				res.QualificationTypeStatus = "Active"
				res.RetryDelayInSeconds = 86400
				res.IsRequestable = true
				So(result.QualificationTypes, ShouldHaveLength, 1)
				So(result.QualificationTypes[0], ShouldResemble, &res)
			})
			Convey("Then the operation succeeded", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then the correct URL was fetched", func() {
				So(srvUrlArgs, ShouldResemble, []string{
					"AWSAccessKeyId=" + FAKE_ACCESS_KEY,
					"AnswerKey=answerKey",
					"AutoGranted=true",
					"AutoGrantedValue=30",
					"Description=description",
					"Operation=UpdateQualificationType",
					"QualificationTypeId=" + QUAL_ID,
					"QualificationTypeStatus=qualificationTypeStatus",
					"RetryDelayInSeconds=10",
					"Service=AWSMechanicalTurkRequester",
					"Signature=" + client.signatureFor("AWSMechanicalTurkRequester", "UpdateQualificationType", srvUrlTimestamp),
					"Test=test",
					"TestDurationInSeconds=20",
					"Timestamp=" + srvUrlTimestamp,
					"Version=2014-08-15",
				})
			})
		})
	})
}
