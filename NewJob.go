package titan

import (
)

type NewJob struct {
    Name  string  `json:"name,omitempty"`
    Image  string  `json:"image,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Delay  int32  `json:"delay,omitempty"`
    Timeout  int32  `json:"timeout,omitempty"`
    Priority  int32  `json:"priority,omitempty"`
    Retries  int32  `json:"retries,omitempty"`
    RetriesDelay  int32  `json:"retries_delay,omitempty"`
    RetryFromId  string  `json:"retry_from_id,omitempty"`
    
}
