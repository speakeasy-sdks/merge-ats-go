// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LinkedAccountSelectiveSyncConfiguration struct {
	// The conditions belonging to a selective sync.
	LinkedAccountConditions []LinkedAccountCondition `json:"linked_account_conditions,omitempty"`
}

func (o *LinkedAccountSelectiveSyncConfiguration) GetLinkedAccountConditions() []LinkedAccountCondition {
	if o == nil {
		return nil
	}
	return o.LinkedAccountConditions
}
