// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateApplicationStageRequest struct {
	// The interview stage to move the application to.
	JobInterviewStage *string `json:"job_interview_stage,omitempty" form:"name=job_interview_stage" multipartForm:"name=job_interview_stage"`
	RemoteUserID      *string `json:"remote_user_id,omitempty" form:"name=remote_user_id" multipartForm:"name=remote_user_id"`
}

func (o *UpdateApplicationStageRequest) GetJobInterviewStage() *string {
	if o == nil {
		return nil
	}
	return o.JobInterviewStage
}

func (o *UpdateApplicationStageRequest) GetRemoteUserID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteUserID
}
