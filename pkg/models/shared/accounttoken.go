// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountToken struct {
	AccountToken string             `json:"account_token"`
	Integration  AccountIntegration `json:"integration"`
}

func (o *AccountToken) GetAccountToken() string {
	if o == nil {
		return ""
	}
	return o.AccountToken
}

func (o *AccountToken) GetIntegration() AccountIntegration {
	if o == nil {
		return AccountIntegration{}
	}
	return o.Integration
}
