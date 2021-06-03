// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1IncidentsIncidentIDRelatedChangeEventsReader is a Reader for the GetV1IncidentsIncidentIDRelatedChangeEvents structure.
type GetV1IncidentsIncidentIDRelatedChangeEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDRelatedChangeEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsIncidentIDRelatedChangeEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDRelatedChangeEventsOK creates a GetV1IncidentsIncidentIDRelatedChangeEventsOK with default headers values
func NewGetV1IncidentsIncidentIDRelatedChangeEventsOK() *GetV1IncidentsIncidentIDRelatedChangeEventsOK {
	return &GetV1IncidentsIncidentIDRelatedChangeEventsOK{}
}

/* GetV1IncidentsIncidentIDRelatedChangeEventsOK describes a response with status code 200, with default header values.

Retrieve all change events that have been associated to the incident
*/
type GetV1IncidentsIncidentIDRelatedChangeEventsOK struct {
	Payload *models.RelatedChangeEventEntityPaginated
}

func (o *GetV1IncidentsIncidentIDRelatedChangeEventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/related_change_events][%d] getV1IncidentsIncidentIdRelatedChangeEventsOK  %+v", 200, o.Payload)
}
func (o *GetV1IncidentsIncidentIDRelatedChangeEventsOK) GetPayload() *models.RelatedChangeEventEntityPaginated {
	return o.Payload
}

func (o *GetV1IncidentsIncidentIDRelatedChangeEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelatedChangeEventEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
