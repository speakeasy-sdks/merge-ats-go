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
        XAccountToken: "eveniet",
        CandidateID: ats.String("occaecati"),
        CreatedAfter: types.MustTimeFromString("2022-11-08T18:10:37.954Z"),
        CreatedBefore: types.MustTimeFromString("2022-04-30T23:44:42.233Z"),
        Cursor: ats.String("reprehenderit"),
        Expand: operations.EeocsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-11-06T03:57:05.427Z"),
        ModifiedBefore: types.MustTimeFromString("2022-04-21T18:58:57.960Z"),
        PageSize: ats.Int64(910994),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatusGenderRaceVeteranStatus.ToPointer(),
        RemoteID: ats.String("vero"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsDisabilityStatus.ToPointer(),
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
        XAccountToken: "iure",
        Expand: operations.EeocsRetrieveExpandCandidate.ToPointer(),
        ID: "0807e2b6-e3ab-4884-9f05-97a60ff2a54a",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsDisabilityStatusGenderRaceVeteranStatus.ToPointer(),
        ShowEnumOrigins: operations.EeocsRetrieveShowEnumOriginsDisabilityStatus.ToPointer(),
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

