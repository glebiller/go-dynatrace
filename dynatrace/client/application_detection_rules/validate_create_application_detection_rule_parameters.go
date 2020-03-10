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

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewValidateCreateApplicationDetectionRuleParams creates a new ValidateCreateApplicationDetectionRuleParams object
// with the default values initialized.
func NewValidateCreateApplicationDetectionRuleParams() *ValidateCreateApplicationDetectionRuleParams {
	var ()
	return &ValidateCreateApplicationDetectionRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateCreateApplicationDetectionRuleParamsWithTimeout creates a new ValidateCreateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateCreateApplicationDetectionRuleParamsWithTimeout(timeout time.Duration) *ValidateCreateApplicationDetectionRuleParams {
	var ()
	return &ValidateCreateApplicationDetectionRuleParams{

		timeout: timeout,
	}
}

// NewValidateCreateApplicationDetectionRuleParamsWithContext creates a new ValidateCreateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateCreateApplicationDetectionRuleParamsWithContext(ctx context.Context) *ValidateCreateApplicationDetectionRuleParams {
	var ()
	return &ValidateCreateApplicationDetectionRuleParams{

		Context: ctx,
	}
}

// NewValidateCreateApplicationDetectionRuleParamsWithHTTPClient creates a new ValidateCreateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateCreateApplicationDetectionRuleParamsWithHTTPClient(client *http.Client) *ValidateCreateApplicationDetectionRuleParams {
	var ()
	return &ValidateCreateApplicationDetectionRuleParams{
		HTTPClient: client,
	}
}

/*ValidateCreateApplicationDetectionRuleParams contains all the parameters to send to the API endpoint
for the validate create application detection rule operation typically these are written to a http.Request
*/
type ValidateCreateApplicationDetectionRuleParams struct {

	/*Body
	  The JSON body of the request. Contains the application detection rule to be validated.

	*/
	Body *dynatrace.ApplicationDetectionRuleConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) WithTimeout(timeout time.Duration) *ValidateCreateApplicationDetectionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) WithContext(ctx context.Context) *ValidateCreateApplicationDetectionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) WithHTTPClient(client *http.Client) *ValidateCreateApplicationDetectionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) WithBody(body *dynatrace.ApplicationDetectionRuleConfig) *ValidateCreateApplicationDetectionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate create application detection rule params
func (o *ValidateCreateApplicationDetectionRuleParams) SetBody(body *dynatrace.ApplicationDetectionRuleConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateCreateApplicationDetectionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
