// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1WebhooksWebhookIDReader is a Reader for the GetV1WebhooksWebhookID structure.
type GetV1WebhooksWebhookIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WebhooksWebhookIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WebhooksWebhookIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1WebhooksWebhookIDOK creates a GetV1WebhooksWebhookIDOK with default headers values
func NewGetV1WebhooksWebhookIDOK() *GetV1WebhooksWebhookIDOK {
	return &GetV1WebhooksWebhookIDOK{}
}

/* GetV1WebhooksWebhookIDOK describes a response with status code 200, with default header values.

Retrieve a specific webhook
*/
type GetV1WebhooksWebhookIDOK struct {
	Payload *models.WebhookEntity
}

func (o *GetV1WebhooksWebhookIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/webhooks/{webhook_id}][%d] getV1WebhooksWebhookIdOK  %+v", 200, o.Payload)
}
func (o *GetV1WebhooksWebhookIDOK) GetPayload() *models.WebhookEntity {
	return o.Payload
}

func (o *GetV1WebhooksWebhookIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
