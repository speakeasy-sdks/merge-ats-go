# ModelOperation

# The ModelOperation Object
### Description
The `ModelOperation` object is used to represent the operations that are currently supported for a given model.

### Usage Example
View what operations are supported for the `Candidate` endpoint.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `AvailableOperations`                        | []*string*                                   | :heavy_check_mark:                           | N/A                                          | ["FETCH","CREATE"]                           |
| `ModelName`                                  | *string*                                     | :heavy_check_mark:                           | N/A                                          | Candidate                                    |
| `RequiredPostParameters`                     | []*string*                                   | :heavy_check_mark:                           | N/A                                          | ["remote_user_id"]                           |
| `SupportedFields`                            | []*string*                                   | :heavy_check_mark:                           | N/A                                          | ["first_name","last_name","company","title"] |