// Code generated by go-swagger; DO NOT EDIT.

package severities

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

// NewGetV1SeveritiesSeveritySlugParams creates a new GetV1SeveritiesSeveritySlugParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SeveritiesSeveritySlugParams() *GetV1SeveritiesSeveritySlugParams {
	return &GetV1SeveritiesSeveritySlugParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SeveritiesSeveritySlugParamsWithTimeout creates a new GetV1SeveritiesSeveritySlugParams object
// with the ability to set a timeout on a request.
func NewGetV1SeveritiesSeveritySlugParamsWithTimeout(timeout time.Duration) *GetV1SeveritiesSeveritySlugParams {
	return &GetV1SeveritiesSeveritySlugParams{
		timeout: timeout,
	}
}

// NewGetV1SeveritiesSeveritySlugParamsWithContext creates a new GetV1SeveritiesSeveritySlugParams object
// with the ability to set a context for a request.
func NewGetV1SeveritiesSeveritySlugParamsWithContext(ctx context.Context) *GetV1SeveritiesSeveritySlugParams {
	return &GetV1SeveritiesSeveritySlugParams{
		Context: ctx,
	}
}

// NewGetV1SeveritiesSeveritySlugParamsWithHTTPClient creates a new GetV1SeveritiesSeveritySlugParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SeveritiesSeveritySlugParamsWithHTTPClient(client *http.Client) *GetV1SeveritiesSeveritySlugParams {
	return &GetV1SeveritiesSeveritySlugParams{
		HTTPClient: client,
	}
}

/* GetV1SeveritiesSeveritySlugParams contains all the parameters to send to the API endpoint
   for the get v1 severities severity slug operation.

   Typically these are written to a http.Request.
*/
type GetV1SeveritiesSeveritySlugParams struct {

	// SeveritySlug.
	SeveritySlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 severities severity slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeveritiesSeveritySlugParams) WithDefaults() *GetV1SeveritiesSeveritySlugParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 severities severity slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeveritiesSeveritySlugParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) WithTimeout(timeout time.Duration) *GetV1SeveritiesSeveritySlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) WithContext(ctx context.Context) *GetV1SeveritiesSeveritySlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) WithHTTPClient(client *http.Client) *GetV1SeveritiesSeveritySlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSeveritySlug adds the severitySlug to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) WithSeveritySlug(severitySlug string) *GetV1SeveritiesSeveritySlugParams {
	o.SetSeveritySlug(severitySlug)
	return o
}

// SetSeveritySlug adds the severitySlug to the get v1 severities severity slug params
func (o *GetV1SeveritiesSeveritySlugParams) SetSeveritySlug(severitySlug string) {
	o.SeveritySlug = severitySlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SeveritiesSeveritySlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param severity_slug
	if err := r.SetPathParam("severity_slug", o.SeveritySlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
