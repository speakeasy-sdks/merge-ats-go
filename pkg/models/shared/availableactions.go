// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AvailableActions - # The AvailableActions Object
// ### Description
// The `Activity` object is used to see all available model/operation combinations for an integration.
//
// ### Usage Example
// Fetch all the actions available for the `Zenefits` integration.
type AvailableActions struct {
	AvailableModelOperations []ModelOperation   `json:"available_model_operations,omitempty"`
	Integration              AccountIntegration `json:"integration"`
	PassthroughAvailable     bool               `json:"passthrough_available"`
}

func (o *AvailableActions) GetAvailableModelOperations() []ModelOperation {
	if o == nil {
		return nil
	}
	return o.AvailableModelOperations
}

func (o *AvailableActions) GetIntegration() AccountIntegration {
	if o == nil {
		return AccountIntegration{}
	}
	return o.Integration
}

func (o *AvailableActions) GetPassthroughAvailable() bool {
	if o == nil {
		return false
	}
	return o.PassthroughAvailable
}
