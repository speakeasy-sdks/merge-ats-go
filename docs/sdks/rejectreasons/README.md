# RejectReasons

### Available Operations

* [List](#list) - Returns a list of `RejectReason` objects.
* [Retrieve](#retrieve) - Returns a `RejectReason` object with the given `id`.

## List

Returns a list of `RejectReason` objects.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RejectReasons.List(ctx, operations.RejectReasonsListRequest{
        XAccountToken: "nam",
        CreatedAfter: types.MustTimeFromString("2021-11-01T04:33:49.270Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-10T02:44:39.117Z"),
        Cursor: mergeatsgo.String("deserunt"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-05-26T20:33:50.122Z"),
        ModifiedBefore: types.MustTimeFromString("2021-03-27T07:55:22.584Z"),
        PageSize: mergeatsgo.Int64(833819),
        RemoteID: mergeatsgo.String("delectus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedRejectReasonList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.RejectReasonsListRequest](../../models/operations/rejectreasonslistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.RejectReasonsListResponse](../../models/operations/rejectreasonslistresponse.md), error**


## Retrieve

Returns a `RejectReason` object with the given `id`.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    xAccountToken := "voluptates"
    id := "0ab7da8a-50ce-4187-b86b-c173d689eee9"
    includeRemoteData := false

    ctx := context.Background()
    res, err := s.RejectReasons.Retrieve(ctx, xAccountToken, id, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }

    if res.RejectReason != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `xAccountToken`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | Token identifying the end user.                                                                  |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `includeRemoteData`                                                                              | **bool*                                                                                          | :heavy_minus_sign:                                                                               | Whether to include the original data Merge fetched from the third-party to produce these models. |


### Response

**[*operations.RejectReasonsRetrieveResponse](../../models/operations/rejectreasonsretrieveresponse.md), error**

