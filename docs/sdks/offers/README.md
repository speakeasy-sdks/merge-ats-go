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
        XAccountToken: "asperiores",
        ApplicationID: ats.String("facilis"),
        CreatedAfter: types.MustTimeFromString("2022-04-17T04:31:41.347Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-22T17:05:01.514Z"),
        CreatorID: ats.String("dolore"),
        Cursor: ats.String("laborum"),
        Expand: operations.OffersListExpandApplication.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-08-01T14:49:09.199Z"),
        ModifiedBefore: types.MustTimeFromString("2022-09-27T03:48:38.490Z"),
        PageSize: ats.Int64(378326),
        RemoteFields: operations.OffersListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("unde"),
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
        XAccountToken: "architecto",
        Expand: operations.OffersRetrieveExpandApplicationCreator.ToPointer(),
        ID: "fe1f08f4-294e-4369-8f44-7f603e8b445e",
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

