# IdStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier representing a specific task. | [optional] [default to null]
**Status** | **string** | States and valid transitions.                   +---------+        +---------&gt; delayed &lt;----------------+                  +----+----+                |                       |                     |                       |                     |                  +----v----+                |        +---------&gt; queued  &lt;----------------+                  +----+----+                *                       |                     *                       |               retry * creates new task                  +----v----+                *                  | running |                *                  +--+-+-+--+                |           +---------|-|-|-----+-------------+       +---|---------+ | +-----|---------+   |       |   |           |       |         |   | +-----v---^-+      +--v-------^+     +--v---^-+ | success   |      | cancelled |     |  error | +-----------+      +-----------+     +--------+  * delayed - has a delay. * queued - Ready to be consumed when it&#39;s turn comes. * running - Currently consumed by a runner which will attempt to process it. * success - (or complete? success/error is common javascript terminology) * error - Something went wrong. In this case more information can be obtained   by inspecting the \&quot;reason\&quot; field.   - timeout   - killed - forcibly killed by worker due to resource restrictions or access     violations.   - bad_exit - exited with non-zero status due to program termination/crash. * cancelled - cancelled via API. More information in the reason field.   - client_request - Request was cancelled by a client.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


