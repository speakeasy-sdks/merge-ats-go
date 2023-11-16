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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    applicationEndpointRequest := shared.ApplicationEndpointRequest{
        Model: shared.ApplicationRequest{
            AppliedAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
            Candidate: mergeatsgo.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
            CreditedTo: mergeatsgo.String("58166795-8d68-4b30-9bfb-bfd402479484"),
            CurrentStage: mergeatsgo.String("d578dfdc-7b0a-4ab6-a2b0-4b40f20eb9ea"),
            IntegrationParams: map[string]interface{}{
                "unique_integration_field": "string",
            },
            Job: mergeatsgo.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            LinkedAccountParams: map[string]interface{}{
                "unique_linked_account_field": "string",
            },
            RejectReason: mergeatsgo.String("59b25f2b-da02-40f5-9656-9fa0db555784"),
            RejectedAt: types.MustTimeFromString("2021-11-15T00:00:00Z"),
            RemoteTemplateID: mergeatsgo.String("92830948203"),
            Source: mergeatsgo.String("Campus recruiting event"),
        },
        RemoteUserID: "string",
    }

    var xAccountToken string = "string"

    var isDebugMode *bool = false

    var runAsync *bool = false

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

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `applicationEndpointRequest`                                                                  | [shared.ApplicationEndpointRequest](../../../pkg/models/shared/applicationendpointrequest.md) | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `xAccountToken`                                                                               | *string*                                                                                      | :heavy_check_mark:                                                                            | Token identifying the end user.                                                               |
| `isDebugMode`                                                                                 | **bool*                                                                                       | :heavy_minus_sign:                                                                            | Whether to include debug fields (such as log file links) in the response.                     |
| `runAsync`                                                                                    | **bool*                                                                                       | :heavy_minus_sign:                                                                            | Whether or not third-party updates should be run asynchronously.                              |


### Response

**[*operations.ApplicationsCreateResponse](../../pkg/models/operations/applicationscreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

Returns a list of `Application` objects.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Applications.List(ctx, operations.ApplicationsListRequest{
        XAccountToken: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ApplicationsListRequest](../../pkg/models/operations/applicationslistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ApplicationsListResponse](../../pkg/models/operations/applicationslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Retrieve

Returns an `Application` object with the given `id`.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "string"

    var id string = "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    var expand *operations.ApplicationsRetrieveQueryParamExpand = operations.ApplicationsRetrieveQueryParamExpandCandidateCurrentStageRejectReason

    var includeRemoteData *bool = false

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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `xAccountToken`                                                                                                            | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | Token identifying the end user.                                                                                            |
| `id`                                                                                                                       | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `expand`                                                                                                                   | [*operations.ApplicationsRetrieveQueryParamExpand](../../../pkg/models/operations/applicationsretrievequeryparamexpand.md) | :heavy_minus_sign:                                                                                                         | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.     |
| `includeRemoteData`                                                                                                        | **bool*                                                                                                                    | :heavy_minus_sign:                                                                                                         | Whether to include the original data Merge fetched from the third-party to produce these models.                           |


### Response

**[*operations.ApplicationsRetrieveResponse](../../pkg/models/operations/applicationsretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RetrievePostMetadata

Returns metadata for `Application` POSTs.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "string"

    var applicationRemoteTemplateID *string = "string"

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

**[*operations.ApplicationsMetaPostRetrieveResponse](../../pkg/models/operations/applicationsmetapostretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateChangeState

Updates the `current_stage` field of an `Application` object

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Applications.UpdateChangeState(ctx, operations.ApplicationsChangeStageCreateRequest{
        UpdateApplicationStageRequest: &shared.UpdateApplicationStageRequest{},
        XAccountToken: "string",
        ID: "c8893fee-92d0-4f72-9f7b-bfee92506581",
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ApplicationsChangeStageCreateRequest](../../pkg/models/operations/applicationschangestagecreaterequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.ApplicationsChangeStageCreateResponse](../../pkg/models/operations/applicationschangestagecreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
