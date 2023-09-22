// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type AccountDetailsRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *AccountDetailsRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type AccountDetailsRetrieveResponse struct {
	AccountDetails *shared.AccountDetails
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
}

func (o *AccountDetailsRetrieveResponse) GetAccountDetails() *shared.AccountDetails {
	if o == nil {
		return nil
	}
	return o.AccountDetails
}

func (o *AccountDetailsRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountDetailsRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountDetailsRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
