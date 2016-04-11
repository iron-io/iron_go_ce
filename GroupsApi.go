package titan

import (
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "github.com/dghubble/sling"
)

type GroupsApi struct {
    Configuration Configuration
}

func NewGroupsApi() *GroupsApi{
    configuration := NewConfiguration()
    return &GroupsApi {
        Configuration: *configuration,
    }
}

func NewGroupsApiWithBasePath(basePath string) *GroupsApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &GroupsApi {
        Configuration: *configuration,
    }
}

/**
 * Get all group names.
 * Get a list of all the groups in the system.
 * @return GroupsWrapper
 */
//func (a GroupsApi) GroupsGet () (GroupsWrapper, error) {
func (a GroupsApi) GroupsGet () (GroupsWrapper, error) {

    _sling := sling.New().Get(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(GroupsWrapper)

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
 * Get information for a group.
 * This gives more details about a job group, such as statistics.
 * @param name name of the group.
 * @return GroupWrapper
 */
//func (a GroupsApi) GroupsNameGet (name string) (GroupWrapper, error) {
func (a GroupsApi) GroupsNameGet (name string) (GroupWrapper, error) {

    _sling := sling.New().Get(a.Configuration.BasePath)

    // create path and map variables
    path := "/v1/groups/{name}"
    path = strings.Replace(path, "{" + "name" + "}", fmt.Sprintf("%v", name), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(GroupWrapper)

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
