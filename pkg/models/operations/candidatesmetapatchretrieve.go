// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type CandidatesMetaPatchRetrieveSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *CandidatesMetaPatchRetrieveSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

type CandidatesMetaPatchRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	ID            string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *CandidatesMetaPatchRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *CandidatesMetaPatchRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type CandidatesMetaPatchRetrieveResponse struct {
	ContentType  string
	MetaResponse *shared.MetaResponse
	StatusCode   int
	RawResponse  *http.Response
}

func (o *CandidatesMetaPatchRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CandidatesMetaPatchRetrieveResponse) GetMetaResponse() *shared.MetaResponse {
	if o == nil {
		return nil
	}
	return o.MetaResponse
}

func (o *CandidatesMetaPatchRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CandidatesMetaPatchRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
