package titan

import (
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "github.com/dghubble/sling"
    "os"
    "time"
)

type JobsApi struct {
    basePath  string
}

func NewJobsApi() *JobsApi{
    return &JobsApi {
        basePath:   "https://localhost:8080/v1",
    }
}

func NewJobsApiWithBasePath(basePath string) *JobsApi{
    return &JobsApi {
        basePath:   basePath,
    }
}

/**
 * Cancel a job.
 * Cancels a job in delayed, queued or running status. The worker may continue to run a running job. reason is set to `client_request`.
 * @param id Job id
 * @param details Human-readable detailed message explaining cancellation reason.
 * @return JobWrapper
 */
//func (a JobsApi) JobIdCancelPost (id string, details string) (JobWrapper, error) {
func (a JobsApi) JobIdCancelPost (id string, details string) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}/cancel"
    path = strings.Replace(path, "{" + "id" + "}", fmt.Sprintf("%v", id), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(details)

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
 * Delete only succeeds if job status is one of `succeeded\n| failed | cancelled`. Cancel a job if it is another state and needs to\nbe deleted.  All information about the job, including the log, is\nirretrievably lost when this is invoked.
 * @param id Job id
 * @return void
 */
//func (a JobsApi) JobIdDelete (id string) (error) {
func (a JobsApi) JobIdDelete (id string) (error) {

    _sling := sling.New().Delete(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}"
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
 * Job is marked as failed if it was in a valid state. Job&#39;s `completed_at` time is initialized.
 * @param id Job id
 * @param reason Reason for job failure.
 * @param details Details of job failure.
 * @param log Output log for the job. Content-Type must be \&quot;text/plain; charset=utf-8\&quot;.
 * @return JobWrapper
 */
//func (a JobsApi) JobIdFailPost (id string, reason string, details string, log *os.File) (JobWrapper, error) {
func (a JobsApi) JobIdFailPost (id string, reason string, details string, log *os.File) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}/fail"
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
        details    string `url:"details,omitempty"`
        log    *os.File `url:"log,omitempty"`
    }
    _sling = _sling.BodyForm(&FormParams{ reason: reason,details: details,log: log })

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
 * @param id Job id
 * @return JobWrapper
 */
//func (a JobsApi) JobIdGet (id string) (JobWrapper, error) {
func (a JobsApi) JobIdGet (id string) (JobWrapper, error) {

    _sling := sling.New().Get(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}"
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
 * @param id Job id
 * @return string
 */
//func (a JobsApi) JobIdLogGet (id string) (string, error) {
func (a JobsApi) JobIdLogGet (id string) (string, error) {

    _sling := sling.New().Get(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}/log"
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
 * Retry a job.
 * The /retry endpoint can be used to force a retry of jobs with status succeeded or cancelled. It can also be used to retry jobs that in the failed state, but whose max_retries field is 0. The retried job will continue to have max_retries = 0.
 * @param id Job id
 * @return JobWrapper
 */
//func (a JobsApi) JobIdRetryPost (id string) (JobWrapper, error) {
func (a JobsApi) JobIdRetryPost (id string) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}/retry"
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
 * Mark job as succeeded.
 * Job status is changed to succeeded if it was in a valid state before. Job&#39;s `completed_at` time is initialized.
 * @param id Job id
 * @param log Output log for the job. Content-Type must be \&quot;text/plain; charset=utf-8\&quot;.
 * @return JobWrapper
 */
//func (a JobsApi) JobIdSuccessPost (id string, log *os.File) (JobWrapper, error) {
func (a JobsApi) JobIdSuccessPost (id string, log *os.File) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}/success"
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
 * Extend job timeout.
 * Consumers can sometimes take a while to run the task after accepting it.  An example is when the runner does not have the docker image locally, it can spend a significant time downloading the image.\nIf the timeout is small, the job may never get to run, or run but not be accepted by Titan. Consumers can touch the job before it times out. Titan will reset the timeout, giving the consumer another timeout seconds to run the job.\nTouch is only valid while the job is in a running state. If touch fails, the runner may stop running the job.
 * @param id Job id
 * @return void
 */
//func (a JobsApi) JobIdTouchPost (id string) (error) {
func (a JobsApi) JobIdTouchPost (id string) (error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/job/{id}/touch"
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
 * Get next job.
 * Gets the next job in the queue, ready for processing. Titan may return &lt;=n jobs. Consumers should start processing jobs in order. Each returned job is set to `status` \&quot;running\&quot; and `started_at` is set to the current time. No other consumer can retrieve this job.
 * @param n Number of jobs to return.
 * @return JobsWrapper
 */
//func (a JobsApi) JobsConsumeGet (n int32) (JobsWrapper, error) {
func (a JobsApi) JobsConsumeGet (n int32) (JobsWrapper, error) {

    _sling := sling.New().Get(a.basePath)

    // create path and map variables
    path := "/v1/jobs/consume"

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
/**
 * Peek at list of jobs.
 * Get a list of active jobs. This endpoint can be used to observe the state of jobs in Titan. To run a job, use /jobs/consume. TODO: Needs pagination support.
 * @param n Number of jobs to return. Titan may return &lt;=n jobs. Titan does not make any guarantees about job ordering, but jobs will not be repeated. To make sure you get unique jobs, use the cursor effectively. TODO: We don&#39;t actually support pagination.
 * @return JobsWrapper
 */
//func (a JobsApi) JobsGet (n int32) (JobsWrapper, error) {
func (a JobsApi) JobsGet (n int32) (JobsWrapper, error) {

    _sling := sling.New().Get(a.basePath)

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
/**
 * Get job list by name.
 * Gets the next job in the queue, ready for processing.
 * @param name Name of this set of jobs.
 * @param createdAfter Will return jobs created after this time. In RFC3339 format.
 * @param n Number of jobs to return.
 * @return JobsWrapper
 */
//func (a JobsApi) JobsNameGet (name string, createdAfter time.Time, n int32) (JobsWrapper, error) {
func (a JobsApi) JobsNameGet (name string, createdAfter time.Time, n int32) (JobsWrapper, error) {

    _sling := sling.New().Get(a.basePath)

    // create path and map variables
    path := "/v1/jobs/{name}"
    path = strings.Replace(path, "{" + "name" + "}", fmt.Sprintf("%v", name), -1)

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
 * Enqueue Job
 * Enqueues job(s). If any of the jobs is invalid, none of the jobs are enqueued.
 * @param body Array of jobs to post.
 * @return JobsWrapper
 */
//func (a JobsApi) JobsPost (body NewJobsWrapper) (JobsWrapper, error) {
func (a JobsApi) JobsPost (body NewJobsWrapper) (JobsWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/jobs"

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
