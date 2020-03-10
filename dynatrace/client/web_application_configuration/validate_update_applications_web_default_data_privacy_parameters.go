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

// NewValidateUpdateApplicationsWebDefaultDataPrivacyParams creates a new ValidateUpdateApplicationsWebDefaultDataPrivacyParams object
// with the default values initialized.
func NewValidateUpdateApplicationsWebDefaultDataPrivacyParams() *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	var ()
	return &ValidateUpdateApplicationsWebDefaultDataPrivacyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateUpdateApplicationsWebDefaultDataPrivacyParamsWithTimeout creates a new ValidateUpdateApplicationsWebDefaultDataPrivacyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateUpdateApplicationsWebDefaultDataPrivacyParamsWithTimeout(timeout time.Duration) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	var ()
	return &ValidateUpdateApplicationsWebDefaultDataPrivacyParams{

		timeout: timeout,
	}
}

// NewValidateUpdateApplicationsWebDefaultDataPrivacyParamsWithContext creates a new ValidateUpdateApplicationsWebDefaultDataPrivacyParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateUpdateApplicationsWebDefaultDataPrivacyParamsWithContext(ctx context.Context) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	var ()
	return &ValidateUpdateApplicationsWebDefaultDataPrivacyParams{

		Context: ctx,
	}
}

// NewValidateUpdateApplicationsWebDefaultDataPrivacyParamsWithHTTPClient creates a new ValidateUpdateApplicationsWebDefaultDataPrivacyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateUpdateApplicationsWebDefaultDataPrivacyParamsWithHTTPClient(client *http.Client) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	var ()
	return &ValidateUpdateApplicationsWebDefaultDataPrivacyParams{
		HTTPClient: client,
	}
}

/*ValidateUpdateApplicationsWebDefaultDataPrivacyParams contains all the parameters to send to the API endpoint
for the validate update applications web default data privacy operation typically these are written to a http.Request
*/
type ValidateUpdateApplicationsWebDefaultDataPrivacyParams struct {

	/*Body
	  JSON body of the request, containing the data privacy settings to validate.

	*/
	Body *dynatrace.ApplicationDataPrivacy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) WithTimeout(timeout time.Duration) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) WithContext(ctx context.Context) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) WithHTTPClient(client *http.Client) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) WithBody(body *dynatrace.ApplicationDataPrivacy) *ValidateUpdateApplicationsWebDefaultDataPrivacyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate update applications web default data privacy params
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) SetBody(body *dynatrace.ApplicationDataPrivacy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateUpdateApplicationsWebDefaultDataPrivacyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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