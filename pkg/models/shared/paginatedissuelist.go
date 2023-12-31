// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaginatedIssueList struct {
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []Issue `json:"results,omitempty"`
}

func (o *PaginatedIssueList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PaginatedIssueList) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *PaginatedIssueList) GetResults() []Issue {
	if o == nil {
		return nil
	}
	return o.Results
}
