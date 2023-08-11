# Applications

### Available Operations

* [Create](#create) - Creates an `Application` object with the given values.
* [CreateChangeState](#createchangestate) - Updates the `current_stage` field of an `Application` object
* [List](#list) - Returns a list of `Application` objects.
* [Retrieve](#retrieve) - Returns an `Application` object with the given `id`.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `Application` POSTs.

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
    operationSecurity := operations.ApplicationsCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.Create(ctx, operations.ApplicationsCreateRequest{
        ApplicationEndpointRequest: shared.ApplicationEndpointRequest{
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
        },
        XAccountToken: "iure",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ApplicationsCreateRequest](../../models/operations/applicationscreaterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ApplicationsCreateSecurity](../../models/operations/applicationscreatesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ApplicationsCreateResponse](../../models/operations/applicationscreateresponse.md), error**


## CreateChangeState

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
    res, err := s.Applications.CreateChangeState(ctx, operations.ApplicationsChangeStageCreateRequest{
        UpdateApplicationStageRequest: &shared.UpdateApplicationStageRequest{
            JobInterviewStage: ats.String("eb10faaa-2352-4c59-9590-7aff1a3a2fa9"),
            RemoteUserID: ats.String("numquam"),
        },
        XAccountToken: "commodi",
        ID: "7739251a-a52c-43f5-ad01-9da1ffe78f09",
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
        XAccountToken: "voluptate",
        CandidateID: ats.String("cum"),
        CreatedAfter: types.MustTimeFromString("2022-12-17T16:42:52.927Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-19T18:36:39.009Z"),
        CreditedToID: ats.String("maiores"),
        CurrentStageID: ats.String("dicta"),
        Cursor: ats.String("corporis"),
        Expand: operations.ApplicationsListExpandCandidateJobCreditedToCurrentStage.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("iusto"),
        ModifiedAfter: types.MustTimeFromString("2022-04-24T15:19:40.519Z"),
        ModifiedBefore: types.MustTimeFromString("2022-02-13T15:01:52.114Z"),
        PageSize: ats.Int64(414263),
        RejectReasonID: ats.String("repudiandae"),
        RemoteID: ats.String("quae"),
        Source: ats.String("ipsum"),
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
    operationSecurity := operations.ApplicationsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.Retrieve(ctx, operations.ApplicationsRetrieveRequest{
        XAccountToken: "quidem",
        Expand: operations.ApplicationsRetrieveExpandCreditedToCurrentStage.ToPointer(),
        ID: "9d488e1e-91e4-450a-92ab-d44269802d50",
        IncludeRemoteData: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Application != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ApplicationsRetrieveRequest](../../models/operations/applicationsretrieverequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.ApplicationsRetrieveSecurity](../../models/operations/applicationsretrievesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


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
    operationSecurity := operations.ApplicationsMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Applications.RetrievePostMetadata(ctx, operations.ApplicationsMetaPostRetrieveRequest{
        XAccountToken: "fugit",
        ApplicationRemoteTemplateID: ats.String("dolorum"),
    }, operationSecurity)
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
| `request`                                                                                                          | [operations.ApplicationsMetaPostRetrieveRequest](../../models/operations/applicationsmetapostretrieverequest.md)   | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `security`                                                                                                         | [operations.ApplicationsMetaPostRetrieveSecurity](../../models/operations/applicationsmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |


### Response

**[*operations.ApplicationsMetaPostRetrieveResponse](../../models/operations/applicationsmetapostretrieveresponse.md), error**

