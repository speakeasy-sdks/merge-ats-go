// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type JobInterviewStagesRetrieveSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *JobInterviewStagesRetrieveSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

// JobInterviewStagesRetrieveExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type JobInterviewStagesRetrieveExpand string

const (
	JobInterviewStagesRetrieveExpandJob JobInterviewStagesRetrieveExpand = "job"
)

func (e JobInterviewStagesRetrieveExpand) ToPointer() *JobInterviewStagesRetrieveExpand {
	return &e
}

func (e *JobInterviewStagesRetrieveExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "job":
		*e = JobInterviewStagesRetrieveExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobInterviewStagesRetrieveExpand: %v", v)
	}
}

type JobInterviewStagesRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	ID            string `pathParam:"style=simple,explode=false,name=id"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobInterviewStagesRetrieveExpand `queryParam:"style=form,explode=true,name=expand"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
}

func (o *JobInterviewStagesRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *JobInterviewStagesRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *JobInterviewStagesRetrieveRequest) GetExpand() *JobInterviewStagesRetrieveExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *JobInterviewStagesRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

type JobInterviewStagesRetrieveResponse struct {
	ContentType       string
	JobInterviewStage *shared.JobInterviewStage
	StatusCode        int
	RawResponse       *http.Response
}

func (o *JobInterviewStagesRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *JobInterviewStagesRetrieveResponse) GetJobInterviewStage() *shared.JobInterviewStage {
	if o == nil {
		return nil
	}
	return o.JobInterviewStage
}

func (o *JobInterviewStagesRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *JobInterviewStagesRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
