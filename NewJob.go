package titan

import (
)

type NewJob struct {
    Image  string  `json:"image,omitempty"`
Username  string  `json:"username,omitempty"`
Password  string  `json:"password,omitempty"`
Payload  string  `json:"payload,omitempty"`
Delay  int32  `json:"delay,omitempty"`
Timeout  int32  `json:"timeout,omitempty"`
Priority  int32  `json:"priority,omitempty"`
MaxRetries  int32  `json:"max_retries,omitempty"`
RetriesDelay  int32  `json:"retries_delay,omitempty"`
}
