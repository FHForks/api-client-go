// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1PostMortemsReportsReportIDReader is a Reader for the PatchV1PostMortemsReportsReportID structure.
type PatchV1PostMortemsReportsReportIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1PostMortemsReportsReportIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1PostMortemsReportsReportIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchV1PostMortemsReportsReportIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1PostMortemsReportsReportIDOK creates a PatchV1PostMortemsReportsReportIDOK with default headers values
func NewPatchV1PostMortemsReportsReportIDOK() *PatchV1PostMortemsReportsReportIDOK {
	return &PatchV1PostMortemsReportsReportIDOK{}
}

/* PatchV1PostMortemsReportsReportIDOK describes a response with status code 200, with default header values.

Update a post mortem report
*/
type PatchV1PostMortemsReportsReportIDOK struct {
	Payload *models.ReportEntity
}

func (o *PatchV1PostMortemsReportsReportIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/post_mortems/reports/{report_id}][%d] patchV1PostMortemsReportsReportIdOK  %+v", 200, o.Payload)
}
func (o *PatchV1PostMortemsReportsReportIDOK) GetPayload() *models.ReportEntity {
	return o.Payload
}

func (o *PatchV1PostMortemsReportsReportIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchV1PostMortemsReportsReportIDBadRequest creates a PatchV1PostMortemsReportsReportIDBadRequest with default headers values
func NewPatchV1PostMortemsReportsReportIDBadRequest() *PatchV1PostMortemsReportsReportIDBadRequest {
	return &PatchV1PostMortemsReportsReportIDBadRequest{}
}

/* PatchV1PostMortemsReportsReportIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchV1PostMortemsReportsReportIDBadRequest struct {
	Payload *models.ErrorEntity
}

func (o *PatchV1PostMortemsReportsReportIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/post_mortems/reports/{report_id}][%d] patchV1PostMortemsReportsReportIdBadRequest  %+v", 400, o.Payload)
}
func (o *PatchV1PostMortemsReportsReportIDBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PatchV1PostMortemsReportsReportIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
