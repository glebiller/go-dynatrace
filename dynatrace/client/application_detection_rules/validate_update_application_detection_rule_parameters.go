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

// NewValidateUpdateApplicationDetectionRuleParams creates a new ValidateUpdateApplicationDetectionRuleParams object
// with the default values initialized.
func NewValidateUpdateApplicationDetectionRuleParams() *ValidateUpdateApplicationDetectionRuleParams {
	var ()
	return &ValidateUpdateApplicationDetectionRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateApplicationDetectionRuleParamsWithTimeout creates a new ValidateUpdateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateApplicationDetectionRuleParamsWithTimeout(timeout time.Duration) *ValidateUpdateApplicationDetectionRuleParams {
	var ()
	return &ValidateUpdateApplicationDetectionRuleParams{

		timeout: timeout,
	}
}

// NewValidateUpdateApplicationDetectionRuleParamsWithContext creates a new ValidateUpdateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateApplicationDetectionRuleParamsWithContext(ctx context.Context) *ValidateUpdateApplicationDetectionRuleParams {
	var ()
	return &ValidateUpdateApplicationDetectionRuleParams{

		Context: ctx,
	}
}

// NewValidateUpdateApplicationDetectionRuleParamsWithHTTPClient creates a new ValidateUpdateApplicationDetectionRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateApplicationDetectionRuleParamsWithHTTPClient(client *http.Client) *ValidateUpdateApplicationDetectionRuleParams {
	var ()
	return &ValidateUpdateApplicationDetectionRuleParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateApplicationDetectionRuleParams contains all the parameters to send to the API endpoint
for the validate update application detection rule operation typically these are written to a http.Request
*/
type ValidateUpdateApplicationDetectionRuleParams struct {

	/*Body
	  The JSON body of the request. Contains the application detection rule to be validated.

	*/
	Body *dynatrace.ApplicationDetectionRuleConfig
	/*ID
	 The ID of the application detection rule to be validated.

	If you set the ID in the body as well, it must match this ID.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) WithTimeout(timeout time.Duration) *ValidateUpdateApplicationDetectionRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) WithContext(ctx context.Context) *ValidateUpdateApplicationDetectionRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) WithHTTPClient(client *http.Client) *ValidateUpdateApplicationDetectionRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) WithBody(body *dynatrace.ApplicationDetectionRuleConfig) *ValidateUpdateApplicationDetectionRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) SetBody(body *dynatrace.ApplicationDetectionRuleConfig) {
	o.Body = body
}

// WithID adds the id to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) WithID(id strfmt.UUID) *ValidateUpdateApplicationDetectionRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate update application detection rule params
func (o *ValidateUpdateApplicationDetectionRuleParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateApplicationDetectionRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}