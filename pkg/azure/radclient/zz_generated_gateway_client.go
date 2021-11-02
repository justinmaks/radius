//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GatewayClient contains the methods for the Gateway group.
// Don't use this type directly, use NewGatewayClient() instead.
type GatewayClient struct {
	ep string
	pl runtime.Pipeline
	subscriptionID string
}

// NewGatewayClient creates a new instance of GatewayClient with the specified values.
func NewGatewayClient(con *arm.Connection, subscriptionID string) *GatewayClient {
	return &GatewayClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a Gateway resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, parameters GatewayResource, options *GatewayBeginCreateOrUpdateOptions) (GatewayCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationName, gatewayName, parameters, options)
	if err != nil {
		return GatewayCreateOrUpdatePollerResponse{}, err
	}
	result := GatewayCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GatewayClient.CreateOrUpdate", "location", resp, 	client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return GatewayCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &GatewayCreateOrUpdatePoller {
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a Gateway resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, parameters GatewayResource, options *GatewayBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationName, gatewayName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	 return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewayClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, parameters GatewayResource, options *GatewayBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/Gateway/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GatewayClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a Gateway resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, options *GatewayBeginDeleteOptions) (GatewayDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationName, gatewayName, options)
	if err != nil {
		return GatewayDeletePollerResponse{}, err
	}
	result := GatewayDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("GatewayClient.Delete", "location", resp, 	client.pl, client.deleteHandleError)
	if err != nil {
		return GatewayDeletePollerResponse{}, err
	}
	result.Poller = &GatewayDeletePoller {
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a Gateway resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, options *GatewayBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationName, gatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	 return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewayClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, options *GatewayBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/Gateway/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GatewayClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a Gateway resource by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayClient) Get(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, options *GatewayGetOptions) (GatewayGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationName, gatewayName, options)
	if err != nil {
		return GatewayGetResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return GatewayGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, gatewayName string, options *GatewayGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/Gateway/{gatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewayClient) getHandleResponse(resp *http.Response) (GatewayGetResponse, error) {
	result := GatewayGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayResource); err != nil {
		return GatewayGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *GatewayClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List the Gateway resources deployed in the application.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayClient) List(ctx context.Context, resourceGroupName string, applicationName string, options *GatewayListOptions) (GatewayListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, applicationName, options)
	if err != nil {
		return GatewayListResponse{}, err
	}
	resp, err := 	client.pl.Do(req)
	if err != nil {
		return GatewayListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *GatewayClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationName string, options *GatewayListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/radiusv3/Application/{applicationName}/Gateway"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationName == "" {
		return nil, errors.New("parameter applicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationName}", url.PathEscape(applicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GatewayClient) listHandleResponse(resp *http.Response) (GatewayListResponse, error) {
	result := GatewayListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayList); err != nil {
		return GatewayListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *GatewayClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
