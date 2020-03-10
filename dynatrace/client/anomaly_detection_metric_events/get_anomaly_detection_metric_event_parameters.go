// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_metric_events

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

// NewGetAnomalyDetectionMetricEventParams creates a new GetAnomalyDetectionMetricEventParams object
// with the default values initialized.
func NewGetAnomalyDetectionMetricEventParams() *GetAnomalyDetectionMetricEventParams {
	var ()
	return &GetAnomalyDetectionMetricEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnomalyDetectionMetricEventParamsWithTimeout creates a new GetAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnomalyDetectionMetricEventParamsWithTimeout(timeout time.Duration) *GetAnomalyDetectionMetricEventParams {
	var ()
	return &GetAnomalyDetectionMetricEventParams{

		timeout: timeout,
	}
}

// NewGetAnomalyDetectionMetricEventParamsWithContext creates a new GetAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnomalyDetectionMetricEventParamsWithContext(ctx context.Context) *GetAnomalyDetectionMetricEventParams {
	var ()
	return &GetAnomalyDetectionMetricEventParams{

		Context: ctx,
	}
}

// NewGetAnomalyDetectionMetricEventParamsWithHTTPClient creates a new GetAnomalyDetectionMetricEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnomalyDetectionMetricEventParamsWithHTTPClient(client *http.Client) *GetAnomalyDetectionMetricEventParams {
	var ()
	return &GetAnomalyDetectionMetricEventParams{
		HTTPClient: client,
	}
}

/*GetAnomalyDetectionMetricEventParams contains all the parameters to send to the API endpoint
for the get anomaly detection metric event operation typically these are written to a http.Request
*/
type GetAnomalyDetectionMetricEventParams struct {

	/*ID
	  The ID of the required metric event.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) WithTimeout(timeout time.Duration) *GetAnomalyDetectionMetricEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) WithContext(ctx context.Context) *GetAnomalyDetectionMetricEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) WithHTTPClient(client *http.Client) *GetAnomalyDetectionMetricEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) WithID(id string) *GetAnomalyDetectionMetricEventParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get anomaly detection metric event params
func (o *GetAnomalyDetectionMetricEventParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnomalyDetectionMetricEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
