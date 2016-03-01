package titan

import (
    "time"
)

type Job struct {
    Image  string  `json:"image,omitempty"`
    Retries  int32  `json:"retries,omitempty"`
    CompletedAt  time.Time  `json:"completed_at,omitempty"`
    Delay  number  `json:"delay,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Name  string  `json:"name,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    StartedAt  time.Time  `json:"started_at,omitempty"`
    Id  string  `json:"id,omitempty"`
    Timeout  number  `json:"timeout,omitempty"`
    Status  string  `json:"status,omitempty"`
    
}
