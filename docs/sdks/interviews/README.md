# Interviews
(*Interviews*)

### Available Operations

* [Create](#create) - Creates a `ScheduledInterview` object with the given values.
* [List](#list) - Returns a list of `ScheduledInterview` objects.
* [Retrieve](#retrieve) - Returns a `ScheduledInterview` object with the given `id`.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `ScheduledInterview` POSTs.

## Create

Creates a `ScheduledInterview` object with the given values.

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
    scheduledInterviewEndpointRequest := shared.ScheduledInterviewEndpointRequest{
        Model: shared.ScheduledInterviewRequest{
            Application: mergeatsgo.String("92e8a369-fffe-430d-b93a-f7e8a16563f1"),
            EndAt: types.MustTimeFromString("2021-10-15T02:00:00Z"),
            IntegrationParams: map[string]interface{}{
                "odio": "bluetooth",
            },
            Interviewers: []string{
                "d642c1fc-6fe0-4724-9bcd-d89dc7fa504e",
            },
            JobInterviewStage: mergeatsgo.String("2f7adb59-3fe6-4b5b-aef6-563f72bd13dc"),
            LinkedAccountParams: map[string]interface{}{
                "perferendis": "overriding",
            },
            Location: mergeatsgo.String("Embarcadero Center 2"),
            Organizer: mergeatsgo.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            StartAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
            Status: shared.ScheduledInterviewRequestStatusScheduled.ToPointer(),
        },
        RemoteUserID: "Fish",
    }
    xAccountToken := "Buckinghamshire"
    isDebugMode := false
    runAsync := false

    ctx := context.Background()
    res, err := s.Interviews.Create(ctx, scheduledInterviewEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }

    if res.ScheduledInterviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `scheduledInterviewEndpointRequest`                                                                  | [shared.ScheduledInterviewEndpointRequest](../../models/shared/scheduledinterviewendpointrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `xAccountToken`                                                                                      | *string*                                                                                             | :heavy_check_mark:                                                                                   | Token identifying the end user.                                                                      |
| `isDebugMode`                                                                                        | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Whether to include debug fields (such as log file links) in the response.                            |
| `runAsync`                                                                                           | **bool*                                                                                              | :heavy_minus_sign:                                                                                   | Whether or not third-party updates should be run asynchronously.                                     |


### Response

**[*operations.InterviewsCreateResponse](../../models/operations/interviewscreateresponse.md), error**


## List

Returns a list of `ScheduledInterview` objects.

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
    res, err := s.Interviews.List(ctx, operations.InterviewsListRequest{
        XAccountToken: "Northeast Metal Canada",
        ApplicationID: mergeatsgo.String("Data Response West"),
        CreatedAfter: types.MustTimeFromString("2023-02-26T13:00:25.189Z"),
        CreatedBefore: types.MustTimeFromString("2023-12-20T19:28:33.339Z"),
        Cursor: mergeatsgo.String("primary"),
        Expand: operations.InterviewsListExpandInterviewersOrganizerApplicationJobInterviewStage.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        JobInterviewStageID: mergeatsgo.String("generate orchid synergies"),
        ModifiedAfter: types.MustTimeFromString("2022-07-07T22:30:40.342Z"),
        ModifiedBefore: types.MustTimeFromString("2021-03-01T19:04:17.136Z"),
        OrganizerID: mergeatsgo.String("empowering virtual"),
        PageSize: mergeatsgo.Int64(73227),
        RemoteFields: operations.InterviewsListRemoteFieldsStatus.ToPointer(),
        RemoteID: mergeatsgo.String("gah accusantium"),
        ShowEnumOrigins: operations.InterviewsListShowEnumOriginsStatus.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedScheduledInterviewList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.InterviewsListRequest](../../models/operations/interviewslistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.InterviewsListResponse](../../models/operations/interviewslistresponse.md), error**


## Retrieve

Returns a `ScheduledInterview` object with the given `id`.

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
    res, err := s.Interviews.Retrieve(ctx, operations.InterviewsRetrieveRequest{
        XAccountToken: "tracksuit Markets",
        Expand: operations.InterviewsRetrieveExpandApplicationJobInterviewStage.ToPointer(),
        ID: "081ad20d-604c-48e9-ab24-1fa379087a15",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: operations.InterviewsRetrieveRemoteFieldsStatus.ToPointer(),
        ShowEnumOrigins: operations.InterviewsRetrieveShowEnumOriginsStatus.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScheduledInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.InterviewsRetrieveRequest](../../models/operations/interviewsretrieverequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.InterviewsRetrieveResponse](../../models/operations/interviewsretrieveresponse.md), error**


## RetrievePostMetadata

Returns metadata for `ScheduledInterview` POSTs.

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

    ctx := context.Background()
    res, err := s.Interviews.RetrievePostMetadata(ctx, xAccountToken)
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

**[*operations.InterviewsMetaPostRetrieveResponse](../../models/operations/interviewsmetapostretrieveresponse.md), error**

