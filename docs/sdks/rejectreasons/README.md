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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.RejectReasonsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.RejectReasons.List(ctx, operations.RejectReasonsListRequest{
        XAccountToken: "voluptatum",
        CreatedAfter: types.MustTimeFromString("2021-02-26T20:36:25.696Z"),
        CreatedBefore: types.MustTimeFromString("2022-11-29T08:01:39.148Z"),
        Cursor: ats.String("blanditiis"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-11-05T22:13:21.002Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-29T16:33:10.615Z"),
        PageSize: ats.Int64(342611),
        RemoteID: ats.String("saepe"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedRejectReasonList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RejectReasonsListRequest](../../models/operations/rejectreasonslistrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.RejectReasonsListSecurity](../../models/operations/rejectreasonslistsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "error"
    id := "04f3b119-4b8a-4bf6-83a7-9f9dfe0ab7da"
    includeRemoteData := false
    operationSecurity := operations.RejectReasonsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.RejectReasons.Retrieve(ctx, operationSecurity, xAccountToken, id, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }

    if res.RejectReason != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `security`                                                                                           | [operations.RejectReasonsRetrieveSecurity](../../models/operations/rejectreasonsretrievesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `xAccountToken`                                                                                      | *string*                                                                                             | :heavy_check_mark:                                                                                   | Token identifying the end user.                                                                      |
| `id`                                                                                                 | *string*                                                                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `includeRemoteData`                                                                                  | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Whether to include the original data Merge fetched from the third-party to produce these models.     |


### Response

**[*operations.RejectReasonsRetrieveResponse](../../models/operations/rejectreasonsretrieveresponse.md), error**

