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
        XAccountToken: "quidem",
        CandidateID: ats.String("reprehenderit"),
        CreatedAfter: types.MustTimeFromString("2020-12-11T19:46:56.272Z"),
        CreatedBefore: types.MustTimeFromString("2021-09-14T12:33:27.065Z"),
        Cursor: ats.String("veniam"),
        Expand: operations.EeocsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-03-18T08:14:24.399Z"),
        ModifiedBefore: types.MustTimeFromString("2022-09-16T12:27:20.507Z"),
        PageSize: ats.Int64(542457),
        RemoteFields: operations.EeocsListRemoteFieldsDisabilityStatusRaceVeteranStatus.ToPointer(),
        RemoteID: ats.String("asperiores"),
        ShowEnumOrigins: operations.EeocsListShowEnumOriginsDisabilityStatusVeteranStatus.ToPointer(),
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
        XAccountToken: "suscipit",
        Expand: operations.EeocsRetrieveExpandCandidate.ToPointer(),
        ID: "bc173d68-9eee-4952-af8d-986e881ead4f",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.EeocsRetrieveRemoteFieldsDisabilityStatus.ToPointer(),
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

