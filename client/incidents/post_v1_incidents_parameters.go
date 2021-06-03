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

// NewPostV1IncidentsParams creates a new PostV1IncidentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsParams() *PostV1IncidentsParams {
	return &PostV1IncidentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsParamsWithTimeout creates a new PostV1IncidentsParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsParamsWithTimeout(timeout time.Duration) *PostV1IncidentsParams {
	return &PostV1IncidentsParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsParamsWithContext creates a new PostV1IncidentsParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsParamsWithContext(ctx context.Context) *PostV1IncidentsParams {
	return &PostV1IncidentsParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsParamsWithHTTPClient creates a new PostV1IncidentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsParamsWithHTTPClient(client *http.Client) *PostV1IncidentsParams {
	return &PostV1IncidentsParams{
		HTTPClient: client,
	}
}

/* PostV1IncidentsParams contains all the parameters to send to the API endpoint
   for the post v1 incidents operation.

   Typically these are written to a http.Request.
*/
type PostV1IncidentsParams struct {

	// V1Incidents.
	V1Incidents *models.PostV1Incidents

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsParams) WithDefaults() *PostV1IncidentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents params
func (o *PostV1IncidentsParams) WithTimeout(timeout time.Duration) *PostV1IncidentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents params
func (o *PostV1IncidentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents params
func (o *PostV1IncidentsParams) WithContext(ctx context.Context) *PostV1IncidentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents params
func (o *PostV1IncidentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents params
func (o *PostV1IncidentsParams) WithHTTPClient(client *http.Client) *PostV1IncidentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents params
func (o *PostV1IncidentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Incidents adds the v1Incidents to the post v1 incidents params
func (o *PostV1IncidentsParams) WithV1Incidents(v1Incidents *models.PostV1Incidents) *PostV1IncidentsParams {
	o.SetV1Incidents(v1Incidents)
	return o
}

// SetV1Incidents adds the v1Incidents to the post v1 incidents params
func (o *PostV1IncidentsParams) SetV1Incidents(v1Incidents *models.PostV1Incidents) {
	o.V1Incidents = v1Incidents
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Incidents != nil {
		if err := r.SetBodyParam(o.V1Incidents); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
