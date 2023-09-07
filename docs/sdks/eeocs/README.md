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
        XAccountToken: "dolorum",
        CandidateID: ats.String("laborum"),
        CreatedAfter: types.MustTimeFromString("2022-04-07T07:44:57.988Z"),
        CreatedBefore: types.MustTimeFromString("2022-07-31T12:04:26.954Z"),
        Cursor: ats.String("nobis"),
        Expand: operations.EeocsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-05-05T09:22:07.208Z"),
        ModifiedBefore: types.MustTimeFromString("2021-11-11T22:59:32.230Z"),
        PageSize: ats.Int64(727044),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatusGender.ToPointer(),
        RemoteID: ats.String("tempora"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsDisabilityStatusGenderRaceVeteranStatus.ToPointer(),
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
        XAccountToken: "explicabo",
        Expand: operations.EeocsRetrieveExpandCandidate.ToPointer(),
        ID: "90747477-8a7b-4d46-ad28-c10ab3cdca42",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsDisabilityStatusRace.ToPointer(),
        ShowEnumOrigins: operations.EeocsRetrieveShowEnumOriginsDisabilityStatusGender.ToPointer(),
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

