# RemoteResponse

# The RemoteResponse Object
### Description
The `RemoteResponse` object is used to represent information returned from a third-party endpoint.

### Usage Example
View the `RemoteResponse` returned from your `DataPassthrough`.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Headers`                                                          | map[string]*interface{}*                                           | :heavy_minus_sign:                                                 | N/A                                                                | [object Object]                                                    |
| `Method`                                                           | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                | GET                                                                |
| `Path`                                                             | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                | /scooters                                                          |
| `Response`                                                         | *interface{}*                                                      | :heavy_check_mark:                                                 | N/A                                                                | [object Object]                                                    |
| `ResponseHeaders`                                                  | map[string]*interface{}*                                           | :heavy_minus_sign:                                                 | N/A                                                                | [object Object]                                                    |
| `ResponseType`                                                     | [*shared.ResponseType](../../../pkg/models/shared/responsetype.md) | :heavy_minus_sign:                                                 | N/A                                                                | JSON                                                               |
| `Status`                                                           | *int64*                                                            | :heavy_check_mark:                                                 | N/A                                                                | 200                                                                |