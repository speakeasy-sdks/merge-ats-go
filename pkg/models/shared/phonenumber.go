// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"time"
)

// PhoneNumberType - The type of phone number.
//
// * `HOME` - HOME
// * `WORK` - WORK
// * `MOBILE` - MOBILE
// * `SKYPE` - SKYPE
// * `OTHER` - OTHER
type PhoneNumberType string

const (
	PhoneNumberTypeHome   PhoneNumberType = "HOME"
	PhoneNumberTypeWork   PhoneNumberType = "WORK"
	PhoneNumberTypeMobile PhoneNumberType = "MOBILE"
	PhoneNumberTypeSkype  PhoneNumberType = "SKYPE"
	PhoneNumberTypeOther  PhoneNumberType = "OTHER"
)

func (e PhoneNumberType) ToPointer() *PhoneNumberType {
	return &e
}

func (e *PhoneNumberType) UnmarshalJSON(data []byte) error {
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
		*e = PhoneNumberType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PhoneNumberType: %v", v)
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
	PhoneNumberType *PhoneNumberType `json:"phone_number_type,omitempty"`
	// The phone number.
	Value *string `json:"value,omitempty"`
}

func (p PhoneNumber) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PhoneNumber) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PhoneNumber) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *PhoneNumber) GetPhoneNumberType() *PhoneNumberType {
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
