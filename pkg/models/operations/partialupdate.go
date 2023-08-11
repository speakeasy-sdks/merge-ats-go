// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type PartialUpdateSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *PartialUpdateSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

type PartialUpdateRequest struct {
	PatchedCandidateEndpointRequest shared.PatchedCandidateEndpointRequest `request:"mediaType=application/json"`
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	ID            string `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `queryParam:"style=form,explode=true,name=is_debug_mode"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool `queryParam:"style=form,explode=true,name=run_async"`
}

func (o *PartialUpdateRequest) GetPatchedCandidateEndpointRequest() shared.PatchedCandidateEndpointRequest {
	if o == nil {
		return shared.PatchedCandidateEndpointRequest{}
	}
	return o.PatchedCandidateEndpointRequest
}

func (o *PartialUpdateRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *PartialUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PartialUpdateRequest) GetIsDebugMode() *bool {
	if o == nil {
		return nil
	}
	return o.IsDebugMode
}

func (o *PartialUpdateRequest) GetRunAsync() *bool {
	if o == nil {
		return nil
	}
	return o.RunAsync
}

type PartialUpdateResponse struct {
	CandidateResponse *shared.CandidateResponse
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}

func (o *PartialUpdateResponse) GetCandidateResponse() *shared.CandidateResponse {
	if o == nil {
		return nil
	}
	return o.CandidateResponse
}

func (o *PartialUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PartialUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PartialUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}