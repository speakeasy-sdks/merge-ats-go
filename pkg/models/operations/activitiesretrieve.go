// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

// QueryParamExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type QueryParamExpand string

const (
	QueryParamExpandUser QueryParamExpand = "user"
)

func (e QueryParamExpand) ToPointer() *QueryParamExpand {
	return &e
}

func (e *QueryParamExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = QueryParamExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamExpand: %v", v)
	}
}

// QueryParamRemoteFields - Deprecated. Use show_enum_origins.
type QueryParamRemoteFields string

const (
	QueryParamRemoteFieldsActivityType           QueryParamRemoteFields = "activity_type"
	QueryParamRemoteFieldsActivityTypeVisibility QueryParamRemoteFields = "activity_type,visibility"
	QueryParamRemoteFieldsVisibility             QueryParamRemoteFields = "visibility"
)

func (e QueryParamRemoteFields) ToPointer() *QueryParamRemoteFields {
	return &e
}

func (e *QueryParamRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "activity_type":
		fallthrough
	case "activity_type,visibility":
		fallthrough
	case "visibility":
		*e = QueryParamRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamRemoteFields: %v", v)
	}
}

// QueryParamShowEnumOrigins - Which fields should be returned in non-normalized form.
type QueryParamShowEnumOrigins string

const (
	QueryParamShowEnumOriginsActivityType           QueryParamShowEnumOrigins = "activity_type"
	QueryParamShowEnumOriginsActivityTypeVisibility QueryParamShowEnumOrigins = "activity_type,visibility"
	QueryParamShowEnumOriginsVisibility             QueryParamShowEnumOrigins = "visibility"
)

func (e QueryParamShowEnumOrigins) ToPointer() *QueryParamShowEnumOrigins {
	return &e
}

func (e *QueryParamShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "activity_type":
		fallthrough
	case "activity_type,visibility":
		fallthrough
	case "visibility":
		*e = QueryParamShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamShowEnumOrigins: %v", v)
	}
}

type ActivitiesRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *QueryParamExpand `queryParam:"style=form,explode=true,name=expand"`
	ID     string            `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *QueryParamRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *QueryParamShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (o *ActivitiesRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *ActivitiesRetrieveRequest) GetExpand() *QueryParamExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *ActivitiesRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivitiesRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *ActivitiesRetrieveRequest) GetRemoteFields() *QueryParamRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *ActivitiesRetrieveRequest) GetShowEnumOrigins() *QueryParamShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

type ActivitiesRetrieveResponse struct {
	Activity *shared.Activity
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ActivitiesRetrieveResponse) GetActivity() *shared.Activity {
	if o == nil {
		return nil
	}
	return o.Activity
}

func (o *ActivitiesRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ActivitiesRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ActivitiesRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
