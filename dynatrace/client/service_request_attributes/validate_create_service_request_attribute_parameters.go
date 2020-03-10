// Code generated by go-swagger; DO NOT EDIT.

package service_request_attributes

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

// NewValidateCreateServiceRequestAttributeParams creates a new ValidateCreateServiceRequestAttributeParams object
// with the default values initialized.
func NewValidateCreateServiceRequestAttributeParams() *ValidateCreateServiceRequestAttributeParams {
	var ()
	return &ValidateCreateServiceRequestAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateServiceRequestAttributeParamsWithTimeout creates a new ValidateCreateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateServiceRequestAttributeParamsWithTimeout(timeout time.Duration) *ValidateCreateServiceRequestAttributeParams {
	var ()
	return &ValidateCreateServiceRequestAttributeParams{

		timeout: timeout,
	}
}

// NewValidateCreateServiceRequestAttributeParamsWithContext creates a new ValidateCreateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateServiceRequestAttributeParamsWithContext(ctx context.Context) *ValidateCreateServiceRequestAttributeParams {
	var ()
	return &ValidateCreateServiceRequestAttributeParams{

		Context: ctx,
	}
}

// NewValidateCreateServiceRequestAttributeParamsWithHTTPClient creates a new ValidateCreateServiceRequestAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateServiceRequestAttributeParamsWithHTTPClient(client *http.Client) *ValidateCreateServiceRequestAttributeParams {
	var ()
	return &ValidateCreateServiceRequestAttributeParams{
		HTTPClient: client,
	}
}

/*ValidateCreateServiceRequestAttributeParams contains all the parameters to send to the API endpoint
for the validate create service request attribute operation typically these are written to a http.Request
*/
type ValidateCreateServiceRequestAttributeParams struct {

	/*Body*/
	Body *dynatrace.RequestAttribute

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) WithTimeout(timeout time.Duration) *ValidateCreateServiceRequestAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) WithContext(ctx context.Context) *ValidateCreateServiceRequestAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) WithHTTPClient(client *http.Client) *ValidateCreateServiceRequestAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) WithBody(body *dynatrace.RequestAttribute) *ValidateCreateServiceRequestAttributeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create service request attribute params
func (o *ValidateCreateServiceRequestAttributeParams) SetBody(body *dynatrace.RequestAttribute) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateServiceRequestAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
