// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// PhoneNumberPhoneNumberType - * `HOME` - HOME
// * `WORK` - WORK
// * `MOBILE` - MOBILE
// * `SKYPE` - SKYPE
// * `OTHER` - OTHER
type PhoneNumberPhoneNumberType string

const (
	PhoneNumberPhoneNumberTypeHome   PhoneNumberPhoneNumberType = "HOME"
	PhoneNumberPhoneNumberTypeWork   PhoneNumberPhoneNumberType = "WORK"
	PhoneNumberPhoneNumberTypeMobile PhoneNumberPhoneNumberType = "MOBILE"
	PhoneNumberPhoneNumberTypeSkype  PhoneNumberPhoneNumberType = "SKYPE"
	PhoneNumberPhoneNumberTypeOther  PhoneNumberPhoneNumberType = "OTHER"
)

func (e PhoneNumberPhoneNumberType) ToPointer() *PhoneNumberPhoneNumberType {
	return &e
}

func (e *PhoneNumberPhoneNumberType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HOME":
		fallthrough
	case "WORK":
		fallthrough
	case "MOBILE":
		fallthrough
	case "SKYPE":
		fallthrough
	case "OTHER":
		*e = PhoneNumberPhoneNumberType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PhoneNumberPhoneNumberType: %v", v)
	}
}

// PhoneNumber - # The PhoneNumber Object
// ### Description
// The `PhoneNumber` object is used to represent a candidate's phone number.
// ### Usage Example
// Fetch from the `GET Candidate` endpoint and view their phone numbers.
type PhoneNumber struct {
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The type of phone number.
	//
	// * `HOME` - HOME
	// * `WORK` - WORK
	// * `MOBILE` - MOBILE
	// * `SKYPE` - SKYPE
	// * `OTHER` - OTHER
	PhoneNumberType *PhoneNumberPhoneNumberType `json:"phone_number_type,omitempty"`
	// The phone number.
	Value *string `json:"value,omitempty"`
}

func (o *PhoneNumber) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *PhoneNumber) GetPhoneNumberType() *PhoneNumberPhoneNumberType {
	if o == nil {
		return nil
	}
	return o.PhoneNumberType
}

func (o *PhoneNumber) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}