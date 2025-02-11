// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1RunbooksRunbookIDReader is a Reader for the GetV1RunbooksRunbookID structure.
type GetV1RunbooksRunbookIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbooksRunbookIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbooksRunbookIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RunbooksRunbookIDOK creates a GetV1RunbooksRunbookIDOK with default headers values
func NewGetV1RunbooksRunbookIDOK() *GetV1RunbooksRunbookIDOK {
	return &GetV1RunbooksRunbookIDOK{}
}

/* GetV1RunbooksRunbookIDOK describes a response with status code 200, with default header values.

Get a runbook and all its configuration
*/
type GetV1RunbooksRunbookIDOK struct {
	Payload *models.RunbookEntity
}

func (o *GetV1RunbooksRunbookIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbooks/{runbook_id}][%d] getV1RunbooksRunbookIdOK  %+v", 200, o.Payload)
}
func (o *GetV1RunbooksRunbookIDOK) GetPayload() *models.RunbookEntity {
	return o.Payload
}

func (o *GetV1RunbooksRunbookIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
