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
    operationSecurity := operations.ActivitiesCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Activities.Create(ctx, operations.ActivitiesCreateRequest{
        ActivityEndpointRequest: shared.ActivityEndpointRequest{
            Model: shared.ActivityRequest{
                ActivityType: shared.ActivityRequestActivityTypeNote.ToPointer(),
                Body: ats.String("Candidate loves integrations!!."),
                Candidate: ats.String("03455bc6-6040-430a-848e-aafacbfdf4fg"),
                IntegrationParams: map[string]interface{}{
                    "unde": "nulla",
                    "corrupti": "illum",
                    "vel": "error",
                    "deserunt": "suscipit",
                },
                LinkedAccountParams: map[string]interface{}{
                    "magnam": "debitis",
                    "ipsa": "delectus",
                },
                Subject: ats.String("Gil Feig's interview"),
                User: ats.String("9d892439-5fab-4dbb-8bd8-34f7f96c7912"),
                Visibility: shared.ActivityRequestVisibilityPrivate.ToPointer(),
            },
            RemoteUserID: "tempora",
        },
        XAccountToken: "suscipit",
        IsDebugMode: ats.Bool(false),
        RunAsync: ats.Bool(false),
    }, operationSecurity)
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
| `request`                                                                                  | [operations.ActivitiesCreateRequest](../../models/operations/activitiescreaterequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ActivitiesCreateSecurity](../../models/operations/activitiescreatesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
        XAccountToken: "molestiae",
        CreatedAfter: types.MustTimeFromString("2020-07-25T16:12:20.938Z"),
        CreatedBefore: types.MustTimeFromString("2022-01-15T14:47:59.325Z"),
        Cursor: ats.String("excepturi"),
        Expand: operations.ActivitiesListExpandUser.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-01-28T03:46:24.500Z"),
        ModifiedBefore: types.MustTimeFromString("2022-10-15T05:10:19.629Z"),
        PageSize: ats.Int64(337396),
        RemoteFields: operations.ActivitiesListRemoteFieldsActivityType.ToPointer(),
        RemoteID: ats.String("deserunt"),
        ShowEnumOrigins: operations.ActivitiesListShowEnumOriginsActivityType.ToPointer(),
        UserID: ats.String("ipsam"),
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
        XAccountToken: "repellendus",
        Expand: operations.ActivitiesRetrieveExpandUser.ToPointer(),
        ID: "fc2ddf7c-c78c-4a1b-a928-fc816742cb73",
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
    operationSecurity := operations.ActivitiesMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Activities.RetrievePostMetadata(ctx, operations.ActivitiesMetaPostRetrieveRequest{
        XAccountToken: "perferendis",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ActivitiesMetaPostRetrieveRequest](../../models/operations/activitiesmetapostretrieverequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.ActivitiesMetaPostRetrieveSecurity](../../models/operations/activitiesmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.ActivitiesMetaPostRetrieveResponse](../../models/operations/activitiesmetapostretrieveresponse.md), error**

