# Tag

# The Tag Object
### Description
The `Tag` object is used to represent a tag for a candidate.
### Usage Example
Fetch from the `LIST Tags` endpoint and view the tags used within a company.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `FieldMappings`                                                                | map[string]*interface{}*                                                       | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `ModifiedAt`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                     | :heavy_minus_sign:                                                             | This is the datetime that this object was last updated by Merge                | 2021-10-16T00:00:00Z                                                           |
| `Name`                                                                         | **string*                                                                      | :heavy_minus_sign:                                                             | The tag's name.                                                                | High-Priority                                                                  |
| `RemoteData`                                                                   | []map[string]*interface{}*                                                     | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `RemoteID`                                                                     | **string*                                                                      | :heavy_minus_sign:                                                             | The third-party API ID of the matching object.                                 | 4567                                                                           |
| `RemoteWasDeleted`                                                             | **bool*                                                                        | :heavy_minus_sign:                                                             | Indicates whether or not this object has been deleted by third party webhooks. |                                                                                |