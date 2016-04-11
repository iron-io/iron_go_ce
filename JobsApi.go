package titan

import (
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "github.com/dghubble/sling"
    "time"
    "os"
)

type JobsApi struct {
    Configuration Configuration
}

func NewJobsApi() *JobsApi{
    configuration := NewConfiguration()
    return &JobsApi {
        Configuration: *configuration,
    }
}

func NewJobsApiWithBasePath(basePath string) *JobsApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &JobsApi {
        Configuration: *configuration,
    }
}

/**
 * Get job list by group name.
 * This will list jobs for a particular group.
 * @param groupName Name of group for this set of jobs.
 * @param createdAfter Will return jobs created after this time. In RFC3339 format.
 * @param n Number of jobs to return.
 * @return JobsWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsGet (groupName string, createdAfter time.Time, n int32) (JobsWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsGet (groupName string, createdAfter time.Time, n int32) (JobsWrapper, error) {

    _sling := sling.New().Get(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)

    _sling = _sling.Path(path)

    type QueryParams struct {
        createdAfter    time.Time `url:"created_after,omitempty"`
n    int32 `url:"n,omitempty"`
}
    _sling = _sling.QueryStruct(&QueryParams{ createdAfter: createdAfter,n: n })
    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(JobsWrapper)

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
 * Cancel a job.
 * Cancels a job in delayed, queued or running status. The worker may continue to run a running job. reason is set to &#x60;client_request&#x60;.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdCancelPost (groupName string, id string) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdCancelPost (groupName string, id string) (JobWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/cancel"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
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
/**
 * Delete the job.
 * Delete only succeeds if job status is one of &#x60;succeeded\n| failed | cancelled&#x60;. Cancel a job if it is another state and needs to\nbe deleted.  All information about the job, including the log, is\nirretrievably lost when this is invoked.\n
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return void
 */
//func (a JobsApi) GroupsGroupNameJobsIdDelete (groupName string, id string) (error) {
func (a JobsApi) GroupsGroupNameJobsIdDelete (groupName string, id string) (error) {

    _sling := sling.New().Delete(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }




  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(nil, &failurePayload)

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

  return err
}
/**
 * Mark job as failed.
 * Job is marked as failed if it was in a valid state. Job&#39;s &#x60;finished_at&#x60; time is initialized.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @param reason Reason for job failure.
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdErrorPost (groupName string, id string, reason string) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdErrorPost (groupName string, id string, reason string) (JobWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/error"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

    type FormParams struct {
        reason    string `url:"reason,omitempty"`
    }
    _sling = _sling.BodyForm(&FormParams{ reason: reason })

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
 * Gets job by id
 * Gets a job by id.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdGet (groupName string, id string) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdGet (groupName string, id string) (JobWrapper, error) {

    _sling := sling.New().Get(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
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
/**
 * Get the log of a completed job.
 * Retrieves the log from log storage.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return string
 */
//func (a JobsApi) GroupsGroupNameJobsIdLogGet (groupName string, id string) (string, error) {
func (a JobsApi) GroupsGroupNameJobsIdLogGet (groupName string, id string) (string, error) {

    _sling := sling.New().Get(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/log"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "text/plain" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(string)

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
 * Send in a log for storage.
 * Logs are sent after a job completes since they may be very large and the runner can process the next job.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @param log Output log for the job. Content-Type must be \&quot;text/plain; charset&#x3D;utf-8\&quot;.
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdLogPost (groupName string, id string, log *os.File) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdLogPost (groupName string, id string, log *os.File) (JobWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/log"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

    type FormParams struct {
        log    *os.File `url:"log,omitempty"`
    }
    _sling = _sling.BodyForm(&FormParams{ log: log })

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
 * DEPRECATED - Update a job
 * Used to update status on job transitions. Eg: from &#39;running&#39; to &#39;success&#39;.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @param body Job data to post
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdPatch (groupName string, id string, body JobWrapper) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdPatch (groupName string, id string, body JobWrapper) (JobWrapper, error) {

    _sling := sling.New().Patch(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
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
 * Retry a job.
 * \&quot;The /retry endpoint can be used to force a retry of jobs\nwith status succeeded or cancelled. It can also be used to retry jobs\nthat in the failed state, but whose max_retries field is 0. The retried\njob will continue to have max_retries &#x3D; 0.\&quot;\n
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdRetryPost (groupName string, id string) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdRetryPost (groupName string, id string) (JobWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/retry"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
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
/**
 * Mark job as started, ie: status &#x3D; &#39;running&#39;
 * Job status is changed to &#39;running&#39; if it was in a valid state before. Job&#39;s &#x60;started_at&#x60; time is initialized.
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return JobWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsIdStartPost (groupName string, id string, body Start) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdStartPost (groupName string, id string, body Start) (JobWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/start"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
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
//func (a JobsApi) GroupsGroupNameJobsIdSuccessPost (groupName string, id string) (JobWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsIdSuccessPost (groupName string, id string) (JobWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/success"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
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
/**
 * Extend job timeout.
 * Consumers can sometimes take a while to run the task after accepting it.  An example is when the runner does not have the docker image locally, it can spend a significant time downloading the image.\nIf the timeout is small, the job may never get to run, or run but not be accepted by Titan. Consumers can touch the job before it times out. Titan will reset the timeout, giving the consumer another timeout seconds to run the job.\nTouch is only valid while the job is in a running state. If touch fails, the runner may stop running the job.\n
 * @param groupName Name of group for this set of jobs.
 * @param id Job id
 * @return void
 */
//func (a JobsApi) GroupsGroupNameJobsIdTouchPost (groupName string, id string) (error) {
func (a JobsApi) GroupsGroupNameJobsIdTouchPost (groupName string, id string) (error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs/{id}/touch"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }




  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(nil, &failurePayload)

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

  return err
}
/**
 * Enqueue Job
 * Enqueues job(s). If any of the jobs is invalid, none of the jobs are enqueued.\n
 * @param groupName name of the group.
 * @param body Array of jobs to post.
 * @return JobsWrapper
 */
//func (a JobsApi) GroupsGroupNameJobsPost (groupName string, body NewJobsWrapper) (JobsWrapper, error) {
func (a JobsApi) GroupsGroupNameJobsPost (groupName string, body NewJobsWrapper) (JobsWrapper, error) {

    _sling := sling.New().Post(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{group_name}/jobs"
    path = strings.Replace(path, "{" + "group_name" + "}", fmt.Sprintf("%v", groupName), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(body)

  var successPayload = new(JobsWrapper)

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
 * Get next job.
 * Gets the next job in the queue, ready for processing. Titan may return &lt;&#x3D;n jobs. Consumers should start processing jobs in order. Each returned job is set to &#x60;status&#x60; \&quot;running\&quot; and &#x60;started_at&#x60; is set to the current time. No other consumer can retrieve this job.
 * @param n Number of jobs to return.
 * @return JobsWrapper
 */
//func (a JobsApi) JobsGet (n int32) (JobsWrapper, error) {
func (a JobsApi) JobsGet (n int32) (JobsWrapper, error) {

    _sling := sling.New().Get(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/jobs"

    _sling = _sling.Path(path)

    type QueryParams struct {
        n    int32 `url:"n,omitempty"`
}
    _sling = _sling.QueryStruct(&QueryParams{ n: n })
    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(JobsWrapper)

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
