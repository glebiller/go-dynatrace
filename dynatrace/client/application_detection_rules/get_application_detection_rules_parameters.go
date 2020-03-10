// Code generated by go-swagger; DO NOT EDIT.

package application_detection_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetApplicationDetectionRulesParams creates a new GetApplicationDetectionRulesParams object
// with the default values initialized.
func NewGetApplicationDetectionRulesParams() *GetApplicationDetectionRulesParams {

	return &GetApplicationDetectionRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationDetectionRulesParamsWithTimeout creates a new GetApplicationDetectionRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationDetectionRulesParamsWithTimeout(timeout time.Duration) *GetApplicationDetectionRulesParams {

	return &GetApplicationDetectionRulesParams{

		timeout: timeout,
	}
}

// NewGetApplicationDetectionRulesParamsWithContext creates a new GetApplicationDetectionRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationDetectionRulesParamsWithContext(ctx context.Context) *GetApplicationDetectionRulesParams {

	return &GetApplicationDetectionRulesParams{

		Context: ctx,
	}
}

// NewGetApplicationDetectionRulesParamsWithHTTPClient creates a new GetApplicationDetectionRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplicationDetectionRulesParamsWithHTTPClient(client *http.Client) *GetApplicationDetectionRulesParams {

	return &GetApplicationDetectionRulesParams{
		HTTPClient: client,
	}
}

/*GetApplicationDetectionRulesParams contains all the parameters to send to the API endpoint
for the get application detection rules operation typically these are written to a http.Request
*/
type GetApplicationDetectionRulesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get application detection rules params
func (o *GetApplicationDetectionRulesParams) WithTimeout(timeout time.Duration) *GetApplicationDetectionRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application detection rules params
func (o *GetApplicationDetectionRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application detection rules params
func (o *GetApplicationDetectionRulesParams) WithContext(ctx context.Context) *GetApplicationDetectionRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application detection rules params
func (o *GetApplicationDetectionRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application detection rules params
func (o *GetApplicationDetectionRulesParams) WithHTTPClient(client *http.Client) *GetApplicationDetectionRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application detection rules params
func (o *GetApplicationDetectionRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationDetectionRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}