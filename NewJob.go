package titan

import (
)

type NewJob struct {
    Name  string  `json:"name,omitempty"`
    Image  string  `json:"image,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Delay  int32  `json:"delay,omitempty"`
    Timeout  int32  `json:"timeout,omitempty"`
    Retries  int32  `json:"retries,omitempty"`
    
}
