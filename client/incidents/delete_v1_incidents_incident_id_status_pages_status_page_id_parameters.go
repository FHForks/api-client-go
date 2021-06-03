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
)

// NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams creates a new DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams() *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	return &DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParamsWithTimeout creates a new DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParamsWithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	return &DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParamsWithContext creates a new DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams object
// with the ability to set a context for a request.
func NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParamsWithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	return &DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams{
		Context: ctx,
	}
}

// NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParamsWithHTTPClient creates a new DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParamsWithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	return &DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams{
		HTTPClient: client,
	}
}

/* DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams contains all the parameters to send to the API endpoint
   for the delete v1 incidents incident Id status pages status page Id operation.

   Typically these are written to a http.Request.
*/
type DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams struct {

	// IncidentID.
	IncidentID string

	// StatusPageID.
	StatusPageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 incidents incident Id status pages status page Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WithDefaults() *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 incidents incident Id status pages status page Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WithIncidentID(incidentID string) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithStatusPageID adds the statusPageID to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WithStatusPageID(statusPageID string) *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams {
	o.SetStatusPageID(statusPageID)
	return o
}

// SetStatusPageID adds the statusPageId to the delete v1 incidents incident Id status pages status page Id params
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) SetStatusPageID(statusPageID string) {
	o.StatusPageID = statusPageID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	// path param status_page_id
	if err := r.SetPathParam("status_page_id", o.StatusPageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
