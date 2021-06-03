// Code generated by go-swagger; DO NOT EDIT.

package severity_matrix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1SeverityMatrixImpactsImpactIDReader is a Reader for the PatchV1SeverityMatrixImpactsImpactID structure.
type PatchV1SeverityMatrixImpactsImpactIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1SeverityMatrixImpactsImpactIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1SeverityMatrixImpactsImpactIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1SeverityMatrixImpactsImpactIDOK creates a PatchV1SeverityMatrixImpactsImpactIDOK with default headers values
func NewPatchV1SeverityMatrixImpactsImpactIDOK() *PatchV1SeverityMatrixImpactsImpactIDOK {
	return &PatchV1SeverityMatrixImpactsImpactIDOK{}
}

/* PatchV1SeverityMatrixImpactsImpactIDOK describes a response with status code 200, with default header values.

Update a specific impact
*/
type PatchV1SeverityMatrixImpactsImpactIDOK struct {
	Payload *models.ImpactEntity
}

func (o *PatchV1SeverityMatrixImpactsImpactIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/severity_matrix/impacts/{impact_id}][%d] patchV1SeverityMatrixImpactsImpactIdOK  %+v", 200, o.Payload)
}
func (o *PatchV1SeverityMatrixImpactsImpactIDOK) GetPayload() *models.ImpactEntity {
	return o.Payload
}

func (o *PatchV1SeverityMatrixImpactsImpactIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImpactEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
