package titan

import (
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "github.com/dghubble/sling"
)

type RunnerApi struct {
    Configuration Configuration
}

func NewRunnerApi() *RunnerApi{
    configuration := NewConfiguration()
    return &RunnerApi {
        Configuration: *configuration,
    }
}

func NewRunnerApiWithBasePath(basePath string) *RunnerApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &RunnerApi {
        Configuration: *configuration,
    }
}

/**
 * Mark job as failed.
 * Job is marked as failed if it was in a valid state. Job&#39;s &#x60;finished_at&#x60; time is initialized.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @param reason Reason for job failure.
 * @return JobWrapper
 */
func (a RunnerApi) GroupsGroupNameJobsIdErrorPost (groupName string, id string, reason string) (JobWrapper, error) {
    // verify the required parameter 'groupName' is set
    if &groupName == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'groupName' when calling RunnerApi->GroupsGroupNameJobsIdErrorPost")
    }
    // verify the required parameter 'id' is set
    if &id == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'id' when calling RunnerApi->GroupsGroupNameJobsIdErrorPost")
    }
    // verify the required parameter 'reason' is set
    if &reason == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'reason' when calling RunnerApi->GroupsGroupNameJobsIdErrorPost")
    }
    _sling := sling.New().Post(a.Configuration.BasePath)

    
    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/error"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // add default headers if any
    for key := range a.Configuration.DefaultHeader {
      _sling = _sling.Set(key, a.Configuration.DefaultHeader[key])
    }
    

    // to determine the Content-Type header
    localVarHttpContentTypes := []string {
        "application/json", 
    }
    //set Content-Type header
    localVarHttpContentType := a.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {    
      _sling = _sling.Set("Content-Type", localVarHttpContentType)
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string {
        "application/json", 
    }
    //set Accept header
    localVarHttpHeaderAccept := a.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {  
        _sling = _sling.Set("Accept", localVarHttpHeaderAccept)
    }

    type FormParams struct {
        Reason    string `url:"reason,omitempty"`
    }
    _sling = _sling.BodyForm(&FormParams{ Reason: reason })

  var successPayload = new(JobWrapper)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * Mark job as started, ie: status &#x3D; &#39;running&#39;
 * Job status is changed to &#39;running&#39; if it was in a valid state before. Job&#39;s &#x60;started_at&#x60; time is initialized.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return JobWrapper
 */
func (a RunnerApi) GroupsGroupNameJobsIdStartPost (groupName string, id string, body Start) (JobWrapper, error) {
    // verify the required parameter 'groupName' is set
    if &groupName == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'groupName' when calling RunnerApi->GroupsGroupNameJobsIdStartPost")
    }
    // verify the required parameter 'id' is set
    if &id == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'id' when calling RunnerApi->GroupsGroupNameJobsIdStartPost")
    }
    // verify the required parameter 'body' is set
    if &body == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'body' when calling RunnerApi->GroupsGroupNameJobsIdStartPost")
    }
    _sling := sling.New().Post(a.Configuration.BasePath)

    
    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/start"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // add default headers if any
    for key := range a.Configuration.DefaultHeader {
      _sling = _sling.Set(key, a.Configuration.DefaultHeader[key])
    }
    

    // to determine the Content-Type header
    localVarHttpContentTypes := []string {
        "application/json", 
    }
    //set Content-Type header
    localVarHttpContentType := a.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {    
      _sling = _sling.Set("Content-Type", localVarHttpContentType)
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string {
        "application/json", 
    }
    //set Accept header
    localVarHttpHeaderAccept := a.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {  
        _sling = _sling.Set("Accept", localVarHttpHeaderAccept)
    }

// body params
    _sling = _sling.BodyJSON(body)

  var successPayload = new(JobWrapper)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * Mark job as succeeded.
 * Job status is changed to succeeded if it was in a valid state before. Job&#39;s &#x60;completed_at&#x60; time is initialized.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return JobWrapper
 */
func (a RunnerApi) GroupsGroupNameJobsIdSuccessPost (groupName string, id string) (JobWrapper, error) {
    // verify the required parameter 'groupName' is set
    if &groupName == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'groupName' when calling RunnerApi->GroupsGroupNameJobsIdSuccessPost")
    }
    // verify the required parameter 'id' is set
    if &id == nil {
        return *new(JobWrapper), errors.New("Missing required parameter 'id' when calling RunnerApi->GroupsGroupNameJobsIdSuccessPost")
    }
    _sling := sling.New().Post(a.Configuration.BasePath)

    
    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/success"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // add default headers if any
    for key := range a.Configuration.DefaultHeader {
      _sling = _sling.Set(key, a.Configuration.DefaultHeader[key])
    }
    

    // to determine the Content-Type header
    localVarHttpContentTypes := []string {
        "application/json", 
    }
    //set Content-Type header
    localVarHttpContentType := a.Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes)
    if localVarHttpContentType != "" {    
      _sling = _sling.Set("Content-Type", localVarHttpContentType)
    }

    // to determine the Accept header
    localVarHttpHeaderAccepts := []string {
        "application/json", 
    }
    //set Accept header
    localVarHttpHeaderAccept := a.Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
    if localVarHttpHeaderAccept != "" {  
        _sling = _sling.Set("Accept", localVarHttpHeaderAccept)
    }


  var successPayload = new(JobWrapper)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
