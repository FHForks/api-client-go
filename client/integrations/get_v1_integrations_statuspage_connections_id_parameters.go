// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetV1IntegrationsStatuspageConnectionsIDParams creates a new GetV1IntegrationsStatuspageConnectionsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsStatuspageConnectionsIDParams() *GetV1IntegrationsStatuspageConnectionsIDParams {
	return &GetV1IntegrationsStatuspageConnectionsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsStatuspageConnectionsIDParamsWithTimeout creates a new GetV1IntegrationsStatuspageConnectionsIDParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsStatuspageConnectionsIDParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsStatuspageConnectionsIDParams {
	return &GetV1IntegrationsStatuspageConnectionsIDParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsStatuspageConnectionsIDParamsWithContext creates a new GetV1IntegrationsStatuspageConnectionsIDParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsStatuspageConnectionsIDParamsWithContext(ctx context.Context) *GetV1IntegrationsStatuspageConnectionsIDParams {
	return &GetV1IntegrationsStatuspageConnectionsIDParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsStatuspageConnectionsIDParamsWithHTTPClient creates a new GetV1IntegrationsStatuspageConnectionsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsStatuspageConnectionsIDParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsStatuspageConnectionsIDParams {
	return &GetV1IntegrationsStatuspageConnectionsIDParams{
		HTTPClient: client,
	}
}

/* GetV1IntegrationsStatuspageConnectionsIDParams contains all the parameters to send to the API endpoint
   for the get v1 integrations statuspage connections Id operation.

   Typically these are written to a http.Request.
*/
type GetV1IntegrationsStatuspageConnectionsIDParams struct {

	/* ID.

	   Connection UUID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations statuspage connections Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) WithDefaults() *GetV1IntegrationsStatuspageConnectionsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations statuspage connections Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsStatuspageConnectionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) WithContext(ctx context.Context) *GetV1IntegrationsStatuspageConnectionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsStatuspageConnectionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) WithID(id string) *GetV1IntegrationsStatuspageConnectionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 integrations statuspage connections Id params
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsStatuspageConnectionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
