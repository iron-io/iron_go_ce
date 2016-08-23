# NewTask

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Name of Docker image to use. This is optional and can be used to override the image defined at the group level. | [default to null]
**Payload** | **string** | Payload for the task. This is what you pass into each task to make it do something. | [optional] [default to null]
**Delay** | **int32** | Number of seconds to wait before queueing the task for consumption for the first time. Must be a positive integer. Tasks with a delay start in state \&quot;delayed\&quot; and transition to \&quot;running\&quot; after delay seconds. | [optional] [default to 0]
**Timeout** | **int32** | Maximum runtime in seconds. If a consumer retrieves the task, but does not change it&#39;s status within timeout seconds, the task is considered failed, with reason timeout (Titan may allow a small grace period). The consumer should also kill the task after timeout seconds. If a consumer tries to change status after Titan has already timed out the task, the consumer will be ignored.  | [optional] [default to 60]
**Priority** | **int32** | Priority of the task. Higher has more priority. 3 levels from 0-2. Tasks at same priority are processed in FIFO order. | [default to null]
**MaxRetries** | **int32** | \&quot;Number of automatic retries this task is allowed.  A retry will be attempted if a task fails. Max 25. Automatic retries are performed by titan when a task reaches a failed state and has &#x60;max_retries&#x60; &gt; 0. A retry is performed by queueing a new task with the same image id and payload. The new task&#39;s max_retries is one less than the original. The new task&#39;s &#x60;retry_of&#x60; field is set to the original Task ID. The old task&#39;s &#x60;retry_at&#x60; field is set to the new Task&#39;s ID.  Titan will delay the new task for retries_delay seconds before queueing it. Cancelled or successful tasks are never automatically retried.\&quot;  | [optional] [default to 0]
**RetriesDelay** | **int32** | Time in seconds to wait before retrying the task. Must be a non-negative integer. | [optional] [default to 60]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


