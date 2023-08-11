// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WarningValidationProblem struct {
	Detail      string                   `json:"detail"`
	ProblemType string                   `json:"problem_type"`
	Source      *ValidationProblemSource `json:"source,omitempty"`
	Title       string                   `json:"title"`
}

func (o *WarningValidationProblem) GetDetail() string {
	if o == nil {
		return ""
	}
	return o.Detail
}

func (o *WarningValidationProblem) GetProblemType() string {
	if o == nil {
		return ""
	}
	return o.ProblemType
}

func (o *WarningValidationProblem) GetSource() *ValidationProblemSource {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *WarningValidationProblem) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}
