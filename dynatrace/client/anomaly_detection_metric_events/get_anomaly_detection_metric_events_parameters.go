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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAnomalyDetectionMetricEventsParams creates a new GetAnomalyDetectionMetricEventsParams object
// with the default values initialized.
func NewGetAnomalyDetectionMetricEventsParams() *GetAnomalyDetectionMetricEventsParams {
	var (
		includeEntityFilterMetricEventsDefault = bool(false)
	)
	return &GetAnomalyDetectionMetricEventsParams{
		IncludeEntityFilterMetricEvents: &includeEntityFilterMetricEventsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAnomalyDetectionMetricEventsParamsWithTimeout creates a new GetAnomalyDetectionMetricEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAnomalyDetectionMetricEventsParamsWithTimeout(timeout time.Duration) *GetAnomalyDetectionMetricEventsParams {
	var (
		includeEntityFilterMetricEventsDefault = bool(false)
	)
	return &GetAnomalyDetectionMetricEventsParams{
		IncludeEntityFilterMetricEvents: &includeEntityFilterMetricEventsDefault,

		timeout: timeout,
	}
}

// NewGetAnomalyDetectionMetricEventsParamsWithContext creates a new GetAnomalyDetectionMetricEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAnomalyDetectionMetricEventsParamsWithContext(ctx context.Context) *GetAnomalyDetectionMetricEventsParams {
	var (
		includeEntityFilterMetricEventsDefault = bool(false)
	)
	return &GetAnomalyDetectionMetricEventsParams{
		IncludeEntityFilterMetricEvents: &includeEntityFilterMetricEventsDefault,

		Context: ctx,
	}
}

// NewGetAnomalyDetectionMetricEventsParamsWithHTTPClient creates a new GetAnomalyDetectionMetricEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAnomalyDetectionMetricEventsParamsWithHTTPClient(client *http.Client) *GetAnomalyDetectionMetricEventsParams {
	var (
		includeEntityFilterMetricEventsDefault = bool(false)
	)
	return &GetAnomalyDetectionMetricEventsParams{
		IncludeEntityFilterMetricEvents: &includeEntityFilterMetricEventsDefault,
		HTTPClient:                      client,
	}
}

/*GetAnomalyDetectionMetricEventsParams contains all the parameters to send to the API endpoint
for the get anomaly detection metric events operation typically these are written to a http.Request
*/
type GetAnomalyDetectionMetricEventsParams struct {

	/*IncludeEntityFilterMetricEvents
	 Flag to include metric events with associated entities to the response.

	Metric events with **entity** filters aren't compatible across environments. If set to `false`, metric events with entity filters will be removed.

	*/
	IncludeEntityFilterMetricEvents *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) WithTimeout(timeout time.Duration) *GetAnomalyDetectionMetricEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) WithContext(ctx context.Context) *GetAnomalyDetectionMetricEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) WithHTTPClient(client *http.Client) *GetAnomalyDetectionMetricEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeEntityFilterMetricEvents adds the includeEntityFilterMetricEvents to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) WithIncludeEntityFilterMetricEvents(includeEntityFilterMetricEvents *bool) *GetAnomalyDetectionMetricEventsParams {
	o.SetIncludeEntityFilterMetricEvents(includeEntityFilterMetricEvents)
	return o
}

// SetIncludeEntityFilterMetricEvents adds the includeEntityFilterMetricEvents to the get anomaly detection metric events params
func (o *GetAnomalyDetectionMetricEventsParams) SetIncludeEntityFilterMetricEvents(includeEntityFilterMetricEvents *bool) {
	o.IncludeEntityFilterMetricEvents = includeEntityFilterMetricEvents
}

// WriteToRequest writes these params to a swagger request
func (o *GetAnomalyDetectionMetricEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeEntityFilterMetricEvents != nil {

		// query param includeEntityFilterMetricEvents
		var qrIncludeEntityFilterMetricEvents bool
		if o.IncludeEntityFilterMetricEvents != nil {
			qrIncludeEntityFilterMetricEvents = *o.IncludeEntityFilterMetricEvents
		}
		qIncludeEntityFilterMetricEvents := swag.FormatBool(qrIncludeEntityFilterMetricEvents)
		if qIncludeEntityFilterMetricEvents != "" {
			if err := r.SetQueryParam("includeEntityFilterMetricEvents", qIncludeEntityFilterMetricEvents); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
