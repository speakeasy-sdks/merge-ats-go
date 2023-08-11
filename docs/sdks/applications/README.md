# Applications

### Available Operations

* [Create](#create) - Creates an `Application` object with the given values.
* [List](#list) - Returns a list of `Application` objects.
* [Retrieve](#retrieve) - Returns an `Application` object with the given `id`.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `Application` POSTs.
* [UpdateChangeState](#updatechangestate) - Updates the `current_stage` field of an `Application` object

## Create

Creates an `Application` object with the given values.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    applicationEndpointRequest := shared.ApplicationEndpointRequest{
        Model: shared.ApplicationRequest{
            AppliedAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
            Candidate: ats.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
            CreditedTo: ats.String("58166795-8d68-4b30-9bfb-bfd402479484"),
            CurrentStage: ats.String("d578dfdc-7b0a-4ab6-a2b0-4b40f20eb9ea"),
            IntegrationParams: map[string]interface{}{
                "natus": "sed",
                "iste": "dolor",
            },
            Job: ats.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            LinkedAccountParams: map[string]interface{}{
                "laboriosam": "hic",
                "saepe": "fuga",
                "in": "corporis",
            },
            RejectReason: ats.String("59b25f2b-da02-40f5-9656-9fa0db555784"),
            RejectedAt: types.MustTimeFromString("2021-11-15T00:00:00Z"),
            RemoteTemplateID: ats.String("92830948203"),
            Source: ats.String("Campus recruiting event"),
        },
        RemoteUserID: "iste",
    }
    xAccountToken := "iure"
    isDebugMode := false
    runAsync := false
    operationSecurity := operations.ApplicationsCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.Create(ctx, operationSecurity, applicationEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `security`                                                                                     | [operations.ApplicationsCreateSecurity](../../models/operations/applicationscreatesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `applicationEndpointRequest`                                                                   | [shared.ApplicationEndpointRequest](../../models/shared/applicationendpointrequest.md)         | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `xAccountToken`                                                                                | *string*                                                                                       | :heavy_check_mark:                                                                             | Token identifying the end user.                                                                |
| `isDebugMode`                                                                                  | **bool*                                                                                        | :heavy_minus_sign:                                                                             | Whether to include debug fields (such as log file links) in the response.                      |
| `runAsync`                                                                                     | **bool*                                                                                        | :heavy_minus_sign:                                                                             | Whether or not third-party updates should be run asynchronously.                               |


### Response

**[*operations.ApplicationsCreateResponse](../../models/operations/applicationscreateresponse.md), error**


## List

Returns a list of `Application` objects.

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
    operationSecurity := operations.ApplicationsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.List(ctx, operations.ApplicationsListRequest{
        XAccountToken: "saepe",
        CandidateID: ats.String("quidem"),
        CreatedAfter: types.MustTimeFromString("2022-12-10T00:25:28.749Z"),
        CreatedBefore: types.MustTimeFromString("2020-12-31T21:22:14.646Z"),
        CreditedToID: ats.String("mollitia"),
        CurrentStageID: ats.String("laborum"),
        Cursor: ats.String("dolores"),
        Expand: operations.ApplicationsListExpandCandidateCurrentStageRejectReason.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("corporis"),
        ModifiedAfter: types.MustTimeFromString("2022-04-01T23:59:21.675Z"),
        ModifiedBefore: types.MustTimeFromString("2022-05-24T03:24:11.703Z"),
        PageSize: ats.Int64(363711),
        RejectReasonID: ats.String("minima"),
        RemoteID: ats.String("excepturi"),
        Source: ats.String("accusantium"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedApplicationList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ApplicationsListRequest](../../models/operations/applicationslistrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ApplicationsListSecurity](../../models/operations/applicationslistsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ApplicationsListResponse](../../models/operations/applicationslistresponse.md), error**


## Retrieve

Returns an `Application` object with the given `id`.

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
    xAccountToken := "iure"
    id := "aff1a3a2-fa94-4677-b925-1aa52c3f5ad0"
    expand := operations.ApplicationsRetrieveExpandCandidateCreditedToCurrentStageRejectReason
    includeRemoteData := false
    operationSecurity := operations.ApplicationsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.Retrieve(ctx, operationSecurity, xAccountToken, id, expand, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }

    if res.Application != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `security`                                                                                                             | [operations.ApplicationsRetrieveSecurity](../../models/operations/applicationsretrievesecurity.md)                     | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |
| `xAccountToken`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Token identifying the end user.                                                                                        |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `expand`                                                                                                               | [*operations.ApplicationsRetrieveExpand](../../models/operations/applicationsretrieveexpand.md)                        | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
| `includeRemoteData`                                                                                                    | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                       |


### Response

**[*operations.ApplicationsRetrieveResponse](../../models/operations/applicationsretrieveresponse.md), error**


## RetrievePostMetadata

Returns metadata for `Application` POSTs.

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
    xAccountToken := "error"
    applicationRemoteTemplateID := "temporibus"
    operationSecurity := operations.ApplicationsMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.RetrievePostMetadata(ctx, operationSecurity, xAccountToken, applicationRemoteTemplateID)
    if err != nil {
        log.Fatal(err)
    }

    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `security`                                                                                                         | [operations.ApplicationsMetaPostRetrieveSecurity](../../models/operations/applicationsmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |
| `xAccountToken`                                                                                                    | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | Token identifying the end user.                                                                                    |
| `applicationRemoteTemplateID`                                                                                      | **string*                                                                                                          | :heavy_minus_sign:                                                                                                 | The template ID associated with the nested application in the request.                                             |


### Response

**[*operations.ApplicationsMetaPostRetrieveResponse](../../models/operations/applicationsmetapostretrieveresponse.md), error**


## UpdateChangeState

Updates the `current_stage` field of an `Application` object

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    operationSecurity := operations.ApplicationsChangeStageCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.UpdateChangeState(ctx, operations.ApplicationsChangeStageCreateRequest{
        UpdateApplicationStageRequest: &shared.UpdateApplicationStageRequest{
            JobInterviewStage: ats.String("a1ffe78f-097b-4007-8f15-471b5e6e13b9"),
            RemoteUserID: ats.String("excepturi"),
        },
        XAccountToken: "pariatur",
        ID: "488e1e91-e450-4ad2-abd4-4269802d502a",
        IsDebugMode: ats.Bool(false),
        RunAsync: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ApplicationsChangeStageCreateRequest](../../models/operations/applicationschangestagecreaterequest.md)   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.ApplicationsChangeStageCreateSecurity](../../models/operations/applicationschangestagecreatesecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.ApplicationsChangeStageCreateResponse](../../models/operations/applicationschangestagecreateresponse.md), error**

