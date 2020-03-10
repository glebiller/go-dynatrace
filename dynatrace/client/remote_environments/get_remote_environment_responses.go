// Code generated by go-swagger; DO NOT EDIT.

package remote_environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetRemoteEnvironmentReader is a Reader for the GetRemoteEnvironment structure.
type GetRemoteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRemoteEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRemoteEnvironmentOK creates a GetRemoteEnvironmentOK with default headers values
func NewGetRemoteEnvironmentOK() *GetRemoteEnvironmentOK {
	return &GetRemoteEnvironmentOK{}
}

/*GetRemoteEnvironmentOK handles this case with default header values.

successful operation
*/
type GetRemoteEnvironmentOK struct {
	Payload *dynatrace.RemoteEnvironmentConfigDto
}

func (o *GetRemoteEnvironmentOK) Error() string {
	return fmt.Sprintf("[GET /remoteEnvironments/{id}][%d] getRemoteEnvironmentOK  %+v", 200, o.Payload)
}

func (o *GetRemoteEnvironmentOK) GetPayload() *dynatrace.RemoteEnvironmentConfigDto {
	return o.Payload
}

func (o *GetRemoteEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.RemoteEnvironmentConfigDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
