# RejectReasons
(*.RejectReasons*)

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
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.RejectReasons.List(ctx, operations.RejectReasonsListRequest{
        XAccountToken: "string",
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
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "string"

    var id string = "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    var includeRemoteData *bool = false

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

