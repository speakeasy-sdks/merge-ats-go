# EmailAddress

# The EmailAddress Object
### Description
The `EmailAddress` object is used to represent a candidate's email address.
### Usage Example
Fetch from the `GET Candidate` endpoint and view their email addresses.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `EmailAddressType`                                                                    | [*shared.EmailAddressType](../../../pkg/models/shared/emailaddresstype.md)            | :heavy_minus_sign:                                                                    | The type of email address.<br/><br/>* `PERSONAL` - PERSONAL<br/>* `WORK` - WORK<br/>* `OTHER` - OTHER | PERSONAL                                                                              |
| `ModifiedAt`                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                            | :heavy_minus_sign:                                                                    | This is the datetime that this object was last updated by Merge                       | 2021-10-16T00:00:00Z                                                                  |
| `Value`                                                                               | **string*                                                                             | :heavy_minus_sign:                                                                    | The email address.                                                                    | merge_is_hiring@merge.dev                                                             |