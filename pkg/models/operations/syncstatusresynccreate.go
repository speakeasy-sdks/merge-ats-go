// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type SyncStatusResyncCreateSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *SyncStatusResyncCreateSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

type SyncStatusResyncCreateRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
}

func (o *SyncStatusResyncCreateRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

type SyncStatusResyncCreateResponse struct {
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
	SyncStatuses []shared.SyncStatus
}

func (o *SyncStatusResyncCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SyncStatusResyncCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SyncStatusResyncCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SyncStatusResyncCreateResponse) GetSyncStatuses() []shared.SyncStatus {
	if o == nil {
		return nil
	}
	return o.SyncStatuses
}