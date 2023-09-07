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
                "molestiae": "modi",
            },
            Job: ats.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            LinkedAccountParams: map[string]interface{}{
                "qui": "impedit",
            },
            RejectReason: ats.String("59b25f2b-da02-40f5-9656-9fa0db555784"),
            RejectedAt: types.MustTimeFromString("2021-11-15T00:00:00Z"),
            RemoteTemplateID: ats.String("92830948203"),
            Source: ats.String("Campus recruiting event"),
        },
        RemoteUserID: "cum",
    }
    xAccountToken := "esse"
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
        XAccountToken: "ipsum",
        CandidateID: ats.String("excepturi"),
        CreatedAfter: types.MustTimeFromString("2022-12-25T03:24:03.949Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-20T13:30:46.463Z"),
        CreditedToID: ats.String("sed"),
        CurrentStageID: ats.String("iste"),
        Cursor: ats.String("dolor"),
        Expand: operations.ApplicationsListExpandCreditedToRejectReason.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("laboriosam"),
        ModifiedAfter: types.MustTimeFromString("2020-04-17T15:42:43.722Z"),
        ModifiedBefore: types.MustTimeFromString("2022-02-06T12:52:33.708Z"),
        PageSize: ats.Int64(359508),
        RejectReasonID: ats.String("iste"),
        RemoteID: ats.String("iure"),
        Source: ats.String("saepe"),
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
    xAccountToken := "quidem"
    id := "10faaa23-52c5-4955-907a-ff1a3a2fa946"
    expand := operations.ApplicationsRetrieveExpandCandidateJobRejectReason
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
    xAccountToken := "molestiae"
    applicationRemoteTemplateID := "velit"
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
            JobInterviewStage: ats.String("9251aa52-c3f5-4ad0-99da-1ffe78f097b0"),
            RemoteUserID: ats.String("doloremque"),
        },
        XAccountToken: "reprehenderit",
        ID: "4f15471b-5e6e-413b-99d4-88e1e91e450a",
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

