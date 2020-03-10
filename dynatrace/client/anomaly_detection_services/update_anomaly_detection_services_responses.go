// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateAnomalyDetectionServicesReader is a Reader for the UpdateAnomalyDetectionServices structure.
type UpdateAnomalyDetectionServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAnomalyDetectionServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateAnomalyDetectionServicesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAnomalyDetectionServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAnomalyDetectionServicesNoContent creates a UpdateAnomalyDetectionServicesNoContent with default headers values
func NewUpdateAnomalyDetectionServicesNoContent() *UpdateAnomalyDetectionServicesNoContent {
	return &UpdateAnomalyDetectionServicesNoContent{}
}

/*UpdateAnomalyDetectionServicesNoContent handles this case with default header values.

Success. Configuration has been updated. Response doesn't have a body.
*/
type UpdateAnomalyDetectionServicesNoContent struct {
}

func (o *UpdateAnomalyDetectionServicesNoContent) Error() string {
	return fmt.Sprintf("[PUT /anomalyDetection/services][%d] updateAnomalyDetectionServicesNoContent ", 204)
}

func (o *UpdateAnomalyDetectionServicesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAnomalyDetectionServicesBadRequest creates a UpdateAnomalyDetectionServicesBadRequest with default headers values
func NewUpdateAnomalyDetectionServicesBadRequest() *UpdateAnomalyDetectionServicesBadRequest {
	return &UpdateAnomalyDetectionServicesBadRequest{}
}

/*UpdateAnomalyDetectionServicesBadRequest handles this case with default header values.

Failed. The input is invalid
*/
type UpdateAnomalyDetectionServicesBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateAnomalyDetectionServicesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /anomalyDetection/services][%d] updateAnomalyDetectionServicesBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAnomalyDetectionServicesBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateAnomalyDetectionServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
