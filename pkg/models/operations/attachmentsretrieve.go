// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type AttachmentsRetrieveSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *AttachmentsRetrieveSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

// AttachmentsRetrieveExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type AttachmentsRetrieveExpand string

const (
	AttachmentsRetrieveExpandCandidate AttachmentsRetrieveExpand = "candidate"
)

func (e AttachmentsRetrieveExpand) ToPointer() *AttachmentsRetrieveExpand {
	return &e
}

func (e *AttachmentsRetrieveExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "candidate":
		*e = AttachmentsRetrieveExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttachmentsRetrieveExpand: %v", v)
	}
}

// AttachmentsRetrieveRemoteFields - Deprecated. Use show_enum_origins.
type AttachmentsRetrieveRemoteFields string

const (
	AttachmentsRetrieveRemoteFieldsAttachmentType AttachmentsRetrieveRemoteFields = "attachment_type"
)

func (e AttachmentsRetrieveRemoteFields) ToPointer() *AttachmentsRetrieveRemoteFields {
	return &e
}

func (e *AttachmentsRetrieveRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "attachment_type":
		*e = AttachmentsRetrieveRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttachmentsRetrieveRemoteFields: %v", v)
	}
}

// AttachmentsRetrieveShowEnumOrigins - Which fields should be returned in non-normalized form.
type AttachmentsRetrieveShowEnumOrigins string

const (
	AttachmentsRetrieveShowEnumOriginsAttachmentType AttachmentsRetrieveShowEnumOrigins = "attachment_type"
)

func (e AttachmentsRetrieveShowEnumOrigins) ToPointer() *AttachmentsRetrieveShowEnumOrigins {
	return &e
}

func (e *AttachmentsRetrieveShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "attachment_type":
		*e = AttachmentsRetrieveShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttachmentsRetrieveShowEnumOrigins: %v", v)
	}
}

type AttachmentsRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *AttachmentsRetrieveExpand `queryParam:"style=form,explode=true,name=expand"`
	ID     string                     `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *AttachmentsRetrieveRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *AttachmentsRetrieveShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (o *AttachmentsRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *AttachmentsRetrieveRequest) GetExpand() *AttachmentsRetrieveExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *AttachmentsRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AttachmentsRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *AttachmentsRetrieveRequest) GetRemoteFields() *AttachmentsRetrieveRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *AttachmentsRetrieveRequest) GetShowEnumOrigins() *AttachmentsRetrieveShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

type AttachmentsRetrieveResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *AttachmentsRetrieveResponse) GetAttachment() *shared.Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}

func (o *AttachmentsRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AttachmentsRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AttachmentsRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}