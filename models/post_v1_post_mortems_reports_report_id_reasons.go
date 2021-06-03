// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1PostMortemsReportsReportIDReasons Add a post mortem reason to the report
//
// swagger:model postV1PostMortemsReportsReportIdReasons
type PostV1PostMortemsReportsReportIDReasons struct {

	// summary
	// Required: true
	Summary *string `json:"summary"`
}

// Validate validates this post v1 post mortems reports report Id reasons
func (m *PostV1PostMortemsReportsReportIDReasons) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1PostMortemsReportsReportIDReasons) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 post mortems reports report Id reasons based on context it is used
func (m *PostV1PostMortemsReportsReportIDReasons) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1PostMortemsReportsReportIDReasons) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1PostMortemsReportsReportIDReasons) UnmarshalBinary(b []byte) error {
	var res PostV1PostMortemsReportsReportIDReasons
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
