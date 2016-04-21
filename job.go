package titan

import (
    "time"
)


type Job struct {
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
    // Unique identifier representing a specific job.
    Id  string  `json:"id,omitempty"`
    // States and valid transitions.\n\n                 +---------+\n       +---------> delayed <----------------+\n                 +----+----+                |\n                      |                     |\n                      |                     |\n                 +----v----+                |\n       +---------> queued  <----------------+\n                 +----+----+                *\n                      |                     *\n                      |               retry * creates new job\n                 +----v----+                *\n                 | running |                *\n                 +--+-+-+--+                |\n          +---------|-|-|-----+-------------+\n      +---|---------+ | +-----|---------+   |\n      |   |           |       |         |   |\n+-----v---^-+      +--v-------^+     +--v---^-+\n| success   |      | cancelled |     |  error |\n+-----------+      +-----------+     +--------+\n\n* delayed - has a delay.\n* queued - Ready to be consumed when it's turn comes.\n* running - Currently consumed by a runner which will attempt to process it.\n* success - (or complete? success/error is common javascript terminology)\n* error - Something went wrong. In this case more information can be obtained\n  by inspecting the \"reason\" field.\n  - timeout\n  - killed - forcibly killed by worker due to resource restrictions or access\n    violations.\n  - bad_exit - exited with non-zero status due to program termination/crash.\n* cancelled - cancelled via API. More information in the reason field.\n  - client_request - Request was cancelled by a client.\n
    Status  string  `json:"status,omitempty"`
    // Group this job belongs to. 
    GroupName  string  `json:"group_name,omitempty"`
    
    Reason  Reason  `json:"reason,omitempty"`
    // Time when job was submitted. Always in UTC.
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    // Time when job started execution. Always in UTC.
    StartedAt  time.Time  `json:"started_at,omitempty"`
    // Time when job completed, whether it was successul or failed. Always in UTC.
    CompletedAt  time.Time  `json:"completed_at,omitempty"`
    // If this field is set, then this job is a retry of the ID in this field.
    RetryOf  string  `json:"retry_of,omitempty"`
    // If this field is set, then this job was retried by the job referenced in this field.
    RetryAt  string  `json:"retry_at,omitempty"`
}
