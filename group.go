package titan

import (
    "time"
)


type Group struct {
    // Name of this group. Must be different than the image name. Can ony contain alphanumeric, -, and _.
    Name  string  `json:"name,omitempty"`
    // Time when image first used/created.
    CreatedAt  time.Time  `json:"created_at,omitempty"`
}
