// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
	"time"
)

type JobInterviewStagesListSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *JobInterviewStagesListSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

// JobInterviewStagesListExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type JobInterviewStagesListExpand string

const (
	JobInterviewStagesListExpandJob JobInterviewStagesListExpand = "job"
)

func (e JobInterviewStagesListExpand) ToPointer() *JobInterviewStagesListExpand {
	return &e
}

func (e *JobInterviewStagesListExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "job":
		*e = JobInterviewStagesListExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobInterviewStagesListExpand: %v", v)
	}
}

type JobInterviewStagesListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobInterviewStagesListExpand `queryParam:"style=form,explode=true,name=expand"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `queryParam:"style=form,explode=true,name=include_deleted_data"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// If provided, will only return interview stages for this job.
	JobID *string `queryParam:"style=form,explode=true,name=job_id"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `queryParam:"style=form,explode=true,name=modified_after"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `queryParam:"style=form,explode=true,name=modified_before"`
	// Number of results to return per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
}

func (o *JobInterviewStagesListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *JobInterviewStagesListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *JobInterviewStagesListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *JobInterviewStagesListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *JobInterviewStagesListRequest) GetExpand() *JobInterviewStagesListExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *JobInterviewStagesListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *JobInterviewStagesListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *JobInterviewStagesListRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *JobInterviewStagesListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *JobInterviewStagesListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *JobInterviewStagesListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *JobInterviewStagesListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

type JobInterviewStagesListResponse struct {
	ContentType                    string
	PaginatedJobInterviewStageList *shared.PaginatedJobInterviewStageList
	StatusCode                     int
	RawResponse                    *http.Response
}

func (o *JobInterviewStagesListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *JobInterviewStagesListResponse) GetPaginatedJobInterviewStageList() *shared.PaginatedJobInterviewStageList {
	if o == nil {
		return nil
	}
	return o.PaginatedJobInterviewStageList
}

func (o *JobInterviewStagesListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *JobInterviewStagesListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
