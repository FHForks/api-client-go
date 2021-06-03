// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExecutionStepEntity execution step entity
//
// swagger:model ExecutionStepEntity
type ExecutionStepEntity struct {

	// config
	Config string `json:"config,omitempty"`

	// executable
	Executable string `json:"executable,omitempty"`

	// execution
	Execution *ExecutionStepExecutionEntity `json:"execution,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// repeats
	Repeats string `json:"repeats,omitempty"`

	// repeats at
	RepeatsAt string `json:"repeats_at,omitempty"`

	// repeats duration
	RepeatsDuration string `json:"repeats_duration,omitempty"`

	// step elements
	StepElements string `json:"step_elements,omitempty"`
}

// Validate validates this execution step entity
func (m *ExecutionStepEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionStepEntity) validateExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.Execution) { // not required
		return nil
	}

	if m.Execution != nil {
		if err := m.Execution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execution")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this execution step entity based on the context it is used
func (m *ExecutionStepEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExecutionStepEntity) contextValidateExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.Execution != nil {
		if err := m.Execution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execution")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExecutionStepEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExecutionStepEntity) UnmarshalBinary(b []byte) error {
	var res ExecutionStepEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
