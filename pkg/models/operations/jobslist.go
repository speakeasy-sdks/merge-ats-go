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

// JobsListExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type JobsListExpand string

const (
	JobsListExpandDepartments                                JobsListExpand = "departments"
	JobsListExpandDepartmentsHiringManagers                  JobsListExpand = "departments,hiring_managers"
	JobsListExpandDepartmentsHiringManagersRecruiters        JobsListExpand = "departments,hiring_managers,recruiters"
	JobsListExpandDepartmentsOffices                         JobsListExpand = "departments,offices"
	JobsListExpandDepartmentsOfficesHiringManagers           JobsListExpand = "departments,offices,hiring_managers"
	JobsListExpandDepartmentsOfficesHiringManagersRecruiters JobsListExpand = "departments,offices,hiring_managers,recruiters"
	JobsListExpandDepartmentsOfficesRecruiters               JobsListExpand = "departments,offices,recruiters"
	JobsListExpandDepartmentsRecruiters                      JobsListExpand = "departments,recruiters"
	JobsListExpandHiringManagers                             JobsListExpand = "hiring_managers"
	JobsListExpandHiringManagersRecruiters                   JobsListExpand = "hiring_managers,recruiters"
	JobsListExpandOffices                                    JobsListExpand = "offices"
	JobsListExpandOfficesHiringManagers                      JobsListExpand = "offices,hiring_managers"
	JobsListExpandOfficesHiringManagersRecruiters            JobsListExpand = "offices,hiring_managers,recruiters"
	JobsListExpandOfficesRecruiters                          JobsListExpand = "offices,recruiters"
	JobsListExpandRecruiters                                 JobsListExpand = "recruiters"
)

func (e JobsListExpand) ToPointer() *JobsListExpand {
	return &e
}

func (e *JobsListExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "departments":
		fallthrough
	case "departments,hiring_managers":
		fallthrough
	case "departments,hiring_managers,recruiters":
		fallthrough
	case "departments,offices":
		fallthrough
	case "departments,offices,hiring_managers":
		fallthrough
	case "departments,offices,hiring_managers,recruiters":
		fallthrough
	case "departments,offices,recruiters":
		fallthrough
	case "departments,recruiters":
		fallthrough
	case "hiring_managers":
		fallthrough
	case "hiring_managers,recruiters":
		fallthrough
	case "offices":
		fallthrough
	case "offices,hiring_managers":
		fallthrough
	case "offices,hiring_managers,recruiters":
		fallthrough
	case "offices,recruiters":
		fallthrough
	case "recruiters":
		*e = JobsListExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobsListExpand: %v", v)
	}
}

// JobsListRemoteFields - Deprecated. Use show_enum_origins.
type JobsListRemoteFields string

const (
	JobsListRemoteFieldsStatus JobsListRemoteFields = "status"
)

func (e JobsListRemoteFields) ToPointer() *JobsListRemoteFields {
	return &e
}

func (e *JobsListRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "status":
		*e = JobsListRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobsListRemoteFields: %v", v)
	}
}

// JobsListShowEnumOrigins - Which fields should be returned in non-normalized form.
type JobsListShowEnumOrigins string

const (
	JobsListShowEnumOriginsStatus JobsListShowEnumOrigins = "status"
)

func (e JobsListShowEnumOrigins) ToPointer() *JobsListShowEnumOrigins {
	return &e
}

func (e *JobsListShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "status":
		*e = JobsListShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobsListShowEnumOrigins: %v", v)
	}
}

// JobsListStatus - If provided, will only return jobs with this status. Options: ('OPEN', 'CLOSED', 'DRAFT', 'ARCHIVED', 'PENDING')
//
// * `OPEN` - OPEN
// * `CLOSED` - CLOSED
// * `DRAFT` - DRAFT
// * `ARCHIVED` - ARCHIVED
// * `PENDING` - PENDING
type JobsListStatus string

const (
	JobsListStatusArchived JobsListStatus = "ARCHIVED"
	JobsListStatusClosed   JobsListStatus = "CLOSED"
	JobsListStatusDraft    JobsListStatus = "DRAFT"
	JobsListStatusOpen     JobsListStatus = "OPEN"
	JobsListStatusPending  JobsListStatus = "PENDING"
)

func (e JobsListStatus) ToPointer() *JobsListStatus {
	return &e
}

func (e *JobsListStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ARCHIVED":
		fallthrough
	case "CLOSED":
		fallthrough
	case "DRAFT":
		fallthrough
	case "OPEN":
		fallthrough
	case "PENDING":
		*e = JobsListStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobsListStatus: %v", v)
	}
}

type JobsListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return jobs with this code.
	Code *string `queryParam:"style=form,explode=true,name=code"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobsListExpand `queryParam:"style=form,explode=true,name=expand"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `queryParam:"style=form,explode=true,name=include_deleted_data"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `queryParam:"style=form,explode=true,name=modified_after"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `queryParam:"style=form,explode=true,name=modified_before"`
	// If provided, will only return jobs for this office; multiple offices can be separated by commas.
	Offices *string `queryParam:"style=form,explode=true,name=offices"`
	// Number of results to return per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *JobsListRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *JobsListShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
	// If provided, will only return jobs with this status. Options: ('OPEN', 'CLOSED', 'DRAFT', 'ARCHIVED', 'PENDING')
	//
	// * `OPEN` - OPEN
	// * `CLOSED` - CLOSED
	// * `DRAFT` - DRAFT
	// * `ARCHIVED` - ARCHIVED
	// * `PENDING` - PENDING
	Status *JobsListStatus `queryParam:"style=form,explode=true,name=status"`
}

func (j JobsListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JobsListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JobsListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *JobsListRequest) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *JobsListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *JobsListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *JobsListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *JobsListRequest) GetExpand() *JobsListExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *JobsListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *JobsListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *JobsListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *JobsListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *JobsListRequest) GetOffices() *string {
	if o == nil {
		return nil
	}
	return o.Offices
}

func (o *JobsListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *JobsListRequest) GetRemoteFields() *JobsListRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *JobsListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *JobsListRequest) GetShowEnumOrigins() *JobsListShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

func (o *JobsListRequest) GetStatus() *JobsListStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type JobsListResponse struct {
	// HTTP response content type for this operation
	ContentType      string
	PaginatedJobList *shared.PaginatedJobList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *JobsListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *JobsListResponse) GetPaginatedJobList() *shared.PaginatedJobList {
	if o == nil {
		return nil
	}
	return o.PaginatedJobList
}

func (o *JobsListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *JobsListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
