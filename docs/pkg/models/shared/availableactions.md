# AvailableActions

# The AvailableActions Object
### Description
The `Activity` object is used to see all available model/operation combinations for an integration.

### Usage Example
Fetch all the actions available for the `Zenefits` integration.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AvailableModelOperations`                                                    | [][shared.ModelOperation](../../../pkg/models/shared/modeloperation.md)       | :heavy_minus_sign:                                                            | N/A                                                                           | [object Object]                                                               |
| `Integration`                                                                 | [shared.AccountIntegration](../../../pkg/models/shared/accountintegration.md) | :heavy_check_mark:                                                            | N/A                                                                           |                                                                               |
| `PassthroughAvailable`                                                        | *bool*                                                                        | :heavy_check_mark:                                                            | N/A                                                                           | true                                                                          |