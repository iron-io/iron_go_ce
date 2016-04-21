package titan

import (
)

// Machine usable reason for job being in this state.\nValid values for error status are `timeout | killed | bad_exit`.\nValid values for cancelled status are `client_request`.\nFor everything else, this is undefined.\n
type Reason struct {
}
