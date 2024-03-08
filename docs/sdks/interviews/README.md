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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    scheduledInterviewEndpointRequest := shared.ScheduledInterviewEndpointRequest{
        Model: shared.ScheduledInterviewRequest{
            Application: mergeatsgo.String("92e8a369-fffe-430d-b93a-f7e8a16563f1"),
            EndAt: types.MustNewTimeFromString("2021-10-15T02:00:00Z"),
            IntegrationParams: map[string]interface{}{
                "unique_integration_field": "unique_integration_field_value",
            },
            Interviewers: []string{
                "f9813dd5-e70b-484c-91d8-00acd6065b07",
                "89a86fcf-d540-4e6b-ac3d-ce07c4ec9b3c",
            },
            JobInterviewStage: mergeatsgo.String("2f7adb59-3fe6-4b5b-aef6-563f72bd13dc"),
            LinkedAccountParams: map[string]interface{}{
                "unique_linked_account_field": "unique_linked_account_field_value",
            },
            Location: mergeatsgo.String("Embarcadero Center 2"),
            Organizer: mergeatsgo.String("52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633"),
            StartAt: types.MustNewTimeFromString("2021-10-15T00:00:00Z"),
            Status: shared.ScheduledInterviewRequestStatusScheduled.ToPointer(),
        },
        RemoteUserID: "<value>",
    }

    var xAccountToken string = "<value>"

    var isDebugMode *bool = mergeatsgo.Bool(false)

    var runAsync *bool = mergeatsgo.Bool(false)

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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `scheduledInterviewEndpointRequest`                                                                      | [shared.ScheduledInterviewEndpointRequest](../../pkg/models/shared/scheduledinterviewendpointrequest.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `xAccountToken`                                                                                          | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Token identifying the end user.                                                                          |
| `isDebugMode`                                                                                            | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | Whether to include debug fields (such as log file links) in the response.                                |
| `runAsync`                                                                                               | **bool*                                                                                                  | :heavy_minus_sign:                                                                                       | Whether or not third-party updates should be run asynchronously.                                         |


### Response

**[*operations.InterviewsCreateResponse](../../pkg/models/operations/interviewscreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Returns a list of `ScheduledInterview` objects.

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
    res, err := s.Interviews.List(ctx, operations.InterviewsListRequest{
        XAccountToken: "<value>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.InterviewsListRequest](../../pkg/models/operations/interviewslistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.InterviewsListResponse](../../pkg/models/operations/interviewslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Returns a `ScheduledInterview` object with the given `id`.

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
    res, err := s.Interviews.Retrieve(ctx, operations.InterviewsRetrieveRequest{
        XAccountToken: "<value>",
        ID: "5fea5659-1081-4ad2-8d60-4c8e92b241fa",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.InterviewsRetrieveRequest](../../pkg/models/operations/interviewsretrieverequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.InterviewsRetrieveResponse](../../pkg/models/operations/interviewsretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RetrievePostMetadata

Returns metadata for `ScheduledInterview` POSTs.

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

**[*operations.InterviewsMetaPostRetrieveResponse](../../pkg/models/operations/interviewsmetapostretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
