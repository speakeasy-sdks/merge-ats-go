// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaginatedScheduledInterviewList struct {
	Next     *string              `json:"next,omitempty"`
	Previous *string              `json:"previous,omitempty"`
	Results  []ScheduledInterview `json:"results,omitempty"`
}

func (o *PaginatedScheduledInterviewList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PaginatedScheduledInterviewList) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *PaginatedScheduledInterviewList) GetResults() []ScheduledInterview {
	if o == nil {
		return nil
	}
	return o.Results
}
