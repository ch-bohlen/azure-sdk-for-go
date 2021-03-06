package security

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IoTSecuritySolutionsAnalyticsRecommendationsClient is the API spec for Microsoft.Security (Azure Security Center)
// resource provider
type IoTSecuritySolutionsAnalyticsRecommendationsClient struct {
	BaseClient
}

// NewIoTSecuritySolutionsAnalyticsRecommendationsClient creates an instance of the
// IoTSecuritySolutionsAnalyticsRecommendationsClient client.
func NewIoTSecuritySolutionsAnalyticsRecommendationsClient(subscriptionID string, ascLocation string) IoTSecuritySolutionsAnalyticsRecommendationsClient {
	return NewIoTSecuritySolutionsAnalyticsRecommendationsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewIoTSecuritySolutionsAnalyticsRecommendationsClientWithBaseURI creates an instance of the
// IoTSecuritySolutionsAnalyticsRecommendationsClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIoTSecuritySolutionsAnalyticsRecommendationsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) IoTSecuritySolutionsAnalyticsRecommendationsClient {
	return IoTSecuritySolutionsAnalyticsRecommendationsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// List security Analytics of a security solution
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// solutionName - the solution manager name
// top - the number of results to retrieve.
func (client IoTSecuritySolutionsAnalyticsRecommendationsClient) List(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result IoTSecurityAggregatedRecommendationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsAnalyticsRecommendationsClient.List")
		defer func() {
			sc := -1
			if result.itsarl.Response.Response != nil {
				sc = result.itsarl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, solutionName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itsarl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "List", resp, "Failure sending request")
		return
	}

	result.itsarl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IoTSecuritySolutionsAnalyticsRecommendationsClient) ListPreparer(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"solutionName":      autorest.Encode("path", solutionName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IoTSecuritySolutionsAnalyticsRecommendationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IoTSecuritySolutionsAnalyticsRecommendationsClient) ListResponder(resp *http.Response) (result IoTSecurityAggregatedRecommendationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client IoTSecuritySolutionsAnalyticsRecommendationsClient) listNextResults(ctx context.Context, lastResults IoTSecurityAggregatedRecommendationList) (result IoTSecurityAggregatedRecommendationList, err error) {
	req, err := lastResults.ioTSecurityAggregatedRecommendationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.IoTSecuritySolutionsAnalyticsRecommendationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client IoTSecuritySolutionsAnalyticsRecommendationsClient) ListComplete(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result IoTSecurityAggregatedRecommendationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IoTSecuritySolutionsAnalyticsRecommendationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, solutionName, top)
	return
}
