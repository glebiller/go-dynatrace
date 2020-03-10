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

// GetApplicationsWebDataPrivacyListReader is a Reader for the GetApplicationsWebDataPrivacyList structure.
type GetApplicationsWebDataPrivacyListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationsWebDataPrivacyListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationsWebDataPrivacyListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationsWebDataPrivacyListOK creates a GetApplicationsWebDataPrivacyListOK with default headers values
func NewGetApplicationsWebDataPrivacyListOK() *GetApplicationsWebDataPrivacyListOK {
	return &GetApplicationsWebDataPrivacyListOK{}
}

/*GetApplicationsWebDataPrivacyListOK handles this case with default header values.

successful operation
*/
type GetApplicationsWebDataPrivacyListOK struct {
	Payload *dynatrace.ApplicationDataPrivacyList
}

func (o *GetApplicationsWebDataPrivacyListOK) Error() string {
	return fmt.Sprintf("[GET /applications/web/dataPrivacy][%d] getApplicationsWebDataPrivacyListOK  %+v", 200, o.Payload)
}

func (o *GetApplicationsWebDataPrivacyListOK) GetPayload() *dynatrace.ApplicationDataPrivacyList {
	return o.Payload
}

func (o *GetApplicationsWebDataPrivacyListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ApplicationDataPrivacyList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
