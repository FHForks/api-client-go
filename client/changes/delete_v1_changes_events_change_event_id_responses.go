// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1ChangesEventsChangeEventIDReader is a Reader for the DeleteV1ChangesEventsChangeEventID structure.
type DeleteV1ChangesEventsChangeEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ChangesEventsChangeEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1ChangesEventsChangeEventIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1ChangesEventsChangeEventIDNoContent creates a DeleteV1ChangesEventsChangeEventIDNoContent with default headers values
func NewDeleteV1ChangesEventsChangeEventIDNoContent() *DeleteV1ChangesEventsChangeEventIDNoContent {
	return &DeleteV1ChangesEventsChangeEventIDNoContent{}
}

/* DeleteV1ChangesEventsChangeEventIDNoContent describes a response with status code 204, with default header values.

Delete a change event
*/
type DeleteV1ChangesEventsChangeEventIDNoContent struct {
}

func (o *DeleteV1ChangesEventsChangeEventIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/changes/events/{change_event_id}][%d] deleteV1ChangesEventsChangeEventIdNoContent ", 204)
}

func (o *DeleteV1ChangesEventsChangeEventIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
