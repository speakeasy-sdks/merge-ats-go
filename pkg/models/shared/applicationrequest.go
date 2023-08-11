// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ApplicationRequest - # The Application Object
// ### Description
// The Application Object is used to represent a candidate's journey through a particular Job's recruiting process. If a Candidate applies for multiple Jobs, there will be a separate Application for each Job if the third-party integration allows it.
//
// ### Usage Example
// Fetch from the `LIST Applications` endpoint and filter by `ID` to show all applications.
type ApplicationRequest struct {
	// When the application was submitted.
	AppliedAt *time.Time `json:"applied_at,omitempty"`
	// The candidate applying.
	Candidate *string `json:"candidate,omitempty"`
	// The user credited for this application.
	CreditedTo *string `json:"credited_to,omitempty"`
	// The application's current stage.
	CurrentStage      *string                `json:"current_stage,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	// The job being applied for.
	Job                 *string                `json:"job,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// The application's reason for rejection.
	RejectReason *string `json:"reject_reason,omitempty"`
	// When the application was rejected.
	RejectedAt       *time.Time `json:"rejected_at,omitempty"`
	RemoteTemplateID *string    `json:"remote_template_id,omitempty"`
	// The application's source.
	Source *string `json:"source,omitempty"`
}

func (o *ApplicationRequest) GetAppliedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AppliedAt
}

func (o *ApplicationRequest) GetCandidate() *string {
	if o == nil {
		return nil
	}
	return o.Candidate
}

func (o *ApplicationRequest) GetCreditedTo() *string {
	if o == nil {
		return nil
	}
	return o.CreditedTo
}

func (o *ApplicationRequest) GetCurrentStage() *string {
	if o == nil {
		return nil
	}
	return o.CurrentStage
}

func (o *ApplicationRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.IntegrationParams
}

func (o *ApplicationRequest) GetJob() *string {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *ApplicationRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.LinkedAccountParams
}

func (o *ApplicationRequest) GetRejectReason() *string {
	if o == nil {
		return nil
	}
	return o.RejectReason
}

func (o *ApplicationRequest) GetRejectedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RejectedAt
}

func (o *ApplicationRequest) GetRemoteTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteTemplateID
}

func (o *ApplicationRequest) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}