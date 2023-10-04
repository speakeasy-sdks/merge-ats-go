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
        XAccountToken: "Northeast Metal Canada",
        CandidateID: mergeatsgo.String("Data Response West"),
        CreatedAfter: types.MustTimeFromString("2023-02-26T13:00:25.189Z"),
        CreatedBefore: types.MustTimeFromString("2023-12-20T19:28:33.339Z"),
        Cursor: mergeatsgo.String("primary"),
        Expand: mergeatsgo.String("Designer hacking"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-11-26T03:10:34.253Z"),
        ModifiedBefore: types.MustTimeFromString("2021-12-18T09:50:13.895Z"),
        PageSize: mergeatsgo.Int64(7468),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatusVeteranStatus.ToPointer(),
        RemoteID: mergeatsgo.String("explicit"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsDisabilityStatusGenderVeteranStatus.ToPointer(),
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
        XAccountToken: "tracksuit Markets",
        Expand: mergeatsgo.String("McKinney"),
        ID: "1ad20d60-4c8e-492b-a41f-a379087a1539",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsGenderVeteranStatus.ToPointer(),
        ShowEnumOrigins: operations.EeocsRetrieveShowEnumOriginsDisabilityStatusVeteranStatus.ToPointer(),
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

