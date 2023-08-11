// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
	"time"
)

type ApplicationsListSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *ApplicationsListSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

// ApplicationsListExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type ApplicationsListExpand string

const (
	ApplicationsListExpandCandidate                                      ApplicationsListExpand = "candidate"
	ApplicationsListExpandCandidateCreditedTo                            ApplicationsListExpand = "candidate,credited_to"
	ApplicationsListExpandCandidateCreditedToCurrentStage                ApplicationsListExpand = "candidate,credited_to,current_stage"
	ApplicationsListExpandCandidateCreditedToCurrentStageRejectReason    ApplicationsListExpand = "candidate,credited_to,current_stage,reject_reason"
	ApplicationsListExpandCandidateCreditedToRejectReason                ApplicationsListExpand = "candidate,credited_to,reject_reason"
	ApplicationsListExpandCandidateCurrentStage                          ApplicationsListExpand = "candidate,current_stage"
	ApplicationsListExpandCandidateCurrentStageRejectReason              ApplicationsListExpand = "candidate,current_stage,reject_reason"
	ApplicationsListExpandCandidateJob                                   ApplicationsListExpand = "candidate,job"
	ApplicationsListExpandCandidateJobCreditedTo                         ApplicationsListExpand = "candidate,job,credited_to"
	ApplicationsListExpandCandidateJobCreditedToCurrentStage             ApplicationsListExpand = "candidate,job,credited_to,current_stage"
	ApplicationsListExpandCandidateJobCreditedToCurrentStageRejectReason ApplicationsListExpand = "candidate,job,credited_to,current_stage,reject_reason"
	ApplicationsListExpandCandidateJobCreditedToRejectReason             ApplicationsListExpand = "candidate,job,credited_to,reject_reason"
	ApplicationsListExpandCandidateJobCurrentStage                       ApplicationsListExpand = "candidate,job,current_stage"
	ApplicationsListExpandCandidateJobCurrentStageRejectReason           ApplicationsListExpand = "candidate,job,current_stage,reject_reason"
	ApplicationsListExpandCandidateJobRejectReason                       ApplicationsListExpand = "candidate,job,reject_reason"
	ApplicationsListExpandCandidateRejectReason                          ApplicationsListExpand = "candidate,reject_reason"
	ApplicationsListExpandCreditedTo                                     ApplicationsListExpand = "credited_to"
	ApplicationsListExpandCreditedToCurrentStage                         ApplicationsListExpand = "credited_to,current_stage"
	ApplicationsListExpandCreditedToCurrentStageRejectReason             ApplicationsListExpand = "credited_to,current_stage,reject_reason"
	ApplicationsListExpandCreditedToRejectReason                         ApplicationsListExpand = "credited_to,reject_reason"
	ApplicationsListExpandCurrentStage                                   ApplicationsListExpand = "current_stage"
	ApplicationsListExpandCurrentStageRejectReason                       ApplicationsListExpand = "current_stage,reject_reason"
	ApplicationsListExpandJob                                            ApplicationsListExpand = "job"
	ApplicationsListExpandJobCreditedTo                                  ApplicationsListExpand = "job,credited_to"
	ApplicationsListExpandJobCreditedToCurrentStage                      ApplicationsListExpand = "job,credited_to,current_stage"
	ApplicationsListExpandJobCreditedToCurrentStageRejectReason          ApplicationsListExpand = "job,credited_to,current_stage,reject_reason"
	ApplicationsListExpandJobCreditedToRejectReason                      ApplicationsListExpand = "job,credited_to,reject_reason"
	ApplicationsListExpandJobCurrentStage                                ApplicationsListExpand = "job,current_stage"
	ApplicationsListExpandJobCurrentStageRejectReason                    ApplicationsListExpand = "job,current_stage,reject_reason"
	ApplicationsListExpandJobRejectReason                                ApplicationsListExpand = "job,reject_reason"
	ApplicationsListExpandRejectReason                                   ApplicationsListExpand = "reject_reason"
)

func (e ApplicationsListExpand) ToPointer() *ApplicationsListExpand {
	return &e
}

func (e *ApplicationsListExpand) UnmarshalJSON(data []byte) error {
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
		*e = ApplicationsListExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ApplicationsListExpand: %v", v)
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
	Expand *ApplicationsListExpand `queryParam:"style=form,explode=true,name=expand"`
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

func (o *ApplicationsListRequest) GetExpand() *ApplicationsListExpand {
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
	ContentType              string
	PaginatedApplicationList *shared.PaginatedApplicationList
	StatusCode               int
	RawResponse              *http.Response
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