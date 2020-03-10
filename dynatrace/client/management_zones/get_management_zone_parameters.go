// Code generated by go-swagger; DO NOT EDIT.

package management_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetManagementZoneParams creates a new GetManagementZoneParams object
// with the default values initialized.
func NewGetManagementZoneParams() *GetManagementZoneParams {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetManagementZoneParams{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetManagementZoneParamsWithTimeout creates a new GetManagementZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetManagementZoneParamsWithTimeout(timeout time.Duration) *GetManagementZoneParams {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetManagementZoneParams{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,

		timeout: timeout,
	}
}

// NewGetManagementZoneParamsWithContext creates a new GetManagementZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetManagementZoneParamsWithContext(ctx context.Context) *GetManagementZoneParams {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetManagementZoneParams{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,

		Context: ctx,
	}
}

// NewGetManagementZoneParamsWithHTTPClient creates a new GetManagementZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetManagementZoneParamsWithHTTPClient(client *http.Client) *GetManagementZoneParams {
	var (
		includeProcessGroupReferencesDefault = bool(false)
	)
	return &GetManagementZoneParams{
		IncludeProcessGroupReferences: &includeProcessGroupReferencesDefault,
		HTTPClient:                    client,
	}
}

/*GetManagementZoneParams contains all the parameters to send to the API endpoint
for the get management zone operation typically these are written to a http.Request
*/
type GetManagementZoneParams struct {

	/*ID
	  The ID of the required management zone.

	*/
	ID string
	/*IncludeProcessGroupReferences
	 Flag to include process group references to the response.

	Process group references aren't compatible across environments. When this flag is set to false, conditions with process group references will be removed. If that leads to a rule having no conditions, the entire rule will be removed.

	*/
	IncludeProcessGroupReferences *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get management zone params
func (o *GetManagementZoneParams) WithTimeout(timeout time.Duration) *GetManagementZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get management zone params
func (o *GetManagementZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get management zone params
func (o *GetManagementZoneParams) WithContext(ctx context.Context) *GetManagementZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get management zone params
func (o *GetManagementZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get management zone params
func (o *GetManagementZoneParams) WithHTTPClient(client *http.Client) *GetManagementZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get management zone params
func (o *GetManagementZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get management zone params
func (o *GetManagementZoneParams) WithID(id string) *GetManagementZoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get management zone params
func (o *GetManagementZoneParams) SetID(id string) {
	o.ID = id
}

// WithIncludeProcessGroupReferences adds the includeProcessGroupReferences to the get management zone params
func (o *GetManagementZoneParams) WithIncludeProcessGroupReferences(includeProcessGroupReferences *bool) *GetManagementZoneParams {
	o.SetIncludeProcessGroupReferences(includeProcessGroupReferences)
	return o
}

// SetIncludeProcessGroupReferences adds the includeProcessGroupReferences to the get management zone params
func (o *GetManagementZoneParams) SetIncludeProcessGroupReferences(includeProcessGroupReferences *bool) {
	o.IncludeProcessGroupReferences = includeProcessGroupReferences
}

// WriteToRequest writes these params to a swagger request
func (o *GetManagementZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.IncludeProcessGroupReferences != nil {

		// query param includeProcessGroupReferences
		var qrIncludeProcessGroupReferences bool
		if o.IncludeProcessGroupReferences != nil {
			qrIncludeProcessGroupReferences = *o.IncludeProcessGroupReferences
		}
		qIncludeProcessGroupReferences := swag.FormatBool(qrIncludeProcessGroupReferences)
		if qIncludeProcessGroupReferences != "" {
			if err := r.SetQueryParam("includeProcessGroupReferences", qIncludeProcessGroupReferences); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
