// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new changes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for changes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1ChangesChangeID(params *DeleteV1ChangesChangeIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ChangesChangeIDNoContent, error)

	DeleteV1ChangesChangeIDIdentitiesIdentityID(params *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ChangesChangeIDIdentitiesIdentityIDNoContent, error)

	DeleteV1ChangesEventsChangeEventID(params *DeleteV1ChangesEventsChangeEventIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ChangesEventsChangeEventIDNoContent, error)

	GetV1Changes(params *GetV1ChangesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesOK, error)

	GetV1ChangesChangeIDIdentities(params *GetV1ChangesChangeIDIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesChangeIDIdentitiesOK, error)

	GetV1ChangesEvents(params *GetV1ChangesEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesEventsOK, error)

	GetV1ChangesEventsChangeEventID(params *GetV1ChangesEventsChangeEventIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesEventsChangeEventIDOK, error)

	PatchV1ChangesChangeID(params *PatchV1ChangesChangeIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ChangesChangeIDOK, error)

	PatchV1ChangesChangeIDIdentitiesIdentityID(params *PatchV1ChangesChangeIDIdentitiesIdentityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ChangesChangeIDIdentitiesIdentityIDOK, error)

	PatchV1ChangesEventsChangeEventID(params *PatchV1ChangesEventsChangeEventIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ChangesEventsChangeEventIDOK, error)

	PostV1Changes(params *PostV1ChangesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ChangesCreated, error)

	PostV1ChangesChangeIDIdentities(params *PostV1ChangesChangeIDIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ChangesChangeIDIdentitiesCreated, error)

	PostV1ChangesEvents(params *PostV1ChangesEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ChangesEventsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteV1ChangesChangeID Archive a change entry
*/
func (a *Client) DeleteV1ChangesChangeID(params *DeleteV1ChangesChangeIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ChangesChangeIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ChangesChangeIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ChangesChangeId",
		Method:             "DELETE",
		PathPattern:        "/v1/changes/{change_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ChangesChangeIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1ChangesChangeIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ChangesChangeId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteV1ChangesChangeIDIdentitiesIdentityID Delete an identity
*/
func (a *Client) DeleteV1ChangesChangeIDIdentitiesIdentityID(params *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ChangesChangeIDIdentitiesIdentityIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ChangesChangeIdIdentitiesIdentityId",
		Method:             "DELETE",
		PathPattern:        "/v1/changes/{change_id}/identities/{identity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ChangesChangeIDIdentitiesIdentityIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1ChangesChangeIDIdentitiesIdentityIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ChangesChangeIdIdentitiesIdentityId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteV1ChangesEventsChangeEventID deletes a change event

  Delete a change event
*/
func (a *Client) DeleteV1ChangesEventsChangeEventID(params *DeleteV1ChangesEventsChangeEventIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ChangesEventsChangeEventIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ChangesEventsChangeEventIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ChangesEventsChangeEventId",
		Method:             "DELETE",
		PathPattern:        "/v1/changes/events/{change_event_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ChangesEventsChangeEventIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1ChangesEventsChangeEventIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ChangesEventsChangeEventId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Changes Lists all changes
*/
func (a *Client) GetV1Changes(params *GetV1ChangesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ChangesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Changes",
		Method:             "GET",
		PathPattern:        "/v1/changes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ChangesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ChangesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Changes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1ChangesChangeIDIdentities Retrieve all identities for the change
*/
func (a *Client) GetV1ChangesChangeIDIdentities(params *GetV1ChangesChangeIDIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesChangeIDIdentitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ChangesChangeIDIdentitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ChangesChangeIdIdentities",
		Method:             "GET",
		PathPattern:        "/v1/changes/{change_id}/identities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ChangesChangeIDIdentitiesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ChangesChangeIDIdentitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ChangesChangeIdIdentities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1ChangesEvents lists change events

  List change events for the organization. Note: Not all information is included on a change event like attachments and related changes. You must fetch a change event separately to retrieve all of the information about it
*/
func (a *Client) GetV1ChangesEvents(params *GetV1ChangesEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ChangesEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ChangesEvents",
		Method:             "GET",
		PathPattern:        "/v1/changes/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ChangesEventsReader{formats: a.formats},
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
	success, ok := result.(*GetV1ChangesEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ChangesEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1ChangesEventsChangeEventID retrieves a change event

  Retrieve a change event
*/
func (a *Client) GetV1ChangesEventsChangeEventID(params *GetV1ChangesEventsChangeEventIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ChangesEventsChangeEventIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ChangesEventsChangeEventIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ChangesEventsChangeEventId",
		Method:             "GET",
		PathPattern:        "/v1/changes/events/{change_event_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ChangesEventsChangeEventIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1ChangesEventsChangeEventIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ChangesEventsChangeEventId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1ChangesChangeID Update a change entry
*/
func (a *Client) PatchV1ChangesChangeID(params *PatchV1ChangesChangeIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ChangesChangeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ChangesChangeIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1ChangesChangeId",
		Method:             "PATCH",
		PathPattern:        "/v1/changes/{change_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ChangesChangeIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ChangesChangeIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1ChangesChangeId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1ChangesChangeIDIdentitiesIdentityID Update an identity
*/
func (a *Client) PatchV1ChangesChangeIDIdentitiesIdentityID(params *PatchV1ChangesChangeIDIdentitiesIdentityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ChangesChangeIDIdentitiesIdentityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ChangesChangeIDIdentitiesIdentityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1ChangesChangeIdIdentitiesIdentityId",
		Method:             "PATCH",
		PathPattern:        "/v1/changes/{change_id}/identities/{identity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ChangesChangeIDIdentitiesIdentityIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ChangesChangeIDIdentitiesIdentityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1ChangesChangeIdIdentitiesIdentityId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchV1ChangesEventsChangeEventID updates a change event

  Update a change event
*/
func (a *Client) PatchV1ChangesEventsChangeEventID(params *PatchV1ChangesEventsChangeEventIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ChangesEventsChangeEventIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ChangesEventsChangeEventIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1ChangesEventsChangeEventId",
		Method:             "PATCH",
		PathPattern:        "/v1/changes/events/{change_event_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ChangesEventsChangeEventIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ChangesEventsChangeEventIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1ChangesEventsChangeEventId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1Changes Create a new change entry
*/
func (a *Client) PostV1Changes(params *PostV1ChangesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ChangesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ChangesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1Changes",
		Method:             "POST",
		PathPattern:        "/v1/changes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ChangesReader{formats: a.formats},
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
	success, ok := result.(*PostV1ChangesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1Changes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1ChangesChangeIDIdentities Create an identity for this change
*/
func (a *Client) PostV1ChangesChangeIDIdentities(params *PostV1ChangesChangeIDIdentitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ChangesChangeIDIdentitiesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ChangesChangeIDIdentitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ChangesChangeIdIdentities",
		Method:             "POST",
		PathPattern:        "/v1/changes/{change_id}/identities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ChangesChangeIDIdentitiesReader{formats: a.formats},
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
	success, ok := result.(*PostV1ChangesChangeIDIdentitiesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ChangesChangeIdIdentities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1ChangesEvents creates a change event

  Create a change event
*/
func (a *Client) PostV1ChangesEvents(params *PostV1ChangesEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ChangesEventsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ChangesEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ChangesEvents",
		Method:             "POST",
		PathPattern:        "/v1/changes/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ChangesEventsReader{formats: a.formats},
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
	success, ok := result.(*PostV1ChangesEventsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ChangesEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
