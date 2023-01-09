//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220315privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DaprInvokeHTTPRouteClient contains the methods for the DaprInvokeHTTPRoute group.
// Don't use this type directly, use NewDaprInvokeHTTPRouteClient() instead.
type DaprInvokeHTTPRouteClient struct {
	host string
	rootScope string
	pl runtime.Pipeline
}

// NewDaprInvokeHTTPRouteClient creates a new instance of DaprInvokeHTTPRouteClient with the specified values.
// rootScope - The scope in which the resource is present. For Azure resource this would be /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDaprInvokeHTTPRouteClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprInvokeHTTPRouteClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DaprInvokeHTTPRouteClient{
		rootScope: rootScope,
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a DaprInvokeHttpRouteResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprInvokeHTTPRouteName - The name of the DaprInvokeHttpRoute link resource
// resource - Resource create parameters.
// options - DaprInvokeHTTPRouteClientCreateOrUpdateOptions contains the optional parameters for the DaprInvokeHTTPRouteClient.CreateOrUpdate
// method.
func (client *DaprInvokeHTTPRouteClient) CreateOrUpdate(ctx context.Context, daprInvokeHTTPRouteName string, resource DaprInvokeHTTPRouteResource, options *DaprInvokeHTTPRouteClientCreateOrUpdateOptions) (DaprInvokeHTTPRouteClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, daprInvokeHTTPRouteName, resource, options)
	if err != nil {
		return DaprInvokeHTTPRouteClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRouteClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DaprInvokeHTTPRouteClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprInvokeHTTPRouteClient) createOrUpdateCreateRequest(ctx context.Context, daprInvokeHTTPRouteName string, resource DaprInvokeHTTPRouteResource, options *DaprInvokeHTTPRouteClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprInvokeHTTPRouteClient) createOrUpdateHandleResponse(resp *http.Response) (DaprInvokeHTTPRouteClientCreateOrUpdateResponse, error) {
	result := DaprInvokeHTTPRouteClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return DaprInvokeHTTPRouteClientCreateOrUpdateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResource); err != nil {
		return DaprInvokeHTTPRouteClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing DaprInvokeHttpRouteResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprInvokeHTTPRouteName - The name of the DaprInvokeHttpRoute link resource
// options - DaprInvokeHTTPRouteClientDeleteOptions contains the optional parameters for the DaprInvokeHTTPRouteClient.Delete
// method.
func (client *DaprInvokeHTTPRouteClient) Delete(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRouteClientDeleteOptions) (DaprInvokeHTTPRouteClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, daprInvokeHTTPRouteName, options)
	if err != nil {
		return DaprInvokeHTTPRouteClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRouteClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return DaprInvokeHTTPRouteClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *DaprInvokeHTTPRouteClient) deleteCreateRequest(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRouteClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *DaprInvokeHTTPRouteClient) deleteHandleResponse(resp *http.Response) (DaprInvokeHTTPRouteClientDeleteResponse, error) {
	result := DaprInvokeHTTPRouteClientDeleteResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return DaprInvokeHTTPRouteClientDeleteResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	return result, nil
}

// Get - Retrieves information about a DaprInvokeHttpRouteResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-15-privatepreview
// daprInvokeHTTPRouteName - The name of the DaprInvokeHttpRoute link resource
// options - DaprInvokeHTTPRouteClientGetOptions contains the optional parameters for the DaprInvokeHTTPRouteClient.Get method.
func (client *DaprInvokeHTTPRouteClient) Get(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRouteClientGetOptions) (DaprInvokeHTTPRouteClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, daprInvokeHTTPRouteName, options)
	if err != nil {
		return DaprInvokeHTTPRouteClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DaprInvokeHTTPRouteClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprInvokeHTTPRouteClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprInvokeHTTPRouteClient) getCreateRequest(ctx context.Context, daprInvokeHTTPRouteName string, options *DaprInvokeHTTPRouteClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/daprInvokeHttpRoutes/{daprInvokeHttpRouteName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if daprInvokeHTTPRouteName == "" {
		return nil, errors.New("parameter daprInvokeHTTPRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{daprInvokeHttpRouteName}", url.PathEscape(daprInvokeHTTPRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprInvokeHTTPRouteClient) getHandleResponse(resp *http.Response) (DaprInvokeHTTPRouteClientGetResponse, error) {
	result := DaprInvokeHTTPRouteClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResource); err != nil {
		return DaprInvokeHTTPRouteClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRootScopePager - Lists information about all DaprInvokeHttpRouteResources in the given root scope
// Generated from API version 2022-03-15-privatepreview
// options - DaprInvokeHTTPRouteClientListByRootScopeOptions contains the optional parameters for the DaprInvokeHTTPRouteClient.ListByRootScope
// method.
func (client *DaprInvokeHTTPRouteClient) NewListByRootScopePager(options *DaprInvokeHTTPRouteClientListByRootScopeOptions) (*runtime.Pager[DaprInvokeHTTPRouteClientListByRootScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[DaprInvokeHTTPRouteClientListByRootScopeResponse]{
		More: func(page DaprInvokeHTTPRouteClientListByRootScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprInvokeHTTPRouteClientListByRootScopeResponse) (DaprInvokeHTTPRouteClientListByRootScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRootScopeCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DaprInvokeHTTPRouteClientListByRootScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DaprInvokeHTTPRouteClientListByRootScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DaprInvokeHTTPRouteClientListByRootScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRootScopeHandleResponse(resp)
		},
	})
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *DaprInvokeHTTPRouteClient) listByRootScopeCreateRequest(ctx context.Context, options *DaprInvokeHTTPRouteClientListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Link/daprInvokeHttpRoutes"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *DaprInvokeHTTPRouteClient) listByRootScopeHandleResponse(resp *http.Response) (DaprInvokeHTTPRouteClientListByRootScopeResponse, error) {
	result := DaprInvokeHTTPRouteClientListByRootScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprInvokeHTTPRouteResourceListResult); err != nil {
		return DaprInvokeHTTPRouteClientListByRootScopeResponse{}, err
	}
	return result, nil
}
