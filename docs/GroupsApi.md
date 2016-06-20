# \GroupsApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGet**](GroupsApi.md#GroupsGet) | **Get** /groups | Get all group names.
[**GroupsNameGet**](GroupsApi.md#GroupsNameGet) | **Get** /groups/{name} | Get information for a group.
[**GroupsNamePut**](GroupsApi.md#GroupsNamePut) | **Put** /groups/{name} | Create/update a job group.
[**GroupsPost**](GroupsApi.md#GroupsPost) | **Post** /groups | Post new group


# **GroupsGet**
> GroupsWrapper GroupsGet()

Get all group names.

Get a list of all the groups in the system.


### Parameters
This endpoint does not need any parameter.

### Return type

[**GroupsWrapper**](GroupsWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameGet**
> GroupWrapper GroupsNameGet($name)

Get information for a group.

This gives more details about a job group, such as statistics.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the group. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNamePut**
> GroupWrapper GroupsNamePut($name, $body)

Create/update a job group.

You can set group level settings here. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the group. | 
 **body** | [**GroupWrapper**](GroupWrapper.md)| Group to post. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsPost**
> GroupWrapper GroupsPost($body)

Post new group

Insert a new job group


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupWrapper**](GroupWrapper.md)| Group to post. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

