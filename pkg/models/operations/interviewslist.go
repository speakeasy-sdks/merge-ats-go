// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"net/http"
	"time"
)

// InterviewsListExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type InterviewsListExpand string

const (
	InterviewsListExpandApplication                                       InterviewsListExpand = "application"
	InterviewsListExpandApplicationJobInterviewStage                      InterviewsListExpand = "application,job_interview_stage"
	InterviewsListExpandInterviewers                                      InterviewsListExpand = "interviewers"
	InterviewsListExpandInterviewersApplication                           InterviewsListExpand = "interviewers,application"
	InterviewsListExpandInterviewersApplicationJobInterviewStage          InterviewsListExpand = "interviewers,application,job_interview_stage"
	InterviewsListExpandInterviewersJobInterviewStage                     InterviewsListExpand = "interviewers,job_interview_stage"
	InterviewsListExpandInterviewersOrganizer                             InterviewsListExpand = "interviewers,organizer"
	InterviewsListExpandInterviewersOrganizerApplication                  InterviewsListExpand = "interviewers,organizer,application"
	InterviewsListExpandInterviewersOrganizerApplicationJobInterviewStage InterviewsListExpand = "interviewers,organizer,application,job_interview_stage"
	InterviewsListExpandInterviewersOrganizerJobInterviewStage            InterviewsListExpand = "interviewers,organizer,job_interview_stage"
	InterviewsListExpandJobInterviewStage                                 InterviewsListExpand = "job_interview_stage"
	InterviewsListExpandOrganizer                                         InterviewsListExpand = "organizer"
	InterviewsListExpandOrganizerApplication                              InterviewsListExpand = "organizer,application"
	InterviewsListExpandOrganizerApplicationJobInterviewStage             InterviewsListExpand = "organizer,application,job_interview_stage"
	InterviewsListExpandOrganizerJobInterviewStage                        InterviewsListExpand = "organizer,job_interview_stage"
)

func (e InterviewsListExpand) ToPointer() *InterviewsListExpand {
	return &e
}

func (e *InterviewsListExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "application":
		fallthrough
	case "application,job_interview_stage":
		fallthrough
	case "interviewers":
		fallthrough
	case "interviewers,application":
		fallthrough
	case "interviewers,application,job_interview_stage":
		fallthrough
	case "interviewers,job_interview_stage":
		fallthrough
	case "interviewers,organizer":
		fallthrough
	case "interviewers,organizer,application":
		fallthrough
	case "interviewers,organizer,application,job_interview_stage":
		fallthrough
	case "interviewers,organizer,job_interview_stage":
		fallthrough
	case "job_interview_stage":
		fallthrough
	case "organizer":
		fallthrough
	case "organizer,application":
		fallthrough
	case "organizer,application,job_interview_stage":
		fallthrough
	case "organizer,job_interview_stage":
		*e = InterviewsListExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InterviewsListExpand: %v", v)
	}
}

type InterviewsListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return interviews for this application.
	ApplicationID *string `queryParam:"style=form,explode=true,name=application_id"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *InterviewsListExpand `queryParam:"style=form,explode=true,name=expand"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `queryParam:"style=form,explode=true,name=include_deleted_data"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// If provided, will only return interviews at this stage.
	JobInterviewStageID *string `queryParam:"style=form,explode=true,name=job_interview_stage_id"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `queryParam:"style=form,explode=true,name=modified_after"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `queryParam:"style=form,explode=true,name=modified_before"`
	// If provided, will only return interviews organized by this user.
	OrganizerID *string `queryParam:"style=form,explode=true,name=organizer_id"`
	// Number of results to return per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// Deprecated. Use show_enum_origins.
	remoteFields *string `const:"status" queryParam:"style=form,explode=true,name=remote_fields"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
	// Which fields should be returned in non-normalized form.
	showEnumOrigins *string `const:"status" queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (i InterviewsListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InterviewsListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InterviewsListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *InterviewsListRequest) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *InterviewsListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *InterviewsListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *InterviewsListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *InterviewsListRequest) GetExpand() *InterviewsListExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *InterviewsListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *InterviewsListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *InterviewsListRequest) GetJobInterviewStageID() *string {
	if o == nil {
		return nil
	}
	return o.JobInterviewStageID
}

func (o *InterviewsListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *InterviewsListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *InterviewsListRequest) GetOrganizerID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizerID
}

func (o *InterviewsListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *InterviewsListRequest) GetRemoteFields() *string {
	return types.String("status")
}

func (o *InterviewsListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *InterviewsListRequest) GetShowEnumOrigins() *string {
	return types.String("status")
}

type InterviewsListResponse struct {
	// HTTP response content type for this operation
	ContentType                     string
	PaginatedScheduledInterviewList *shared.PaginatedScheduledInterviewList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *InterviewsListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InterviewsListResponse) GetPaginatedScheduledInterviewList() *shared.PaginatedScheduledInterviewList {
	if o == nil {
		return nil
	}
	return o.PaginatedScheduledInterviewList
}

func (o *InterviewsListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InterviewsListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
