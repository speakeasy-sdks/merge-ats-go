# CandidatesListRequest


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `XAccountToken`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Token identifying the end user.                                                                                        |
| `CreatedAfter`                                                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                                             | :heavy_minus_sign:                                                                                                     | If provided, will only return objects created after this datetime.                                                     |
| `CreatedBefore`                                                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                                                             | :heavy_minus_sign:                                                                                                     | If provided, will only return objects created before this datetime.                                                    |
| `Cursor`                                                                                                               | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | The pagination cursor value.                                                                                           |
| `EmailAddresses`                                                                                                       | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | If provided, will only return candidates with these email addresses; multiple addresses can be separated by commas.    |
| `Expand`                                                                                                               | [*CandidatesListExpand](../../models/operations/candidateslistexpand.md)                                               | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
| `FirstName`                                                                                                            | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | If provided, will only return candidates with this first name.                                                         |
| `IncludeDeletedData`                                                                                                   | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include data that was marked as deleted by third party webhooks.                                            |
| `IncludeRemoteData`                                                                                                    | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                       |
| `LastName`                                                                                                             | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | If provided, will only return candidates with this last name.                                                          |
| `ModifiedAfter`                                                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                                                             | :heavy_minus_sign:                                                                                                     | If provided, only objects synced by Merge after this date time will be returned.                                       |
| `ModifiedBefore`                                                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                                                             | :heavy_minus_sign:                                                                                                     | If provided, only objects synced by Merge before this date time will be returned.                                      |
| `PageSize`                                                                                                             | **int64*                                                                                                               | :heavy_minus_sign:                                                                                                     | Number of results to return per page.                                                                                  |
| `RemoteID`                                                                                                             | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | The API provider's ID for the given object.                                                                            |
| `Tags`                                                                                                                 | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | If provided, will only return candidates with these tags; multiple tags can be separated by commas.                    |