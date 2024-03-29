# Activities
(*Activities*)

### Available Operations

* [Create](#create) - Creates an `Activity` object with the given values.
* [List](#list) - Returns a list of `Activity` objects.
* [Retrieve](#retrieve) - Returns an `Activity` object with the given `id`.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `Activity` POSTs.

## Create

Creates an `Activity` object with the given values.

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
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    activityEndpointRequest := shared.ActivityEndpointRequest{
        Model: shared.ActivityRequest{
            ActivityType: shared.ActivityRequestActivityTypeNote.ToPointer(),
            Body: mergeatsgo.String("Candidate loves integrations!!."),
            Candidate: mergeatsgo.String("03455bc6-6040-430a-848e-aafacbfdf4fg"),
            IntegrationParams: map[string]interface{}{
                "unique_integration_field": "unique_integration_field_value",
            },
            LinkedAccountParams: map[string]interface{}{
                "unique_linked_account_field": "unique_linked_account_field_value",
            },
            Subject: mergeatsgo.String("Gil Feig's interview"),
            User: mergeatsgo.String("9d892439-5fab-4dbb-8bd8-34f7f96c7912"),
            Visibility: shared.ActivityRequestVisibilityPrivate.ToPointer(),
        },
        RemoteUserID: "<value>",
    }

    var xAccountToken string = "<value>"

    var isDebugMode *bool = mergeatsgo.Bool(false)

    var runAsync *bool = mergeatsgo.Bool(false)

    ctx := context.Background()
    res, err := s.Activities.Create(ctx, activityEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }
    if res.ActivityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `activityEndpointRequest`                                                            | [shared.ActivityEndpointRequest](../../pkg/models/shared/activityendpointrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `xAccountToken`                                                                      | *string*                                                                             | :heavy_check_mark:                                                                   | Token identifying the end user.                                                      |
| `isDebugMode`                                                                        | **bool*                                                                              | :heavy_minus_sign:                                                                   | Whether to include debug fields (such as log file links) in the response.            |
| `runAsync`                                                                           | **bool*                                                                              | :heavy_minus_sign:                                                                   | Whether or not third-party updates should be run asynchronously.                     |


### Response

**[*operations.ActivitiesCreateResponse](../../pkg/models/operations/activitiescreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Returns a list of `Activity` objects.

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
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Activities.List(ctx, operations.ActivitiesListRequest{
        XAccountToken: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedActivityList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ActivitiesListRequest](../../pkg/models/operations/activitieslistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ActivitiesListResponse](../../pkg/models/operations/activitieslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Returns an `Activity` object with the given `id`.

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
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Activities.Retrieve(ctx, operations.ActivitiesRetrieveRequest{
        XAccountToken: "<value>",
        ID: "5fea5659-1081-4ad2-8d60-4c8e92b241fa",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Activity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ActivitiesRetrieveRequest](../../pkg/models/operations/activitiesretrieverequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ActivitiesRetrieveResponse](../../pkg/models/operations/activitiesretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RetrievePostMetadata

Returns metadata for `Activity` POSTs.

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
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var xAccountToken string = "<value>"

    ctx := context.Background()
    res, err := s.Activities.RetrievePostMetadata(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xAccountToken`                                       | *string*                                              | :heavy_check_mark:                                    | Token identifying the end user.                       |


### Response

**[*operations.ActivitiesMetaPostRetrieveResponse](../../pkg/models/operations/activitiesmetapostretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
