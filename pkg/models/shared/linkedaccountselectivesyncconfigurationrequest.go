// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LinkedAccountSelectiveSyncConfigurationRequest struct {
	// The conditions belonging to a selective sync.
	LinkedAccountConditions []LinkedAccountConditionRequest `json:"linked_account_conditions"`
}

func (o *LinkedAccountSelectiveSyncConfigurationRequest) GetLinkedAccountConditions() []LinkedAccountConditionRequest {
	if o == nil {
		return []LinkedAccountConditionRequest{}
	}
	return o.LinkedAccountConditions
}
