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
    res, err := s.Offers.List(ctx, operations.OffersListRequest{
        XAccountToken: "perferendis",
        ApplicationID: mergeatsgo.String("corrupti"),
        CreatedAfter: types.MustTimeFromString("2022-03-06T01:38:10.905Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-29T20:03:41.130Z"),
        CreatorID: mergeatsgo.String("eius"),
        Cursor: mergeatsgo.String("necessitatibus"),
        Expand: operations.OffersListExpandApplication.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-06-03T07:58:02.123Z"),
        ModifiedBefore: types.MustTimeFromString("2021-01-18T16:49:49.451Z"),
        PageSize: mergeatsgo.Int64(271653),
        RemoteFields: mergeatsgo.String("tempora"),
        RemoteID: mergeatsgo.String("voluptate"),
        ShowEnumOrigins: mergeatsgo.String("reiciendis"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.OffersListRequest](../../models/operations/offerslistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.OffersListResponse](../../models/operations/offerslistresponse.md), error**


## Retrieve

Returns an `Offer` object with the given `id`.

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

    ctx := context.Background()
    res, err := s.Offers.Retrieve(ctx, operations.OffersRetrieveRequest{
        XAccountToken: "ex",
        Expand: operations.OffersRetrieveExpandApplication.ToPointer(),
        ID: "3e8b445e-80ca-455e-bd20-e457e1858b6a",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: mergeatsgo.String("voluptatum"),
        ShowEnumOrigins: mergeatsgo.String("error"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.OffersRetrieveRequest](../../models/operations/offersretrieverequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.OffersRetrieveResponse](../../models/operations/offersretrieveresponse.md), error**

