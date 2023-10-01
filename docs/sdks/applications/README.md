# Applications
(*Applications*)

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
    applicationEndpointRequest := shared.ApplicationEndpointRequest{
        Model: shared.ApplicationRequest{
            AppliedAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
            Candidate: mergeatsgo.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
            CreditedTo: mergeatsgo.String("58166795-8d68-4b30-9bfb-bfd402479484"),
            CurrentStage: mergeatsgo.String("d578dfdc-7b0a-4ab6-a2b0-4b40f20eb9ea"),
            IntegrationParams: map[string]interface{}{
                "odio": "bluetooth",
            },
            Job: mergeatsgo.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            LinkedAccountParams: map[string]interface{}{
                "nulla": "Money",
            },
            RejectReason: mergeatsgo.String("59b25f2b-da02-40f5-9656-9fa0db555784"),
            RejectedAt: types.MustTimeFromString("2021-11-15T00:00:00Z"),
            RemoteTemplateID: mergeatsgo.String("92830948203"),
            Source: mergeatsgo.String("Campus recruiting event"),
        },
        RemoteUserID: "Cambridgeshire grey technology",
    }
    xAccountToken := "East"
    isDebugMode := false
    runAsync := false

    ctx := context.Background()
    res, err := s.Applications.Create(ctx, applicationEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `applicationEndpointRequest`                                                           | [shared.ApplicationEndpointRequest](../../models/shared/applicationendpointrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `xAccountToken`                                                                        | *string*                                                                               | :heavy_check_mark:                                                                     | Token identifying the end user.                                                        |
| `isDebugMode`                                                                          | **bool*                                                                                | :heavy_minus_sign:                                                                     | Whether to include debug fields (such as log file links) in the response.              |
| `runAsync`                                                                             | **bool*                                                                                | :heavy_minus_sign:                                                                     | Whether or not third-party updates should be run asynchronously.                       |


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
    res, err := s.Applications.List(ctx, operations.ApplicationsListRequest{
        XAccountToken: "Northeast Metal Canada",
        CandidateID: mergeatsgo.String("Data Response West"),
        CreatedAfter: types.MustTimeFromString("2023-02-26T13:00:25.189Z"),
        CreatedBefore: types.MustTimeFromString("2023-12-20T19:28:33.339Z"),
        CreditedToID: mergeatsgo.String("primary"),
        CurrentStageID: mergeatsgo.String("Designer hacking"),
        Cursor: mergeatsgo.String("synergies backing"),
        Expand: operations.ApplicationsListExpandCandidateJobCreditedToRejectReason.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        JobID: mergeatsgo.String("virtual"),
        ModifiedAfter: types.MustTimeFromString("2021-03-22T04:25:28.253Z"),
        ModifiedBefore: types.MustTimeFromString("2022-09-19T00:01:59.827Z"),
        PageSize: mergeatsgo.Int64(931165),
        RejectReasonID: mergeatsgo.String("accusantium defensive"),
        RemoteID: mergeatsgo.String("green Smart"),
        Source: mergeatsgo.String("North"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedApplicationList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ApplicationsListRequest](../../models/operations/applicationslistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    xAccountToken := "till"
    id := "56591081-ad20-4d60-8c8e-92b241fa3790"
    expand := operations.ApplicationsRetrieveExpandCreditedTo
    includeRemoteData := false

    ctx := context.Background()
    res, err := s.Applications.Retrieve(ctx, xAccountToken, id, expand, includeRemoteData)
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
    xAccountToken := "Borders"
    applicationRemoteTemplateID := "Home"

    ctx := context.Background()
    res, err := s.Applications.RetrievePostMetadata(ctx, xAccountToken, applicationRemoteTemplateID)
    if err != nil {
        log.Fatal(err)
    }

    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `xAccountToken`                                                        | *string*                                                               | :heavy_check_mark:                                                     | Token identifying the end user.                                        |
| `applicationRemoteTemplateID`                                          | **string*                                                              | :heavy_minus_sign:                                                     | The template ID associated with the nested application in the request. |


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
    res, err := s.Applications.UpdateChangeState(ctx, operations.ApplicationsChangeStageCreateRequest{
        UpdateApplicationStageRequest: &shared.UpdateApplicationStageRequest{
            JobInterviewStage: mergeatsgo.String("c8893fee-92d0-4f72-9f7b-bfee92506581"),
            RemoteUserID: mergeatsgo.String("Nissan Dollar"),
        },
        XAccountToken: "Internal",
        ID: "ed279139-548e-4b48-9cd8-42f688a13cc6",
        IsDebugMode: mergeatsgo.Bool(false),
        RunAsync: mergeatsgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ApplicationsChangeStageCreateRequest](../../models/operations/applicationschangestagecreaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.ApplicationsChangeStageCreateResponse](../../models/operations/applicationschangestagecreateresponse.md), error**

