// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetMaintenanceWindowsReader is a Reader for the GetMaintenanceWindows structure.
type GetMaintenanceWindowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMaintenanceWindowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMaintenanceWindowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMaintenanceWindowsOK creates a GetMaintenanceWindowsOK with default headers values
func NewGetMaintenanceWindowsOK() *GetMaintenanceWindowsOK {
	return &GetMaintenanceWindowsOK{}
}

/*GetMaintenanceWindowsOK handles this case with default header values.

successful operation
*/
type GetMaintenanceWindowsOK struct {
	Payload *dynatrace.StubList
}

func (o *GetMaintenanceWindowsOK) Error() string {
	return fmt.Sprintf("[GET /maintenanceWindows][%d] getMaintenanceWindowsOK  %+v", 200, o.Payload)
}

func (o *GetMaintenanceWindowsOK) GetPayload() *dynatrace.StubList {
	return o.Payload
}

func (o *GetMaintenanceWindowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.StubList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
