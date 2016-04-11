package titan

import (
    "time"
)

type Start struct {
    StartedAt  time.Time  `json:"started_at,omitempty"`
}
