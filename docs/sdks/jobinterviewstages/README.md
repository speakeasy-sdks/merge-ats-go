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
        XAccountToken: "exercitationem",
        CreatedAfter: types.MustTimeFromString("2022-09-18T07:06:33.190Z"),
        CreatedBefore: types.MustTimeFromString("2022-02-17T01:08:30.498Z"),
        Cursor: ats.String("vero"),
        Expand: operations.JobInterviewStagesListExpandJob.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("est"),
        ModifiedAfter: types.MustTimeFromString("2022-08-13T07:34:51.264Z"),
        ModifiedBefore: types.MustTimeFromString("2020-03-31T16:32:36.866Z"),
        PageSize: ats.Int64(759283),
        RemoteID: ats.String("occaecati"),
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
        XAccountToken: "nemo",
        Expand: operations.JobInterviewStagesRetrieveExpandJob.ToPointer(),
        ID: "78a64584-273a-4841-8d16-2309fb092992",
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

