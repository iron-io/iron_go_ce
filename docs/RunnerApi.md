# \RunnerApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsNameTasksIdErrorPost**](RunnerApi.md#GroupsNameTasksIdErrorPost) | **Post** /groups/{name}/tasks/{id}/error | Mark task as failed.
[**GroupsNameTasksIdStartPost**](RunnerApi.md#GroupsNameTasksIdStartPost) | **Post** /groups/{name}/tasks/{id}/start | Mark task as started, ie: status &#x3D; &#39;running&#39;
[**GroupsNameTasksIdSuccessPost**](RunnerApi.md#GroupsNameTasksIdSuccessPost) | **Post** /groups/{name}/tasks/{id}/success | Mark task as succeeded.


# **GroupsNameTasksIdErrorPost**
> TaskWrapper GroupsNameTasksIdErrorPost($name, $id, $body)

Mark task as failed.

Task is marked as failed if it was in a valid state. Task's `finished_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 
 **body** | [**Complete**](Complete.md)|  | 

### Return type

[**TaskWrapper**](TaskWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdStartPost**
> TaskWrapper GroupsNameTasksIdStartPost($name, $id, $body)

Mark task as started, ie: status = 'running'

Task status is changed to 'running' if it was in a valid state before. Task's `started_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 
 **body** | [**Start**](Start.md)|  | 

### Return type

[**TaskWrapper**](TaskWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdSuccessPost**
> TaskWrapper GroupsNameTasksIdSuccessPost($name, $id, $body)

Mark task as succeeded.

Task status is changed to succeeded if it was in a valid state before. Task's `completed_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 
 **body** | [**Complete**](Complete.md)|  | 

### Return type

[**TaskWrapper**](TaskWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

