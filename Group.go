package titan

import (
    "time"
)

type Group struct {
    Name  string  `json:"name,omitempty"`
Image  string  `json:"image,omitempty"`
Username  string  `json:"username,omitempty"`
Password  string  `json:"password,omitempty"`
CreatedAt  time.Time  `json:"created_at,omitempty"`
}
