// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WebhookReceiver struct {
	Event    string  `json:"event"`
	IsActive bool    `json:"is_active"`
	Key      *string `json:"key,omitempty"`
}

func (o *WebhookReceiver) GetEvent() string {
	if o == nil {
		return ""
	}
	return o.Event
}

func (o *WebhookReceiver) GetIsActive() bool {
	if o == nil {
		return false
	}
	return o.IsActive
}

func (o *WebhookReceiver) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}