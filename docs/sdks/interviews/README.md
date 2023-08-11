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
    operationSecurity := operations.InterviewsCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Interviews.Create(ctx, operations.InterviewsCreateRequest{
        ScheduledInterviewEndpointRequest: shared.ScheduledInterviewEndpointRequest{
            Model: shared.ScheduledInterviewRequest{
                Application: ats.String("92e8a369-fffe-430d-b93a-f7e8a16563f1"),
                EndAt: types.MustTimeFromString("2021-10-15T02:00:00Z"),
                IntegrationParams: map[string]interface{}{
                    "quaerat": "molestiae",
                    "ex": "ut",
                    "culpa": "adipisci",
                },
                Interviewers: []string{
                    "865e7956-f925-41a5-a9da-660ff57bfaad",
                    "4f9efc1b-4512-4c10-b264-8dc2f615199e",
                    "bfd0e9fe-6c63-42ca-baed-0117996312fd",
                    "e0477177-8ff6-41d0-9747-6360a15db6a6",
                },
                JobInterviewStage: ats.String("2f7adb59-3fe6-4b5b-aef6-563f72bd13dc"),
                LinkedAccountParams: map[string]interface{}{
                    "perferendis": "eum",
                    "voluptas": "iste",
                },
                Location: ats.String("Embarcadero Center 2"),
                Organizer: ats.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
                StartAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
                Status: shared.ScheduledInterviewRequestStatusScheduled.ToPointer(),
            },
            RemoteUserID: "id",
        },
        XAccountToken: "ab",
        IsDebugMode: ats.Bool(false),
        RunAsync: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ScheduledInterviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.InterviewsCreateRequest](../../models/operations/interviewscreaterequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.InterviewsCreateSecurity](../../models/operations/interviewscreatesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
        XAccountToken: "error",
        ApplicationID: ats.String("possimus"),
        CreatedAfter: types.MustTimeFromString("2021-01-15T12:05:30.894Z"),
        CreatedBefore: types.MustTimeFromString("2021-07-19T18:28:47.948Z"),
        Cursor: ats.String("ad"),
        Expand: operations.InterviewsListExpandInterviewersOrganizerApplicationJobInterviewStage.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobInterviewStageID: ats.String("enim"),
        ModifiedAfter: types.MustTimeFromString("2022-03-02T20:09:12.443Z"),
        ModifiedBefore: types.MustTimeFromString("2022-03-23T19:57:49.518Z"),
        OrganizerID: ats.String("ex"),
        PageSize: ats.Int64(281153),
        RemoteFields: operations.InterviewsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("ad"),
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
        XAccountToken: "expedita",
        Expand: operations.InterviewsRetrieveExpandApplication.ToPointer(),
        ID: "8b61891b-aa0f-4e1a-9e00-8e6f8c5f350d",
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
    operationSecurity := operations.InterviewsMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Interviews.RetrievePostMetadata(ctx, operations.InterviewsMetaPostRetrieveRequest{
        XAccountToken: "totam",
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
| `request`                                                                                                      | [operations.InterviewsMetaPostRetrieveRequest](../../models/operations/interviewsmetapostretrieverequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.InterviewsMetaPostRetrieveSecurity](../../models/operations/interviewsmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.InterviewsMetaPostRetrieveResponse](../../models/operations/interviewsmetapostretrieveresponse.md), error**

