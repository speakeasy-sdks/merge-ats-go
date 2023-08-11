# Offices

### Available Operations

* [List](#list) - Returns a list of `Office` objects.
* [Retrieve](#retrieve) - Returns an `Office` object with the given `id`.

## List

Returns a list of `Office` objects.

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
    operationSecurity := operations.OfficesListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Offices.List(ctx, operations.OfficesListRequest{
        XAccountToken: "maxime",
        CreatedAfter: types.MustTimeFromString("2022-11-03T21:52:43.365Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-24T21:57:47.411Z"),
        Cursor: ats.String("libero"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-05-03T21:20:05.777Z"),
        ModifiedBefore: types.MustTimeFromString("2021-04-17T05:51:37.071Z"),
        PageSize: ats.Int64(637462),
        RemoteID: ats.String("quos"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedOfficeList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.OfficesListRequest](../../models/operations/officeslistrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.OfficesListSecurity](../../models/operations/officeslistsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.OfficesListResponse](../../models/operations/officeslistresponse.md), error**


## Retrieve

Returns an `Office` object with the given `id`.

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
    operationSecurity := operations.OfficesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Offices.Retrieve(ctx, operations.OfficesRetrieveRequest{
        XAccountToken: "placeat",
        ID: "070e1084-cb06-472d-9ad8-79eeb9665b85",
        IncludeRemoteData: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Office != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.OfficesRetrieveRequest](../../models/operations/officesretrieverequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.OfficesRetrieveSecurity](../../models/operations/officesretrievesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.OfficesRetrieveResponse](../../models/operations/officesretrieveresponse.md), error**

