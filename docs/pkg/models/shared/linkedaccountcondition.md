# LinkedAccountCondition


## Fields

| Field                                                                                                                           | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `CommonModel`                                                                                                                   | **string*                                                                                                                       | :heavy_minus_sign:                                                                                                              | The common model for a specific condition.                                                                                      |
| `ConditionSchemaID`                                                                                                             | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | The ID indicating which condition schema to use for a specific condition.                                                       |
| `FieldName`                                                                                                                     | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | The name of the field on the common model that this condition corresponds to, if they conceptually match. e.g. "location_type". |
| `NativeName`                                                                                                                    | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | User-facing *native condition* name. e.g. "Skip Manager".                                                                       |
| `Operator`                                                                                                                      | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | The operator for a specific condition.                                                                                          |
| `Value`                                                                                                                         | *interface{}*                                                                                                                   | :heavy_minus_sign:                                                                                                              | The value for a condition.                                                                                                      |