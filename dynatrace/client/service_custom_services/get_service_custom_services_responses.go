// Code generated by go-swagger; DO NOT EDIT.

package service_custom_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetServiceCustomServicesReader is a Reader for the GetServiceCustomServices structure.
type GetServiceCustomServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCustomServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCustomServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceCustomServicesOK creates a GetServiceCustomServicesOK with default headers values
func NewGetServiceCustomServicesOK() *GetServiceCustomServicesOK {
	return &GetServiceCustomServicesOK{}
}

/*GetServiceCustomServicesOK handles this case with default header values.

successful operation
*/
type GetServiceCustomServicesOK struct {
	Payload *dynatrace.StubList
}

func (o *GetServiceCustomServicesOK) Error() string {
	return fmt.Sprintf("[GET /service/customServices/{technology}][%d] getServiceCustomServicesOK  %+v", 200, o.Payload)
}

func (o *GetServiceCustomServicesOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetServiceCustomServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}