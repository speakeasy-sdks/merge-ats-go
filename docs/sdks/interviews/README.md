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
                    "beatae": "dolores",
                },
                Interviewers: []string{
                    "63f94e29-e973-4e92-aa57-a15be3e06080",
                    "7e2b6e3a-b884-45f0-997a-60ff2a54a31e",
                },
                JobInterviewStage: ats.String("2f7adb59-3fe6-4b5b-aef6-563f72bd13dc"),
                LinkedAccountParams: map[string]interface{}{
                    "quaerat": "molestiae",
                    "ex": "ut",
                    "culpa": "adipisci",
                },
                Location: ats.String("Embarcadero Center 2"),
                Organizer: ats.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
                StartAt: types.MustTimeFromString("2021-10-15T00:00:00Z"),
                Status: shared.ScheduledInterviewRequestStatusScheduled.ToPointer(),
            },
            RemoteUserID: "debitis",
        },
        XAccountToken: "laudantium",
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
        XAccountToken: "eum",
        ApplicationID: ats.String("nemo"),
        CreatedAfter: types.MustTimeFromString("2021-08-19T02:39:25.517Z"),
        CreatedBefore: types.MustTimeFromString("2022-04-29T15:23:46.044Z"),
        Cursor: ats.String("eum"),
        Expand: operations.InterviewsListExpandOrganizerJobInterviewStage.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobInterviewStageID: ats.String("provident"),
        ModifiedAfter: types.MustTimeFromString("2022-08-24T14:31:36.261Z"),
        ModifiedBefore: types.MustTimeFromString("2022-05-06T04:30:23.999Z"),
        OrganizerID: ats.String("nostrum"),
        PageSize: ats.Int64(652309),
        RemoteFields: operations.InterviewsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("provident"),
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
        XAccountToken: "possimus",
        Expand: operations.InterviewsRetrieveExpandInterviewersOrganizerJobInterviewStage.ToPointer(),
        ID: "660ff57b-faad-44f9-afc1-b4512c103264",
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
        XAccountToken: "rem",
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

