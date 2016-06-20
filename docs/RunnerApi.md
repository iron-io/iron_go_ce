# \RunnerApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsNameJobsIdErrorPost**](RunnerApi.md#GroupsNameJobsIdErrorPost) | **Post** /groups/{name}/jobs/{id}/error | Mark job as failed.
[**GroupsNameJobsIdStartPost**](RunnerApi.md#GroupsNameJobsIdStartPost) | **Post** /groups/{name}/jobs/{id}/start | Mark job as started, ie: status &#x3D; &#39;running&#39;
[**GroupsNameJobsIdSuccessPost**](RunnerApi.md#GroupsNameJobsIdSuccessPost) | **Post** /groups/{name}/jobs/{id}/success | Mark job as succeeded.


# **GroupsNameJobsIdErrorPost**
> JobWrapper GroupsNameJobsIdErrorPost($name, $id, $body)

Mark job as failed.

Job is marked as failed if it was in a valid state. Job's `finished_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 
 **body** | [**Complete**](Complete.md)|  | 

### Return type

[**JobWrapper**](JobWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameJobsIdStartPost**
> JobWrapper GroupsNameJobsIdStartPost($name, $id, $body)

Mark job as started, ie: status = 'running'

Job status is changed to 'running' if it was in a valid state before. Job's `started_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 
 **body** | [**Start**](Start.md)|  | 

### Return type

[**JobWrapper**](JobWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameJobsIdSuccessPost**
> JobWrapper GroupsNameJobsIdSuccessPost($name, $id, $body)

Mark job as succeeded.

Job status is changed to succeeded if it was in a valid state before. Job's `completed_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 
 **body** | [**Complete**](Complete.md)|  | 

### Return type

[**JobWrapper**](JobWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

