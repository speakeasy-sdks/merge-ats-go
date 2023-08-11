// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MetaResponse struct {
	HasConditionalParams           bool                   `json:"has_conditional_params"`
	HasRequiredLinkedAccountParams bool                   `json:"has_required_linked_account_params"`
	RemoteFieldClasses             map[string]interface{} `json:"remote_field_classes,omitempty"`
	RequestSchema                  map[string]interface{} `json:"request_schema"`
	Status                         *LinkedAccountStatus   `json:"status,omitempty"`
}

func (o *MetaResponse) GetHasConditionalParams() bool {
	if o == nil {
		return false
	}
	return o.HasConditionalParams
}

func (o *MetaResponse) GetHasRequiredLinkedAccountParams() bool {
	if o == nil {
		return false
	}
	return o.HasRequiredLinkedAccountParams
}

func (o *MetaResponse) GetRemoteFieldClasses() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RemoteFieldClasses
}

func (o *MetaResponse) GetRequestSchema() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.RequestSchema
}

func (o *MetaResponse) GetStatus() *LinkedAccountStatus {
	if o == nil {
		return nil
	}
	return o.Status
}