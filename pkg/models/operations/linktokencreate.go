// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type LinkTokenCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	LinkToken   *shared.LinkToken
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
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
