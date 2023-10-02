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
        XAccountToken: "Northeast Metal Canada",
        ApplicationID: mergeatsgo.String("Data Response West"),
        CreatedAfter: types.MustTimeFromString("2023-02-26T13:00:25.189Z"),
        CreatedBefore: types.MustTimeFromString("2023-12-20T19:28:33.339Z"),
        CreatorID: mergeatsgo.String("primary"),
        Cursor: mergeatsgo.String("Designer hacking"),
        Expand: operations.OffersListExpandApplicationCreator.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-12-18T09:50:13.895Z"),
        ModifiedBefore: types.MustTimeFromString("2021-01-09T04:15:41.822Z"),
        PageSize: mergeatsgo.Int64(504966),
        RemoteFields: mergeatsgo.String("explicit"),
        RemoteID: mergeatsgo.String("virtual"),
        ShowEnumOrigins: mergeatsgo.String("itaque"),
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
        XAccountToken: "tracksuit Markets",
        Expand: operations.OffersRetrieveExpandApplication.ToPointer(),
        ID: "081ad20d-604c-48e9-ab24-1fa379087a15",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: mergeatsgo.String("Mobility"),
        ShowEnumOrigins: mergeatsgo.String("recklessly program"),
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

