# Offers
(*Offers*)

### Available Operations

* [List](#list) - Returns a list of `Offer` objects.
* [Retrieve](#retrieve) - Returns an `Offer` object with the given `id`.

## List

Returns a list of `Offer` objects.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Offers.List(ctx, operations.OffersListRequest{
        XAccountToken: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedOfferList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.OffersListRequest](../../pkg/models/operations/offerslistrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.OffersListResponse](../../pkg/models/operations/offerslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Retrieve

Returns an `Offer` object with the given `id`.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Offers.Retrieve(ctx, operations.OffersRetrieveRequest{
        XAccountToken: "string",
        ID: "5fea5659-1081-4ad2-8d60-4c8e92b241fa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Offer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.OffersRetrieveRequest](../../pkg/models/operations/offersretrieverequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.OffersRetrieveResponse](../../pkg/models/operations/offersretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
