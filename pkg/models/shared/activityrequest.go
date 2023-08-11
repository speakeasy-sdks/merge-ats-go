// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ActivityRequestActivityType - * `NOTE` - NOTE
// * `EMAIL` - EMAIL
// * `OTHER` - OTHER
type ActivityRequestActivityType string

const (
	ActivityRequestActivityTypeNote  ActivityRequestActivityType = "NOTE"
	ActivityRequestActivityTypeEmail ActivityRequestActivityType = "EMAIL"
	ActivityRequestActivityTypeOther ActivityRequestActivityType = "OTHER"
)

func (e ActivityRequestActivityType) ToPointer() *ActivityRequestActivityType {
	return &e
}

func (e *ActivityRequestActivityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NOTE":
		fallthrough
	case "EMAIL":
		fallthrough
	case "OTHER":
		*e = ActivityRequestActivityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivityRequestActivityType: %v", v)
	}
}

// ActivityRequestVisibility - * `ADMIN_ONLY` - ADMIN_ONLY
// * `PUBLIC` - PUBLIC
// * `PRIVATE` - PRIVATE
type ActivityRequestVisibility string

const (
	ActivityRequestVisibilityAdminOnly ActivityRequestVisibility = "ADMIN_ONLY"
	ActivityRequestVisibilityPublic    ActivityRequestVisibility = "PUBLIC"
	ActivityRequestVisibilityPrivate   ActivityRequestVisibility = "PRIVATE"
)

func (e ActivityRequestVisibility) ToPointer() *ActivityRequestVisibility {
	return &e
}

func (e *ActivityRequestVisibility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADMIN_ONLY":
		fallthrough
	case "PUBLIC":
		fallthrough
	case "PRIVATE":
		*e = ActivityRequestVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivityRequestVisibility: %v", v)
	}
}

// ActivityRequest - # The Activity Object
// ### Description
// The `Activity` object is used to represent an activity for a candidate performed by a user.
// ### Usage Example
// Fetch from the `LIST Activities` endpoint and filter by `ID` to show all activities.
type ActivityRequest struct {
	// The activity's type.
	//
	// * `NOTE` - NOTE
	// * `EMAIL` - EMAIL
	// * `OTHER` - OTHER
	ActivityType *ActivityRequestActivityType `json:"activity_type,omitempty"`
	// The activity's body.
	Body *string `json:"body,omitempty"`
	// The activity’s candidate.
	Candidate           *string                `json:"candidate,omitempty"`
	IntegrationParams   map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// The activity's subject.
	Subject *string `json:"subject,omitempty"`
	// The user that performed the action.
	User *string `json:"user,omitempty"`
	// The activity's visibility.
	//
	// * `ADMIN_ONLY` - ADMIN_ONLY
	// * `PUBLIC` - PUBLIC
	// * `PRIVATE` - PRIVATE
	Visibility *ActivityRequestVisibility `json:"visibility,omitempty"`
}

func (o *ActivityRequest) GetActivityType() *ActivityRequestActivityType {
	if o == nil {
		return nil
	}
	return o.ActivityType
}

func (o *ActivityRequest) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *ActivityRequest) GetCandidate() *string {
	if o == nil {
		return nil
	}
	return o.Candidate
}

func (o *ActivityRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.IntegrationParams
}

func (o *ActivityRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.LinkedAccountParams
}

func (o *ActivityRequest) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *ActivityRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *ActivityRequest) GetVisibility() *ActivityRequestVisibility {
	if o == nil {
		return nil
	}
	return o.Visibility
}
