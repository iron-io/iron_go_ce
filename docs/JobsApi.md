# \JobsApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGroupNameJobsGet**](JobsApi.md#GroupsGroupNameJobsGet) | **Get** /groups/{group_name}/jobs | Get job list by group name.
[**GroupsGroupNameJobsIdCancelPost**](JobsApi.md#GroupsGroupNameJobsIdCancelPost) | **Post** /groups/{group_name}/jobs/{id}/cancel | Cancel a job.
[**GroupsGroupNameJobsIdDelete**](JobsApi.md#GroupsGroupNameJobsIdDelete) | **Delete** /groups/{group_name}/jobs/{id} | Delete the job.
[**GroupsGroupNameJobsIdErrorPost**](JobsApi.md#GroupsGroupNameJobsIdErrorPost) | **Post** /groups/{group_name}/jobs/{id}/error | Mark job as failed.
[**GroupsGroupNameJobsIdGet**](JobsApi.md#GroupsGroupNameJobsIdGet) | **Get** /groups/{group_name}/jobs/{id} | Gets job by id
[**GroupsGroupNameJobsIdLogGet**](JobsApi.md#GroupsGroupNameJobsIdLogGet) | **Get** /groups/{group_name}/jobs/{id}/log | Get the log of a completed job.
[**GroupsGroupNameJobsIdLogPost**](JobsApi.md#GroupsGroupNameJobsIdLogPost) | **Post** /groups/{group_name}/jobs/{id}/log | Send in a log for storage.
[**GroupsGroupNameJobsIdRetryPost**](JobsApi.md#GroupsGroupNameJobsIdRetryPost) | **Post** /groups/{group_name}/jobs/{id}/retry | Retry a job.
[**GroupsGroupNameJobsIdStartPost**](JobsApi.md#GroupsGroupNameJobsIdStartPost) | **Post** /groups/{group_name}/jobs/{id}/start | Mark job as started, ie: status &#x3D; &#39;running&#39;
[**GroupsGroupNameJobsIdSuccessPost**](JobsApi.md#GroupsGroupNameJobsIdSuccessPost) | **Post** /groups/{group_name}/jobs/{id}/success | Mark job as succeeded.
[**GroupsGroupNameJobsIdTouchPost**](JobsApi.md#GroupsGroupNameJobsIdTouchPost) | **Post** /groups/{group_name}/jobs/{id}/touch | Extend job timeout.
[**GroupsGroupNameJobsPost**](JobsApi.md#GroupsGroupNameJobsPost) | **Post** /groups/{group_name}/jobs | Enqueue Job
[**JobsGet**](JobsApi.md#JobsGet) | **Get** /jobs | Get next job.


# **GroupsGroupNameJobsGet**
> JobsWrapper GroupsGroupNameJobsGet($groupName, $createdAfter, $n)

Get job list by group name.

This will list jobs for a particular group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **createdAfter** | **time.Time**| Will return jobs created after this time. In RFC3339 format. | [optional] 
 **n** | **int32**| Number of jobs to return. | [optional] 

### Return type

[**JobsWrapper**](JobsWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupNameJobsIdCancelPost**
> JobWrapper GroupsGroupNameJobsIdCancelPost($groupName, $id)

Cancel a job.

Cancels a job in delayed, queued or running status. The worker may continue to run a running job. reason is set to `client_request`.


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

# **GroupsGroupNameJobsIdDelete**
> GroupsGroupNameJobsIdDelete($groupName, $id)

Delete the job.

Delete only succeeds if job status is one of `succeeded\n| failed | cancelled`. Cancel a job if it is another state and needs to\nbe deleted.  All information about the job, including the log, is\nirretrievably lost when this is invoked.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **GroupsGroupNameJobsIdGet**
> JobWrapper GroupsGroupNameJobsIdGet($groupName, $id)

Gets job by id

Gets a job by id.


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

# **GroupsGroupNameJobsIdLogGet**
> string GroupsGroupNameJobsIdLogGet($groupName, $id)

Get the log of a completed job.

Retrieves the log from log storage.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupNameJobsIdLogPost**
> JobWrapper GroupsGroupNameJobsIdLogPost($groupName, $id, $log)

Send in a log for storage.

Logs are sent after a job completes since they may be very large and the runner can process the next job.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 
 **log** | ***os.File**| Output log for the job. Content-Type must be \&quot;text/plain; charset&#x3D;utf-8\&quot;. | 

### Return type

[**JobWrapper**](JobWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupNameJobsIdRetryPost**
> JobWrapper GroupsGroupNameJobsIdRetryPost($groupName, $id)

Retry a job.

\"The /retry endpoint can be used to force a retry of jobs\nwith status succeeded or cancelled. It can also be used to retry jobs\nthat in the failed state, but whose max_retries field is 0. The retried\njob will continue to have max_retries = 0.\"\n


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

# **GroupsGroupNameJobsIdTouchPost**
> GroupsGroupNameJobsIdTouchPost($groupName, $id)

Extend job timeout.

Consumers can sometimes take a while to run the task after accepting it.  An example is when the runner does not have the docker image locally, it can spend a significant time downloading the image.\nIf the timeout is small, the job may never get to run, or run but not be accepted by Titan. Consumers can touch the job before it times out. Titan will reset the timeout, giving the consumer another timeout seconds to run the job.\nTouch is only valid while the job is in a running state. If touch fails, the runner may stop running the job.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| Name of group for this set of jobs. | 
 **id** | **string**| Job id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGroupNameJobsPost**
> JobsWrapper GroupsGroupNameJobsPost($groupName, $body)

Enqueue Job

Enqueues job(s). If any of the jobs is invalid, none of the jobs are enqueued.\n


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string**| name of the group. | 
 **body** | [**NewJobsWrapper**](NewJobsWrapper.md)| Array of jobs to post. | 

### Return type

[**JobsWrapper**](JobsWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsGet**
> JobsWrapper JobsGet($n)

Get next job.

Gets the next job in the queue, ready for processing. Titan may return <=n jobs. Consumers should start processing jobs in order. Each returned job is set to `status` \"running\" and `started_at` is set to the current time. No other consumer can retrieve this job.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32**| Number of jobs to return. | [optional] [default to 1]

### Return type

[**JobsWrapper**](JobsWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

