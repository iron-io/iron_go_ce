package swagger

import (
    "time"
)

type Image struct {
    Image  string  `json:"image,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    
}
