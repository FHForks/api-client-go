// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1IncidentsIncidentIDNotes Add a note to an incident
//
// swagger:model postV1IncidentsIncidentIdNotes
type PostV1IncidentsIncidentIDNotes struct {

	// body
	// Required: true
	Body *string `json:"body"`

	// status pages
	StatusPages []*PostV1IncidentsIncidentIDNotesStatusPagesItems0 `json:"status_pages"`

	// visibility
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this post v1 incidents incident Id notes
func (m *PostV1IncidentsIncidentIDNotes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusPages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDNotes) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsIncidentIDNotes) validateStatusPages(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusPages) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusPages); i++ {
		if swag.IsZero(m.StatusPages[i]) { // not required
			continue
		}

		if m.StatusPages[i] != nil {
			if err := m.StatusPages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 incidents incident Id notes based on the context it is used
func (m *PostV1IncidentsIncidentIDNotes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatusPages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDNotes) contextValidateStatusPages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusPages); i++ {

		if m.StatusPages[i] != nil {
			if err := m.StatusPages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDNotes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDNotes) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsIncidentIDNotes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1IncidentsIncidentIDNotesStatusPagesItems0 post v1 incidents incident ID notes status pages items0
//
// swagger:model PostV1IncidentsIncidentIDNotesStatusPagesItems0
type PostV1IncidentsIncidentIDNotesStatusPagesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// integration slug
	// Required: true
	IntegrationSlug *string `json:"integration_slug"`
}

// Validate validates this post v1 incidents incident ID notes status pages items0
func (m *PostV1IncidentsIncidentIDNotesStatusPagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDNotesStatusPagesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsIncidentIDNotesStatusPagesItems0) validateIntegrationSlug(formats strfmt.Registry) error {

	if err := validate.Required("integration_slug", "body", m.IntegrationSlug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 incidents incident ID notes status pages items0 based on context it is used
func (m *PostV1IncidentsIncidentIDNotesStatusPagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDNotesStatusPagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDNotesStatusPagesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsIncidentIDNotesStatusPagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
