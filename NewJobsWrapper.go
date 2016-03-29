package titan

import (
)

type NewJobsWrapper struct {
    Jobs  []NewJobWithImage  `json:"jobs,omitempty"`
    
}
