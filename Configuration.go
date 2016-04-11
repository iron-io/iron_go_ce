package titan

import (

)

type Configuration struct {
    UserName  string  `json:"userName,omitempty"`
    ApiKey  string  `json:"apiKey,omitempty"`
    Debug  bool  `json:"debug,omitempty"`
    DebugFile  string  `json:"debugFile,omitempty"`
    OAuthToken  string  `json:"oAuthToken,omitempty"`
    Timeout  int  `json:"timeout,omitempty"`
    BasePath  string  `json:"basePath,omitempty"`
    Host  string  `json:"host,omitempty"`
    Scheme  string  `json:"scheme,omitempty"`
}

func NewConfiguration() *Configuration {
    return &Configuration{
        BasePath: "https://localhost:8080/v1",
        UserName: "",
        Debug: false,
        }
}