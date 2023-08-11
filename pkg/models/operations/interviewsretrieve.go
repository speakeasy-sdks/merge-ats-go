// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"net/http"
)

type InterviewsRetrieveSecurity struct {
	TokenAuth string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *InterviewsRetrieveSecurity) GetTokenAuth() string {
	if o == nil {
		return ""
	}
	return o.TokenAuth
}

// InterviewsRetrieveExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type InterviewsRetrieveExpand string

const (
	InterviewsRetrieveExpandApplication                                       InterviewsRetrieveExpand = "application"
	InterviewsRetrieveExpandApplicationJobInterviewStage                      InterviewsRetrieveExpand = "application,job_interview_stage"
	InterviewsRetrieveExpandInterviewers                                      InterviewsRetrieveExpand = "interviewers"
	InterviewsRetrieveExpandInterviewersApplication                           InterviewsRetrieveExpand = "interviewers,application"
	InterviewsRetrieveExpandInterviewersApplicationJobInterviewStage          InterviewsRetrieveExpand = "interviewers,application,job_interview_stage"
	InterviewsRetrieveExpandInterviewersJobInterviewStage                     InterviewsRetrieveExpand = "interviewers,job_interview_stage"
	InterviewsRetrieveExpandInterviewersOrganizer                             InterviewsRetrieveExpand = "interviewers,organizer"
	InterviewsRetrieveExpandInterviewersOrganizerApplication                  InterviewsRetrieveExpand = "interviewers,organizer,application"
	InterviewsRetrieveExpandInterviewersOrganizerApplicationJobInterviewStage InterviewsRetrieveExpand = "interviewers,organizer,application,job_interview_stage"
	InterviewsRetrieveExpandInterviewersOrganizerJobInterviewStage            InterviewsRetrieveExpand = "interviewers,organizer,job_interview_stage"
	InterviewsRetrieveExpandJobInterviewStage                                 InterviewsRetrieveExpand = "job_interview_stage"
	InterviewsRetrieveExpandOrganizer                                         InterviewsRetrieveExpand = "organizer"
	InterviewsRetrieveExpandOrganizerApplication                              InterviewsRetrieveExpand = "organizer,application"
	InterviewsRetrieveExpandOrganizerApplicationJobInterviewStage             InterviewsRetrieveExpand = "organizer,application,job_interview_stage"
	InterviewsRetrieveExpandOrganizerJobInterviewStage                        InterviewsRetrieveExpand = "organizer,job_interview_stage"
)

func (e InterviewsRetrieveExpand) ToPointer() *InterviewsRetrieveExpand {
	return &e
}

func (e *InterviewsRetrieveExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "application":
		fallthrough
	case "application,job_interview_stage":
		fallthrough
	case "interviewers":
		fallthrough
	case "interviewers,application":
		fallthrough
	case "interviewers,application,job_interview_stage":
		fallthrough
	case "interviewers,job_interview_stage":
		fallthrough
	case "interviewers,organizer":
		fallthrough
	case "interviewers,organizer,application":
		fallthrough
	case "interviewers,organizer,application,job_interview_stage":
		fallthrough
	case "interviewers,organizer,job_interview_stage":
		fallthrough
	case "job_interview_stage":
		fallthrough
	case "organizer":
		fallthrough
	case "organizer,application":
		fallthrough
	case "organizer,application,job_interview_stage":
		fallthrough
	case "organizer,job_interview_stage":
		*e = InterviewsRetrieveExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InterviewsRetrieveExpand: %v", v)
	}
}

// InterviewsRetrieveRemoteFields - Deprecated. Use show_enum_origins.
type InterviewsRetrieveRemoteFields string

const (
	InterviewsRetrieveRemoteFieldsStatus InterviewsRetrieveRemoteFields = "status"
)

func (e InterviewsRetrieveRemoteFields) ToPointer() *InterviewsRetrieveRemoteFields {
	return &e
}

func (e *InterviewsRetrieveRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "status":
		*e = InterviewsRetrieveRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InterviewsRetrieveRemoteFields: %v", v)
	}
}

// InterviewsRetrieveShowEnumOrigins - Which fields should be returned in non-normalized form.
type InterviewsRetrieveShowEnumOrigins string

const (
	InterviewsRetrieveShowEnumOriginsStatus InterviewsRetrieveShowEnumOrigins = "status"
)

func (e InterviewsRetrieveShowEnumOrigins) ToPointer() *InterviewsRetrieveShowEnumOrigins {
	return &e
}

func (e *InterviewsRetrieveShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "status":
		*e = InterviewsRetrieveShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InterviewsRetrieveShowEnumOrigins: %v", v)
	}
}

type InterviewsRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *InterviewsRetrieveExpand `queryParam:"style=form,explode=true,name=expand"`
	ID     string                    `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *InterviewsRetrieveRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *InterviewsRetrieveShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (o *InterviewsRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *InterviewsRetrieveRequest) GetExpand() *InterviewsRetrieveExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *InterviewsRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InterviewsRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *InterviewsRetrieveRequest) GetRemoteFields() *InterviewsRetrieveRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *InterviewsRetrieveRequest) GetShowEnumOrigins() *InterviewsRetrieveShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

type InterviewsRetrieveResponse struct {
	ContentType        string
	ScheduledInterview *shared.ScheduledInterview
	StatusCode         int
	RawResponse        *http.Response
}

func (o *InterviewsRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InterviewsRetrieveResponse) GetScheduledInterview() *shared.ScheduledInterview {
	if o == nil {
		return nil
	}
	return o.ScheduledInterview
}

func (o *InterviewsRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InterviewsRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}