// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Reason - * `GENERAL_CUSTOMER_REQUEST` - GENERAL_CUSTOMER_REQUEST
// * `GDPR` - GDPR
// * `OTHER` - OTHER
type Reason string

const (
	ReasonGeneralCustomerRequest Reason = "GENERAL_CUSTOMER_REQUEST"
	ReasonGdpr                   Reason = "GDPR"
	ReasonOther                  Reason = "OTHER"
)

func (e Reason) ToPointer() *Reason {
	return &e
}

func (e *Reason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GENERAL_CUSTOMER_REQUEST":
		fallthrough
	case "GDPR":
		fallthrough
	case "OTHER":
		*e = Reason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Reason: %v", v)
	}
}

type IgnoreCommonModelRequest struct {
	Message *string `json:"message,omitempty"`
	Reason  Reason  `json:"reason"`
}

func (o *IgnoreCommonModelRequest) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *IgnoreCommonModelRequest) GetReason() Reason {
	if o == nil {
		return Reason("")
	}
	return o.Reason
}
