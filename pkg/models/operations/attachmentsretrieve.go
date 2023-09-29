// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"net/http"
)

type AttachmentsRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	expand *string `const:"candidate" queryParam:"style=form,explode=true,name=expand"`
	ID     string  `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	remoteFields *string `const:"attachment_type" queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	showEnumOrigins *string `const:"attachment_type" queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (a AttachmentsRetrieveRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttachmentsRetrieveRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AttachmentsRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *AttachmentsRetrieveRequest) GetExpand() *string {
	return types.String("candidate")
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

func (o *AttachmentsRetrieveRequest) GetRemoteFields() *string {
	return types.String("attachment_type")
}

func (o *AttachmentsRetrieveRequest) GetShowEnumOrigins() *string {
	return types.String("attachment_type")
}

type AttachmentsRetrieveResponse struct {
	Attachment *shared.Attachment
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
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
