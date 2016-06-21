/* 
 * Titan API
 *
 * The ultimate, language agnostic, container based job processing framework.
 *
 * OpenAPI spec version: 0.4.9
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

package titan

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

type RunnerApi struct {
	Configuration Configuration
}

func NewRunnerApi() *RunnerApi {
	configuration := NewConfiguration()
	return &RunnerApi{
		Configuration: *configuration,
	}
}

func NewRunnerApiWithBasePath(basePath string) *RunnerApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &RunnerApi{
		Configuration: *configuration,
	}
}

/**
 * Mark job as failed.
 * Job is marked as failed if it was in a valid state. Job&#39;s &#x60;finished_at&#x60; time is initialized.
 *
 * @param name Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return *JobWrapper
 */
func (a RunnerApi) GroupsNameJobsIdErrorPost(name string, id string, body Complete) (*JobWrapper, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/groups/{name}/jobs/{id}/error"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// verify the required parameter 'name' is set
	if &name == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'name' when calling RunnerApi->GroupsNameJobsIdErrorPost")
	}
	// verify the required parameter 'id' is set
	if &id == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'id' when calling RunnerApi->GroupsNameJobsIdErrorPost")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'body' when calling RunnerApi->GroupsNameJobsIdErrorPost")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(JobWrapper)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Mark job as started, ie: status &#x3D; &#39;running&#39;
 * Job status is changed to &#39;running&#39; if it was in a valid state before. Job&#39;s &#x60;started_at&#x60; time is initialized.
 *
 * @param name Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return *JobWrapper
 */
func (a RunnerApi) GroupsNameJobsIdStartPost(name string, id string, body Start) (*JobWrapper, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/groups/{name}/jobs/{id}/start"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// verify the required parameter 'name' is set
	if &name == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'name' when calling RunnerApi->GroupsNameJobsIdStartPost")
	}
	// verify the required parameter 'id' is set
	if &id == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'id' when calling RunnerApi->GroupsNameJobsIdStartPost")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'body' when calling RunnerApi->GroupsNameJobsIdStartPost")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(JobWrapper)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Mark job as succeeded.
 * Job status is changed to succeeded if it was in a valid state before. Job&#39;s &#x60;completed_at&#x60; time is initialized.
 *
 * @param name Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return *JobWrapper
 */
func (a RunnerApi) GroupsNameJobsIdSuccessPost(name string, id string, body Complete) (*JobWrapper, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/groups/{name}/jobs/{id}/success"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// verify the required parameter 'name' is set
	if &name == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'name' when calling RunnerApi->GroupsNameJobsIdSuccessPost")
	}
	// verify the required parameter 'id' is set
	if &id == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'id' when calling RunnerApi->GroupsNameJobsIdSuccessPost")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		return new(JobWrapper), nil, errors.New("Missing required parameter 'body' when calling RunnerApi->GroupsNameJobsIdSuccessPost")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload = new(JobWrapper)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

