// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type CandidatesMetaPostRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *CandidatesMetaPostRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type CandidatesMetaPostRetrieveResponse struct {
	// HTTP response content type for this operation
	ContentType  string
	MetaResponse *shared.MetaResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CandidatesMetaPostRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CandidatesMetaPostRetrieveResponse) GetMetaResponse() *shared.MetaResponse {
	if o == nil {
		return nil
	}
	return o.MetaResponse
}

func (o *CandidatesMetaPostRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CandidatesMetaPostRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
