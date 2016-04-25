package titan

import (
    "time"
)

type Complete struct {
    CompletedAt  time.Time  `json:"completed_at,omitempty"`
    Reason  string  `json:"reason,omitempty"`
    Error_  string  `json:"error,omitempty"`
    
}
