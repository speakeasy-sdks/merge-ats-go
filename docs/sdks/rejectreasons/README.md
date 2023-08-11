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
        XAccountToken: "itaque",
        CreatedAfter: types.MustTimeFromString("2022-08-10T23:51:31.143Z"),
        CreatedBefore: types.MustTimeFromString("2022-04-10T19:32:47.130Z"),
        Cursor: ats.String("laborum"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-10-11T01:40:09.903Z"),
        ModifiedBefore: types.MustTimeFromString("2022-12-22T23:09:45.932Z"),
        PageSize: ats.Int64(39992),
        RemoteID: ats.String("consequatur"),
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
        XAccountToken: "officia",
        ID: "e6b6bc9b-8f75-49ea-855a-9741d3113529",
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

