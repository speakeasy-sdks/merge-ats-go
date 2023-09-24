// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"time"
)

// ScheduledInterviewRequestStatus - * `SCHEDULED` - SCHEDULED
// * `AWAITING_FEEDBACK` - AWAITING_FEEDBACK
// * `COMPLETE` - COMPLETE
type ScheduledInterviewRequestStatus string

const (
	ScheduledInterviewRequestStatusScheduled        ScheduledInterviewRequestStatus = "SCHEDULED"
	ScheduledInterviewRequestStatusAwaitingFeedback ScheduledInterviewRequestStatus = "AWAITING_FEEDBACK"
	ScheduledInterviewRequestStatusComplete         ScheduledInterviewRequestStatus = "COMPLETE"
)

func (e ScheduledInterviewRequestStatus) ToPointer() *ScheduledInterviewRequestStatus {
	return &e
}

func (e *ScheduledInterviewRequestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SCHEDULED":
		fallthrough
	case "AWAITING_FEEDBACK":
		fallthrough
	case "COMPLETE":
		*e = ScheduledInterviewRequestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduledInterviewRequestStatus: %v", v)
	}
}

// ScheduledInterviewRequest - # The ScheduledInterview Object
// ### Description
// The `ScheduledInterview` object is used to represent a scheduled interview for a given candidate’s application to a job. An `Application` can have multiple `ScheduledInterview`s depending on the particular hiring process.
// ### Usage Example
// Fetch from the `LIST ScheduledInterviews` endpoint and filter by `interviewers` to show all office locations.
type ScheduledInterviewRequest struct {
	// The application being interviewed.
	Application *string `json:"application,omitempty"`
	// When the interview was ended.
	EndAt             *time.Time             `json:"end_at,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	// Array of `RemoteUser` IDs.
	Interviewers []string `json:"interviewers,omitempty"`
	// The stage of the interview.
	JobInterviewStage   *string                `json:"job_interview_stage,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// The interview's location.
	Location *string `json:"location,omitempty"`
	// The user organizing the interview.
	Organizer *string `json:"organizer,omitempty"`
	// When the interview was started.
	StartAt *time.Time `json:"start_at,omitempty"`
	// The interview's status.
	//
	// * `SCHEDULED` - SCHEDULED
	// * `AWAITING_FEEDBACK` - AWAITING_FEEDBACK
	// * `COMPLETE` - COMPLETE
	Status *ScheduledInterviewRequestStatus `json:"status,omitempty"`
}

func (s ScheduledInterviewRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScheduledInterviewRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ScheduledInterviewRequest) GetApplication() *string {
	if o == nil {
		return nil
	}
	return o.Application
}

func (o *ScheduledInterviewRequest) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *ScheduledInterviewRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.IntegrationParams
}

func (o *ScheduledInterviewRequest) GetInterviewers() []string {
	if o == nil {
		return nil
	}
	return o.Interviewers
}

func (o *ScheduledInterviewRequest) GetJobInterviewStage() *string {
	if o == nil {
		return nil
	}
	return o.JobInterviewStage
}

func (o *ScheduledInterviewRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.LinkedAccountParams
}

func (o *ScheduledInterviewRequest) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *ScheduledInterviewRequest) GetOrganizer() *string {
	if o == nil {
		return nil
	}
	return o.Organizer
}

func (o *ScheduledInterviewRequest) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *ScheduledInterviewRequest) GetStatus() *ScheduledInterviewRequestStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
