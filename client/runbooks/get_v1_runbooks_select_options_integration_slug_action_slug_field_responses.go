// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldReader is a Reader for the GetV1RunbooksSelectOptionsIntegrationSlugActionSlugField structure.
type GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK creates a GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK with default headers values
func NewGetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK() *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK {
	return &GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK{}
}

/* GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK describes a response with status code 200, with default header values.

get SelectOption(s)
*/
type GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK struct {
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbooks/select_options/{integration_slug}/{action_slug}/{field}][%d] getV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK ", 200)
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
