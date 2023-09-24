// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type PassthroughCreateRequest struct {
	DataPassthroughRequest shared.DataPassthroughRequest `request:"mediaType=application/json"`
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *PassthroughCreateRequest) GetDataPassthroughRequest() shared.DataPassthroughRequest {
	if o == nil {
		return shared.DataPassthroughRequest{}
	}
	return o.DataPassthroughRequest
}

func (o *PassthroughCreateRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type PassthroughCreateResponse struct {
	ContentType    string
	RemoteResponse *shared.RemoteResponse
	StatusCode     int
	RawResponse    *http.Response
}

func (o *PassthroughCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PassthroughCreateResponse) GetRemoteResponse() *shared.RemoteResponse {
	if o == nil {
		return nil
	}
	return o.RemoteResponse
}

func (o *PassthroughCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PassthroughCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
