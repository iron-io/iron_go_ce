# \RunnerApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGroupNameJobsIdErrorPost**](RunnerApi.md#GroupsGroupNameJobsIdErrorPost) | **Post** /groups/{group_name}/jobs/{id}/error | Mark job as failed.
[**GroupsGroupNameJobsIdStartPost**](RunnerApi.md#GroupsGroupNameJobsIdStartPost) | **Post** /groups/{group_name}/jobs/{id}/start | Mark job as started, ie: status &#x3D; &#39;running&#39;
[**GroupsGroupNameJobsIdSuccessPost**](RunnerApi.md#GroupsGroupNameJobsIdSuccessPost) | **Post** /groups/{group_name}/jobs/{id}/success | Mark job as succeeded.


# **GroupsGroupNameJobsIdErrorPost**
> JobWrapper GroupsGroupNameJobsIdErrorPost($groupName, $id, $reason)

Mark job as failed.

Job is marked as failed if it was in a valid state. Job's `finished_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 
 **reason** | **string**| Reason for job failure. | 

### Return type

[**JobWrapper**](JobWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupNameJobsIdStartPost**
> JobWrapper GroupsGroupNameJobsIdStartPost($groupName, $id, $body)

Mark job as started, ie: status = 'running'

Job status is changed to 'running' if it was in a valid state before. Job's `started_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
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

# **GroupsGroupNameJobsIdSuccessPost**
> JobWrapper GroupsGroupNameJobsIdSuccessPost($groupName, $id)

Mark job as succeeded.

Job status is changed to succeeded if it was in a valid state before. Job's `completed_at` time is initialized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 

### Return type

[**JobWrapper**](JobWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

