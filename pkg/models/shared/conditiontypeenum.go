// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ConditionTypeEnum - * `BOOLEAN` - BOOLEAN
// * `DATE` - DATE
// * `DATE_TIME` - DATE_TIME
// * `INTEGER` - INTEGER
// * `FLOAT` - FLOAT
// * `STRING` - STRING
// * `LIST_OF_STRINGS` - LIST_OF_STRINGS
type ConditionTypeEnum string

const (
	ConditionTypeEnumBoolean       ConditionTypeEnum = "BOOLEAN"
	ConditionTypeEnumDate          ConditionTypeEnum = "DATE"
	ConditionTypeEnumDateTime      ConditionTypeEnum = "DATE_TIME"
	ConditionTypeEnumInteger       ConditionTypeEnum = "INTEGER"
	ConditionTypeEnumFloat         ConditionTypeEnum = "FLOAT"
	ConditionTypeEnumString        ConditionTypeEnum = "STRING"
	ConditionTypeEnumListOfStrings ConditionTypeEnum = "LIST_OF_STRINGS"
)

func (e ConditionTypeEnum) ToPointer() *ConditionTypeEnum {
	return &e
}

func (e *ConditionTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BOOLEAN":
		fallthrough
	case "DATE":
		fallthrough
	case "DATE_TIME":
		fallthrough
	case "INTEGER":
		fallthrough
	case "FLOAT":
		fallthrough
	case "STRING":
		fallthrough
	case "LIST_OF_STRINGS":
		*e = ConditionTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConditionTypeEnum: %v", v)
	}
}
