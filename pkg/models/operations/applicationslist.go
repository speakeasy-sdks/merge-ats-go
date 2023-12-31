// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"net/http"
	"time"
)

// ApplicationsListQueryParamExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type ApplicationsListQueryParamExpand string

const (
	ApplicationsListQueryParamExpandCandidate                                      ApplicationsListQueryParamExpand = "candidate"
	ApplicationsListQueryParamExpandCandidateCreditedTo                            ApplicationsListQueryParamExpand = "candidate,credited_to"
	ApplicationsListQueryParamExpandCandidateCreditedToCurrentStage                ApplicationsListQueryParamExpand = "candidate,credited_to,current_stage"
	ApplicationsListQueryParamExpandCandidateCreditedToCurrentStageRejectReason    ApplicationsListQueryParamExpand = "candidate,credited_to,current_stage,reject_reason"
	ApplicationsListQueryParamExpandCandidateCreditedToRejectReason                ApplicationsListQueryParamExpand = "candidate,credited_to,reject_reason"
	ApplicationsListQueryParamExpandCandidateCurrentStage                          ApplicationsListQueryParamExpand = "candidate,current_stage"
	ApplicationsListQueryParamExpandCandidateCurrentStageRejectReason              ApplicationsListQueryParamExpand = "candidate,current_stage,reject_reason"
	ApplicationsListQueryParamExpandCandidateJob                                   ApplicationsListQueryParamExpand = "candidate,job"
	ApplicationsListQueryParamExpandCandidateJobCreditedTo                         ApplicationsListQueryParamExpand = "candidate,job,credited_to"
	ApplicationsListQueryParamExpandCandidateJobCreditedToCurrentStage             ApplicationsListQueryParamExpand = "candidate,job,credited_to,current_stage"
	ApplicationsListQueryParamExpandCandidateJobCreditedToCurrentStageRejectReason ApplicationsListQueryParamExpand = "candidate,job,credited_to,current_stage,reject_reason"
	ApplicationsListQueryParamExpandCandidateJobCreditedToRejectReason             ApplicationsListQueryParamExpand = "candidate,job,credited_to,reject_reason"
	ApplicationsListQueryParamExpandCandidateJobCurrentStage                       ApplicationsListQueryParamExpand = "candidate,job,current_stage"
	ApplicationsListQueryParamExpandCandidateJobCurrentStageRejectReason           ApplicationsListQueryParamExpand = "candidate,job,current_stage,reject_reason"
	ApplicationsListQueryParamExpandCandidateJobRejectReason                       ApplicationsListQueryParamExpand = "candidate,job,reject_reason"
	ApplicationsListQueryParamExpandCandidateRejectReason                          ApplicationsListQueryParamExpand = "candidate,reject_reason"
	ApplicationsListQueryParamExpandCreditedTo                                     ApplicationsListQueryParamExpand = "credited_to"
	ApplicationsListQueryParamExpandCreditedToCurrentStage                         ApplicationsListQueryParamExpand = "credited_to,current_stage"
	ApplicationsListQueryParamExpandCreditedToCurrentStageRejectReason             ApplicationsListQueryParamExpand = "credited_to,current_stage,reject_reason"
	ApplicationsListQueryParamExpandCreditedToRejectReason                         ApplicationsListQueryParamExpand = "credited_to,reject_reason"
	ApplicationsListQueryParamExpandCurrentStage                                   ApplicationsListQueryParamExpand = "current_stage"
	ApplicationsListQueryParamExpandCurrentStageRejectReason                       ApplicationsListQueryParamExpand = "current_stage,reject_reason"
	ApplicationsListQueryParamExpandJob                                            ApplicationsListQueryParamExpand = "job"
	ApplicationsListQueryParamExpandJobCreditedTo                                  ApplicationsListQueryParamExpand = "job,credited_to"
	ApplicationsListQueryParamExpandJobCreditedToCurrentStage                      ApplicationsListQueryParamExpand = "job,credited_to,current_stage"
	ApplicationsListQueryParamExpandJobCreditedToCurrentStageRejectReason          ApplicationsListQueryParamExpand = "job,credited_to,current_stage,reject_reason"
	ApplicationsListQueryParamExpandJobCreditedToRejectReason                      ApplicationsListQueryParamExpand = "job,credited_to,reject_reason"
	ApplicationsListQueryParamExpandJobCurrentStage                                ApplicationsListQueryParamExpand = "job,current_stage"
	ApplicationsListQueryParamExpandJobCurrentStageRejectReason                    ApplicationsListQueryParamExpand = "job,current_stage,reject_reason"
	ApplicationsListQueryParamExpandJobRejectReason                                ApplicationsListQueryParamExpand = "job,reject_reason"
	ApplicationsListQueryParamExpandRejectReason                                   ApplicationsListQueryParamExpand = "reject_reason"
)

