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

// EeocsListQueryParamExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type EeocsListQueryParamExpand string

const (
	EeocsListQueryParamExpandCandidate EeocsListQueryParamExpand = "candidate"
)

func (e EeocsListQueryParamExpand) ToPointer() *EeocsListQueryParamExpand {
	return &e
}

func (e *EeocsListQueryParamExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "candidate":
		*e = EeocsListQueryParamExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EeocsListQueryParamExpand: %v", v)
	}
}

// EeocsListQueryParamRemoteFields - Deprecated. Use show_enum_origins.
type EeocsListQueryParamRemoteFields string

const (
	EeocsListQueryParamRemoteFieldsDisabilityStatus                        EeocsListQueryParamRemoteFields = "disability_status"
	EeocsListQueryParamRemoteFieldsDisabilityStatusGender                  EeocsListQueryParamRemoteFields = "disability_status,gender"
	EeocsListQueryParamRemoteFieldsDisabilityStatusGenderRace              EeocsListQueryParamRemoteFields = "disability_status,gender,race"
	EeocsListQueryParamRemoteFieldsDisabilityStatusGenderRaceVeteranStatus EeocsListQueryParamRemoteFields = "disability_status,gender,race,veteran_status"
	EeocsListQueryParamRemoteFieldsDisabilityStatusGenderVeteranStatus     EeocsListQueryParamRemoteFields = "disability_status,gender,veteran_status"
	EeocsListQueryParamRemoteFieldsDisabilityStatusRace                    EeocsListQueryParamRemoteFields = "disability_status,race"
	EeocsListQueryParamRemoteFieldsDisabilityStatusRaceVeteranStatus       EeocsListQueryParamRemoteFields = "disability_status,race,veteran_status"
	EeocsListQueryParamRemoteFieldsDisabilityStatusVeteranStatus           EeocsListQueryParamRemoteFields = "disability_status,veteran_status"
	EeocsListQueryParamRemoteFieldsGender                                  EeocsListQueryParamRemoteFields = "gender"
	EeocsListQueryParamRemoteFieldsGenderRace                              EeocsListQueryParamRemoteFields = "gender,race"
	EeocsListQueryParamRemoteFieldsGenderRaceVeteranStatus                 EeocsListQueryParamRemoteFields = "gender,race,veteran_status"
	EeocsListQueryParamRemoteFieldsGenderVeteranStatus                     EeocsListQueryParamRemoteFields = "gender,veteran_status"
	EeocsListQueryParamRemoteFieldsRace                                    EeocsListQueryParamRemoteFields = "race"
	EeocsListQueryParamRemoteFieldsRaceVeteranStatus                       EeocsListQueryParamRemoteFields = "race,veteran_status"
	EeocsListQueryParamRemoteFieldsVeteranStatus                           EeocsListQueryParamRemoteFields = "veteran_status"
)

func (e EeocsListQueryParamRemoteFields) ToPointer() *EeocsListQueryParamRemoteFields {
	return &e
}

func (e *EeocsListQueryParamRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disability_status":
		fallthrough
	case "disability_status,gender":
		fallthrough
	case "disability_status,gender,race":
		fallthrough
	case "disability_status,gender,race,veteran_status":
		fallthrough
	case "disability_status,gender,veteran_status":
		fallthrough
	case "disability_status,race":
		fallthrough
	case "disability_status,race,veteran_status":
		fallthrough
	case "disability_status,veteran_status":
		fallthrough
	case "gender":
		fallthrough
	case "gender,race":
		fallthrough
	case "gender,race,veteran_status":
		fallthrough
	case "gender,veteran_status":
		fallthrough
	case "race":
		fallthrough
	case "race,veteran_status":
		fallthrough
	case "veteran_status":
		*e = EeocsListQueryParamRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EeocsListQueryParamRemoteFields: %v", v)
	}
}

// EeocsListQueryParamShowEnumOrigins - Which fields should be returned in non-normalized form.
type EeocsListQueryParamShowEnumOrigins string

