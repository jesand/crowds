package amt

import (
	"fmt"
	amtgen "github.com/jesand/crowds/amt/gen/mechanicalturk.amazonaws.com/AWSMechanicalTurk/2014-08-15/AWSMechanicalTurkRequester.xsd_go"
	xsdt "github.com/metaleap/go-xsd/types"
	"sort"
	"strings"
)

// ApproveAssignment approves the results of a completed assignment.
func (client AmtClient) ApproveAssignment(assignmentId,
	requesterFeedback string) (amtgen.TxsdApproveAssignmentResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdApproveAssignment
		args     amtgen.TApproveAssignmentRequest
		response amtgen.TxsdApproveAssignmentResponse
	)
	args.AssignmentId = xsdt.String(assignmentId)
	args.RequesterFeedback = xsdt.String(requesterFeedback)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("ApproveAssignment", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// ApproveRejectedAssignment approves an assignment that was previously
// rejected.
func (client AmtClient) ApproveRejectedAssignment(assignmentId,
	requesterFeedback string) (amtgen.TxsdApproveRejectedAssignmentResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdApproveRejectedAssignment
		args     amtgen.TApproveRejectedAssignmentRequest
		response amtgen.TxsdApproveRejectedAssignmentResponse
	)
	args.AssignmentId = xsdt.String(assignmentId)
	args.RequesterFeedback = xsdt.String(requesterFeedback)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("ApproveRejectedAssignment", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// AssignQualification gives a Worker a Qualification.
func (client AmtClient) AssignQualification(qualificationTypeId,
	workerId string, integerValue int, sendNotification bool) (
	amtgen.TxsdAssignQualificationResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdAssignQualification
		args     amtgen.TAssignQualificationRequest
		response amtgen.TxsdAssignQualificationResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.WorkerId = xsdt.String(workerId)
	args.IntegerValue = xsdt.Int(integerValue)
	args.SendNotification = xsdt.Boolean(sendNotification)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("AssignQualification", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// BlockWorker allows you to prevent a Worker from working on your HITs.
func (client AmtClient) BlockWorker(workerId, reason string) (
	amtgen.TxsdBlockWorkerResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdBlockWorker
		args     amtgen.TBlockWorkerRequest
		response amtgen.TxsdBlockWorkerResponse
	)
	args.WorkerId = xsdt.String(workerId)
	args.Reason = xsdt.String(reason)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("BlockWorker", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// ChangeHITTypeOfHIT allows you to change the HITType properties of a HIT.
func (client AmtClient) ChangeHITTypeOfHIT(hitId, hitTypeId string) (
	amtgen.TxsdChangeHITTypeOfHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdChangeHITTypeOfHIT
		args     amtgen.TChangeHITTypeOfHITRequest
		response amtgen.TxsdChangeHITTypeOfHITResponse
	)
	args.HITId = xsdt.String(hitId)
	args.HITTypeId = xsdt.String(hitTypeId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("ChangeHITTypeOfHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// CreateHIT creates a new Human Intelligence Task (HIT) without a HITTypeId.
func (client AmtClient) CreateHIT(title, description, question string,
	hitLayoutId string, hitLayoutParameters map[string]string,
	reward float32, assignmentDurationInSeconds,
	lifetimeInSeconds, maxAssignments, autoApprovalDelayInSeconds int,
	keywords []string,
	qualificationRequirements []*amtgen.TQualificationRequirement,
	assignmentReviewPolicy, hitReviewPolicy *amtgen.TReviewPolicy,
	requesterAnnotation, uniqueRequestToken string) (
	amtgen.TxsdCreateHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdCreateHIT
		args     amtgen.TCreateHITRequest
		response amtgen.TxsdCreateHITResponse
	)
	args.Title = xsdt.String(title)
	args.Description = xsdt.String(description)
	args.Question = xsdt.String(question)
	args.HITLayoutId = xsdt.String(hitLayoutId)
	var hitLayoutParameterOrder []string
	for name, _ := range hitLayoutParameters {
		hitLayoutParameterOrder = append(hitLayoutParameterOrder, name)
	}
	sort.Strings(hitLayoutParameterOrder)
	for _, name := range hitLayoutParameterOrder {
		value := hitLayoutParameters[name]
		var param amtgen.THITLayoutParameter
		param.Name = xsdt.String(name)
		param.Value = xsdt.String(value)
		args.HITLayoutParameters = append(args.HITLayoutParameters, &param)
	}
	args.Reward = &amtgen.TPrice{}
	args.Reward.Amount = xsdt.Decimal(fmt.Sprint(reward))
	args.Reward.CurrencyCode = CURRENCY_USD
	args.AssignmentDurationInSeconds = xsdt.Long(assignmentDurationInSeconds)
	args.LifetimeInSeconds = xsdt.Long(lifetimeInSeconds)
	args.MaxAssignments = xsdt.Int(maxAssignments)
	args.AutoApprovalDelayInSeconds = xsdt.Long(autoApprovalDelayInSeconds)
	args.Keywords = xsdt.String(strings.Join(keywords, ","))
	args.QualificationRequirements = qualificationRequirements
	args.AssignmentReviewPolicy = assignmentReviewPolicy
	args.HITReviewPolicy = hitReviewPolicy
	args.RequesterAnnotation = xsdt.String(requesterAnnotation)
	args.UniqueRequestToken = xsdt.String(uniqueRequestToken)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("CreateHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// CreateHITFromHITTypeId creates a new Human Intelligence Task (HIT) from a
// HITTypeId.
func (client AmtClient) CreateHITFromHITTypeId(hitTypeId, question string,
	hitLayoutId string, hitLayoutParameters map[string]string,
	lifetimeInSeconds, maxAssignments int,
	assignmentReviewPolicy, hitReviewPolicy *amtgen.TReviewPolicy,
	requesterAnnotation, uniqueRequestToken string) (
	amtgen.TxsdCreateHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdCreateHIT
		args     amtgen.TCreateHITRequest
		response amtgen.TxsdCreateHITResponse
	)
	args.HITTypeId = xsdt.String(hitTypeId)
	args.Question = xsdt.String(question)
	args.HITLayoutId = xsdt.String(hitLayoutId)
	var hitLayoutParameterOrder []string
	for name, _ := range hitLayoutParameters {
		hitLayoutParameterOrder = append(hitLayoutParameterOrder, name)
	}
	sort.Strings(hitLayoutParameterOrder)
	for _, name := range hitLayoutParameterOrder {
		value := hitLayoutParameters[name]
		var param amtgen.THITLayoutParameter
		param.Name = xsdt.String(name)
		param.Value = xsdt.String(value)
		args.HITLayoutParameters = append(args.HITLayoutParameters, &param)
	}
	args.LifetimeInSeconds = xsdt.Long(lifetimeInSeconds)
	args.MaxAssignments = xsdt.Int(maxAssignments)
	args.AssignmentReviewPolicy = assignmentReviewPolicy
	args.HITReviewPolicy = hitReviewPolicy
	args.RequesterAnnotation = xsdt.String(requesterAnnotation)
	args.UniqueRequestToken = xsdt.String(uniqueRequestToken)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("CreateHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// CreateQualificationType creates a new Qualification type.
func (client AmtClient) CreateQualificationType(name, description string,
	keywords []string, retryDelayInSeconds int,
	qualificationTypeStatus, test, answerKey string,
	testDurationInSeconds int, autoGranted bool,
	autoGrantedValue int) (amtgen.TxsdCreateQualificationTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdCreateQualificationType
		args     amtgen.TCreateQualificationTypeRequest
		response amtgen.TxsdCreateQualificationTypeResponse
	)
	args.Name = xsdt.String(name)
	args.Description = xsdt.String(description)
	args.Keywords = xsdt.String(strings.Join(keywords, ","))
	args.RetryDelayInSeconds = xsdt.Long(retryDelayInSeconds)
	args.QualificationTypeStatus = amtgen.TQualificationTypeStatus(
		qualificationTypeStatus)
	args.Test = xsdt.String(test)
	args.AnswerKey = xsdt.String(answerKey)
	args.TestDurationInSeconds = xsdt.Long(testDurationInSeconds)
	args.AutoGranted = xsdt.Boolean(autoGranted)
	args.AutoGrantedValue = xsdt.Int(autoGrantedValue)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("CreateQualificationType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// DisableHIT removes a HIT from the Amazon Mechanical Turk marketplace.
func (client AmtClient) DisableHIT(hitId string) (
	amtgen.TxsdDisableHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdDisableHIT
		args     amtgen.TDisableHITRequest
		response amtgen.TxsdDisableHITResponse
	)
	args.HITId = xsdt.String(hitId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("DisableHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// DisposeHIT disposes of a HIT that is no longer needed.
func (client AmtClient) DisposeHIT(hitId string) (
	amtgen.TxsdDisposeHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdDisposeHIT
		args     amtgen.TDisposeHITRequest
		response amtgen.TxsdDisposeHITResponse
	)
	args.HITId = xsdt.String(hitId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("DisposeHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// DisposeQualificationType disposes of a HIT that is no longer needed.
func (client AmtClient) DisposeQualificationType(qualificationTypeId string) (
	amtgen.TxsdDisposeQualificationTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdDisposeQualificationType
		args     amtgen.TDisposeQualificationTypeRequest
		response amtgen.TxsdDisposeQualificationTypeResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("DisposeQualificationType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// ExtendHIT increases the maximum number of assignments, or extends the
// expiration date, of an existing HIT.
func (client AmtClient) ExtendHIT(hitId string,
	maxAssignmentsIncrement, expirationIncrementInSeconds int,
	uniqueRequestToken string) (
	amtgen.TxsdExtendHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdExtendHIT
		args     amtgen.TExtendHITRequest
		response amtgen.TxsdExtendHITResponse
	)
	args.HITId = xsdt.String(hitId)
	args.MaxAssignmentsIncrement = xsdt.Int(maxAssignmentsIncrement)
	args.ExpirationIncrementInSeconds = xsdt.Long(expirationIncrementInSeconds)
	args.UniqueRequestToken = xsdt.String(uniqueRequestToken)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("ExtendHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// ForceExpireHIT causes a HIT to expire immediately, as if the
// LifetimeInSeconds parameter of the HIT had elapsed.
func (client AmtClient) ForceExpireHIT(hitId string) (
	amtgen.TxsdForceExpireHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdForceExpireHIT
		args     amtgen.TForceExpireHITRequest
		response amtgen.TxsdForceExpireHITResponse
	)
	args.HITId = xsdt.String(hitId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("ForceExpireHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetAccountBalance causes a HIT to expire immediately, as if the
// LifetimeInSeconds parameter of the HIT had elapsed.
func (client AmtClient) GetAccountBalance() (
	amtgen.TxsdGetAccountBalanceResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetAccountBalance
		args     amtgen.TGetAccountBalanceRequest
		response amtgen.TxsdGetAccountBalanceResponse
	)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetAccountBalance", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetAssignment retrieves an assignment with an AssignmentStatus value of
// Submitted, Approved, or Rejected.
func (client AmtClient) GetAssignment(assignmentId string) (
	amtgen.TxsdGetAssignmentResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetAssignment
		args     amtgen.TGetAssignmentRequest
		response amtgen.TxsdGetAssignmentResponse
	)
	args.AssignmentId = xsdt.String(assignmentId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetAssignment", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetAssignmentsForHIT retrieves completed assignments for a HIT.
func (client AmtClient) GetAssignmentsForHIT(hitId string,
	assignmentStatuses []string, sortProperty string, sortAscending bool,
	pageSize, pageNumber int) (
	amtgen.TxsdGetAssignmentsForHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetAssignmentsForHIT
		args     amtgen.TGetAssignmentsForHITRequest
		response amtgen.TxsdGetAssignmentsForHITResponse
	)
	args.HITId = xsdt.String(hitId)
	for _, status := range assignmentStatuses {
		args.AssignmentStatuses = append(args.AssignmentStatuses,
			amtgen.TAssignmentStatus(status))
	}
	args.SortProperty = amtgen.TGetAssignmentsForHITSortProperty(sortProperty)
	if sortAscending {
		args.SortDirection = amtgen.TSortDirection("Ascending")
	} else {
		args.SortDirection = amtgen.TSortDirection("Descending")
	}
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetAssignmentsForHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetBlockedWorkers retrieves a list of Workers who are blocked from working
// on your HITs.
func (client AmtClient) GetBlockedWorkers(pageSize, pageNumber int) (
	amtgen.TxsdGetBlockedWorkersResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetBlockedWorkers
		args     amtgen.TGetBlockedWorkersRequest
		response amtgen.TxsdGetBlockedWorkersResponse
	)
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetBlockedWorkers", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetBonusPayments retrieves the amounts of bonuses you have paid to Workers
// for a given HIT or assignment.
func (client AmtClient) GetBonusPayments(hitId, assignmentId string,
	pageSize, pageNumber int) (
	amtgen.TxsdGetBonusPaymentsResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetBonusPayments
		args     amtgen.TGetBonusPaymentsRequest
		response amtgen.TxsdGetBonusPaymentsResponse
	)
	args.HITId = xsdt.String(hitId)
	args.AssignmentId = xsdt.String(assignmentId)
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetBonusPayments", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetFileUploadURL generates and returns a temporary URL.
func (client AmtClient) GetFileUploadURL(assignmentId,
	questionIdentifier string) (
	amtgen.TxsdGetFileUploadURLResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetFileUploadURL
		args     amtgen.TGetFileUploadURLRequest
		response amtgen.TxsdGetFileUploadURLResponse
	)
	args.AssignmentId = xsdt.String(assignmentId)
	args.QuestionIdentifier = xsdt.String(questionIdentifier)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetFileUploadURL", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetHIT retrieves the details of the specified HIT.
func (client AmtClient) GetHIT(hitId string) (
	amtgen.TxsdGetHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetHIT
		args     amtgen.TGetHITRequest
		response amtgen.TxsdGetHITResponse
	)
	args.HITId = xsdt.String(hitId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetHITsForQualificationType returns the HITs that use the given Qualification
// type for a Qualification requirement.
func (client AmtClient) GetHITsForQualificationType(qualificationTypeId string,
	pageSize, pageNumber int) (
	amtgen.TxsdGetHITsForQualificationTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetHITsForQualificationType
		args     amtgen.TGetHITsForQualificationTypeRequest
		response amtgen.TxsdGetHITsForQualificationTypeResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetHITsForQualificationType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetQualificationsForQualificationType returns all of the Qualifications
// granted to Workers for a given Qualification type.
func (client AmtClient) GetQualificationsForQualificationType(
	qualificationTypeId string, isGranted bool,
	pageSize, pageNumber int) (
	amtgen.TxsdGetQualificationsForQualificationTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetQualificationsForQualificationType
		args     amtgen.TGetQualificationsForQualificationTypeRequest
		response amtgen.TxsdGetQualificationsForQualificationTypeResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	if isGranted {
		args.Status = amtgen.TQualificationStatus("Granted")
	} else {
		args.Status = amtgen.TQualificationStatus("Revoked")
	}
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetQualificationsForQualificationType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetQualificationRequests returns all of the Qualifications
// granted to Workers for a given Qualification type.
func (client AmtClient) GetQualificationRequests(
	qualificationTypeId, sortProperty string, sortAscending bool,
	pageSize, pageNumber int) (
	amtgen.TxsdGetQualificationRequestsResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetQualificationRequests
		args     amtgen.TGetQualificationRequestsRequest
		response amtgen.TxsdGetQualificationRequestsResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.SortProperty = amtgen.TGetQualificationRequestsSortProperty(
		sortProperty)
	if sortAscending {
		args.SortDirection = amtgen.TSortDirection("Ascending")
	} else {
		args.SortDirection = amtgen.TSortDirection("Descending")
	}
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetQualificationRequests", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetQualificationScore returns the value of a Worker's Qualification for a
// given Qualification type.
func (client AmtClient) GetQualificationScore(
	qualificationTypeId, subjectId string) (
	amtgen.TxsdGetQualificationScoreResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetQualificationScore
		args     amtgen.TGetQualificationScoreRequest
		response amtgen.TxsdGetQualificationScoreResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.SubjectId = xsdt.String(subjectId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetQualificationScore", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetQualificationType retrieves information about a Qualification type using
// its ID.
func (client AmtClient) GetQualificationType(qualificationTypeId string) (
	amtgen.TxsdGetQualificationTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetQualificationType
		args     amtgen.TGetQualificationTypeRequest
		response amtgen.TxsdGetQualificationTypeResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetQualificationType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetRequesterStatistic retrieves statistics about you (the Requester calling
// the operation).
func (client AmtClient) GetRequesterStatistic(statistic, timePeriod string,
	count int) (amtgen.TxsdGetRequesterStatisticResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetRequesterStatistic
		args     amtgen.TGetRequesterStatisticRequest
		response amtgen.TxsdGetRequesterStatisticResponse
	)
	args.Statistic = amtgen.TRequesterStatistic(statistic)
	args.TimePeriod = amtgen.TimePeriod(timePeriod)
	args.Count = xsdt.Int(count)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetRequesterStatistic", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetRequesterWorkerStatistic retrieves statistics about a specific Worker who
// has completed Human Intelligence Tasks (HITs) for you.
func (client AmtClient) GetRequesterWorkerStatistic(statistic, workerId,
	timePeriod string, count int) (amtgen.TxsdGetRequesterWorkerStatisticResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetRequesterWorkerStatistic
		args     amtgen.TGetRequesterWorkerStatisticRequest
		response amtgen.TxsdGetRequesterWorkerStatisticResponse
	)
	args.Statistic = amtgen.TRequesterStatistic(statistic)
	args.WorkerId = xsdt.String(workerId)
	args.TimePeriod = amtgen.TimePeriod(timePeriod)
	args.Count = xsdt.Int(count)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetRequesterWorkerStatistic", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetReviewableHITs retrieves the HITs with Status equal to Reviewable or
// Status equal to Reviewing that belong to the Requester calling the operation.
func (client AmtClient) GetReviewableHITs(hitTypeId, status,
	sortProperty string, sortAscending bool,
	pageSize, pageNumber int) (amtgen.TxsdGetReviewableHITsResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetReviewableHITs
		args     amtgen.TGetReviewableHITsRequest
		response amtgen.TxsdGetReviewableHITsResponse
	)
	args.HITTypeId = xsdt.String(hitTypeId)
	args.Status = amtgen.TReviewableHITStatus(status)
	args.SortProperty = amtgen.TGetReviewableHITsSortProperty(sortProperty)
	if sortAscending {
		args.SortDirection = amtgen.TSortDirection("Ascending")
	} else {
		args.SortDirection = amtgen.TSortDirection("Descending")
	}
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetReviewableHITs", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GetReviewResultsForHIT retrieves the computed results and the actions taken
// in the course of executing your Review Policies during a CreateHIT operation.
func (client AmtClient) GetReviewResultsForHIT(hitId string,
	policyLevels []string,
	retrieveActions, retrieveResults bool,
	pageSize, pageNumber int) (amtgen.TxsdGetReviewResultsForHITResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGetReviewResultsForHIT
		args     amtgen.TGetReviewResultsForHITRequest
		response amtgen.TxsdGetReviewResultsForHITResponse
	)
	args.HITId = xsdt.String(hitId)
	for _, level := range policyLevels {
		args.PolicyLevels = append(args.PolicyLevels,
			amtgen.TReviewPolicyLevel(level))
	}
	args.RetrieveActions = xsdt.Boolean(retrieveActions)
	args.RetrieveResults = xsdt.Boolean(retrieveResults)
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GetReviewResultsForHIT", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GrantBonus issues a payment of money from your account to a Worker.
func (client AmtClient) GrantBonus(workerId, assignmentId string,
	bonusAmount float32, reason, uniqueRequestToken string) (
	amtgen.TxsdGrantBonusResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGrantBonus
		args     amtgen.TGrantBonusRequest
		response amtgen.TxsdGrantBonusResponse
	)
	args.WorkerId = xsdt.String(workerId)
	args.AssignmentId = xsdt.String(assignmentId)
	args.BonusAmount = &amtgen.TPrice{}
	args.BonusAmount.Amount = xsdt.Decimal(fmt.Sprint(bonusAmount))
	args.BonusAmount.CurrencyCode = CURRENCY_USD
	args.Reason = xsdt.String(reason)
	args.UniqueRequestToken = xsdt.String(uniqueRequestToken)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GrantBonus", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// GrantQualification issues a payment of money from your account to a Worker.
func (client AmtClient) GrantQualification(qualificationRequestId string,
	integerValue int) (
	amtgen.TxsdGrantQualificationResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdGrantQualification
		args     amtgen.TGrantQualificationRequest
		response amtgen.TxsdGrantQualificationResponse
	)
	args.QualificationRequestId = xsdt.String(qualificationRequestId)
	args.IntegerValue = xsdt.Int(integerValue)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("GrantQualification", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// NotifyWorkers sends an email to one or more Workers that you specify with
// the Worker ID.
func (client AmtClient) NotifyWorkers(subject, messageText string,
	workerIds []string) (amtgen.TxsdNotifyWorkersResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdNotifyWorkers
		args     amtgen.TNotifyWorkersRequest
		response amtgen.TxsdNotifyWorkersResponse
	)
	args.Subject = xsdt.String(subject)
	args.MessageText = xsdt.String(messageText)
	for _, workerId := range workerIds {
		args.WorkerIds = append(args.WorkerIds, xsdt.String(workerId))
	}
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("NotifyWorkers", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// RegisterHITType creates a new HIT type.
func (client AmtClient) RegisterHITType(title, description string,
	reward float32, assignmentDurationInSeconds, autoApprovalDelayInSeconds int,
	keywords []string,
	qualificationRequirements []*amtgen.TQualificationRequirement) (
	amtgen.TxsdRegisterHITTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdRegisterHITType
		args     amtgen.TRegisterHITTypeRequest
		response amtgen.TxsdRegisterHITTypeResponse
	)
	args.Title = xsdt.String(title)
	args.Description = xsdt.String(description)
	args.Reward = &amtgen.TPrice{}
	args.Reward.Amount = xsdt.Decimal(fmt.Sprint(reward))
	args.Reward.CurrencyCode = CURRENCY_USD
	args.AssignmentDurationInSeconds = xsdt.Long(assignmentDurationInSeconds)
	args.AutoApprovalDelayInSeconds = xsdt.Long(autoApprovalDelayInSeconds)
	args.Keywords = xsdt.String(strings.Join(keywords, ","))
	args.QualificationRequirements = qualificationRequirements
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("RegisterHITType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// RejectAssignment rejects the results of a completed assignment.
func (client AmtClient) RejectAssignment(assignmentId,
	requesterFeedback string) (
	amtgen.TxsdRejectAssignmentResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdRejectAssignment
		args     amtgen.TRejectAssignmentRequest
		response amtgen.TxsdRejectAssignmentResponse
	)
	args.AssignmentId = xsdt.String(assignmentId)
	args.RequesterFeedback = xsdt.String(requesterFeedback)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("RejectAssignment", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// RejectQualificationRequest rejects a user's request for a Qualification.
func (client AmtClient) RejectQualificationRequest(qualificationRequestId,
	reason string) (
	amtgen.TxsdRejectQualificationRequestResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdRejectQualificationRequest
		args     amtgen.TRejectQualificationRequestRequest
		response amtgen.TxsdRejectQualificationRequestResponse
	)
	args.QualificationRequestId = xsdt.String(qualificationRequestId)
	args.Reason = xsdt.String(reason)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("RejectQualificationRequest", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// RevokeQualification revokes a previously granted Qualification from a user.
func (client AmtClient) RevokeQualification(subjectId, qualificationTypeId,
	reason string) (
	amtgen.TxsdRevokeQualificationResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdRevokeQualification
		args     amtgen.TRevokeQualificationRequest
		response amtgen.TxsdRevokeQualificationResponse
	)
	args.SubjectId = xsdt.String(subjectId)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.Reason = xsdt.String(reason)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("RevokeQualification", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// SearchHITs returns all of a Requester's HITs, on behalf of the Requester.
func (client AmtClient) SearchHITs(sortProperty string, sortAscending bool,
	pageSize, pageNumber int) (
	amtgen.TxsdSearchHITsResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdSearchHITs
		args     amtgen.TSearchHITsRequest
		response amtgen.TxsdSearchHITsResponse
	)
	args.SortProperty = amtgen.TSearchHITsSortProperty(sortProperty)
	if sortAscending {
		args.SortDirection = amtgen.TSortDirection("Ascending")
	} else {
		args.SortDirection = amtgen.TSortDirection("Descending")
	}
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("SearchHITs", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// SearchQualificationTypes searches for Qualification types using the specified
// search query, and returns a list of Qualification types.
func (client AmtClient) SearchQualificationTypes(
	query, sortProperty string, sortAscending bool,
	pageSize, pageNumber int, mustBeRequestable, mustBeOwnedByCaller bool) (
	amtgen.TxsdSearchQualificationTypesResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdSearchQualificationTypes
		args     amtgen.TSearchQualificationTypesRequest
		response amtgen.TxsdSearchQualificationTypesResponse
	)
	args.Query = xsdt.String(query)
	args.SortProperty = amtgen.TSearchQualificationTypesSortProperty(sortProperty)
	if sortAscending {
		args.SortDirection = amtgen.TSortDirection("Ascending")
	} else {
		args.SortDirection = amtgen.TSortDirection("Descending")
	}
	args.PageSize = xsdt.Int(pageSize)
	args.PageNumber = xsdt.Int(pageNumber)
	args.MustBeRequestable = xsdt.Boolean(mustBeRequestable)
	args.MustBeOwnedByCaller = xsdt.Boolean(mustBeOwnedByCaller)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("SearchQualificationTypes", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// SendTestEventNotification causes Amazon Mechanical Turk to send a
// notification message as if a HIT event occurred, according to the provided
// notification specification.
func (client AmtClient) SendTestEventNotification(
	notification *amtgen.TNotificationSpecification, testEventType string) (
	amtgen.TxsdSendTestEventNotificationResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdSendTestEventNotification
		args     amtgen.TSendTestEventNotificationRequest
		response amtgen.TxsdSendTestEventNotificationResponse
	)
	args.Notification = notification
	args.TestEventType = amtgen.TEventType(testEventType)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("SendTestEventNotification", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// SetHITAsReviewing updates the status of a HIT. If the status is Reviewable,
// this operation updates the status to Reviewing, or reverts a Reviewing HIT
// back to the Reviewable status.
func (client AmtClient) SetHITAsReviewing(hitID string, revert bool) (
	amtgen.TxsdSetHITAsReviewingResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdSetHITAsReviewing
		args     amtgen.TSetHITAsReviewingRequest
		response amtgen.TxsdSetHITAsReviewingResponse
	)
	args.HITId = xsdt.String(hitID)
	args.Revert = xsdt.Boolean(revert)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("SetHITAsReviewing", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// SetHITTypeNotification creates, updates, disables or re-enables notifications
// for a HIT type.
func (client AmtClient) SetHITTypeNotification(hitTypeID string,
	notification *amtgen.TNotificationSpecification, active bool) (
	amtgen.TxsdSetHITTypeNotificationResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdSetHITTypeNotification
		args     amtgen.TSetHITTypeNotificationRequest
		response amtgen.TxsdSetHITTypeNotificationResponse
	)
	args.HITTypeId = xsdt.String(hitTypeID)
	args.Notification = notification
	args.Active = xsdt.Boolean(active)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("SetHITTypeNotification", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// UnblockWorker allows you to reinstate a blocked Worker to work on your HITs.
func (client AmtClient) UnblockWorker(workerId, reason string) (
	amtgen.TxsdUnblockWorkerResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdUnblockWorker
		args     amtgen.TUnblockWorkerRequest
		response amtgen.TxsdUnblockWorkerResponse
	)
	args.WorkerId = xsdt.String(workerId)
	args.Reason = xsdt.String(reason)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("UnblockWorker", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// UpdateQualificationScore changes the value of a Qualification previously
// granted to a Worker.
func (client AmtClient) UpdateQualificationScore(qualificationTypeId,
	subjectId string, integerValue int) (
	amtgen.TxsdUpdateQualificationScoreResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdUpdateQualificationScore
		args     amtgen.TUpdateQualificationScoreRequest
		response amtgen.TxsdUpdateQualificationScoreResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.SubjectId = xsdt.String(subjectId)
	args.IntegerValue = xsdt.Int(integerValue)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("UpdateQualificationScore", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}

// UpdateQualificationType modifies the attributes of an existing Qualification
// type.
func (client AmtClient) UpdateQualificationType(qualificationTypeId string,
	retryDelayInSeconds int, qualificationTypeStatus, description, test,
	answerKey string, testDurationInSeconds int, autoGranted bool,
	autoGrantedValue int) (
	amtgen.TxsdUpdateQualificationTypeResponse, error) {

	// Prepare the request
	var (
		request  amtgen.TxsdUpdateQualificationType
		args     amtgen.TUpdateQualificationTypeRequest
		response amtgen.TxsdUpdateQualificationTypeResponse
	)
	args.QualificationTypeId = xsdt.String(qualificationTypeId)
	args.RetryDelayInSeconds = xsdt.Long(retryDelayInSeconds)
	args.QualificationTypeStatus = amtgen.TQualificationTypeStatus(
		qualificationTypeStatus)
	args.Description = xsdt.String(description)
	args.Test = xsdt.String(test)
	args.AnswerKey = xsdt.String(answerKey)
	args.TestDurationInSeconds = xsdt.Long(testDurationInSeconds)
	args.AutoGranted = xsdt.Boolean(autoGranted)
	args.AutoGrantedValue = xsdt.Int(autoGrantedValue)
	request.Requests = append(request.Requests, &args)

	// Send the request
	req, err := client.signRequest("UpdateQualificationType", &request)
	if err == nil {
		err = client.sendRequest(req, &response)
	}
	return response, err
}