func (e ApplicationsListQueryParamExpand) ToPointer() *ApplicationsListQueryParamExpand {
	return &e
}

func (e *ApplicationsListQueryParamExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "candidate":
		fallthrough
	case "candidate,credited_to":
		fallthrough
	case "candidate,credited_to,current_stage":
		fallthrough
	case "candidate,credited_to,current_stage,reject_reason":
		fallthrough
	case "candidate,credited_to,reject_reason":
		fallthrough
	case "candidate,current_stage":
		fallthrough
	case "candidate,current_stage,reject_reason":
		fallthrough
	case "candidate,job":
		fallthrough
	case "candidate,job,credited_to":
		fallthrough
	case "candidate,job,credited_to,current_stage":
		fallthrough
	case "candidate,job,credited_to,current_stage,reject_reason":
		fallthrough
	case "candidate,job,credited_to,reject_reason":
		fallthrough
	case "candidate,job,current_stage":
		fallthrough
	case "candidate,job,current_stage,reject_reason":
		fallthrough
	case "candidate,job,reject_reason":
		fallthrough
	case "candidate,reject_reason":
		fallthrough
	case "credited_to":
		fallthrough
	case "credited_to,current_stage":
		fallthrough
	case "credited_to,current_stage,reject_reason":
		fallthrough
	case "credited_to,reject_reason":
		fallthrough
	case "current_stage":
		fallthrough
	case "current_stage,reject_reason":
		fallthrough
	case "job":
		fallthrough
	case "job,credited_to":
		fallthrough
	case "job,credited_to,current_stage":
		fallthrough
	case "job,credited_to,current_stage,reject_reason":
		fallthrough
	case "job,credited_to,reject_reason":
		fallthrough
	case "job,current_stage":
		fallthrough
	case "job,current_stage,reject_reason":
		fallthrough
	case "job,reject_reason":
		fallthrough
	case "reject_reason":
		*e = ApplicationsListQueryParamExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplicationsListQueryParamExpand: %v", v)
	}
}

type ApplicationsListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return applications for this candidate.
	CandidateID *string `queryParam:"style=form,explode=true,name=candidate_id"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// If provided, will only return applications credited to this user.
	CreditedToID *string `queryParam:"style=form,explode=true,name=credited_to_id"`
	// If provided, will only return applications at this interview stage.
	CurrentStageID *string `queryParam:"style=form,explode=true,name=current_stage_id"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ApplicationsListQueryParamExpand `queryParam:"style=form,explode=true,name=expand"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `queryParam:"style=form,explode=true,name=include_deleted_data"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// If provided, will only return applications for this job.
	JobID *string `queryParam:"style=form,explode=true,name=job_id"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `queryParam:"style=form,explode=true,name=modified_after"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `queryParam:"style=form,explode=true,name=modified_before"`
	// Number of results to return per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// If provided, will only return applications with this reject reason.
	RejectReasonID *string `queryParam:"style=form,explode=true,name=reject_reason_id"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
	// If provided, will only return applications with this source.
	Source *string `queryParam:"style=form,explode=true,name=source"`
}

func (a ApplicationsListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationsListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationsListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *ApplicationsListRequest) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *ApplicationsListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *ApplicationsListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *ApplicationsListRequest) GetCreditedToID() *string {
	if o == nil {
		return nil
	}
	return o.CreditedToID
}

func (o *ApplicationsListRequest) GetCurrentStageID() *string {
	if o == nil {
		return nil
	}
	return o.CurrentStageID
}

func (o *ApplicationsListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ApplicationsListRequest) GetExpand() *ApplicationsListQueryParamExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *ApplicationsListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *ApplicationsListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *ApplicationsListRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *ApplicationsListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *ApplicationsListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *ApplicationsListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ApplicationsListRequest) GetRejectReasonID() *string {
	if o == nil {
		return nil
	}
	return o.RejectReasonID
}

func (o *ApplicationsListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *ApplicationsListRequest) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

type ApplicationsListResponse struct {
	// HTTP response content type for this operation
	ContentType              string
	PaginatedApplicationList *shared.PaginatedApplicationList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ApplicationsListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplicationsListResponse) GetPaginatedApplicationList() *shared.PaginatedApplicationList {
	if o == nil {
		return nil
	}
	return o.PaginatedApplicationList
}

func (o *ApplicationsListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplicationsListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
