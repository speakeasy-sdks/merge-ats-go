// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type InterviewsCreateRequest struct {
	ScheduledInterviewEndpointRequest shared.ScheduledInterviewEndpointRequest `request:"mediaType=application/json"`
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `queryParam:"style=form,explode=true,name=is_debug_mode"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool `queryParam:"style=form,explode=true,name=run_async"`
}

func (o *InterviewsCreateRequest) GetScheduledInterviewEndpointRequest() shared.ScheduledInterviewEndpointRequest {
	if o == nil {
		return shared.ScheduledInterviewEndpointRequest{}
	}
	return o.ScheduledInterviewEndpointRequest
}

func (o *InterviewsCreateRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *InterviewsCreateRequest) GetIsDebugMode() *bool {
	if o == nil {
		return nil
	}
	return o.IsDebugMode
}

func (o *InterviewsCreateRequest) GetRunAsync() *bool {
	if o == nil {
		return nil
	}
	return o.RunAsync
}

type InterviewsCreateResponse struct {
	ContentType                string
	ScheduledInterviewResponse *shared.ScheduledInterviewResponse
	StatusCode                 int
	RawResponse                *http.Response
}

func (o *InterviewsCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InterviewsCreateResponse) GetScheduledInterviewResponse() *shared.ScheduledInterviewResponse {
	if o == nil {
		return nil
	}
	return o.ScheduledInterviewResponse
}

func (o *InterviewsCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InterviewsCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
