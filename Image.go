package titan

import (
    "time"
)

type Image struct {
    Name  string  `json:"name,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    
}
