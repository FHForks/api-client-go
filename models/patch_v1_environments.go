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

// PatchV1Environments Update an environment
//
// swagger:model patchV1Environments
type PatchV1Environments struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this patch v1 environments
func (m *PatchV1Environments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1Environments) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 environments based on context it is used
func (m *PatchV1Environments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1Environments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1Environments) UnmarshalBinary(b []byte) error {
	var res PatchV1Environments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
