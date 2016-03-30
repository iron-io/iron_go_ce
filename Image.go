package swagger

import (
    "time"
)

type Image struct {
    Name  string  `json:"name,omitempty"`
    Image  string  `json:"image,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    
}
