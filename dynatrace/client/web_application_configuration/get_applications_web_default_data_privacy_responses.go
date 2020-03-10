// Code generated by go-swagger; DO NOT EDIT.

package web_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetApplicationsWebDefaultDataPrivacyReader is a Reader for the GetApplicationsWebDefaultDataPrivacy structure.
type GetApplicationsWebDefaultDataPrivacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsWebDefaultDataPrivacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsWebDefaultDataPrivacyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationsWebDefaultDataPrivacyOK creates a GetApplicationsWebDefaultDataPrivacyOK with default headers values
func NewGetApplicationsWebDefaultDataPrivacyOK() *GetApplicationsWebDefaultDataPrivacyOK {
	return &GetApplicationsWebDefaultDataPrivacyOK{}
}

/*GetApplicationsWebDefaultDataPrivacyOK handles this case with default header values.

successful operation
*/
type GetApplicationsWebDefaultDataPrivacyOK struct {
	Payload *dynatrace.ApplicationDataPrivacy
}

func (o *GetApplicationsWebDefaultDataPrivacyOK) Error() string {
	return fmt.Sprintf("[GET /applications/web/default/dataPrivacy][%d] getApplicationsWebDefaultDataPrivacyOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsWebDefaultDataPrivacyOK) GetPayload() *dynatrace.ApplicationDataPrivacy {
	return o.Payload
}

func (o *GetApplicationsWebDefaultDataPrivacyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ApplicationDataPrivacy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
