// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DebugModelLogSummary struct {
	Method     string `json:"method"`
	StatusCode int64  `json:"status_code"`
	URL        string `json:"url"`
}

func (o *DebugModelLogSummary) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *DebugModelLogSummary) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DebugModelLogSummary) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}