// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1TicketingTicketsReader is a Reader for the GetV1TicketingTickets structure.
type GetV1TicketingTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TicketingTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TicketingTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TicketingTicketsOK creates a GetV1TicketingTicketsOK with default headers values
func NewGetV1TicketingTicketsOK() *GetV1TicketingTicketsOK {
	return &GetV1TicketingTicketsOK{}
}

/* GetV1TicketingTicketsOK describes a response with status code 200, with default header values.

List all of the functionalities that have been added to the organiation
*/
type GetV1TicketingTicketsOK struct {
	Payload *models.TicketEntity
}

func (o *GetV1TicketingTicketsOK) Error() string {
	return fmt.Sprintf("[GET /v1/ticketing/tickets][%d] getV1TicketingTicketsOK  %+v", 200, o.Payload)
}
func (o *GetV1TicketingTicketsOK) GetPayload() *models.TicketEntity {
	return o.Payload
}

func (o *GetV1TicketingTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
