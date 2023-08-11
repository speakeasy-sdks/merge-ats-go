// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaginatedEEOCList struct {
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []Eeoc  `json:"results,omitempty"`
}

func (o *PaginatedEEOCList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *PaginatedEEOCList) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *PaginatedEEOCList) GetResults() []Eeoc {
	if o == nil {
		return nil
	}
	return o.Results
}