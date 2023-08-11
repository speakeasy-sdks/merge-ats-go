# JobInterviewStages

### Available Operations

* [List](#list) - Returns a list of `JobInterviewStage` objects.
* [Retrieve](#retrieve) - Returns a `JobInterviewStage` object with the given `id`.

## List

Returns a list of `JobInterviewStage` objects.

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
    operationSecurity := operations.JobInterviewStagesListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.JobInterviewStages.List(ctx, operations.JobInterviewStagesListRequest{
        XAccountToken: "iusto",
        CreatedAfter: types.MustTimeFromString("2022-06-29T10:30:00.284Z"),
        CreatedBefore: types.MustTimeFromString("2020-02-06T20:25:31.567Z"),
        Cursor: ats.String("vel"),
        Expand: operations.JobInterviewStagesListExpandJob.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("architecto"),
        ModifiedAfter: types.MustTimeFromString("2022-11-18T13:59:08.408Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-08T10:43:52.536Z"),
        PageSize: ats.Int64(271252),
        RemoteID: ats.String("esse"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedJobInterviewStageList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.JobInterviewStagesListRequest](../../models/operations/jobinterviewstageslistrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.JobInterviewStagesListSecurity](../../models/operations/jobinterviewstageslistsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.JobInterviewStagesListResponse](../../models/operations/jobinterviewstageslistresponse.md), error**


## Retrieve

Returns a `JobInterviewStage` object with the given `id`.

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
    operationSecurity := operations.JobInterviewStagesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.JobInterviewStages.Retrieve(ctx, operations.JobInterviewStagesRetrieveRequest{
        XAccountToken: "ex",
        Expand: operations.JobInterviewStagesRetrieveExpandJob.ToPointer(),
        ID: "360a15db-6a66-4065-9a1a-deaab5851d6c",
        IncludeRemoteData: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobInterviewStage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.JobInterviewStagesRetrieveRequest](../../models/operations/jobinterviewstagesretrieverequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.JobInterviewStagesRetrieveSecurity](../../models/operations/jobinterviewstagesretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.JobInterviewStagesRetrieveResponse](../../models/operations/jobinterviewstagesretrieveresponse.md), error**

