// Code generated by go-swagger; DO NOT EDIT.

package web_application_configuration

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

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewUpdateApplicationsWebDefaultParams creates a new UpdateApplicationsWebDefaultParams object
// with the default values initialized.
func NewUpdateApplicationsWebDefaultParams() *UpdateApplicationsWebDefaultParams {
	var ()
	return &UpdateApplicationsWebDefaultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateApplicationsWebDefaultParamsWithTimeout creates a new UpdateApplicationsWebDefaultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateApplicationsWebDefaultParamsWithTimeout(timeout time.Duration) *UpdateApplicationsWebDefaultParams {
	var ()
	return &UpdateApplicationsWebDefaultParams{

		timeout: timeout,
	}
}

// NewUpdateApplicationsWebDefaultParamsWithContext creates a new UpdateApplicationsWebDefaultParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateApplicationsWebDefaultParamsWithContext(ctx context.Context) *UpdateApplicationsWebDefaultParams {
	var ()
	return &UpdateApplicationsWebDefaultParams{

		Context: ctx,
	}
}

// NewUpdateApplicationsWebDefaultParamsWithHTTPClient creates a new UpdateApplicationsWebDefaultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateApplicationsWebDefaultParamsWithHTTPClient(client *http.Client) *UpdateApplicationsWebDefaultParams {
	var ()
	return &UpdateApplicationsWebDefaultParams{
		HTTPClient: client,
	}
}

/*UpdateApplicationsWebDefaultParams contains all the parameters to send to the API endpoint
for the update applications web default operation typically these are written to a http.Request
*/
type UpdateApplicationsWebDefaultParams struct {

	/*Body
	  JSON body of the request, containing the new parameters of the default web application.

	*/
	Body *dynatrace.WebApplicationConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) WithTimeout(timeout time.Duration) *UpdateApplicationsWebDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) WithContext(ctx context.Context) *UpdateApplicationsWebDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) WithHTTPClient(client *http.Client) *UpdateApplicationsWebDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) WithBody(body *dynatrace.WebApplicationConfig) *UpdateApplicationsWebDefaultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update applications web default params
func (o *UpdateApplicationsWebDefaultParams) SetBody(body *dynatrace.WebApplicationConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateApplicationsWebDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
