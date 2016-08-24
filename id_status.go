/* 
 * IronWorker CE API
 *
 * The ultimate, language agnostic, container based task processing framework.
 *
 * OpenAPI spec version: 0.5.5
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package worker

type IdStatus struct {

	// Unique identifier representing a specific task.
	Id string `json:"id,omitempty"`

	// States and valid transitions.                   +---------+        +---------> delayed <----------------+                  +----+----+                |                       |                     |                       |                     |                  +----v----+                |        +---------> queued  <----------------+                  +----+----+                *                       |                     *                       |               retry * creates new task                  +----v----+                *                  | running |                *                  +--+-+-+--+                |           +---------|-|-|-----+-------------+       +---|---------+ | +-----|---------+   |       |   |           |       |         |   | +-----v---^-+      +--v-------^+     +--v---^-+ | success   |      | cancelled |     |  error | +-----------+      +-----------+     +--------+  * delayed - has a delay. * queued - Ready to be consumed when it's turn comes. * running - Currently consumed by a runner which will attempt to process it. * success - (or complete? success/error is common javascript terminology) * error - Something went wrong. In this case more information can be obtained   by inspecting the \"reason\" field.   - timeout   - killed - forcibly killed by worker due to resource restrictions or access     violations.   - bad_exit - exited with non-zero status due to program termination/crash. * cancelled - cancelled via API. More information in the reason field.   - client_request - Request was cancelled by a client. 
	Status string `json:"status,omitempty"`
}
