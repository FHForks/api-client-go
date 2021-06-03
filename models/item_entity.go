// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ItemEntity item entity
//
// swagger:model ItemEntity
type ItemEntity struct {

	// condition id
	ConditionID string `json:"condition_id,omitempty"`

	// impact id
	ImpactID string `json:"impact_id,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`
}

// Validate validates this item entity
func (m *ItemEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this item entity based on context it is used
func (m *ItemEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemEntity) UnmarshalBinary(b []byte) error {
	var res ItemEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
