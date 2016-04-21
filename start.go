package titan

import (
    "time"
)


type Start struct {
    // Time when job started execution. Always in UTC.
    StartedAt  time.Time  `json:"started_at,omitempty"`
}
