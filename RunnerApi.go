package titan

import (
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "github.com/dghubble/sling"
)

type RunnerApi struct {
    basePath  string
}

func NewRunnerApi() *RunnerApi{
    return &RunnerApi {
        basePath:   "https://localhost:8080/v1",
    }
}

func NewRunnerApiWithBasePath(basePath string) *RunnerApi{
    return &RunnerApi {
        basePath:   basePath,
    }
}

/**
 * Mark job as failed.
 * Job is marked as failed if it was in a valid state. Job&#39;s `finished_at` time is initialized.
 * @param name Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return JobWrapper
 */
//func (a RunnerApi) GroupsNameJobsIdErrorPost (name string, id string, body Complete) (JobWrapper, error) {
func (a RunnerApi) GroupsNameJobsIdErrorPost (name string, id string, body Complete) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/groups/{name}/jobs/{id}/error"
    path = strings.Replace(path, "{" + "name" + "}", fmt.Sprintf("%v", name), -1)
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
 * Mark job as started, ie: status = &#39;running&#39;
 * Job status is changed to &#39;running&#39; if it was in a valid state before. Job&#39;s `started_at` time is initialized.
 * @param name Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return JobWrapper
 */
//func (a RunnerApi) GroupsNameJobsIdStartPost (name string, id string, body Start) (JobWrapper, error) {
func (a RunnerApi) GroupsNameJobsIdStartPost (name string, id string, body Start) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/groups/{name}/jobs/{id}/start"
    path = strings.Replace(path, "{" + "name" + "}", fmt.Sprintf("%v", name), -1)
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
 * Job status is changed to succeeded if it was in a valid state before. Job&#39;s `completed_at` time is initialized.
 * @param name Name of group for this set of jobs.
 * @param id Job id
 * @param body 
 * @return JobWrapper
 */
//func (a RunnerApi) GroupsNameJobsIdSuccessPost (name string, id string, body Complete) (JobWrapper, error) {
func (a RunnerApi) GroupsNameJobsIdSuccessPost (name string, id string, body Complete) (JobWrapper, error) {

    _sling := sling.New().Post(a.basePath)

    // create path and map variables
    path := "/v1/groups/{name}/jobs/{id}/success"
    path = strings.Replace(path, "{" + "name" + "}", fmt.Sprintf("%v", name), -1)
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
