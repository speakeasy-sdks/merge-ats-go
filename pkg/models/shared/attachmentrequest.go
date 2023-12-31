// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AttachmentRequestAttachmentType - The attachment's type.
//
// * `RESUME` - RESUME
// * `COVER_LETTER` - COVER_LETTER
// * `OFFER_LETTER` - OFFER_LETTER
// * `OTHER` - OTHER
type AttachmentRequestAttachmentType string

const (
	AttachmentRequestAttachmentTypeResume      AttachmentRequestAttachmentType = "RESUME"
	AttachmentRequestAttachmentTypeCoverLetter AttachmentRequestAttachmentType = "COVER_LETTER"
	AttachmentRequestAttachmentTypeOfferLetter AttachmentRequestAttachmentType = "OFFER_LETTER"
	AttachmentRequestAttachmentTypeOther       AttachmentRequestAttachmentType = "OTHER"
)

func (e AttachmentRequestAttachmentType) ToPointer() *AttachmentRequestAttachmentType {
	return &e
}

func (e *AttachmentRequestAttachmentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RESUME":
		fallthrough
	case "COVER_LETTER":
		fallthrough
	case "OFFER_LETTER":
		fallthrough
	case "OTHER":
		*e = AttachmentRequestAttachmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttachmentRequestAttachmentType: %v", v)
	}
}

// AttachmentRequest - # The Attachment Object
// ### Description
// The `Attachment` object is used to represent a file attached to a candidate.
// ### Usage Example
// Fetch from the `LIST Attachments` endpoint and view attachments accessible by a company.
type AttachmentRequest struct {
	// The attachment's type.
	//
	// * `RESUME` - RESUME
	// * `COVER_LETTER` - COVER_LETTER
	// * `OFFER_LETTER` - OFFER_LETTER
	// * `OTHER` - OTHER
	AttachmentType *AttachmentRequestAttachmentType `json:"attachment_type,omitempty"`
	Candidate      *string                          `json:"candidate,omitempty"`
	// The attachment's name.
	FileName *string `json:"file_name,omitempty"`
	// The attachment's url.
	FileURL             *string                `json:"file_url,omitempty"`
	IntegrationParams   map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
}

func (o *AttachmentRequest) GetAttachmentType() *AttachmentRequestAttachmentType {
	if o == nil {
		return nil
	}
	return o.AttachmentType
}

func (o *AttachmentRequest) GetCandidate() *string {
	if o == nil {
		return nil
	}
	return o.Candidate
}

func (o *AttachmentRequest) GetFileName() *string {
	if o == nil {
		return nil
	}
	return o.FileName
}

func (o *AttachmentRequest) GetFileURL() *string {
	if o == nil {
		return nil
	}
	return o.FileURL
}

func (o *AttachmentRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.IntegrationParams
}

func (o *AttachmentRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.LinkedAccountParams
}
