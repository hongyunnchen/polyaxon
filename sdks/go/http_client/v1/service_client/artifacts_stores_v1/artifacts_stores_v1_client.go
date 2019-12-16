// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package artifacts_stores_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new artifacts stores v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for artifacts stores v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateArtifactsStore lists runs
*/
func (a *Client) CreateArtifactsStore(params *CreateArtifactsStoreParams, authInfo runtime.ClientAuthInfoWriter) (*CreateArtifactsStoreOK, *CreateArtifactsStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateArtifactsStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateArtifactsStore",
		Method:             "POST",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateArtifactsStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateArtifactsStoreOK:
		return value, nil, nil
	case *CreateArtifactsStoreNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteArtifactsStore patches run
*/
func (a *Client) DeleteArtifactsStore(params *DeleteArtifactsStoreParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteArtifactsStoreOK, *DeleteArtifactsStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteArtifactsStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteArtifactsStore",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteArtifactsStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteArtifactsStoreOK:
		return value, nil, nil
	case *DeleteArtifactsStoreNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetArtifactsStore creates new run
*/
func (a *Client) GetArtifactsStore(params *GetArtifactsStoreParams, authInfo runtime.ClientAuthInfoWriter) (*GetArtifactsStoreOK, *GetArtifactsStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetArtifactsStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetArtifactsStore",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetArtifactsStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetArtifactsStoreOK:
		return value, nil, nil
	case *GetArtifactsStoreNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListArtifactsStoreNames lists bookmarked runs for user
*/
func (a *Client) ListArtifactsStoreNames(params *ListArtifactsStoreNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListArtifactsStoreNamesOK, *ListArtifactsStoreNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListArtifactsStoreNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListArtifactsStoreNames",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListArtifactsStoreNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListArtifactsStoreNamesOK:
		return value, nil, nil
	case *ListArtifactsStoreNamesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListArtifactsStores lists archived runs for user
*/
func (a *Client) ListArtifactsStores(params *ListArtifactsStoresParams, authInfo runtime.ClientAuthInfoWriter) (*ListArtifactsStoresOK, *ListArtifactsStoresNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListArtifactsStoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListArtifactsStores",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListArtifactsStoresReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListArtifactsStoresOK:
		return value, nil, nil
	case *ListArtifactsStoresNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchArtifactsStore updates run
*/
func (a *Client) PatchArtifactsStore(params *PatchArtifactsStoreParams, authInfo runtime.ClientAuthInfoWriter) (*PatchArtifactsStoreOK, *PatchArtifactsStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchArtifactsStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchArtifactsStore",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores/{artifact_store.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchArtifactsStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchArtifactsStoreOK:
		return value, nil, nil
	case *PatchArtifactsStoreNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateArtifactsStore gets run
*/
func (a *Client) UpdateArtifactsStore(params *UpdateArtifactsStoreParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateArtifactsStoreOK, *UpdateArtifactsStoreNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateArtifactsStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateArtifactsStore",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/artifacts_stores/{artifact_store.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateArtifactsStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateArtifactsStoreOK:
		return value, nil, nil
	case *UpdateArtifactsStoreNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadArtifact uploads artifact to a store
*/
func (a *Client) UploadArtifact(params *UploadArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*UploadArtifactOK, *UploadArtifactNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadArtifact",
		Method:             "POST",
		PathPattern:        "/api/v1/catalogs/{owner}/artifacts_stores/{uuid}/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UploadArtifactOK:
		return value, nil, nil
	case *UploadArtifactNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for artifacts_stores_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}