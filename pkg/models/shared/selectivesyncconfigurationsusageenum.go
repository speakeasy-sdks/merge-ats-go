// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// SelectiveSyncConfigurationsUsageEnum - * `IN_NEXT_SYNC` - IN_NEXT_SYNC
// * `IN_LAST_SYNC` - IN_LAST_SYNC
type SelectiveSyncConfigurationsUsageEnum string

const (
	SelectiveSyncConfigurationsUsageEnumInNextSync SelectiveSyncConfigurationsUsageEnum = "IN_NEXT_SYNC"
	SelectiveSyncConfigurationsUsageEnumInLastSync SelectiveSyncConfigurationsUsageEnum = "IN_LAST_SYNC"
)

func (e SelectiveSyncConfigurationsUsageEnum) ToPointer() *SelectiveSyncConfigurationsUsageEnum {
	return &e
}

func (e *SelectiveSyncConfigurationsUsageEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IN_NEXT_SYNC":
		fallthrough
	case "IN_LAST_SYNC":
		*e = SelectiveSyncConfigurationsUsageEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SelectiveSyncConfigurationsUsageEnum: %v", v)
	}
}
