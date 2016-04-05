package titan

import (
    "time"
)

type Group struct {
    Name  string  `json:"name,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    
}
