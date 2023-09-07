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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    activityEndpointRequest := shared.ActivityEndpointRequest{
        Model: shared.ActivityRequest{
            ActivityType: shared.ActivityRequestActivityTypeNote.ToPointer(),
            Body: ats.String("Candidate loves integrations!!."),
            Candidate: ats.String("03455bc6-6040-430a-848e-aafacbfdf4fg"),
            IntegrationParams: map[string]interface{}{
                "quibusdam": "unde",
            },
            LinkedAccountParams: map[string]interface{}{
                "nulla": "corrupti",
            },
            Subject: ats.String("Gil Feig's interview"),
            User: ats.String("9d892439-5fab-4dbb-8bd8-34f7f96c7912"),
            Visibility: shared.ActivityRequestVisibilityPrivate.ToPointer(),
        },
        RemoteUserID: "illum",
    }
    xAccountToken := "vel"
    isDebugMode := false
    runAsync := false
    operationSecurity := operations.ActivitiesCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Activities.Create(ctx, operationSecurity, activityEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }

    if res.ActivityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `security`                                                                                 | [operations.ActivitiesCreateSecurity](../../models/operations/activitiescreatesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `activityEndpointRequest`                                                                  | [shared.ActivityEndpointRequest](../../models/shared/activityendpointrequest.md)           | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `xAccountToken`                                                                            | *string*                                                                                   | :heavy_check_mark:                                                                         | Token identifying the end user.                                                            |
| `isDebugMode`                                                                              | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Whether to include debug fields (such as log file links) in the response.                  |
| `runAsync`                                                                                 | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Whether or not third-party updates should be run asynchronously.                           |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.ActivitiesListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Activities.List(ctx, operations.ActivitiesListRequest{
        XAccountToken: "error",
        CreatedAfter: types.MustTimeFromString("2022-03-26T09:37:56.283Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-14T09:35:47.986Z"),
        Cursor: ats.String("debitis"),
        Expand: operations.ActivitiesListExpandUser.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-01-14T06:18:51.036Z"),
        ModifiedBefore: types.MustTimeFromString("2022-08-14T01:03:07.567Z"),
        PageSize: ats.Int64(477665),
        RemoteFields: operations.ActivitiesListRemoteFieldsVisibility.ToPointer(),
        RemoteID: ats.String("placeat"),
        ShowEnumOrigins: operations.ActivitiesListShowEnumOriginsActivityTypeVisibility.ToPointer(),
        UserID: ats.String("iusto"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedActivityList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ActivitiesListRequest](../../models/operations/activitieslistrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ActivitiesListSecurity](../../models/operations/activitieslistsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.ActivitiesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Activities.Retrieve(ctx, operations.ActivitiesRetrieveRequest{
        XAccountToken: "excepturi",
        Expand: operations.ActivitiesRetrieveExpandUser.ToPointer(),
        ID: "6ed151a0-5dfc-42dd-b7cc-78ca1ba928fc",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.ActivitiesRetrieveRemoteFieldsActivityTypeVisibility.ToPointer(),
        ShowEnumOrigins: operations.ActivitiesRetrieveShowEnumOriginsActivityType.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Activity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ActivitiesRetrieveRequest](../../models/operations/activitiesretrieverequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ActivitiesRetrieveSecurity](../../models/operations/activitiesretrievesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "commodi"
    operationSecurity := operations.ActivitiesMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Activities.RetrievePostMetadata(ctx, operationSecurity, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `security`                                                                                                     | [operations.ActivitiesMetaPostRetrieveSecurity](../../models/operations/activitiesmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `xAccountToken`                                                                                                | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Token identifying the end user.                                                                                |


### Response

**[*operations.ActivitiesMetaPostRetrieveResponse](../../models/operations/activitiesmetapostretrieveresponse.md), error**

