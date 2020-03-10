// Code generated by go-swagger; DO NOT EDIT.

package management_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetManagementZoneReader is a Reader for the GetManagementZone structure.
type GetManagementZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetManagementZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetManagementZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagementZoneOK creates a GetManagementZoneOK with default headers values
func NewGetManagementZoneOK() *GetManagementZoneOK {
	return &GetManagementZoneOK{}
}

/*GetManagementZoneOK handles this case with default header values.

successful operation
*/
type GetManagementZoneOK struct {
	Payload *dynatrace.ManagementZone
}

func (o *GetManagementZoneOK) Error() string {
	return fmt.Sprintf("[GET /managementZones/{id}][%d] getManagementZoneOK  %+v", 200, o.Payload)
}

func (o *GetManagementZoneOK) GetPayload() *dynatrace.ManagementZone {
	return o.Payload
}

func (o *GetManagementZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ManagementZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}