# Eeocs
(*Eeocs*)

### Available Operations

* [List](#list) - Returns a list of `EEOC` objects.
* [Retrieve](#retrieve) - Returns an `EEOC` object with the given `id`.

## List

Returns a list of `EEOC` objects.

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
    res, err := s.Eeocs.List(ctx, operations.EeocsListRequest{
        XAccountToken: "nulla",
        CandidateID: mergeatsgo.String("voluptas"),
        CreatedAfter: types.MustTimeFromString("2022-10-22T12:27:01.822Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-29T12:13:01.368Z"),
        Cursor: mergeatsgo.String("explicabo"),
        Expand: mergeatsgo.String("provident"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-07-11T02:03:34.375Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-06T21:41:34.842Z"),
        PageSize: mergeatsgo.Int64(262118),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatusRaceVeteranStatus.ToPointer(),
        RemoteID: mergeatsgo.String("esse"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsDisabilityStatusVeteranStatus.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedEEOCList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.EeocsListRequest](../../models/operations/eeocslistrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.EeocsListResponse](../../models/operations/eeocslistresponse.md), error**


## Retrieve

Returns an `EEOC` object with the given `id`.

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
    res, err := s.Eeocs.Retrieve(ctx, operations.EeocsRetrieveRequest{
        XAccountToken: "fuga",
        Expand: mergeatsgo.String("reprehenderit"),
        ID: "bd466d28-c10a-4b3c-9ca4-251904e523c7",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsRaceVeteranStatus.ToPointer(),
        ShowEnumOrigins: operations.EeocsRetrieveShowEnumOriginsDisabilityStatus.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Eeoc != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.EeocsRetrieveRequest](../../models/operations/eeocsretrieverequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.EeocsRetrieveResponse](../../models/operations/eeocsretrieveresponse.md), error**

