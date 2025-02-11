// Code generated by go-swagger; DO NOT EDIT.

package incident_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1IncidentRolesReader is a Reader for the PostV1IncidentRoles structure.
type PostV1IncidentRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IncidentRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1IncidentRolesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1IncidentRolesCreated creates a PostV1IncidentRolesCreated with default headers values
func NewPostV1IncidentRolesCreated() *PostV1IncidentRolesCreated {
	return &PostV1IncidentRolesCreated{}
}

/* PostV1IncidentRolesCreated describes a response with status code 201, with default header values.

Create an incident role
*/
type PostV1IncidentRolesCreated struct {
	Payload *models.IncidentRoleEntity
}

func (o *PostV1IncidentRolesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/incident_roles][%d] postV1IncidentRolesCreated  %+v", 201, o.Payload)
}
func (o *PostV1IncidentRolesCreated) GetPayload() *models.IncidentRoleEntity {
	return o.Payload
}

func (o *PostV1IncidentRolesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentRoleEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
