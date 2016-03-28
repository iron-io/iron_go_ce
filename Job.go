package titan

import (
    "time"
)

type Job struct {
    Name  string  `json:"name,omitempty"`
    Image  string  `json:"image,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Delay  int32  `json:"delay,omitempty"`
    Timeout  int32  `json:"timeout,omitempty"`
    Priority  int32  `json:"priority,omitempty"`
    Retries  int32  `json:"retries,omitempty"`
    RetriesDelay  int32  `json:"retries_delay,omitempty"`
    RetryFromId  string  `json:"retry_from_id,omitempty"`
    Id  string  `json:"id,omitempty"`
    Status  string  `json:"status,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    Error_  string  `json:"error,omitempty"`
    StartedAt  time.Time  `json:"started_at,omitempty"`
    CompletedAt  time.Time  `json:"completed_at,omitempty"`
    RetryId  string  `json:"retry_id,omitempty"`
    
}
