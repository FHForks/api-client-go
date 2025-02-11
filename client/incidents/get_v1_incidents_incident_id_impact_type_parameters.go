// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetV1IncidentsIncidentIDImpactTypeParams creates a new GetV1IncidentsIncidentIDImpactTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDImpactTypeParams() *GetV1IncidentsIncidentIDImpactTypeParams {
	return &GetV1IncidentsIncidentIDImpactTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDImpactTypeParamsWithTimeout creates a new GetV1IncidentsIncidentIDImpactTypeParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDImpactTypeParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDImpactTypeParams {
	return &GetV1IncidentsIncidentIDImpactTypeParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDImpactTypeParamsWithContext creates a new GetV1IncidentsIncidentIDImpactTypeParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDImpactTypeParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDImpactTypeParams {
	return &GetV1IncidentsIncidentIDImpactTypeParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDImpactTypeParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDImpactTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDImpactTypeParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDImpactTypeParams {
	return &GetV1IncidentsIncidentIDImpactTypeParams{
		HTTPClient: client,
	}
}

/* GetV1IncidentsIncidentIDImpactTypeParams contains all the parameters to send to the API endpoint
   for the get v1 incidents incident Id impact type operation.

   Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDImpactTypeParams struct {

	// IncidentID.
	//
	// Format: int32
	IncidentID int32

	// Type.
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id impact type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WithDefaults() *GetV1IncidentsIncidentIDImpactTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id impact type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDImpactTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDImpactTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDImpactTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDImpactTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WithIncidentID(incidentID int32) *GetV1IncidentsIncidentIDImpactTypeParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) SetIncidentID(incidentID int32) {
	o.IncidentID = incidentID
}

// WithType adds the typeVar to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WithType(typeVar string) *GetV1IncidentsIncidentIDImpactTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get v1 incidents incident Id impact type params
func (o *GetV1IncidentsIncidentIDImpactTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDImpactTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", swag.FormatInt32(o.IncidentID)); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
