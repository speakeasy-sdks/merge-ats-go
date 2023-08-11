// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ApplicationResponse struct {
	Errors []ErrorValidationProblem `json:"errors"`
	Logs   []DebugModeLog           `json:"logs,omitempty"`
	// # The Application Object
	// ### Description
	// The Application Object is used to represent a candidate's journey through a particular Job's recruiting process. If a Candidate applies for multiple Jobs, there will be a separate Application for each Job if the third-party integration allows it.
	//
	// ### Usage Example
	// Fetch from the `LIST Applications` endpoint and filter by `ID` to show all applications.
	Model    Application                `json:"model"`
	Warnings []WarningValidationProblem `json:"warnings"`
}

func (o *ApplicationResponse) GetErrors() []ErrorValidationProblem {
	if o == nil {
		return []ErrorValidationProblem{}
	}
	return o.Errors
}

func (o *ApplicationResponse) GetLogs() []DebugModeLog {
	if o == nil {
		return nil
	}
	return o.Logs
}

func (o *ApplicationResponse) GetModel() Application {
	if o == nil {
		return Application{}
	}
	return o.Model
}

func (o *ApplicationResponse) GetWarnings() []WarningValidationProblem {
	if o == nil {
		return []WarningValidationProblem{}
	}
	return o.Warnings
}