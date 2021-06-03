// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1IncidentsIncidentIDRelatedChangeEvents Update a change event
//
// swagger:model patchV1IncidentsIncidentIdRelatedChangeEvents
type PatchV1IncidentsIncidentIDRelatedChangeEvents struct {

	// type
	// Enum: [caused fixed suspect dismissed]
	Type string `json:"type,omitempty"`

	// A short description about why this change event is related
	Why string `json:"why,omitempty"`
}

// Validate validates this patch v1 incidents incident Id related change events
func (m *PatchV1IncidentsIncidentIDRelatedChangeEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchV1IncidentsIncidentIdRelatedChangeEventsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["caused","fixed","suspect","dismissed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchV1IncidentsIncidentIdRelatedChangeEventsTypeTypePropEnum = append(patchV1IncidentsIncidentIdRelatedChangeEventsTypeTypePropEnum, v)
	}
}

const (

	// PatchV1IncidentsIncidentIDRelatedChangeEventsTypeCaused captures enum value "caused"
	PatchV1IncidentsIncidentIDRelatedChangeEventsTypeCaused string = "caused"

	// PatchV1IncidentsIncidentIDRelatedChangeEventsTypeFixed captures enum value "fixed"
	PatchV1IncidentsIncidentIDRelatedChangeEventsTypeFixed string = "fixed"

	// PatchV1IncidentsIncidentIDRelatedChangeEventsTypeSuspect captures enum value "suspect"
	PatchV1IncidentsIncidentIDRelatedChangeEventsTypeSuspect string = "suspect"

	// PatchV1IncidentsIncidentIDRelatedChangeEventsTypeDismissed captures enum value "dismissed"
	PatchV1IncidentsIncidentIDRelatedChangeEventsTypeDismissed string = "dismissed"
)

// prop value enum
func (m *PatchV1IncidentsIncidentIDRelatedChangeEvents) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchV1IncidentsIncidentIdRelatedChangeEventsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchV1IncidentsIncidentIDRelatedChangeEvents) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 incidents incident Id related change events based on context it is used
func (m *PatchV1IncidentsIncidentIDRelatedChangeEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDRelatedChangeEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDRelatedChangeEvents) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDRelatedChangeEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
