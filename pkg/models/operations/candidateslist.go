// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
	"time"
)

type CandidatesListSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *CandidatesListSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

// CandidatesListExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type CandidatesListExpand string

const (
	CandidatesListExpandApplications            CandidatesListExpand = "applications"
	CandidatesListExpandApplicationsAttachments CandidatesListExpand = "applications,attachments"
	CandidatesListExpandAttachments             CandidatesListExpand = "attachments"
)

func (e CandidatesListExpand) ToPointer() *CandidatesListExpand {
	return &e
}

func (e *CandidatesListExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "applications":
		fallthrough
	case "applications,attachments":
		fallthrough
	case "attachments":
		*e = CandidatesListExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CandidatesListExpand: %v", v)
	}
}

type CandidatesListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// If provided, will only return candidates with these email addresses; multiple addresses can be separated by commas.
	EmailAddresses *string `queryParam:"style=form,explode=true,name=email_addresses"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *CandidatesListExpand `queryParam:"style=form,explode=true,name=expand"`
	// If provided, will only return candidates with this first name.
	FirstName *string `queryParam:"style=form,explode=true,name=first_name"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `queryParam:"style=form,explode=true,name=include_deleted_data"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// If provided, will only return candidates with this last name.
	LastName *string `queryParam:"style=form,explode=true,name=last_name"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `queryParam:"style=form,explode=true,name=modified_after"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `queryParam:"style=form,explode=true,name=modified_before"`
	// Number of results to return per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
	// If provided, will only return candidates with these tags; multiple tags can be separated by commas.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

func (o *CandidatesListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *CandidatesListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *CandidatesListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *CandidatesListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *CandidatesListRequest) GetEmailAddresses() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *CandidatesListRequest) GetExpand() *CandidatesListExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *CandidatesListRequest) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *CandidatesListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *CandidatesListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *CandidatesListRequest) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *CandidatesListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *CandidatesListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *CandidatesListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *CandidatesListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *CandidatesListRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type CandidatesListResponse struct {
	ContentType            string
	PaginatedCandidateList *shared.PaginatedCandidateList
	StatusCode             int
	RawResponse            *http.Response
}

func (o *CandidatesListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CandidatesListResponse) GetPaginatedCandidateList() *shared.PaginatedCandidateList {
	if o == nil {
		return nil
	}
	return o.PaginatedCandidateList
}

func (o *CandidatesListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CandidatesListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
