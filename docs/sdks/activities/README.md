# Activities

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
    activityEndpointRequest := shared.ActivityEndpointRequest{
        Model: shared.ActivityRequest{
            ActivityType: shared.ActivityRequestActivityTypeNote.ToPointer(),
            Body: mergeatsgo.String("Candidate loves integrations!!."),
            Candidate: mergeatsgo.String("03455bc6-6040-430a-848e-aafacbfdf4fg"),
            IntegrationParams: map[string]interface{}{
                "unde": "nulla",
            },
            LinkedAccountParams: map[string]interface{}{
                "corrupti": "illum",
            },
            Subject: mergeatsgo.String("Gil Feig's interview"),
            User: mergeatsgo.String("9d892439-5fab-4dbb-8bd8-34f7f96c7912"),
            Visibility: shared.ActivityRequestVisibilityPrivate.ToPointer(),
        },
        RemoteUserID: "vel",
    }
    xAccountToken := "error"
    isDebugMode := false
    runAsync := false

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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `activityEndpointRequest`                                                        | [shared.ActivityEndpointRequest](../../models/shared/activityendpointrequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `xAccountToken`                                                                  | *string*                                                                         | :heavy_check_mark:                                                               | Token identifying the end user.                                                  |
| `isDebugMode`                                                                    | **bool*                                                                          | :heavy_minus_sign:                                                               | Whether to include debug fields (such as log file links) in the response.        |
| `runAsync`                                                                       | **bool*                                                                          | :heavy_minus_sign:                                                               | Whether or not third-party updates should be run asynchronously.                 |


### Response

**[*operations.ActivitiesCreateResponse](../../models/operations/activitiescreateresponse.md), error**


## List

Returns a list of `Activity` objects.

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
    res, err := s.Activities.List(ctx, operations.ActivitiesListRequest{
        XAccountToken: "deserunt",
        CreatedAfter: types.MustTimeFromString("2022-07-25T06:44:09.184Z"),
        CreatedBefore: types.MustTimeFromString("2022-02-09T12:04:06.508Z"),
        Cursor: mergeatsgo.String("ipsa"),
        Expand: mergeatsgo.String("delectus"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-08-14T01:03:07.567Z"),
        ModifiedBefore: types.MustTimeFromString("2022-03-18T00:29:19.137Z"),
        PageSize: mergeatsgo.Int64(812169),
        RemoteFields: operations.ActivitiesListRemoteFieldsActivityTypeVisibility.ToPointer(),
        RemoteID: mergeatsgo.String("iusto"),
        ShowEnumOrigins: operations.ActivitiesListShowEnumOriginsActivityTypeVisibility.ToPointer(),
        UserID: mergeatsgo.String("nisi"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ActivitiesListRequest](../../models/operations/activitieslistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ActivitiesListResponse](../../models/operations/activitieslistresponse.md), error**


## Retrieve

Returns an `Activity` object with the given `id`.

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
    res, err := s.Activities.Retrieve(ctx, operations.ActivitiesRetrieveRequest{
        XAccountToken: "recusandae",
        Expand: mergeatsgo.String("temporibus"),
        ID: "151a05df-c2dd-4f7c-878c-a1ba928fc816",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: operations.ActivitiesRetrieveRemoteFieldsActivityTypeVisibility.ToPointer(),
        ShowEnumOrigins: operations.ActivitiesRetrieveShowEnumOriginsActivityType.ToPointer(),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ActivitiesRetrieveRequest](../../models/operations/activitiesretrieverequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ActivitiesRetrieveResponse](../../models/operations/activitiesretrieveresponse.md), error**


## RetrievePostMetadata

Returns metadata for `Activity` POSTs.

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
    xAccountToken := "qui"

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

**[*operations.ActivitiesMetaPostRetrieveResponse](../../models/operations/activitiesmetapostretrieveresponse.md), error**

