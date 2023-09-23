// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

// CandidatesRetrieveExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type CandidatesRetrieveExpand string

const (
	CandidatesRetrieveExpandApplications            CandidatesRetrieveExpand = "applications"
	CandidatesRetrieveExpandApplicationsAttachments CandidatesRetrieveExpand = "applications,attachments"
	CandidatesRetrieveExpandAttachments             CandidatesRetrieveExpand = "attachments"
)

func (e CandidatesRetrieveExpand) ToPointer() *CandidatesRetrieveExpand {
	return &e
}

func (e *CandidatesRetrieveExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "applications":
		fallthrough
	case "applications,attachments":
		fallthrough
	case "attachments":
		*e = CandidatesRetrieveExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CandidatesRetrieveExpand: %v", v)
	}
}

type CandidatesRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	ID            string `pathParam:"style=simple,explode=false,name=id"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *CandidatesRetrieveExpand `queryParam:"style=form,explode=true,name=expand"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
}

func (o *CandidatesRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *CandidatesRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CandidatesRetrieveRequest) GetExpand() *CandidatesRetrieveExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *CandidatesRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

type CandidatesRetrieveResponse struct {
	Candidate   *shared.Candidate
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *CandidatesRetrieveResponse) GetCandidate() *shared.Candidate {
	if o == nil {
		return nil
	}
	return o.Candidate
}

func (o *CandidatesRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CandidatesRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CandidatesRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
