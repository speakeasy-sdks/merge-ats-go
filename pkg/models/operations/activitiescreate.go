// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type ActivitiesCreateRequest struct {
	ActivityEndpointRequest shared.ActivityEndpointRequest `request:"mediaType=application/json"`
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `queryParam:"style=form,explode=true,name=is_debug_mode"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool `queryParam:"style=form,explode=true,name=run_async"`
}

func (o *ActivitiesCreateRequest) GetActivityEndpointRequest() shared.ActivityEndpointRequest {
	if o == nil {
		return shared.ActivityEndpointRequest{}
	}
	return o.ActivityEndpointRequest
}

func (o *ActivitiesCreateRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *ActivitiesCreateRequest) GetIsDebugMode() *bool {
	if o == nil {
		return nil
	}
	return o.IsDebugMode
}

func (o *ActivitiesCreateRequest) GetRunAsync() *bool {
	if o == nil {
		return nil
	}
	return o.RunAsync
}

type ActivitiesCreateResponse struct {
	ActivityResponse *shared.ActivityResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ActivitiesCreateResponse) GetActivityResponse() *shared.ActivityResponse {
	if o == nil {
		return nil
	}
	return o.ActivityResponse
}

func (o *ActivitiesCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ActivitiesCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ActivitiesCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
