package sql

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
	"net/http"
)

// JobStepExecutionsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type JobStepExecutionsClient struct {
	BaseClient
}

// NewJobStepExecutionsClient creates an instance of the JobStepExecutionsClient client.
func NewJobStepExecutionsClient(subscriptionID string) JobStepExecutionsClient {
	return NewJobStepExecutionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobStepExecutionsClientWithBaseURI creates an instance of the JobStepExecutionsClient client.
func NewJobStepExecutionsClientWithBaseURI(baseURI string, subscriptionID string) JobStepExecutionsClient {
	return JobStepExecutionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a step execution of a job execution.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
// jobExecutionID - the unique id of the job execution
// stepName - the name of the step.
func (client JobStepExecutionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID, stepName string) (result JobExecution, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobStepExecutionsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobExecutionId":    autorest.Encode("path", jobExecutionID),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/steps/{stepName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepExecutionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobStepExecutionsClient) GetResponder(resp *http.Response) (result JobExecution, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByJobExecution lists the step executions of a job execution.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
// jobExecutionID - the id of the job execution
// createTimeMin - if specified, only job executions created at or after the specified time are included.
// createTimeMax - if specified, only job executions created before the specified time are included.
// endTimeMin - if specified, only job executions completed at or after the specified time are included.
// endTimeMax - if specified, only job executions completed before the specified time are included.
// isActive - if specified, only active or only completed job executions are included.
// skip - the number of elements in the collection to skip.
// top - the number of elements to return from the collection.
func (client JobStepExecutionsClient) ListByJobExecution(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (result JobExecutionListResultPage, err error) {
	result.fn = client.listByJobExecutionNextResults
	req, err := client.ListByJobExecutionPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, createTimeMin, createTimeMax, endTimeMin, endTimeMax, isActive, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "ListByJobExecution", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByJobExecutionSender(req)
	if err != nil {
		result.jelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "ListByJobExecution", resp, "Failure sending request")
		return
	}

	result.jelr, err = client.ListByJobExecutionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "ListByJobExecution", resp, "Failure responding to request")
	}

	return
}

// ListByJobExecutionPreparer prepares the ListByJobExecution request.
func (client JobStepExecutionsClient) ListByJobExecutionPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobExecutionId":    autorest.Encode("path", jobExecutionID),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if createTimeMin != nil {
		queryParameters["createTimeMin"] = autorest.Encode("query", *createTimeMin)
	}
	if createTimeMax != nil {
		queryParameters["createTimeMax"] = autorest.Encode("query", *createTimeMax)
	}
	if endTimeMin != nil {
		queryParameters["endTimeMin"] = autorest.Encode("query", *endTimeMin)
	}
	if endTimeMax != nil {
		queryParameters["endTimeMax"] = autorest.Encode("query", *endTimeMax)
	}
	if isActive != nil {
		queryParameters["isActive"] = autorest.Encode("query", *isActive)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/steps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByJobExecutionSender sends the ListByJobExecution request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepExecutionsClient) ListByJobExecutionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByJobExecutionResponder handles the response to the ListByJobExecution request. The method always
// closes the http.Response Body.
func (client JobStepExecutionsClient) ListByJobExecutionResponder(resp *http.Response) (result JobExecutionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByJobExecutionNextResults retrieves the next set of results, if any.
func (client JobStepExecutionsClient) listByJobExecutionNextResults(lastResults JobExecutionListResult) (result JobExecutionListResult, err error) {
	req, err := lastResults.jobExecutionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "listByJobExecutionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByJobExecutionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "listByJobExecutionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByJobExecutionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepExecutionsClient", "listByJobExecutionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByJobExecutionComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobStepExecutionsClient) ListByJobExecutionComplete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (result JobExecutionListResultIterator, err error) {
	result.page, err = client.ListByJobExecution(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, createTimeMin, createTimeMax, endTimeMin, endTimeMax, isActive, skip, top)
	return
}
