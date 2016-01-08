// Automatically generated by MockGen. DO NOT EDIT!
// Source: amt.go

package amt

import (
	gomock "github.com/golang/mock/gomock"
	AWSMechanicalTurkRequester_xsd_go "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurk/2014-08-15/AWSMechanicalTurkRequester.xsd_go"
)

// Mock of AmtClient interface
type MockAmtClient struct {
	ctrl     *gomock.Controller
	recorder *_MockAmtClientRecorder
}

// Recorder for MockAmtClient (not exported)
type _MockAmtClientRecorder struct {
	mock *MockAmtClient
}

func NewMockAmtClient(ctrl *gomock.Controller) *MockAmtClient {
	mock := &MockAmtClient{ctrl: ctrl}
	mock.recorder = &_MockAmtClientRecorder{mock}
	return mock
}

func (_m *MockAmtClient) EXPECT() *_MockAmtClientRecorder {
	return _m.recorder
}

func (_m *MockAmtClient) ApproveAssignment(assignmentId string, requesterFeedback string) (AWSMechanicalTurkRequester_xsd_go.TxsdApproveAssignmentResponse, error) {
	ret := _m.ctrl.Call(_m, "ApproveAssignment", assignmentId, requesterFeedback)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdApproveAssignmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) ApproveAssignment(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ApproveAssignment", arg0, arg1)
}

func (_m *MockAmtClient) ApproveRejectedAssignment(assignmentId string, requesterFeedback string) (AWSMechanicalTurkRequester_xsd_go.TxsdApproveRejectedAssignmentResponse, error) {
	ret := _m.ctrl.Call(_m, "ApproveRejectedAssignment", assignmentId, requesterFeedback)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdApproveRejectedAssignmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) ApproveRejectedAssignment(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ApproveRejectedAssignment", arg0, arg1)
}

func (_m *MockAmtClient) AssignQualification(qualificationTypeId string, workerId string, integerValue int, sendNotification bool) (AWSMechanicalTurkRequester_xsd_go.TxsdAssignQualificationResponse, error) {
	ret := _m.ctrl.Call(_m, "AssignQualification", qualificationTypeId, workerId, integerValue, sendNotification)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdAssignQualificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) AssignQualification(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssignQualification", arg0, arg1, arg2, arg3)
}

func (_m *MockAmtClient) BlockWorker(workerId string, reason string) (AWSMechanicalTurkRequester_xsd_go.TxsdBlockWorkerResponse, error) {
	ret := _m.ctrl.Call(_m, "BlockWorker", workerId, reason)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdBlockWorkerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) BlockWorker(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BlockWorker", arg0, arg1)
}

func (_m *MockAmtClient) ChangeHITTypeOfHIT(hitId string, hitTypeId string) (AWSMechanicalTurkRequester_xsd_go.TxsdChangeHITTypeOfHITResponse, error) {
	ret := _m.ctrl.Call(_m, "ChangeHITTypeOfHIT", hitId, hitTypeId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdChangeHITTypeOfHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) ChangeHITTypeOfHIT(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ChangeHITTypeOfHIT", arg0, arg1)
}

func (_m *MockAmtClient) CreateHIT(title string, description string, question string, hitLayoutId string, hitLayoutParameters map[string]string, reward float32, assignmentDurationInSeconds int, lifetimeInSeconds int, maxAssignments int, autoApprovalDelayInSeconds int, keywords []string, qualificationRequirements []*AWSMechanicalTurkRequester_xsd_go.TQualificationRequirement, assignmentReviewPolicy *AWSMechanicalTurkRequester_xsd_go.TReviewPolicy, hitReviewPolicy *AWSMechanicalTurkRequester_xsd_go.TReviewPolicy, requesterAnnotation string, uniqueRequestToken string) (AWSMechanicalTurkRequester_xsd_go.TxsdCreateHITResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateHIT", title, description, question, hitLayoutId, hitLayoutParameters, reward, assignmentDurationInSeconds, lifetimeInSeconds, maxAssignments, autoApprovalDelayInSeconds, keywords, qualificationRequirements, assignmentReviewPolicy, hitReviewPolicy, requesterAnnotation, uniqueRequestToken)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdCreateHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) CreateHIT(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateHIT", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
}

func (_m *MockAmtClient) CreateHITFromArgs(args AWSMechanicalTurkRequester_xsd_go.TCreateHITRequest) (AWSMechanicalTurkRequester_xsd_go.TxsdCreateHITResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateHITFromArgs", args)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdCreateHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) CreateHITFromArgs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateHITFromArgs", arg0)
}

