// Code generated by go-swagger; DO NOT EDIT.

package notes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1NotesReader is a Reader for the PostV1Notes structure.
type PostV1NotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1NotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1NotesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1NotesCreated creates a PostV1NotesCreated with default headers values
func NewPostV1NotesCreated() *PostV1NotesCreated {
	return &PostV1NotesCreated{}
}

/* PostV1NotesCreated describes a response with status code 201, with default header values.

Add a note to an incident
*/
type PostV1NotesCreated struct {
}

func (o *PostV1NotesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/notes][%d] postV1NotesCreated ", 201)
}

func (o *PostV1NotesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
