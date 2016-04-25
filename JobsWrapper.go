package titan

import (
)

type JobsWrapper struct {
    Jobs  []Job  `json:"jobs,omitempty"`
    Error_  ErrorBody  `json:"error,omitempty"`
    
}
