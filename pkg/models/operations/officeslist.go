// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"net/http"
	"time"
)

type OfficesListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `queryParam:"style=form,explode=true,name=include_deleted_data"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `queryParam:"style=form,explode=true,name=modified_after"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `queryParam:"style=form,explode=true,name=modified_before"`
	// Number of results to return per page.
	PageSize *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
}

func (o OfficesListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OfficesListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OfficesListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *OfficesListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *OfficesListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *OfficesListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *OfficesListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *OfficesListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *OfficesListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *OfficesListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *OfficesListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *OfficesListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

type OfficesListResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	PaginatedOfficeList *shared.PaginatedOfficeList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *OfficesListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *OfficesListResponse) GetPaginatedOfficeList() *shared.PaginatedOfficeList {
	if o == nil {
		return nil
	}
	return o.PaginatedOfficeList
}

func (o *OfficesListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *OfficesListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
