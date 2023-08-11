// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type LinkTokenCreateSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *LinkTokenCreateSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

type LinkTokenCreateResponse struct {
	ContentType string
	LinkToken   *shared.LinkToken
	StatusCode  int
	RawResponse *http.Response
}

func (o *LinkTokenCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LinkTokenCreateResponse) GetLinkToken() *shared.LinkToken {
	if o == nil {
		return nil
	}
	return o.LinkToken
}

func (o *LinkTokenCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LinkTokenCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
