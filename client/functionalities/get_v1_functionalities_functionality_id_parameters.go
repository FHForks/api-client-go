// Code generated by go-swagger; DO NOT EDIT.

package functionalities

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

// NewGetV1FunctionalitiesFunctionalityIDParams creates a new GetV1FunctionalitiesFunctionalityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1FunctionalitiesFunctionalityIDParams() *GetV1FunctionalitiesFunctionalityIDParams {
	return &GetV1FunctionalitiesFunctionalityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1FunctionalitiesFunctionalityIDParamsWithTimeout creates a new GetV1FunctionalitiesFunctionalityIDParams object
// with the ability to set a timeout on a request.
func NewGetV1FunctionalitiesFunctionalityIDParamsWithTimeout(timeout time.Duration) *GetV1FunctionalitiesFunctionalityIDParams {
	return &GetV1FunctionalitiesFunctionalityIDParams{
		timeout: timeout,
	}
}

// NewGetV1FunctionalitiesFunctionalityIDParamsWithContext creates a new GetV1FunctionalitiesFunctionalityIDParams object
// with the ability to set a context for a request.
func NewGetV1FunctionalitiesFunctionalityIDParamsWithContext(ctx context.Context) *GetV1FunctionalitiesFunctionalityIDParams {
	return &GetV1FunctionalitiesFunctionalityIDParams{
		Context: ctx,
	}
}

// NewGetV1FunctionalitiesFunctionalityIDParamsWithHTTPClient creates a new GetV1FunctionalitiesFunctionalityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1FunctionalitiesFunctionalityIDParamsWithHTTPClient(client *http.Client) *GetV1FunctionalitiesFunctionalityIDParams {
	return &GetV1FunctionalitiesFunctionalityIDParams{
		HTTPClient: client,
	}
}

/* GetV1FunctionalitiesFunctionalityIDParams contains all the parameters to send to the API endpoint
   for the get v1 functionalities functionality Id operation.

   Typically these are written to a http.Request.
*/
type GetV1FunctionalitiesFunctionalityIDParams struct {

	// FunctionalityID.
	FunctionalityID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 functionalities functionality Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FunctionalitiesFunctionalityIDParams) WithDefaults() *GetV1FunctionalitiesFunctionalityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 functionalities functionality Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FunctionalitiesFunctionalityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) WithTimeout(timeout time.Duration) *GetV1FunctionalitiesFunctionalityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) WithContext(ctx context.Context) *GetV1FunctionalitiesFunctionalityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) WithHTTPClient(client *http.Client) *GetV1FunctionalitiesFunctionalityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFunctionalityID adds the functionalityID to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) WithFunctionalityID(functionalityID string) *GetV1FunctionalitiesFunctionalityIDParams {
	o.SetFunctionalityID(functionalityID)
	return o
}

// SetFunctionalityID adds the functionalityId to the get v1 functionalities functionality Id params
func (o *GetV1FunctionalitiesFunctionalityIDParams) SetFunctionalityID(functionalityID string) {
	o.FunctionalityID = functionalityID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1FunctionalitiesFunctionalityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param functionality_id
	if err := r.SetPathParam("functionality_id", o.FunctionalityID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
