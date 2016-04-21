# Job

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Name of image to use. | [default to null]
**Payload** | **string** | Payload for the job. This is what you pass into each job to make it do something. | [optional] [default to null]
**Delay** | **int32** | Number of seconds to wait before queueing the job for consumption for the first time. Must be a positive integer. Jobs with a delay start in state \&quot;delayed\&quot; and transition to \&quot;running\&quot; after delay seconds. | [optional] [default to 0]
**Timeout** | **int32** | Maximum runtime in seconds. If a consumer retrieves the\njob, but does not change it&#39;s status within timeout seconds, the job\nis considered failed, with reason timeout (Titan may allow a small\ngrace period). The consumer should also kill the job after timeout\nseconds. If a consumer tries to change status after Titan has already\ntimed out the job, the consumer will be ignored.\n | [optional] [default to 60]
**Priority** | **int32** | Priority of the job. Higher has more priority. 3 levels from 0-2. Jobs at same priority are processed in FIFO order. | [default to null]
**MaxRetries** | **int32** | \&quot;Number of automatic retries this job is allowed. A retry will be attempted if a task fails. Max 25. Automatic retries are performed by titan when a task reaches a failed state and has &#x60;max_retries&#x60; &gt; 0. A retry is performed by queueing a new job with the same image id and payload. The new job&#39;s max_retries is one less than the original. The new job&#39;s &#x60;retry_of&#x60; field is set to the original Job ID.  Titan will delay the new job for retries_delay seconds before queueing it. Cancelled or successful tasks are never automatically retried.\&quot;\n | [optional] [default to 0]
**RetriesDelay** | **int32** | Time in seconds to wait before retrying the job. Must be a non-negative integer. | [optional] [default to 60]
**Id** | **string** | Unique identifier representing a specific job. | [optional] [default to null]
**Status** | **string** | States and valid transitions.\n\n                 +---------+\n       +---------&gt; delayed &lt;----------------+\n                 +----+----+                |\n                      |                     |\n                      |                     |\n                 +----v----+                |\n       +---------&gt; queued  &lt;----------------+\n                 +----+----+                *\n                      |                     *\n                      |               retry * creates new job\n                 +----v----+                *\n                 | running |                *\n                 +--+-+-+--+                |\n          +---------|-|-|-----+-------------+\n      +---|---------+ | +-----|---------+   |\n      |   |           |       |         |   |\n+-----v---^-+      +--v-------^+     +--v---^-+\n| success   |      | cancelled |     |  error |\n+-----------+      +-----------+     +--------+\n\n* delayed - has a delay.\n* queued - Ready to be consumed when it&#39;s turn comes.\n* running - Currently consumed by a runner which will attempt to process it.\n* success - (or complete? success/error is common javascript terminology)\n* error - Something went wrong. In this case more information can be obtained\n  by inspecting the \&quot;reason\&quot; field.\n  - timeout\n  - killed - forcibly killed by worker due to resource restrictions or access\n    violations.\n  - bad_exit - exited with non-zero status due to program termination/crash.\n* cancelled - cancelled via API. More information in the reason field.\n  - client_request - Request was cancelled by a client.\n | [optional] [default to null]
**GroupName** | **string** | Group this job belongs to.  | [optional] [default to null]
**Reason** | [**Reason**](Reason.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Time when job was submitted. Always in UTC. | [optional] [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | Time when job started execution. Always in UTC. | [optional] [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | Time when job completed, whether it was successul or failed. Always in UTC. | [optional] [default to null]
**RetryOf** | **string** | If this field is set, then this job is a retry of the ID in this field. | [optional] [default to null]
**RetryAt** | **string** | If this field is set, then this job was retried by the job referenced in this field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


