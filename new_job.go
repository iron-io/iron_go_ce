package titan

import (
)


type NewJob struct {
    // Name of image to use.
    Image  string  `json:"image,omitempty"`
    // Payload for the job. This is what you pass into each job to make it do something.
    Payload  string  `json:"payload,omitempty"`
    // Number of seconds to wait before queueing the job for consumption for the first time. Must be a positive integer. Jobs with a delay start in state \"delayed\" and transition to \"running\" after delay seconds.
    Delay  int32  `json:"delay,omitempty"`
    // Maximum runtime in seconds. If a consumer retrieves the\njob, but does not change it's status within timeout seconds, the job\nis considered failed, with reason timeout (Titan may allow a small\ngrace period). The consumer should also kill the job after timeout\nseconds. If a consumer tries to change status after Titan has already\ntimed out the job, the consumer will be ignored.\n
    Timeout  int32  `json:"timeout,omitempty"`
    // Priority of the job. Higher has more priority. 3 levels from 0-2. Jobs at same priority are processed in FIFO order.
    Priority  int32  `json:"priority,omitempty"`
    // \"Number of automatic retries this job is allowed. A retry will be attempted if a task fails. Max 25. Automatic retries are performed by titan when a task reaches a failed state and has `max_retries` > 0. A retry is performed by queueing a new job with the same image id and payload. The new job's max_retries is one less than the original. The new job's `retry_of` field is set to the original Job ID.  Titan will delay the new job for retries_delay seconds before queueing it. Cancelled or successful tasks are never automatically retried.\"\n
    MaxRetries  int32  `json:"max_retries,omitempty"`
    // Time in seconds to wait before retrying the job. Must be a non-negative integer.
    RetriesDelay  int32  `json:"retries_delay,omitempty"`
}
