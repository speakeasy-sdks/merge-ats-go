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
        XAccountToken: "nisi",
        CreatedAfter: types.MustTimeFromString("2020-05-14T04:22:15.063Z"),
        CreatedBefore: types.MustTimeFromString("2021-08-14T17:27:02.277Z"),
        Cursor: ats.String("sequi"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2020-03-31T16:32:36.866Z"),
        ModifiedBefore: types.MustTimeFromString("2021-04-06T05:59:09.931Z"),
        PageSize: ats.Int64(364544),
        RemoteID: ats.String("voluptate"),
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
    xAccountToken := "blanditiis"
    id := "a6458427-3a84-418d-9623-09fb0929921a"
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

