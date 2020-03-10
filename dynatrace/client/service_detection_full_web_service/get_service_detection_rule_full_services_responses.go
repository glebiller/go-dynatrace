// Code generated by go-swagger; DO NOT EDIT.

package service_detection_full_web_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetServiceDetectionRuleFullServicesReader is a Reader for the GetServiceDetectionRuleFullServices structure.
type GetServiceDetectionRuleFullServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceDetectionRuleFullServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceDetectionRuleFullServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceDetectionRuleFullServicesOK creates a GetServiceDetectionRuleFullServicesOK with default headers values
func NewGetServiceDetectionRuleFullServicesOK() *GetServiceDetectionRuleFullServicesOK {
	return &GetServiceDetectionRuleFullServicesOK{}
}

/*GetServiceDetectionRuleFullServicesOK handles this case with default header values.

Success. The response contains the ordered list of rules.
*/
type GetServiceDetectionRuleFullServicesOK struct {
	Payload *dynatrace.StubList
}

func (o *GetServiceDetectionRuleFullServicesOK) Error() string {
	return fmt.Sprintf("[GET /service/detectionRules/FULL_WEB_SERVICE][%d] getServiceDetectionRuleFullServicesOK  %+v", 200, o.Payload)
}

func (o *GetServiceDetectionRuleFullServicesOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetServiceDetectionRuleFullServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
