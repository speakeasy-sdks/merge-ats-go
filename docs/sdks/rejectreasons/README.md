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
        XAccountToken: "fugit",
        CreatedAfter: types.MustTimeFromString("2022-09-29T05:24:35.816Z"),
        CreatedBefore: types.MustTimeFromString("2022-02-16T03:38:31.705Z"),
        Cursor: ats.String("officia"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-03-29T07:53:10.726Z"),
        ModifiedBefore: types.MustTimeFromString("2021-12-03T18:34:18.310Z"),
        PageSize: ats.Int64(185518),
        RemoteID: ats.String("expedita"),
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
    operationSecurity := operations.RejectReasonsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.RejectReasons.Retrieve(ctx, operations.RejectReasonsRetrieveRequest{
        XAccountToken: "voluptatum",
        ID: "95c537c6-454e-4fb0-b348-96c3ca5acfbe",
        IncludeRemoteData: ats.Bool(false),
    }, operationSecurity)
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
| `request`                                                                                            | [operations.RejectReasonsRetrieveRequest](../../models/operations/rejectreasonsretrieverequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.RejectReasonsRetrieveSecurity](../../models/operations/rejectreasonsretrievesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.RejectReasonsRetrieveResponse](../../models/operations/rejectreasonsretrieveresponse.md), error**

