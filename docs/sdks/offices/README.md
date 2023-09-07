# Offices

### Available Operations

* [List](#list) - Returns a list of `Office` objects.
* [Retrieve](#retrieve) - Returns an `Office` object with the given `id`.

## List

Returns a list of `Office` objects.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.OfficesListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Offices.List(ctx, operations.OfficesListRequest{
        XAccountToken: "rem",
        CreatedAfter: types.MustTimeFromString("2022-04-02T00:47:15.232Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-02T12:56:56.327Z"),
        Cursor: ats.String("minima"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2020-02-07T00:03:24.642Z"),
        ModifiedBefore: types.MustTimeFromString("2022-06-30T09:48:42.630Z"),
        PageSize: ats.Int64(48690),
        RemoteID: ats.String("saepe"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedOfficeList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.OfficesListRequest](../../models/operations/officeslistrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.OfficesListSecurity](../../models/operations/officeslistsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.OfficesListResponse](../../models/operations/officeslistresponse.md), error**


## Retrieve

Returns an `Office` object with the given `id`.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "numquam"
    id := "57e1858b-6a89-4fbe-ba5a-a8e4824d0ab4"
    includeRemoteData := false
    operationSecurity := operations.OfficesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Offices.Retrieve(ctx, operationSecurity, xAccountToken, id, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }

    if res.Office != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `security`                                                                                       | [operations.OfficesRetrieveSecurity](../../models/operations/officesretrievesecurity.md)         | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |
| `xAccountToken`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | Token identifying the end user.                                                                  |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `includeRemoteData`                                                                              | **bool*                                                                                          | :heavy_minus_sign:                                                                               | Whether to include the original data Merge fetched from the third-party to produce these models. |


### Response

**[*operations.OfficesRetrieveResponse](../../models/operations/officesretrieveresponse.md), error**

