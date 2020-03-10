// Code generated by go-swagger; DO NOT EDIT.

package service_detection_full_web_request

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

// NewCreateServiceDetectionRuleFullRequestParams creates a new CreateServiceDetectionRuleFullRequestParams object
// with the default values initialized.
func NewCreateServiceDetectionRuleFullRequestParams() *CreateServiceDetectionRuleFullRequestParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullRequestParams{
		Position: &positionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceDetectionRuleFullRequestParamsWithTimeout creates a new CreateServiceDetectionRuleFullRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateServiceDetectionRuleFullRequestParamsWithTimeout(timeout time.Duration) *CreateServiceDetectionRuleFullRequestParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullRequestParams{
		Position: &positionDefault,

		timeout: timeout,
	}
}

// NewCreateServiceDetectionRuleFullRequestParamsWithContext creates a new CreateServiceDetectionRuleFullRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateServiceDetectionRuleFullRequestParamsWithContext(ctx context.Context) *CreateServiceDetectionRuleFullRequestParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullRequestParams{
		Position: &positionDefault,

		Context: ctx,
	}
}

// NewCreateServiceDetectionRuleFullRequestParamsWithHTTPClient creates a new CreateServiceDetectionRuleFullRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateServiceDetectionRuleFullRequestParamsWithHTTPClient(client *http.Client) *CreateServiceDetectionRuleFullRequestParams {
	var (
		positionDefault = string("APPEND")
	)
	return &CreateServiceDetectionRuleFullRequestParams{
		Position:   &positionDefault,
		HTTPClient: client,
	}
}

/*CreateServiceDetectionRuleFullRequestParams contains all the parameters to send to the API endpoint
for the create service detection rule full request operation typically these are written to a http.Request
*/
type CreateServiceDetectionRuleFullRequestParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the new service detection rule.

	 You must not specify the ID of the rule!

	The **order** field is ignored in this request. To enforce a particular order, use the `PUT /service/detectionRules/FULL_WEB_REQUEST/reorder` request.

	*/
	Body *dynatrace.FullWebRequestRule
	/*Position
	 The position of the new rule:

	* `APPEND`: at the bottom of the rule list.
	* `PREPEND`: at the top of the rule list.

	If not set, the `APPEND` is used.

	*/
	Position *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) WithTimeout(timeout time.Duration) *CreateServiceDetectionRuleFullRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) WithContext(ctx context.Context) *CreateServiceDetectionRuleFullRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) WithHTTPClient(client *http.Client) *CreateServiceDetectionRuleFullRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) WithBody(body *dynatrace.FullWebRequestRule) *CreateServiceDetectionRuleFullRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) SetBody(body *dynatrace.FullWebRequestRule) {
	o.Body = body
}

// WithPosition adds the position to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) WithPosition(position *string) *CreateServiceDetectionRuleFullRequestParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the create service detection rule full request params
func (o *CreateServiceDetectionRuleFullRequestParams) SetPosition(position *string) {
	o.Position = position
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceDetectionRuleFullRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Position != nil {

		// query param position
		var qrPosition string
		if o.Position != nil {
			qrPosition = *o.Position
		}
		qPosition := qrPosition
		if qPosition != "" {
			if err := r.SetQueryParam("position", qPosition); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}