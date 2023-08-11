// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type AsyncPassthroughCreateJSONSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *AsyncPassthroughCreateJSONSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

type AsyncPassthroughCreateJSONRequest struct {
	DataPassthroughRequest shared.DataPassthroughRequest `request:"mediaType=application/json"`
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *AsyncPassthroughCreateJSONRequest) GetDataPassthroughRequest() shared.DataPassthroughRequest {
	if o == nil {
		return shared.DataPassthroughRequest{}
	}
	return o.DataPassthroughRequest
}

func (o *AsyncPassthroughCreateJSONRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type AsyncPassthroughCreateJSONResponse struct {
	AsyncPassthroughReciept *shared.AsyncPassthroughReciept
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
}

func (o *AsyncPassthroughCreateJSONResponse) GetAsyncPassthroughReciept() *shared.AsyncPassthroughReciept {
	if o == nil {
		return nil
	}
	return o.AsyncPassthroughReciept
}

func (o *AsyncPassthroughCreateJSONResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AsyncPassthroughCreateJSONResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AsyncPassthroughCreateJSONResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
