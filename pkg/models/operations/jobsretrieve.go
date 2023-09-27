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

// JobsRetrieveExpand - Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
type JobsRetrieveExpand string

const (
	JobsRetrieveExpandDepartments                                JobsRetrieveExpand = "departments"
	JobsRetrieveExpandDepartmentsHiringManagers                  JobsRetrieveExpand = "departments,hiring_managers"
	JobsRetrieveExpandDepartmentsHiringManagersRecruiters        JobsRetrieveExpand = "departments,hiring_managers,recruiters"
	JobsRetrieveExpandDepartmentsOffices                         JobsRetrieveExpand = "departments,offices"
	JobsRetrieveExpandDepartmentsOfficesHiringManagers           JobsRetrieveExpand = "departments,offices,hiring_managers"
	JobsRetrieveExpandDepartmentsOfficesHiringManagersRecruiters JobsRetrieveExpand = "departments,offices,hiring_managers,recruiters"
	JobsRetrieveExpandDepartmentsOfficesRecruiters               JobsRetrieveExpand = "departments,offices,recruiters"
	JobsRetrieveExpandDepartmentsRecruiters                      JobsRetrieveExpand = "departments,recruiters"
	JobsRetrieveExpandHiringManagers                             JobsRetrieveExpand = "hiring_managers"
	JobsRetrieveExpandHiringManagersRecruiters                   JobsRetrieveExpand = "hiring_managers,recruiters"
	JobsRetrieveExpandOffices                                    JobsRetrieveExpand = "offices"
	JobsRetrieveExpandOfficesHiringManagers                      JobsRetrieveExpand = "offices,hiring_managers"
	JobsRetrieveExpandOfficesHiringManagersRecruiters            JobsRetrieveExpand = "offices,hiring_managers,recruiters"
	JobsRetrieveExpandOfficesRecruiters                          JobsRetrieveExpand = "offices,recruiters"
	JobsRetrieveExpandRecruiters                                 JobsRetrieveExpand = "recruiters"
)

func (e JobsRetrieveExpand) ToPointer() *JobsRetrieveExpand {
	return &e
}

func (e *JobsRetrieveExpand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "departments":
		fallthrough
	case "departments,hiring_managers":
		fallthrough
	case "departments,hiring_managers,recruiters":
		fallthrough
	case "departments,offices":
		fallthrough
	case "departments,offices,hiring_managers":
		fallthrough
	case "departments,offices,hiring_managers,recruiters":
		fallthrough
	case "departments,offices,recruiters":
		fallthrough
	case "departments,recruiters":
		fallthrough
	case "hiring_managers":
		fallthrough
	case "hiring_managers,recruiters":
		fallthrough
	case "offices":
		fallthrough
	case "offices,hiring_managers":
		fallthrough
	case "offices,hiring_managers,recruiters":
		fallthrough
	case "offices,recruiters":
		fallthrough
	case "recruiters":
		*e = JobsRetrieveExpand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobsRetrieveExpand: %v", v)
	}
}

type JobsRetrieveRequest struct {
	// Token identifying the end user.
	XAccountToken string `header:"style=simple,explode=false,name=X-Account-Token"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *JobsRetrieveExpand `queryParam:"style=form,explode=true,name=expand"`
	ID     string              `pathParam:"style=simple,explode=false,name=id"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `queryParam:"style=form,explode=true,name=include_remote_data"`
	// Deprecated. Use show_enum_origins.
	remoteFields *string `const:"status" queryParam:"style=form,explode=true,name=remote_fields"`
	// Which fields should be returned in non-normalized form.
	showEnumOrigins *string `const:"status" queryParam:"style=form,explode=true,name=show_enum_origins"`
}

func (j JobsRetrieveRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JobsRetrieveRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JobsRetrieveRequest) GetXAccountToken() string {
	if o == nil {
		return ""
	}
	return o.XAccountToken
}

func (o *JobsRetrieveRequest) GetExpand() *JobsRetrieveExpand {
	if o == nil {
		return nil
	}
	return o.Expand
}

func (o *JobsRetrieveRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *JobsRetrieveRequest) GetIncludeRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRemoteData
}

func (o *JobsRetrieveRequest) GetRemoteFields() *string {
	return types.String("status")
}

func (o *JobsRetrieveRequest) GetShowEnumOrigins() *string {
	return types.String("status")
}

type JobsRetrieveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Job         *shared.Job
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *JobsRetrieveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *JobsRetrieveResponse) GetJob() *shared.Job {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *JobsRetrieveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *JobsRetrieveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
