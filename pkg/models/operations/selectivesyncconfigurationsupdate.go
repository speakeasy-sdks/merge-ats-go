// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type SelectiveSyncConfigurationsUpdateRequest struct {
	LinkedAccountSelectiveSyncConfigurationListRequest shared.LinkedAccountSelectiveSyncConfigurationListRequest `request:"mediaType=application/json"`
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *SelectiveSyncConfigurationsUpdateRequest) GetLinkedAccountSelectiveSyncConfigurationListRequest() shared.LinkedAccountSelectiveSyncConfigurationListRequest {
	if o == nil {
		return shared.LinkedAccountSelectiveSyncConfigurationListRequest{}
	}
	return o.LinkedAccountSelectiveSyncConfigurationListRequest
}

func (o *SelectiveSyncConfigurationsUpdateRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type SelectiveSyncConfigurationsUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Classes     []shared.LinkedAccountSelectiveSyncConfiguration
}

func (o *SelectiveSyncConfigurationsUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SelectiveSyncConfigurationsUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SelectiveSyncConfigurationsUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SelectiveSyncConfigurationsUpdateResponse) GetClasses() []shared.LinkedAccountSelectiveSyncConfiguration {
	if o == nil {
		return nil
	}
	return o.Classes
}
