package azuredata

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SQLManagedInstancesClient is the the AzureData management API provides a RESTful set of web APIs to manage Azure
// Data Resources.
type SQLManagedInstancesClient struct {
	BaseClient
}

// NewSQLManagedInstancesClient creates an instance of the SQLManagedInstancesClient client.
func NewSQLManagedInstancesClient(subscriptionID string, subscriptionID1 string) SQLManagedInstancesClient {
	return NewSQLManagedInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewSQLManagedInstancesClientWithBaseURI creates an instance of the SQLManagedInstancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSQLManagedInstancesClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) SQLManagedInstancesClient {
	return SQLManagedInstancesClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// Create creates or replaces a SQL Managed Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLManagedInstanceName - the name of SQL Managed Instances
// parameters - the SQL Managed Instance to be created or updated.
func (client SQLManagedInstancesClient) Create(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string, parameters SQLManagedInstance) (result SQLManagedInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, SQLManagedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SQLManagedInstancesClient) CreatePreparer(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string, parameters SQLManagedInstance) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlManagedInstanceName": autorest.Encode("path", SQLManagedInstanceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlManagedInstances/{sqlManagedInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SQLManagedInstancesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SQLManagedInstancesClient) CreateResponder(resp *http.Response) (result SQLManagedInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a SQL Managed Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLManagedInstanceName - the name of Sql Managed Instances
func (client SQLManagedInstancesClient) Delete(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, SQLManagedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLManagedInstancesClient) DeletePreparer(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlManagedInstanceName": autorest.Encode("path", SQLManagedInstanceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlManagedInstances/{sqlManagedInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLManagedInstancesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLManagedInstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves a SQL Managed Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLManagedInstanceName - name of SQL Managed Instance
func (client SQLManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string) (result SQLManagedInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, SQLManagedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLManagedInstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlManagedInstanceName": autorest.Encode("path", SQLManagedInstanceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlManagedInstances/{sqlManagedInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLManagedInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLManagedInstancesClient) GetResponder(resp *http.Response) (result SQLManagedInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client SQLManagedInstancesClient) List(ctx context.Context) (result SQLManagedInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.List")
		defer func() {
			sc := -1
			if result.smilr.Response.Response != nil {
				sc = result.smilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.smilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.smilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SQLManagedInstancesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlManagedInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SQLManagedInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SQLManagedInstancesClient) ListResponder(resp *http.Response) (result SQLManagedInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SQLManagedInstancesClient) listNextResults(ctx context.Context, lastResults SQLManagedInstanceListResult) (result SQLManagedInstanceListResult, err error) {
	req, err := lastResults.sQLManagedInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLManagedInstancesClient) ListComplete(ctx context.Context) (result SQLManagedInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets all sqlManagedInstances in a resource group.
// Parameters:
// resourceGroupName - the name of the Azure resource group
func (client SQLManagedInstancesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result SQLManagedInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.smilr.Response.Response != nil {
				sc = result.smilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.smilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.smilr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SQLManagedInstancesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlManagedInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLManagedInstancesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SQLManagedInstancesClient) ListByResourceGroupResponder(resp *http.Response) (result SQLManagedInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SQLManagedInstancesClient) listByResourceGroupNextResults(ctx context.Context, lastResults SQLManagedInstanceListResult) (result SQLManagedInstanceListResult, err error) {
	req, err := lastResults.sQLManagedInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLManagedInstancesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result SQLManagedInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates a SQL Managed Instance resource
// Parameters:
// resourceGroupName - the name of the Azure resource group
// SQLManagedInstanceName - name of sqlManagedInstance
// parameters - the SQL Managed Instance.
func (client SQLManagedInstancesClient) Update(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string, parameters SQLManagedInstanceUpdate) (result SQLManagedInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLManagedInstancesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, SQLManagedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLManagedInstancesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SQLManagedInstancesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, SQLManagedInstanceName string, parameters SQLManagedInstanceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlManagedInstanceName": autorest.Encode("path", SQLManagedInstanceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlManagedInstances/{sqlManagedInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SQLManagedInstancesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SQLManagedInstancesClient) UpdateResponder(resp *http.Response) (result SQLManagedInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
