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
        XAccountToken: "quos",
        CandidateID: ats.String("voluptatibus"),
        CreatedAfter: types.MustTimeFromString("2022-09-23T08:26:30.493Z"),
        CreatedBefore: types.MustTimeFromString("2022-01-11T22:08:32.388Z"),
        Cursor: ats.String("ex"),
        Expand: operations.EeocsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-10-02T07:53:52.364Z"),
        ModifiedBefore: types.MustTimeFromString("2021-06-26T01:49:52.614Z"),
        PageSize: ats.Int64(708609),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatusGenderVeteranStatus.ToPointer(),
        RemoteID: ats.String("incidunt"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsDisabilityStatusRace.ToPointer(),
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
        XAccountToken: "debitis",
        Expand: operations.EeocsRetrieveExpandCandidate.ToPointer(),
        ID: "80ca55ef-d20e-4457-a185-8b6a89fbe3a5",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsGenderRace.ToPointer(),
        ShowEnumOrigins: operations.EeocsRetrieveShowEnumOriginsGenderRaceVeteranStatus.ToPointer(),
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

