/* 
 * IronWorker CE API
 *
 * The ultimate, language agnostic, container based task processing framework.
 *
 * OpenAPI spec version: 0.5.2
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

import (
	"time"
)

type Group struct {

	// Name of this group. Must be different than the image name. Can ony contain alphanumeric, -, and _.
	Name string `json:"name,omitempty"`

	// Time when image first used/created.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Name of Docker image to use in this group. You should include the image tag, which should be a version number, to be more accurate. Can be overridden on a per task basis with task.image.
	Image string `json:"image,omitempty"`

	// User defined environment variables that will be passed in to each task in this group.
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// The maximum number of tasks that will run at the exact same time in this group.
	MaxConcurrency int32 `json:"max_concurrency,omitempty"`
}
