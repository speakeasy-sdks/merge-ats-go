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
        XAccountToken: "placeat",
        ApplicationID: ats.String("perspiciatis"),
        CreatedAfter: types.MustTimeFromString("2021-12-04T19:37:52.532Z"),
        CreatedBefore: types.MustTimeFromString("2021-08-20T03:23:43.072Z"),
        CreatorID: ats.String("ullam"),
        Cursor: ats.String("unde"),
        Expand: operations.OffersListExpandCreator.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-06-17T22:10:29.339Z"),
        ModifiedBefore: types.MustTimeFromString("2022-08-23T12:57:35.673Z"),
        PageSize: ats.Int64(668234),
        RemoteFields: operations.OffersListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("error"),
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
        XAccountToken: "esse",
        Expand: operations.OffersRetrieveExpandApplication.ToPointer(),
        ID: "1d311352-965b-4b8a-b202-611435e139db",
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

