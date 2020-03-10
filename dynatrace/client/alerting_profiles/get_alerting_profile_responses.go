// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetAlertingProfileReader is a Reader for the GetAlertingProfile structure.
type GetAlertingProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlertingProfileOK creates a GetAlertingProfileOK with default headers values
func NewGetAlertingProfileOK() *GetAlertingProfileOK {
	return &GetAlertingProfileOK{}
}

/*GetAlertingProfileOK handles this case with default header values.

successful operation
*/
type GetAlertingProfileOK struct {
	Payload *dynatrace.AlertingProfile
}

func (o *GetAlertingProfileOK) Error() string {
	return fmt.Sprintf("[GET /alertingProfiles/{id}][%d] getAlertingProfileOK  %+v", 200, o.Payload)
}

func (o *GetAlertingProfileOK) GetPayload() *dynatrace.AlertingProfile {
	return o.Payload
}

func (o *GetAlertingProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.AlertingProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
