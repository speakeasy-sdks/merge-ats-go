// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

// UsersRetrieveRemoteFields - Deprecated. Use show_enum_origins.
type UsersRetrieveRemoteFields string

const (
	UsersRetrieveRemoteFieldsAccessRole UsersRetrieveRemoteFields = "access_role"
)

func (e UsersRetrieveRemoteFields) ToPointer() *UsersRetrieveRemoteFields {
	return &e
}

func (e *UsersRetrieveRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_role":
		*e = UsersRetrieveRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UsersRetrieveRemoteFields: %v", v)
	}
}

// UsersRetrieveShowEnumOrigins - Which fields should be returned in non-normalized form.
type UsersRetrieveShowEnumOrigins string

const (
	UsersRetrieveShowEnumOriginsAccessRole UsersRetrieveShowEnumOrigins = "access_role"
)

func (e UsersRetrieveShowEnumOrigins) ToPointer() *UsersRetrieveShowEnumOrigins {
	return &e
}

func (e *UsersRetrieveShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_role":
		*e = UsersRetrieveShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UsersRetrieveShowEnumOrigins: %v", v)
	}
}

type UsersRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	ID            string `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *UsersRetrieveRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *UsersRetrieveShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (o *UsersRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *UsersRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UsersRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *UsersRetrieveRequest) GetRemoteFields() *UsersRetrieveRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *UsersRetrieveRequest) GetShowEnumOrigins() *UsersRetrieveShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

type UsersRetrieveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	RemoteUser  *shared.RemoteUser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UsersRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UsersRetrieveResponse) GetRemoteUser() *shared.RemoteUser {
	if o == nil {
		return nil
	}
	return o.RemoteUser
}

func (o *UsersRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UsersRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
