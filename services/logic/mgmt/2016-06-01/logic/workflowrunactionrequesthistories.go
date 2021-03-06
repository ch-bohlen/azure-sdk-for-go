package logic

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

// WorkflowRunActionRequestHistoriesClient is the REST API for Azure Logic Apps.
type WorkflowRunActionRequestHistoriesClient struct {
	BaseClient
}

// NewWorkflowRunActionRequestHistoriesClient creates an instance of the WorkflowRunActionRequestHistoriesClient
// client.
func NewWorkflowRunActionRequestHistoriesClient(subscriptionID string) WorkflowRunActionRequestHistoriesClient {
	return NewWorkflowRunActionRequestHistoriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunActionRequestHistoriesClientWithBaseURI creates an instance of the
// WorkflowRunActionRequestHistoriesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkflowRunActionRequestHistoriesClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunActionRequestHistoriesClient {
	return WorkflowRunActionRequestHistoriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow run request history.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// runName - the workflow run name.
// actionName - the workflow action name.
// requestHistoryName - the request history name.
func (client WorkflowRunActionRequestHistoriesClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, requestHistoryName string) (result RequestHistory, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRequestHistoriesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, runName, actionName, requestHistoryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowRunActionRequestHistoriesClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, requestHistoryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":         autorest.Encode("path", actionName),
		"requestHistoryName": autorest.Encode("path", requestHistoryName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"runName":            autorest.Encode("path", runName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"workflowName":       autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories/{requestHistoryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionRequestHistoriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRequestHistoriesClient) GetResponder(resp *http.Response) (result RequestHistory, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list a workflow run request history.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// runName - the workflow run name.
// actionName - the workflow action name.
func (client WorkflowRunActionRequestHistoriesClient) List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result RequestHistoryListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRequestHistoriesClient.List")
		defer func() {
			sc := -1
			if result.rhlr.Response.Response != nil {
				sc = result.rhlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workflowName, runName, actionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rhlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "List", resp, "Failure sending request")
		return
	}

	result.rhlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowRunActionRequestHistoriesClient) ListPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":        autorest.Encode("path", actionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionRequestHistoriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRequestHistoriesClient) ListResponder(resp *http.Response) (result RequestHistoryListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkflowRunActionRequestHistoriesClient) listNextResults(ctx context.Context, lastResults RequestHistoryListResult) (result RequestHistoryListResult, err error) {
	req, err := lastResults.requestHistoryListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRequestHistoriesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowRunActionRequestHistoriesClient) ListComplete(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result RequestHistoryListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunActionRequestHistoriesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workflowName, runName, actionName)
	return
}
