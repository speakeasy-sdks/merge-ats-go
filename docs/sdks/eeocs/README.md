# Eeocs

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.EeocsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Eeocs.List(ctx, operations.EeocsListRequest{
        XAccountToken: "eius",
        CandidateID: ats.String("necessitatibus"),
        CreatedAfter: types.MustTimeFromString("2022-08-05T13:00:56.741Z"),
        CreatedBefore: types.MustTimeFromString("2021-11-23T23:35:18.899Z"),
        Cursor: ats.String("voluptatibus"),
        Expand: operations.EeocsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-09-23T08:26:30.493Z"),
        ModifiedBefore: types.MustTimeFromString("2022-01-11T22:08:32.388Z"),
        PageSize: ats.Int64(401713),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatus.ToPointer(),
        RemoteID: ats.String("non"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsRaceVeteranStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedEEOCList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.EeocsListRequest](../../models/operations/eeocslistrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.EeocsListSecurity](../../models/operations/eeocslistsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.EeocsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Eeocs.Retrieve(ctx, operations.EeocsRetrieveRequest{
        XAccountToken: "praesentium",
        Expand: operations.EeocsRetrieveExpandCandidate.ToPointer(),
        ID: "b445e80c-a55e-4fd2-8e45-7e1858b6a89f",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsGenderRaceVeteranStatus.ToPointer(),
        ShowEnumOrigins: operations.EeocsRetrieveShowEnumOriginsRaceVeteranStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Eeoc != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.EeocsRetrieveRequest](../../models/operations/eeocsretrieverequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.EeocsRetrieveSecurity](../../models/operations/eeocsretrievesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.EeocsRetrieveResponse](../../models/operations/eeocsretrieveresponse.md), error**

