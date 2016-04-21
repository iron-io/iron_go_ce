package titan

import (
)


type IdStatus struct {
    // Unique identifier representing a specific job.
    Id  string  `json:"id,omitempty"`
    // States and valid transitions.\n\n                 +---------+\n       +---------> delayed <----------------+\n                 +----+----+                |\n                      |                     |\n                      |                     |\n                 +----v----+                |\n       +---------> queued  <----------------+\n                 +----+----+                *\n                      |                     *\n                      |               retry * creates new job\n                 +----v----+                *\n                 | running |                *\n                 +--+-+-+--+                |\n          +---------|-|-|-----+-------------+\n      +---|---------+ | +-----|---------+   |\n      |   |           |       |         |   |\n+-----v---^-+      +--v-------^+     +--v---^-+\n| success   |      | cancelled |     |  error |\n+-----------+      +-----------+     +--------+\n\n* delayed - has a delay.\n* queued - Ready to be consumed when it's turn comes.\n* running - Currently consumed by a runner which will attempt to process it.\n* success - (or complete? success/error is common javascript terminology)\n* error - Something went wrong. In this case more information can be obtained\n  by inspecting the \"reason\" field.\n  - timeout\n  - killed - forcibly killed by worker due to resource restrictions or access\n    violations.\n  - bad_exit - exited with non-zero status due to program termination/crash.\n* cancelled - cancelled via API. More information in the reason field.\n  - client_request - Request was cancelled by a client.\n
    Status  string  `json:"status,omitempty"`
}
