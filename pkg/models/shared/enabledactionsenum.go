// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// EnabledActionsEnum - * `READ` - READ
// * `WRITE` - WRITE
type EnabledActionsEnum string

const (
	EnabledActionsEnumRead  EnabledActionsEnum = "READ"
	EnabledActionsEnumWrite EnabledActionsEnum = "WRITE"
)

func (e EnabledActionsEnum) ToPointer() *EnabledActionsEnum {
	return &e
}

func (e *EnabledActionsEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "READ":
		fallthrough
	case "WRITE":
		*e = EnabledActionsEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EnabledActionsEnum: %v", v)
	}
}
