// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schedules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schedules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Schedules(params *GetV1SchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SchedulesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetV1Schedules lists all schedules

  List all known schedules in FireHydrant as pulled from external sources
*/
func (a *Client) GetV1Schedules(params *GetV1SchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SchedulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Schedules",
		Method:             "GET",
		PathPattern:        "/v1/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SchedulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1SchedulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Schedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
