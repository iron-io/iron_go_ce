package titan

import (
    "strings"
    "fmt"
    "encoding/json"
    "errors"
    "github.com/dghubble/sling"
)

type ImagesApi struct {
    basePath  string
}

func NewImagesApi() *ImagesApi{
    return &ImagesApi {
        basePath:   "https://localhost:8080/v1",
    }
}

func NewImagesApiWithBasePath(basePath string) *ImagesApi{
    return &ImagesApi {
        basePath:   basePath,
    }
}

/**
 * Get all image names.
 * TODO: Using images for lack of a better term. See https://github.com/iron-io/titan/issues/43 for discussion.
 * @return ImagesWrapper
 */
//func (a ImagesApi) ImagesGet () (ImagesWrapper, error) {
func (a ImagesApi) ImagesGet () (ImagesWrapper, error) {

    _sling := sling.New().Get(a.basePath)

    // create path and map variables
    path := "/v1/images"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(ImagesWrapper)

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
 * Get image by name.
 * NOT IMPLEMENTED YET. This gives more details about on image, such as statistics and what not.
 * @param name Name of the image.
 * @return ImageWrapper
 */
//func (a ImagesApi) ImagesNameGet (name string) (ImageWrapper, error) {
func (a ImagesApi) ImagesNameGet (name string) (ImageWrapper, error) {

    _sling := sling.New().Get(a.basePath)

    // create path and map variables
    path := "/v1/images/{name}"
    path = strings.Replace(path, "{" + "name" + "}", fmt.Sprintf("%v", name), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(ImageWrapper)

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
