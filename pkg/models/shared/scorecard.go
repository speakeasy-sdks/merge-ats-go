// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"time"
)

// OverallRecommendation - * `DEFINITELY_NO` - DEFINITELY_NO
// * `NO` - NO
// * `YES` - YES
// * `STRONG_YES` - STRONG_YES
// * `NO_DECISION` - NO_DECISION
type OverallRecommendation string

const (
	OverallRecommendationDefinitelyNo OverallRecommendation = "DEFINITELY_NO"
	OverallRecommendationNo           OverallRecommendation = "NO"
	OverallRecommendationYes          OverallRecommendation = "YES"
	OverallRecommendationStrongYes    OverallRecommendation = "STRONG_YES"
	OverallRecommendationNoDecision   OverallRecommendation = "NO_DECISION"
)

func (e OverallRecommendation) ToPointer() *OverallRecommendation {
	return &e
}

func (e *OverallRecommendation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DEFINITELY_NO":
		fallthrough
	case "NO":
		fallthrough
	case "YES":
		fallthrough
	case "STRONG_YES":
		fallthrough
	case "NO_DECISION":
		*e = OverallRecommendation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OverallRecommendation: %v", v)
	}
}

// Scorecard - # The Scorecard Object
// ### Description
// The `Scorecard` object is used to represent an interviewer's candidate recommendation based on a particular interview.
// ### Usage Example
// Fetch from the `LIST Scorecards` endpoint and filter by `application` to show all scorecard for an applicant.
type Scorecard struct {
	// The application being scored.
	Application   *string                `json:"application,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	ID            *string                `json:"id,omitempty"`
	// The interview being scored.
	Interview *string `json:"interview,omitempty"`
	// The interviewer doing the scoring.
	Interviewer *string `json:"interviewer,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The inteviewer's recommendation.
	//
	// * `DEFINITELY_NO` - DEFINITELY_NO
	// * `NO` - NO
	// * `YES` - YES
	// * `STRONG_YES` - STRONG_YES
	// * `NO_DECISION` - NO_DECISION
	OverallRecommendation *OverallRecommendation `json:"overall_recommendation,omitempty"`
	// When the third party's scorecard was created.
	RemoteCreatedAt *time.Time   `json:"remote_created_at,omitempty"`
	RemoteData      []RemoteData `json:"remote_data,omitempty"`
	// The third-party API ID of the matching object.
	RemoteID *string `json:"remote_id,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	// When the scorecard was submitted.
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
}

func (s Scorecard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Scorecard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Scorecard) GetApplication() *string {
	if o == nil {
		return nil
	}
	return o.Application
}

func (o *Scorecard) GetFieldMappings() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *Scorecard) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Scorecard) GetInterview() *string {
	if o == nil {
		return nil
	}
	return o.Interview
}

func (o *Scorecard) GetInterviewer() *string {
	if o == nil {
		return nil
	}
	return o.Interviewer
}

func (o *Scorecard) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *Scorecard) GetOverallRecommendation() *OverallRecommendation {
	if o == nil {
		return nil
	}
	return o.OverallRecommendation
}

func (o *Scorecard) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *Scorecard) GetRemoteData() []RemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *Scorecard) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *Scorecard) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}

func (o *Scorecard) GetSubmittedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.SubmittedAt
}
