// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"net/http"
)

// ActivitiesRetrieveRemoteFields - Deprecated. Use show_enum_origins.
type ActivitiesRetrieveRemoteFields string

const (
	ActivitiesRetrieveRemoteFieldsActivityType           ActivitiesRetrieveRemoteFields = "activity_type"
	ActivitiesRetrieveRemoteFieldsActivityTypeVisibility ActivitiesRetrieveRemoteFields = "activity_type,visibility"
	ActivitiesRetrieveRemoteFieldsVisibility             ActivitiesRetrieveRemoteFields = "visibility"
)

func (e ActivitiesRetrieveRemoteFields) ToPointer() *ActivitiesRetrieveRemoteFields {
	return &e
}

func (e *ActivitiesRetrieveRemoteFields) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "activity_type":
		fallthrough
	case "activity_type,visibility":
		fallthrough
	case "visibility":
		*e = ActivitiesRetrieveRemoteFields(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivitiesRetrieveRemoteFields: %v", v)
	}
}

// ActivitiesRetrieveShowEnumOrigins - Which fields should be returned in non-normalized form.
type ActivitiesRetrieveShowEnumOrigins string

const (
	ActivitiesRetrieveShowEnumOriginsActivityType           ActivitiesRetrieveShowEnumOrigins = "activity_type"
	ActivitiesRetrieveShowEnumOriginsActivityTypeVisibility ActivitiesRetrieveShowEnumOrigins = "activity_type,visibility"
	ActivitiesRetrieveShowEnumOriginsVisibility             ActivitiesRetrieveShowEnumOrigins = "visibility"
)

func (e ActivitiesRetrieveShowEnumOrigins) ToPointer() *ActivitiesRetrieveShowEnumOrigins {
	return &e
}

func (e *ActivitiesRetrieveShowEnumOrigins) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "activity_type":
		fallthrough
	case "activity_type,visibility":
		fallthrough
	case "visibility":
		*e = ActivitiesRetrieveShowEnumOrigins(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActivitiesRetrieveShowEnumOrigins: %v", v)
	}
}

type ActivitiesRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	expand *string `const:"user" queryParam:"style=form,explode=true,name=expand"`
	ID     string  `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *ActivitiesRetrieveRemoteFields `queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *ActivitiesRetrieveShowEnumOrigins `queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (a ActivitiesRetrieveRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ActivitiesRetrieveRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ActivitiesRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *ActivitiesRetrieveRequest) GetExpand() *string {
	return types.String("user")
}

func (o *ActivitiesRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ActivitiesRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *ActivitiesRetrieveRequest) GetRemoteFields() *ActivitiesRetrieveRemoteFields {
	if o == nil {
		return nil
	}
	return o.RemoteFields
}

func (o *ActivitiesRetrieveRequest) GetShowEnumOrigins() *ActivitiesRetrieveShowEnumOrigins {
	if o == nil {
		return nil
	}
	return o.ShowEnumOrigins
}

type ActivitiesRetrieveResponse struct {
	Activity    *shared.Activity
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *ActivitiesRetrieveResponse) GetActivity() *shared.Activity {
	if o == nil {
		return nil
	}
	return o.Activity
}

func (o *ActivitiesRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ActivitiesRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ActivitiesRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
