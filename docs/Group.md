# Group

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of this group. Must be different than the image name. Can ony contain alphanumeric, -, and _. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Time when image first used/created. | [optional] [default to null]
**Image** | **string** | Name of Docker image to use in this group. You should include the image tag, which should be a version number, to be more accurate. Can be overridden on a per job basis with job.image. | [optional] [default to null]
**EnvVars** | **map[string]string** | User defined environment variables that will be passed in to each job in this group. | [optional] [default to null]
**MaxConcurrency** | **int32** | The maximum number of jobs that will run at the exact same time in this group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


