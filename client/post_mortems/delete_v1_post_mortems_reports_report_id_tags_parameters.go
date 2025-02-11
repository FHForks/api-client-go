// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

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

// NewDeleteV1PostMortemsReportsReportIDTagsParams creates a new DeleteV1PostMortemsReportsReportIDTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1PostMortemsReportsReportIDTagsParams() *DeleteV1PostMortemsReportsReportIDTagsParams {
	return &DeleteV1PostMortemsReportsReportIDTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDTagsParamsWithTimeout creates a new DeleteV1PostMortemsReportsReportIDTagsParams object
// with the ability to set a timeout on a request.
func NewDeleteV1PostMortemsReportsReportIDTagsParamsWithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDTagsParams {
	return &DeleteV1PostMortemsReportsReportIDTagsParams{
		timeout: timeout,
	}
}

// NewDeleteV1PostMortemsReportsReportIDTagsParamsWithContext creates a new DeleteV1PostMortemsReportsReportIDTagsParams object
// with the ability to set a context for a request.
func NewDeleteV1PostMortemsReportsReportIDTagsParamsWithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDTagsParams {
	return &DeleteV1PostMortemsReportsReportIDTagsParams{
		Context: ctx,
	}
}

// NewDeleteV1PostMortemsReportsReportIDTagsParamsWithHTTPClient creates a new DeleteV1PostMortemsReportsReportIDTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1PostMortemsReportsReportIDTagsParamsWithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDTagsParams {
	return &DeleteV1PostMortemsReportsReportIDTagsParams{
		HTTPClient: client,
	}
}

/* DeleteV1PostMortemsReportsReportIDTagsParams contains all the parameters to send to the API endpoint
   for the delete v1 post mortems reports report Id tags operation.

   Typically these are written to a http.Request.
*/
type DeleteV1PostMortemsReportsReportIDTagsParams struct {

	// ReportID.
	ReportID string

	// Tags.
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 post mortems reports report Id tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WithDefaults() *DeleteV1PostMortemsReportsReportIDTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 post mortems reports report Id tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WithTimeout(timeout time.Duration) *DeleteV1PostMortemsReportsReportIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WithContext(ctx context.Context) *DeleteV1PostMortemsReportsReportIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WithHTTPClient(client *http.Client) *DeleteV1PostMortemsReportsReportIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportID adds the reportID to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WithReportID(reportID string) *DeleteV1PostMortemsReportsReportIDTagsParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WithTags adds the tags to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WithTags(tags []string) *DeleteV1PostMortemsReportsReportIDTagsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete v1 post mortems reports report Id tags params
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	if o.Tags != nil {

		// binding items for tags
		joinedTags := o.bindParamTags(reg)

		// form array param tags
		if err := r.SetFormParam("tags", joinedTags...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteV1PostMortemsReportsReportIDTags binds the parameter tags
func (o *DeleteV1PostMortemsReportsReportIDTagsParams) bindParamTags(formats strfmt.Registry) []string {
	tagsIR := o.Tags

	var tagsIC []string
	for _, tagsIIR := range tagsIR { // explode []string

		tagsIIV := tagsIIR // string as string
		tagsIC = append(tagsIC, tagsIIV)
	}

	// items.CollectionFormat: ""
	tagsIS := swag.JoinByFormat(tagsIC, "")

	return tagsIS
}
