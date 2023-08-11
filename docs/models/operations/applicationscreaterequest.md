# ApplicationsCreateRequest


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ApplicationEndpointRequest`                                                           | [shared.ApplicationEndpointRequest](../../models/shared/applicationendpointrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `XAccountToken`                                                                        | *string*                                                                               | :heavy_check_mark:                                                                     | Token identifying the end user.                                                        |
| `IsDebugMode`                                                                          | **bool*                                                                                | :heavy_minus_sign:                                                                     | Whether to include debug fields (such as log file links) in the response.              |
| `RunAsync`                                                                             | **bool*                                                                                | :heavy_minus_sign:                                                                     | Whether or not third-party updates should be run asynchronously.                       |