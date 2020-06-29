// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PrivateEndpointConnectionsOperations contains the methods for the PrivateEndpointConnections group.
type PrivateEndpointConnectionsOperations interface {
	// Delete - Deletes the specified private endpoint connection associated with the storage account.
	Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (*http.Response, error)
	// Get - Gets the specified private endpoint connection associated with the storage account.
	Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (*PrivateEndpointConnectionResponse, error)
	// Put - Update the state of specified private endpoint connection associated with the storage account.
	Put(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (*PrivateEndpointConnectionResponse, error)
}

// privateEndpointConnectionsOperations implements the PrivateEndpointConnectionsOperations interface.
type privateEndpointConnectionsOperations struct {
	*Client
	subscriptionID string
}

// Delete - Deletes the specified private endpoint connection associated with the storage account.
func (client *privateEndpointConnectionsOperations) Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (*http.Response, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, accountName, privateEndpointConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteCreateRequest creates the Delete request.
func (client *privateEndpointConnectionsOperations) deleteCreateRequest(resourceGroupName string, accountName string, privateEndpointConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-06-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *privateEndpointConnectionsOperations) deleteHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteHandleError handles the Delete error response.
func (client *privateEndpointConnectionsOperations) deleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified private endpoint connection associated with the storage account.
func (client *privateEndpointConnectionsOperations) Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string) (*PrivateEndpointConnectionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, accountName, privateEndpointConnectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *privateEndpointConnectionsOperations) getCreateRequest(resourceGroupName string, accountName string, privateEndpointConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-06-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *privateEndpointConnectionsOperations) getHandleResponse(resp *azcore.Response) (*PrivateEndpointConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := PrivateEndpointConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointConnection)
}

// getHandleError handles the Get error response.
func (client *privateEndpointConnectionsOperations) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put - Update the state of specified private endpoint connection associated with the storage account.
func (client *privateEndpointConnectionsOperations) Put(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (*PrivateEndpointConnectionResponse, error) {
	req, err := client.putCreateRequest(resourceGroupName, accountName, privateEndpointConnectionName, properties)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putCreateRequest creates the Put request.
func (client *privateEndpointConnectionsOperations) putCreateRequest(resourceGroupName string, accountName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-06-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(properties)
}

// putHandleResponse handles the Put response.
func (client *privateEndpointConnectionsOperations) putHandleResponse(resp *azcore.Response) (*PrivateEndpointConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putHandleError(resp)
	}
	result := PrivateEndpointConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PrivateEndpointConnection)
}

// putHandleError handles the Put error response.
func (client *privateEndpointConnectionsOperations) putHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}