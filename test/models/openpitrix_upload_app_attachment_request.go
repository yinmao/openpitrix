// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixUploadAppAttachmentRequest openpitrix upload app attachment request
// swagger:model openpitrixUploadAppAttachmentRequest
type OpenpitrixUploadAppAttachmentRequest struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// attachment content
	AttachmentContent strfmt.Base64 `json:"attachment_content,omitempty"`

	// only for screenshot, range: [0, 5]
	Sequence int64 `json:"sequence,omitempty"`

	// optional: icon/screenshot
	Type OpenpitrixUploadAppAttachmentRequestType `json:"type,omitempty"`
}

// Validate validates this openpitrix upload app attachment request
func (m *OpenpitrixUploadAppAttachmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixUploadAppAttachmentRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixUploadAppAttachmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixUploadAppAttachmentRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixUploadAppAttachmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
