package swagger

import (
)

type JobSubmissionsWrapper struct {
    Jobs  []JobSubmissionWithImage  `json:"jobs,omitempty"`
    
}
