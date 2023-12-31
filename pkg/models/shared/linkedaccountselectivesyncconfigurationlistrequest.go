// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LinkedAccountSelectiveSyncConfigurationListRequest struct {
	// The selective syncs associated with a linked account.
	SyncConfigurations []LinkedAccountSelectiveSyncConfigurationRequest `json:"sync_configurations"`
}

func (o *LinkedAccountSelectiveSyncConfigurationListRequest) GetSyncConfigurations() []LinkedAccountSelectiveSyncConfigurationRequest {
	if o == nil {
		return []LinkedAccountSelectiveSyncConfigurationRequest{}
	}
	return o.SyncConfigurations
}
