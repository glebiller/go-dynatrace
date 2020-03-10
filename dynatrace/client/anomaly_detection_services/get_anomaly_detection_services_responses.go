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

// GetAnomalyDetectionServicesReader is a Reader for the GetAnomalyDetectionServices structure.
type GetAnomalyDetectionServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnomalyDetectionServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnomalyDetectionServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAnomalyDetectionServicesOK creates a GetAnomalyDetectionServicesOK with default headers values
func NewGetAnomalyDetectionServicesOK() *GetAnomalyDetectionServicesOK {
	return &GetAnomalyDetectionServicesOK{}
}

/*GetAnomalyDetectionServicesOK handles this case with default header values.

successful operation
*/
type GetAnomalyDetectionServicesOK struct {
	Payload *dynatrace.ServiceAnomalyDetectionConfig
}

func (o *GetAnomalyDetectionServicesOK) Error() string {
	return fmt.Sprintf("[GET /anomalyDetection/services][%d] getAnomalyDetectionServicesOK  %+v", 200, o.Payload)
}

func (o *GetAnomalyDetectionServicesOK) GetPayload() *dynatrace.ServiceAnomalyDetectionConfig {
	return o.Payload
}

func (o *GetAnomalyDetectionServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ServiceAnomalyDetectionConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}