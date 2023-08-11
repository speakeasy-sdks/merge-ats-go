# Interviews

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    scheduledInterviewEndpointRequest := shared.ScheduledInterviewEndpointRequest{
        Model: shared.ScheduledInterviewRequest{
            Application: ats.String("92e8a369-fffe-430d-b93a-f7e8a16563f1"),
            EndAt: types.MustTimeFromString("2021-10-15T02:00:00Z"),
            IntegrationParams: map[string]interface{}{
                "nostrum": "officia",
                "dolorum": "corrupti",
                "accusamus": "tempora",
            },
            Interviewers: []string{
                "24d0ab40-7508-48e5-9862-065e904f3b11",
                "94b8abf6-03a7-49f9-9fe0-ab7da8a50ce1",
                "87f86bc1-73d6-489e-ae95-26f8d986e881",
            },
            JobInterviewStage: ats.String("2f7adb59-3fe6-4b5b-aef6-563f72bd13dc"),
            LinkedAccountParams: map[string]interface{}{
                "dolorum": "repellendus",
                "labore": "reiciendis",
                "doloremque": "repudiandae",
                "dicta": "accusantium",
            },
            Location: ats.String("Embarcadero Center 2"),
            Organizer: ats.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            StartAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
            Status: shared.ScheduledInterviewRequestStatusScheduled.ToPointer(),
        },
        RemoteUserID: "beatae",
    }
    xAccountToken := "dolores"
    isDebugMode := false
    runAsync := false
    operationSecurity := operations.InterviewsCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Interviews.Create(ctx, operationSecurity, scheduledInterviewEndpointRequest, xAccountToken, isDebugMode, runAsync)
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
| `security`                                                                                           | [operations.InterviewsCreateSecurity](../../models/operations/interviewscreatesecurity.md)           | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.InterviewsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Interviews.List(ctx, operations.InterviewsListRequest{
        XAccountToken: "enim",
        ApplicationID: ats.String("laboriosam"),
        CreatedAfter: types.MustTimeFromString("2022-01-18T11:13:47.798Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-26T23:29:35.541Z"),
        Cursor: ats.String("saepe"),
        Expand: operations.InterviewsListExpandInterviewers.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobInterviewStageID: ats.String("occaecati"),
        ModifiedAfter: types.MustTimeFromString("2021-03-17T06:13:21.698Z"),
        ModifiedBefore: types.MustTimeFromString("2022-10-06T02:44:18.773Z"),
        OrganizerID: ats.String("eveniet"),
        PageSize: ats.Int64(580887),
        RemoteFields: operations.InterviewsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("consequuntur"),
        ShowEnumOrigins: operations.InterviewsListShowEnumOriginsStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedScheduledInterviewList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.InterviewsListRequest](../../models/operations/interviewslistrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.InterviewsListSecurity](../../models/operations/interviewslistsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.InterviewsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Interviews.Retrieve(ctx, operations.InterviewsRetrieveRequest{
        XAccountToken: "fugit",
        Expand: operations.InterviewsRetrieveExpandInterviewersOrganizerJobInterviewStage.ToPointer(),
        ID: "57a15be3-e060-4807-a2b6-e3ab8845f059",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.InterviewsRetrieveRemoteFieldsStatus.ToPointer(),
        ShowEnumOrigins: operations.InterviewsRetrieveShowEnumOriginsStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ScheduledInterview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.InterviewsRetrieveRequest](../../models/operations/interviewsretrieverequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.InterviewsRetrieveSecurity](../../models/operations/interviewsretrievesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "nihil"
    operationSecurity := operations.InterviewsMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Interviews.RetrievePostMetadata(ctx, operationSecurity, xAccountToken)
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
| `security`                                                                                                     | [operations.InterviewsMetaPostRetrieveSecurity](../../models/operations/interviewsmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `xAccountToken`                                                                                                | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Token identifying the end user.                                                                                |


### Response

**[*operations.InterviewsMetaPostRetrieveResponse](../../models/operations/interviewsmetapostretrieveresponse.md), error**