func (_m *MockAmtClient) CreateHITFromHITTypeId(hitTypeId string, question string, hitLayoutId string, hitLayoutParameters map[string]string, lifetimeInSeconds int, maxAssignments int, assignmentReviewPolicy *AWSMechanicalTurkRequester_xsd_go.TReviewPolicy, hitReviewPolicy *AWSMechanicalTurkRequester_xsd_go.TReviewPolicy, requesterAnnotation string, uniqueRequestToken string) (AWSMechanicalTurkRequester_xsd_go.TxsdCreateHITResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateHITFromHITTypeId", hitTypeId, question, hitLayoutId, hitLayoutParameters, lifetimeInSeconds, maxAssignments, assignmentReviewPolicy, hitReviewPolicy, requesterAnnotation, uniqueRequestToken)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdCreateHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) CreateHITFromHITTypeId(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateHITFromHITTypeId", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

func (_m *MockAmtClient) CreateQualificationType(name string, description string, keywords []string, retryDelayInSeconds int, qualificationTypeStatus string, test string, answerKey string, testDurationInSeconds int, autoGranted bool, autoGrantedValue int) (AWSMechanicalTurkRequester_xsd_go.TxsdCreateQualificationTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateQualificationType", name, description, keywords, retryDelayInSeconds, qualificationTypeStatus, test, answerKey, testDurationInSeconds, autoGranted, autoGrantedValue)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdCreateQualificationTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) CreateQualificationType(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateQualificationType", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

func (_m *MockAmtClient) DisableHIT(hitId string) (AWSMechanicalTurkRequester_xsd_go.TxsdDisableHITResponse, error) {
	ret := _m.ctrl.Call(_m, "DisableHIT", hitId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdDisableHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) DisableHIT(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableHIT", arg0)
}

func (_m *MockAmtClient) DisposeHIT(hitId string) (AWSMechanicalTurkRequester_xsd_go.TxsdDisposeHITResponse, error) {
	ret := _m.ctrl.Call(_m, "DisposeHIT", hitId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdDisposeHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) DisposeHIT(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisposeHIT", arg0)
}

func (_m *MockAmtClient) DisposeQualificationType(qualificationTypeId string) (AWSMechanicalTurkRequester_xsd_go.TxsdDisposeQualificationTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "DisposeQualificationType", qualificationTypeId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdDisposeQualificationTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) DisposeQualificationType(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisposeQualificationType", arg0)
}

func (_m *MockAmtClient) ExtendHIT(hitId string, maxAssignmentsIncrement int, expirationIncrementInSeconds int, uniqueRequestToken string) (AWSMechanicalTurkRequester_xsd_go.TxsdExtendHITResponse, error) {
	ret := _m.ctrl.Call(_m, "ExtendHIT", hitId, maxAssignmentsIncrement, expirationIncrementInSeconds, uniqueRequestToken)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdExtendHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) ExtendHIT(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExtendHIT", arg0, arg1, arg2, arg3)
}

func (_m *MockAmtClient) ForceExpireHIT(hitId string) (AWSMechanicalTurkRequester_xsd_go.TxsdForceExpireHITResponse, error) {
	ret := _m.ctrl.Call(_m, "ForceExpireHIT", hitId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdForceExpireHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) ForceExpireHIT(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForceExpireHIT", arg0)
}

func (_m *MockAmtClient) GetAccountBalance() (AWSMechanicalTurkRequester_xsd_go.TxsdGetAccountBalanceResponse, error) {
	ret := _m.ctrl.Call(_m, "GetAccountBalance")
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetAccountBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetAccountBalance() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccountBalance")
}

func (_m *MockAmtClient) GetAssignment(assignmentId string) (AWSMechanicalTurkRequester_xsd_go.TxsdGetAssignmentResponse, error) {
	ret := _m.ctrl.Call(_m, "GetAssignment", assignmentId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetAssignmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetAssignment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAssignment", arg0)
}

func (_m *MockAmtClient) GetAssignmentsForHIT(hitId string, assignmentStatuses []string, sortProperty string, sortAscending bool, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetAssignmentsForHITResponse, error) {
	ret := _m.ctrl.Call(_m, "GetAssignmentsForHIT", hitId, assignmentStatuses, sortProperty, sortAscending, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetAssignmentsForHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetAssignmentsForHIT(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAssignmentsForHIT", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockAmtClient) GetBlockedWorkers(pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetBlockedWorkersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetBlockedWorkers", pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetBlockedWorkersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetBlockedWorkers(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBlockedWorkers", arg0, arg1)
}

func (_m *MockAmtClient) GetBonusPayments(hitId string, assignmentId string, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetBonusPaymentsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetBonusPayments", hitId, assignmentId, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetBonusPaymentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetBonusPayments(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBonusPayments", arg0, arg1, arg2, arg3)
}

func (_m *MockAmtClient) GetFileUploadURL(assignmentId string, questionIdentifier string) (AWSMechanicalTurkRequester_xsd_go.TxsdGetFileUploadURLResponse, error) {
	ret := _m.ctrl.Call(_m, "GetFileUploadURL", assignmentId, questionIdentifier)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetFileUploadURLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetFileUploadURL(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFileUploadURL", arg0, arg1)
}

func (_m *MockAmtClient) GetHIT(hitId string) (AWSMechanicalTurkRequester_xsd_go.TxsdGetHITResponse, error) {
	ret := _m.ctrl.Call(_m, "GetHIT", hitId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetHIT(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHIT", arg0)
}

func (_m *MockAmtClient) GetHITsForQualificationType(qualificationTypeId string, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetHITsForQualificationTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "GetHITsForQualificationType", qualificationTypeId, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetHITsForQualificationTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetHITsForQualificationType(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHITsForQualificationType", arg0, arg1, arg2)
}

func (_m *MockAmtClient) GetQualificationRequests(qualificationTypeId string, sortProperty string, sortAscending bool, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationRequestsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetQualificationRequests", qualificationTypeId, sortProperty, sortAscending, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationRequestsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetQualificationRequests(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetQualificationRequests", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockAmtClient) GetQualificationScore(qualificationTypeId string, subjectId string) (AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationScoreResponse, error) {
	ret := _m.ctrl.Call(_m, "GetQualificationScore", qualificationTypeId, subjectId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationScoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetQualificationScore(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetQualificationScore", arg0, arg1)
}

func (_m *MockAmtClient) GetQualificationsForQualificationType(qualificationTypeId string, isGranted bool, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationsForQualificationTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "GetQualificationsForQualificationType", qualificationTypeId, isGranted, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationsForQualificationTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetQualificationsForQualificationType(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetQualificationsForQualificationType", arg0, arg1, arg2, arg3)
}

func (_m *MockAmtClient) GetQualificationType(qualificationTypeId string) (AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "GetQualificationType", qualificationTypeId)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetQualificationTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetQualificationType(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetQualificationType", arg0)
}

func (_m *MockAmtClient) GetRequesterStatistic(statistic string, timePeriod string, count int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetRequesterStatisticResponse, error) {
	ret := _m.ctrl.Call(_m, "GetRequesterStatistic", statistic, timePeriod, count)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetRequesterStatisticResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetRequesterStatistic(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequesterStatistic", arg0, arg1, arg2)
}

func (_m *MockAmtClient) GetRequesterWorkerStatistic(statistic string, workerId string, timePeriod string, count int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetRequesterWorkerStatisticResponse, error) {
	ret := _m.ctrl.Call(_m, "GetRequesterWorkerStatistic", statistic, workerId, timePeriod, count)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetRequesterWorkerStatisticResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetRequesterWorkerStatistic(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequesterWorkerStatistic", arg0, arg1, arg2, arg3)
}

func (_m *MockAmtClient) GetReviewableHITs(hitTypeId string, status string, sortProperty string, sortAscending bool, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetReviewableHITsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetReviewableHITs", hitTypeId, status, sortProperty, sortAscending, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetReviewableHITsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetReviewableHITs(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReviewableHITs", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockAmtClient) GetReviewResultsForHIT(hitId string, policyLevels []string, retrieveActions bool, retrieveResults bool, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdGetReviewResultsForHITResponse, error) {
	ret := _m.ctrl.Call(_m, "GetReviewResultsForHIT", hitId, policyLevels, retrieveActions, retrieveResults, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGetReviewResultsForHITResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GetReviewResultsForHIT(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetReviewResultsForHIT", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockAmtClient) GrantBonus(workerId string, assignmentId string, bonusAmount float32, reason string, uniqueRequestToken string) (AWSMechanicalTurkRequester_xsd_go.TxsdGrantBonusResponse, error) {
	ret := _m.ctrl.Call(_m, "GrantBonus", workerId, assignmentId, bonusAmount, reason, uniqueRequestToken)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGrantBonusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GrantBonus(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GrantBonus", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockAmtClient) GrantQualification(qualificationRequestId string, integerValue int) (AWSMechanicalTurkRequester_xsd_go.TxsdGrantQualificationResponse, error) {
	ret := _m.ctrl.Call(_m, "GrantQualification", qualificationRequestId, integerValue)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdGrantQualificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) GrantQualification(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GrantQualification", arg0, arg1)
}

func (_m *MockAmtClient) NotifyWorkers(subject string, messageText string, workerIds []string) (AWSMechanicalTurkRequester_xsd_go.TxsdNotifyWorkersResponse, error) {
	ret := _m.ctrl.Call(_m, "NotifyWorkers", subject, messageText, workerIds)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdNotifyWorkersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) NotifyWorkers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NotifyWorkers", arg0, arg1, arg2)
}

func (_m *MockAmtClient) RegisterHITType(title string, description string, reward float32, assignmentDurationInSeconds int, autoApprovalDelayInSeconds int, keywords []string, qualificationRequirements []*AWSMechanicalTurkRequester_xsd_go.TQualificationRequirement) (AWSMechanicalTurkRequester_xsd_go.TxsdRegisterHITTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "RegisterHITType", title, description, reward, assignmentDurationInSeconds, autoApprovalDelayInSeconds, keywords, qualificationRequirements)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdRegisterHITTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) RegisterHITType(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterHITType", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockAmtClient) RegisterHITTypeFromArgs(args AWSMechanicalTurkRequester_xsd_go.TRegisterHITTypeRequest) (AWSMechanicalTurkRequester_xsd_go.TxsdRegisterHITTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "RegisterHITTypeFromArgs", args)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdRegisterHITTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) RegisterHITTypeFromArgs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterHITTypeFromArgs", arg0)
}

func (_m *MockAmtClient) RejectAssignment(assignmentId string, requesterFeedback string) (AWSMechanicalTurkRequester_xsd_go.TxsdRejectAssignmentResponse, error) {
	ret := _m.ctrl.Call(_m, "RejectAssignment", assignmentId, requesterFeedback)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdRejectAssignmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) RejectAssignment(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RejectAssignment", arg0, arg1)
}

func (_m *MockAmtClient) RejectQualificationRequest(qualificationRequestId string, reason string) (AWSMechanicalTurkRequester_xsd_go.TxsdRejectQualificationRequestResponse, error) {
	ret := _m.ctrl.Call(_m, "RejectQualificationRequest", qualificationRequestId, reason)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdRejectQualificationRequestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) RejectQualificationRequest(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RejectQualificationRequest", arg0, arg1)
}

func (_m *MockAmtClient) RevokeQualification(subjectId string, qualificationTypeId string, reason string) (AWSMechanicalTurkRequester_xsd_go.TxsdRevokeQualificationResponse, error) {
	ret := _m.ctrl.Call(_m, "RevokeQualification", subjectId, qualificationTypeId, reason)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdRevokeQualificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) RevokeQualification(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeQualification", arg0, arg1, arg2)
}

func (_m *MockAmtClient) SearchHITs(sortProperty string, sortAscending bool, pageSize int, pageNumber int) (AWSMechanicalTurkRequester_xsd_go.TxsdSearchHITsResponse, error) {
	ret := _m.ctrl.Call(_m, "SearchHITs", sortProperty, sortAscending, pageSize, pageNumber)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdSearchHITsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) SearchHITs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SearchHITs", arg0, arg1, arg2, arg3)
}

func (_m *MockAmtClient) SearchQualificationTypes(query string, sortProperty string, sortAscending bool, pageSize int, pageNumber int, mustBeRequestable bool, mustBeOwnedByCaller bool) (AWSMechanicalTurkRequester_xsd_go.TxsdSearchQualificationTypesResponse, error) {
	ret := _m.ctrl.Call(_m, "SearchQualificationTypes", query, sortProperty, sortAscending, pageSize, pageNumber, mustBeRequestable, mustBeOwnedByCaller)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdSearchQualificationTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) SearchQualificationTypes(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SearchQualificationTypes", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockAmtClient) SendTestEventNotification(notification *AWSMechanicalTurkRequester_xsd_go.TNotificationSpecification, testEventType string) (AWSMechanicalTurkRequester_xsd_go.TxsdSendTestEventNotificationResponse, error) {
	ret := _m.ctrl.Call(_m, "SendTestEventNotification", notification, testEventType)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdSendTestEventNotificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) SendTestEventNotification(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTestEventNotification", arg0, arg1)
}

func (_m *MockAmtClient) SetHITAsReviewing(hitID string, revert bool) (AWSMechanicalTurkRequester_xsd_go.TxsdSetHITAsReviewingResponse, error) {
	ret := _m.ctrl.Call(_m, "SetHITAsReviewing", hitID, revert)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdSetHITAsReviewingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) SetHITAsReviewing(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetHITAsReviewing", arg0, arg1)
}

func (_m *MockAmtClient) SetHITTypeNotification(hitTypeID string, notification *AWSMechanicalTurkRequester_xsd_go.TNotificationSpecification, active bool) (AWSMechanicalTurkRequester_xsd_go.TxsdSetHITTypeNotificationResponse, error) {
	ret := _m.ctrl.Call(_m, "SetHITTypeNotification", hitTypeID, notification, active)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdSetHITTypeNotificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) SetHITTypeNotification(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetHITTypeNotification", arg0, arg1, arg2)
}

func (_m *MockAmtClient) UnblockWorker(workerId string, reason string) (AWSMechanicalTurkRequester_xsd_go.TxsdUnblockWorkerResponse, error) {
	ret := _m.ctrl.Call(_m, "UnblockWorker", workerId, reason)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdUnblockWorkerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) UnblockWorker(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnblockWorker", arg0, arg1)
}

func (_m *MockAmtClient) UpdateQualificationScore(qualificationTypeId string, subjectId string, integerValue int) (AWSMechanicalTurkRequester_xsd_go.TxsdUpdateQualificationScoreResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateQualificationScore", qualificationTypeId, subjectId, integerValue)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdUpdateQualificationScoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) UpdateQualificationScore(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateQualificationScore", arg0, arg1, arg2)
}

func (_m *MockAmtClient) UpdateQualificationType(qualificationTypeId string, retryDelayInSeconds int, qualificationTypeStatus string, description string, test string, answerKey string, testDurationInSeconds int, autoGranted bool, autoGrantedValue int) (AWSMechanicalTurkRequester_xsd_go.TxsdUpdateQualificationTypeResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateQualificationType", qualificationTypeId, retryDelayInSeconds, qualificationTypeStatus, description, test, answerKey, testDurationInSeconds, autoGranted, autoGrantedValue)
	ret0, _ := ret[0].(AWSMechanicalTurkRequester_xsd_go.TxsdUpdateQualificationTypeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAmtClientRecorder) UpdateQualificationType(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateQualificationType", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}
