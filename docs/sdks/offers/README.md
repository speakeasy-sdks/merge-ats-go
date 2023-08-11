# Offers

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.OffersListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Offers.List(ctx, operations.OffersListRequest{
        XAccountToken: "eius",
        ApplicationID: ats.String("aspernatur"),
        CreatedAfter: types.MustTimeFromString("2022-10-19T19:28:46.304Z"),
        CreatedBefore: types.MustTimeFromString("2021-12-21T17:45:55.733Z"),
        CreatorID: ats.String("incidunt"),
        Cursor: ats.String("quasi"),
        Expand: operations.OffersListExpandApplicationCreator.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-08-25T11:14:44.041Z"),
        ModifiedBefore: types.MustTimeFromString("2022-11-03T15:45:18.212Z"),
        PageSize: ats.Int64(233078),
        RemoteFields: operations.OffersListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("aperiam"),
        ShowEnumOrigins: operations.OffersListShowEnumOriginsStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedOfferList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.OffersListRequest](../../models/operations/offerslistrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.OffersListSecurity](../../models/operations/offerslistsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.OffersRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Offers.Retrieve(ctx, operations.OffersRetrieveRequest{
        XAccountToken: "cupiditate",
        Expand: operations.OffersRetrieveExpandCreator.ToPointer(),
        ID: "b0929921-aefb-49f5-8c4d-86e68e4be056",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.OffersRetrieveRemoteFieldsStatus.ToPointer(),
        ShowEnumOrigins: operations.OffersRetrieveShowEnumOriginsStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Offer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.OffersRetrieveRequest](../../models/operations/offersretrieverequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.OffersRetrieveSecurity](../../models/operations/offersretrievesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.OffersRetrieveResponse](../../models/operations/offersretrieveresponse.md), error**

