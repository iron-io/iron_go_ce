package titan

import (
    "time"
)

type Job struct {
    Image  string  `json:"image,omitempty"`
    RetryFromId  string  `json:"retry_from_id,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    Priority  int32  `json:"priority,omitempty"`
    RetriesDelay  int32  `json:"retries_delay,omitempty"`
    Error  string  `json:"error,omitempty"`
    Timeout  int32  `json:"timeout,omitempty"`
    Retries  int32  `json:"retries,omitempty"`
    CompletedAt  time.Time  `json:"completed_at,omitempty"`
    Delay  int32  `json:"delay,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Name  string  `json:"name,omitempty"`
    StartedAt  time.Time  `json:"started_at,omitempty"`
    Id  string  `json:"id,omitempty"`
    RetryId  string  `json:"retry_id,omitempty"`
    Status  string  `json:"status,omitempty"`
    
}
