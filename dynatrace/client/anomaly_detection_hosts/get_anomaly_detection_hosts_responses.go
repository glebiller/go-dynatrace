// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetAnomalyDetectionHostsReader is a Reader for the GetAnomalyDetectionHosts structure.
type GetAnomalyDetectionHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnomalyDetectionHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnomalyDetectionHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAnomalyDetectionHostsOK creates a GetAnomalyDetectionHostsOK with default headers values
func NewGetAnomalyDetectionHostsOK() *GetAnomalyDetectionHostsOK {
	return &GetAnomalyDetectionHostsOK{}
}

/*GetAnomalyDetectionHostsOK handles this case with default header values.

successful operation
*/
type GetAnomalyDetectionHostsOK struct {
	Payload *dynatrace.HostsAnomalyDetectionConfig
}

func (o *GetAnomalyDetectionHostsOK) Error() string {
	return fmt.Sprintf("[GET /anomalyDetection/hosts][%d] getAnomalyDetectionHostsOK  %+v", 200, o.Payload)
}

func (o *GetAnomalyDetectionHostsOK) GetPayload() *dynatrace.HostsAnomalyDetectionConfig {
	return o.Payload
}

func (o *GetAnomalyDetectionHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.HostsAnomalyDetectionConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