const (
	EeocsListQueryParamShowEnumOriginsDisabilityStatus                        EeocsListQueryParamShowEnumOrigins = "disability_status"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusGender                  EeocsListQueryParamShowEnumOrigins = "disability_status,gender"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusGenderRace              EeocsListQueryParamShowEnumOrigins = "disability_status,gender,race"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusGenderRaceVeteranStatus EeocsListQueryParamShowEnumOrigins = "disability_status,gender,race,veteran_status"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusGenderVeteranStatus     EeocsListQueryParamShowEnumOrigins = "disability_status,gender,veteran_status"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusRace                    EeocsListQueryParamShowEnumOrigins = "disability_status,race"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusRaceVeteranStatus       EeocsListQueryParamShowEnumOrigins = "disability_status,race,veteran_status"
	EeocsListQueryParamShowEnumOriginsDisabilityStatusVeteranStatus           EeocsListQueryParamShowEnumOrigins = "disability_status,veteran_status"
	EeocsListQueryParamShowEnumOriginsGender                                  EeocsListQueryParamShowEnumOrigins = "gender"
	EeocsListQueryParamShowEnumOriginsGenderRace                              EeocsListQueryParamShowEnumOrigins = "gender,race"
	EeocsListQueryParamShowEnumOriginsGenderRaceVeteranStatus                 EeocsListQueryParamShowEnumOrigins = "gender,race,veteran_status"
	EeocsListQueryParamShowEnumOriginsGenderVeteranStatus                     EeocsListQueryParamShowEnumOrigins = "gender,veteran_status"
	EeocsListQueryParamShowEnumOriginsRace                                    EeocsListQueryParamShowEnumOrigins = "race"
	EeocsListQueryParamShowEnumOriginsRaceVeteranStatus                       EeocsListQueryParamShowEnumOrigins = "race,veteran_status"
	EeocsListQueryParamShowEnumOriginsVeteranStatus                           EeocsListQueryParamShowEnumOrigins = "veteran_status"
)

func (e EeocsListQueryParamShowEnumOrigins) ToPointer() *EeocsListQueryParamShowEnumOrigins {
	return &e
}

func (e *EeocsListQueryParamShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disability_status":
		fallthrough
	case "disability_status,gender":
		fallthrough
	case "disability_status,gender,race":
		fallthrough
	case "disability_status,gender,race,veteran_status":
		fallthrough
	case "disability_status,gender,veteran_status":
		fallthrough
	case "disability_status,race":
		fallthrough
	case "disability_status,race,veteran_status":
		fallthrough
	case "disability_status,veteran_status":
		fallthrough
	case "gender":
		fallthrough
	case "gender,race":
		fallthrough
	case "gender,race,veteran_status":
		fallthrough
	case "gender,veteran_status":
		fallthrough
	case "race":
		fallthrough
	case "race,veteran_status":
		fallthrough
	case "veteran_status":
		*e = EeocsListQueryParamShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EeocsListQueryParamShowEnumOrigins: %v", v)
	}
}

type EeocsListRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// If provided, will only return EEOC info for this candidate.
	CandidateID *string `queryParam:"style=form,explode=true,name=candidate_id"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *EeocsListQueryParamExpand `queryParam:"style=form,explode=true,name=expand"`
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
	// Deprecated. Use show_enum_origins.
	RemoteFields *EeocsListQueryParamRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// The API provider's ID for the given object.
	RemoteID *string `queryParam:"style=form,explode=true,name=remote_id"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *EeocsListQueryParamShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (e EeocsListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EeocsListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EeocsListRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *EeocsListRequest) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *EeocsListRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *EeocsListRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *EeocsListRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *EeocsListRequest) GetExpand() *EeocsListQueryParamExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *EeocsListRequest) GetIncludeDeletedData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeletedData
}

func (o *EeocsListRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *EeocsListRequest) GetModifiedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAfter
}

func (o *EeocsListRequest) GetModifiedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedBefore
}

func (o *EeocsListRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *EeocsListRequest) GetRemoteFields() *EeocsListQueryParamRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *EeocsListRequest) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *EeocsListRequest) GetShowEnumOrigins() *EeocsListQueryParamShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

type EeocsListResponse struct {
	// HTTP response content type for this operation
	ContentType       string
	PaginatedEEOCList *shared.PaginatedEEOCList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *EeocsListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EeocsListResponse) GetPaginatedEEOCList() *shared.PaginatedEEOCList {
	if o == nil {
		return nil
	}
	return o.PaginatedEEOCList
}

func (o *EeocsListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EeocsListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
