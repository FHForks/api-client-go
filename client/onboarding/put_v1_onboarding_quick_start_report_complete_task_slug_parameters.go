// Code generated by go-swagger; DO NOT EDIT.

package onboarding

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

// NewPutV1OnboardingQuickStartReportCompleteTaskSlugParams creates a new PutV1OnboardingQuickStartReportCompleteTaskSlugParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1OnboardingQuickStartReportCompleteTaskSlugParams() *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	return &PutV1OnboardingQuickStartReportCompleteTaskSlugParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1OnboardingQuickStartReportCompleteTaskSlugParamsWithTimeout creates a new PutV1OnboardingQuickStartReportCompleteTaskSlugParams object
// with the ability to set a timeout on a request.
func NewPutV1OnboardingQuickStartReportCompleteTaskSlugParamsWithTimeout(timeout time.Duration) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	return &PutV1OnboardingQuickStartReportCompleteTaskSlugParams{
		timeout: timeout,
	}
}

// NewPutV1OnboardingQuickStartReportCompleteTaskSlugParamsWithContext creates a new PutV1OnboardingQuickStartReportCompleteTaskSlugParams object
// with the ability to set a context for a request.
func NewPutV1OnboardingQuickStartReportCompleteTaskSlugParamsWithContext(ctx context.Context) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	return &PutV1OnboardingQuickStartReportCompleteTaskSlugParams{
		Context: ctx,
	}
}

// NewPutV1OnboardingQuickStartReportCompleteTaskSlugParamsWithHTTPClient creates a new PutV1OnboardingQuickStartReportCompleteTaskSlugParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1OnboardingQuickStartReportCompleteTaskSlugParamsWithHTTPClient(client *http.Client) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	return &PutV1OnboardingQuickStartReportCompleteTaskSlugParams{
		HTTPClient: client,
	}
}

/* PutV1OnboardingQuickStartReportCompleteTaskSlugParams contains all the parameters to send to the API endpoint
   for the put v1 onboarding quick start report complete task slug operation.

   Typically these are written to a http.Request.
*/
type PutV1OnboardingQuickStartReportCompleteTaskSlugParams struct {

	// TaskSlug.
	//
	// Format: int32
	TaskSlug int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 onboarding quick start report complete task slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) WithDefaults() *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 onboarding quick start report complete task slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) WithTimeout(timeout time.Duration) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) WithContext(ctx context.Context) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) WithHTTPClient(client *http.Client) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskSlug adds the taskSlug to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) WithTaskSlug(taskSlug int32) *PutV1OnboardingQuickStartReportCompleteTaskSlugParams {
	o.SetTaskSlug(taskSlug)
	return o
}

// SetTaskSlug adds the taskSlug to the put v1 onboarding quick start report complete task slug params
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) SetTaskSlug(taskSlug int32) {
	o.TaskSlug = taskSlug
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1OnboardingQuickStartReportCompleteTaskSlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param task_slug
	if err := r.SetPathParam("task_slug", swag.FormatInt32(o.TaskSlug)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
