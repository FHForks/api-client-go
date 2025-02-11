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

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1IncidentsIncidentIDTeamAssignmentsParams creates a new PostV1IncidentsIncidentIDTeamAssignmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDTeamAssignmentsParams() *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	return &PostV1IncidentsIncidentIDTeamAssignmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDTeamAssignmentsParamsWithTimeout creates a new PostV1IncidentsIncidentIDTeamAssignmentsParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDTeamAssignmentsParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	return &PostV1IncidentsIncidentIDTeamAssignmentsParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDTeamAssignmentsParamsWithContext creates a new PostV1IncidentsIncidentIDTeamAssignmentsParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDTeamAssignmentsParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	return &PostV1IncidentsIncidentIDTeamAssignmentsParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDTeamAssignmentsParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDTeamAssignmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDTeamAssignmentsParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	return &PostV1IncidentsIncidentIDTeamAssignmentsParams{
		HTTPClient: client,
	}
}

/* PostV1IncidentsIncidentIDTeamAssignmentsParams contains all the parameters to send to the API endpoint
   for the post v1 incidents incident Id team assignments operation.

   Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDTeamAssignmentsParams struct {

	// V1IncidentsIncidentIDTeamAssignments.
	V1IncidentsIncidentIDTeamAssignments *models.PostV1IncidentsIncidentIDTeamAssignments

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id team assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WithDefaults() *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id team assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IncidentsIncidentIDTeamAssignments adds the v1IncidentsIncidentIDTeamAssignments to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WithV1IncidentsIncidentIDTeamAssignments(v1IncidentsIncidentIDTeamAssignments *models.PostV1IncidentsIncidentIDTeamAssignments) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	o.SetV1IncidentsIncidentIDTeamAssignments(v1IncidentsIncidentIDTeamAssignments)
	return o
}

// SetV1IncidentsIncidentIDTeamAssignments adds the v1IncidentsIncidentIdTeamAssignments to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) SetV1IncidentsIncidentIDTeamAssignments(v1IncidentsIncidentIDTeamAssignments *models.PostV1IncidentsIncidentIDTeamAssignments) {
	o.V1IncidentsIncidentIDTeamAssignments = v1IncidentsIncidentIDTeamAssignments
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDTeamAssignmentsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id team assignments params
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDTeamAssignmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1IncidentsIncidentIDTeamAssignments != nil {
		if err := r.SetBodyParam(o.V1IncidentsIncidentIDTeamAssignments); err != nil {
			return err
		}
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
