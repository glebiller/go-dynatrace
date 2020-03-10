// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

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

// NewValidateUpdateAlertingProfileParams creates a new ValidateUpdateAlertingProfileParams object
// with the default values initialized.
func NewValidateUpdateAlertingProfileParams() *ValidateUpdateAlertingProfileParams {
	var ()
	return &ValidateUpdateAlertingProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateAlertingProfileParamsWithTimeout creates a new ValidateUpdateAlertingProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateAlertingProfileParamsWithTimeout(timeout time.Duration) *ValidateUpdateAlertingProfileParams {
	var ()
	return &ValidateUpdateAlertingProfileParams{

		timeout: timeout,
	}
}

// NewValidateUpdateAlertingProfileParamsWithContext creates a new ValidateUpdateAlertingProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateAlertingProfileParamsWithContext(ctx context.Context) *ValidateUpdateAlertingProfileParams {
	var ()
	return &ValidateUpdateAlertingProfileParams{

		Context: ctx,
	}
}

// NewValidateUpdateAlertingProfileParamsWithHTTPClient creates a new ValidateUpdateAlertingProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateAlertingProfileParamsWithHTTPClient(client *http.Client) *ValidateUpdateAlertingProfileParams {
	var ()
	return &ValidateUpdateAlertingProfileParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateAlertingProfileParams contains all the parameters to send to the API endpoint
for the validate update alerting profile operation typically these are written to a http.Request
*/
type ValidateUpdateAlertingProfileParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the alerting profile to be validated.

	*/
	Body *dynatrace.AlertingProfile
	/*ID
	  The ID of the alerting profile to be validated.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) WithTimeout(timeout time.Duration) *ValidateUpdateAlertingProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) WithContext(ctx context.Context) *ValidateUpdateAlertingProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) WithHTTPClient(client *http.Client) *ValidateUpdateAlertingProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) WithBody(body *dynatrace.AlertingProfile) *ValidateUpdateAlertingProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) SetBody(body *dynatrace.AlertingProfile) {
	o.Body = body
}

// WithID adds the id to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) WithID(id strfmt.UUID) *ValidateUpdateAlertingProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate update alerting profile params
func (o *ValidateUpdateAlertingProfileParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateAlertingProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
