# \TasksApi

All URIs are relative to *https://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsNameTasksGet**](TasksApi.md#GroupsNameTasksGet) | **Get** /groups/{name}/tasks | Get task list by group name.
[**GroupsNameTasksIdCancelPost**](TasksApi.md#GroupsNameTasksIdCancelPost) | **Post** /groups/{name}/tasks/{id}/cancel | Cancel a task.
[**GroupsNameTasksIdDelete**](TasksApi.md#GroupsNameTasksIdDelete) | **Delete** /groups/{name}/tasks/{id} | Delete the task.
[**GroupsNameTasksIdErrorPost**](TasksApi.md#GroupsNameTasksIdErrorPost) | **Post** /groups/{name}/tasks/{id}/error | Mark task as failed.
[**GroupsNameTasksIdGet**](TasksApi.md#GroupsNameTasksIdGet) | **Get** /groups/{name}/tasks/{id} | Gets task by id
[**GroupsNameTasksIdLogGet**](TasksApi.md#GroupsNameTasksIdLogGet) | **Get** /groups/{name}/tasks/{id}/log | Get the log of a completed task.
[**GroupsNameTasksIdLogPost**](TasksApi.md#GroupsNameTasksIdLogPost) | **Post** /groups/{name}/tasks/{id}/log | Send in a log for storage.
[**GroupsNameTasksIdRetryPost**](TasksApi.md#GroupsNameTasksIdRetryPost) | **Post** /groups/{name}/tasks/{id}/retry | Retry a task.
[**GroupsNameTasksIdStartPost**](TasksApi.md#GroupsNameTasksIdStartPost) | **Post** /groups/{name}/tasks/{id}/start | Mark task as started, ie: status &#x3D; &#39;running&#39;
[**GroupsNameTasksIdSuccessPost**](TasksApi.md#GroupsNameTasksIdSuccessPost) | **Post** /groups/{name}/tasks/{id}/success | Mark task as succeeded.
[**GroupsNameTasksIdTouchPost**](TasksApi.md#GroupsNameTasksIdTouchPost) | **Post** /groups/{name}/tasks/{id}/touch | Extend task timeout.
[**GroupsNameTasksPost**](TasksApi.md#GroupsNameTasksPost) | **Post** /groups/{name}/tasks | Enqueue task
[**TasksGet**](TasksApi.md#TasksGet) | **Get** /tasks | Get next task.


# **GroupsNameTasksGet**
> TasksWrapper GroupsNameTasksGet($name, $createdAfter, $n, $cursor)

Get task list by group name.

This will list tasks for a particular group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **createdAfter** | **time.Time**| Will return tasks created after this time. In RFC3339 format. | [optional] 
 **n** | **int32**| Number of tasks to return per page. Default is 50. Max is 1000. | [optional] 
 **cursor** | **string**| Pass this in from a previous query to paginate results. | [optional] 

### Return type

[**TasksWrapper**](TasksWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdCancelPost**
> TaskWrapper GroupsNameTasksIdCancelPost($name, $id)

Cancel a task.

Cancels a task in delayed, queued or running status. The worker may continue to run a running task. reason is set to `client_request`. The task's completed_at field is set to the current time on the taskserver.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 

### Return type

[**TaskWrapper**](TaskWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdDelete**
> GroupsNameTasksIdDelete($name, $id)

Delete the task.

Delete only succeeds if task status is one of `succeeded | failed | cancelled`. Cancel a task if it is another state and needs to be deleted.  All information about the task, including the log, is irretrievably lost when this is invoked. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **GroupsNameTasksIdGet**
> TaskWrapper GroupsNameTasksIdGet($name, $id)

Gets task by id

Gets a task by id.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| task id | 

### Return type

[**TaskWrapper**](TaskWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdLogGet**
> string GroupsNameTasksIdLogGet($name, $id)

Get the log of a completed task.

Retrieves the log from log storage.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdLogPost**
> TaskWrapper GroupsNameTasksIdLogPost($name, $id, $log)

Send in a log for storage.

Logs are sent after a task completes since they may be very large and the runner can process the next task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 
 **log** | ***os.File**| Output log for the task. Content-Type must be \&quot;text/plain; charset&#x3D;utf-8\&quot;. | 

### Return type

[**TaskWrapper**](TaskWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksIdRetryPost**
> TaskWrapper GroupsNameTasksIdRetryPost($name, $id)

Retry a task.

\"The /retry endpoint can be used to force a retry of tasks with status succeeded, cancelled or failed. The retried task has the same attributes. max_retries is not modified.\" 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 

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

# **GroupsNameTasksIdTouchPost**
> GroupsNameTasksIdTouchPost($name, $id)

Extend task timeout.

Consumers can sometimes take a while to run the task after accepting it.  An example is when the runner does not have the docker image locally, it can spend a significant time downloading the image. If the timeout is small, the task may never get to run, or run but not be accepted by Titan. Consumers can touch the task before it times out. Titan will reset the timeout, giving the consumer another timeout seconds to run the task. Touch is only valid while the task is in a running state. If touch fails, the runner may stop running the task. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| Name of group for this set of tasks. | 
 **id** | **string**| Task id | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNameTasksPost**
> TasksWrapper GroupsNameTasksPost($name, $body)

Enqueue task

Enqueues task(s). If any of the tasks is invalid, none of the tasks are enqueued. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the group. | 
 **body** | [**NewTasksWrapper**](NewTasksWrapper.md)| Array of tasks to post. | 

### Return type

[**TasksWrapper**](TasksWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TasksGet**
> TasksWrapper TasksGet($n)

Get next task.

Gets the next task in the queue, ready for processing. Titan may return <=n tasks. Consumers should start processing tasks in order. Each returned task is set to `status` \"running\" and `started_at` is set to the current time. No other consumer can retrieve this task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32**| Number of tasks to return. | [optional] [default to 1]

### Return type

[**TasksWrapper**](TasksWrapper.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

