// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type AttachmentsMetaPostRetrieveSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *AttachmentsMetaPostRetrieveSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

type AttachmentsMetaPostRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *AttachmentsMetaPostRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type AttachmentsMetaPostRetrieveResponse struct {
	ContentType  string
	MetaResponse *shared.MetaResponse
	StatusCode   int
	RawResponse  *http.Response
}

func (o *AttachmentsMetaPostRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AttachmentsMetaPostRetrieveResponse) GetMetaResponse() *shared.MetaResponse {
	if o == nil {
		return nil
	}
	return o.MetaResponse
}

func (o *AttachmentsMetaPostRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AttachmentsMetaPostRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
