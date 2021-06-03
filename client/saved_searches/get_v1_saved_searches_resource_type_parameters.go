// Code generated by go-swagger; DO NOT EDIT.

package saved_searches

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

// NewGetV1SavedSearchesResourceTypeParams creates a new GetV1SavedSearchesResourceTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SavedSearchesResourceTypeParams() *GetV1SavedSearchesResourceTypeParams {
	return &GetV1SavedSearchesResourceTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SavedSearchesResourceTypeParamsWithTimeout creates a new GetV1SavedSearchesResourceTypeParams object
// with the ability to set a timeout on a request.
func NewGetV1SavedSearchesResourceTypeParamsWithTimeout(timeout time.Duration) *GetV1SavedSearchesResourceTypeParams {
	return &GetV1SavedSearchesResourceTypeParams{
		timeout: timeout,
	}
}

// NewGetV1SavedSearchesResourceTypeParamsWithContext creates a new GetV1SavedSearchesResourceTypeParams object
// with the ability to set a context for a request.
func NewGetV1SavedSearchesResourceTypeParamsWithContext(ctx context.Context) *GetV1SavedSearchesResourceTypeParams {
	return &GetV1SavedSearchesResourceTypeParams{
		Context: ctx,
	}
}

// NewGetV1SavedSearchesResourceTypeParamsWithHTTPClient creates a new GetV1SavedSearchesResourceTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SavedSearchesResourceTypeParamsWithHTTPClient(client *http.Client) *GetV1SavedSearchesResourceTypeParams {
	return &GetV1SavedSearchesResourceTypeParams{
		HTTPClient: client,
	}
}

/* GetV1SavedSearchesResourceTypeParams contains all the parameters to send to the API endpoint
   for the get v1 saved searches resource type operation.

   Typically these are written to a http.Request.
*/
type GetV1SavedSearchesResourceTypeParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   Filter saved searches with a query on their name
	*/
	Query *string

	// ResourceType.
	ResourceType string

	/* UserID.

	   The user ID used to filter saved searches.
	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 saved searches resource type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SavedSearchesResourceTypeParams) WithDefaults() *GetV1SavedSearchesResourceTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 saved searches resource type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SavedSearchesResourceTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithTimeout(timeout time.Duration) *GetV1SavedSearchesResourceTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithContext(ctx context.Context) *GetV1SavedSearchesResourceTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithHTTPClient(client *http.Client) *GetV1SavedSearchesResourceTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithPage(page *int32) *GetV1SavedSearchesResourceTypeParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithPerPage(perPage *int32) *GetV1SavedSearchesResourceTypeParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithQuery(query *string) *GetV1SavedSearchesResourceTypeParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetQuery(query *string) {
	o.Query = query
}

// WithResourceType adds the resourceType to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithResourceType(resourceType string) *GetV1SavedSearchesResourceTypeParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithUserID adds the userID to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) WithUserID(userID *string) *GetV1SavedSearchesResourceTypeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get v1 saved searches resource type params
func (o *GetV1SavedSearchesResourceTypeParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SavedSearchesResourceTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
