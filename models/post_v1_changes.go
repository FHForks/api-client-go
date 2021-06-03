// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostV1Changes Create a new change entry
//
// swagger:model postV1Changes
type PostV1Changes struct {

	// description
	Description string `json:"description,omitempty"`

	// A labels hash of keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this post v1 changes
func (m *PostV1Changes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post v1 changes based on context it is used
func (m *PostV1Changes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1Changes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1Changes) UnmarshalBinary(b []byte) error {
	var res PostV1Changes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
