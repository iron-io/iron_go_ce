package titan

import (
)

type NewJob struct {
    Name  string  `json:"name,omitempty"`
    Image  string  `json:"image,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Delay  number  `json:"delay,omitempty"`
    Timeout  number  `json:"timeout,omitempty"`
    Retries  int32  `json:"retries,omitempty"`
    
}
