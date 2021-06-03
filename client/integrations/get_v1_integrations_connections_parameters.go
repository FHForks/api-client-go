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

// NewGetV1IntegrationsConnectionsParams creates a new GetV1IntegrationsConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsConnectionsParams() *GetV1IntegrationsConnectionsParams {
	return &GetV1IntegrationsConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsConnectionsParamsWithTimeout creates a new GetV1IntegrationsConnectionsParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsConnectionsParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsConnectionsParams {
	return &GetV1IntegrationsConnectionsParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsConnectionsParamsWithContext creates a new GetV1IntegrationsConnectionsParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsConnectionsParamsWithContext(ctx context.Context) *GetV1IntegrationsConnectionsParams {
	return &GetV1IntegrationsConnectionsParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsConnectionsParamsWithHTTPClient creates a new GetV1IntegrationsConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsConnectionsParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsConnectionsParams {
	return &GetV1IntegrationsConnectionsParams{
		HTTPClient: client,
	}
}

/* GetV1IntegrationsConnectionsParams contains all the parameters to send to the API endpoint
   for the get v1 integrations connections operation.

   Typically these are written to a http.Request.
*/
type GetV1IntegrationsConnectionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsConnectionsParams) WithDefaults() *GetV1IntegrationsConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations connections params
func (o *GetV1IntegrationsConnectionsParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations connections params
func (o *GetV1IntegrationsConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations connections params
func (o *GetV1IntegrationsConnectionsParams) WithContext(ctx context.Context) *GetV1IntegrationsConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations connections params
func (o *GetV1IntegrationsConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations connections params
func (o *GetV1IntegrationsConnectionsParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations connections params
func (o *GetV1IntegrationsConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
