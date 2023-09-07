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
        XAccountToken: "minima",
        CreatedAfter: types.MustTimeFromString("2022-10-08T01:09:40.281Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-21T17:17:20.623Z"),
        Cursor: ats.String("temporibus"),
        Expand: operations.JobInterviewStagesListExpandJob.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("accusantium"),
        ModifiedAfter: types.MustTimeFromString("2022-12-20T14:34:53.149Z"),
        ModifiedBefore: types.MustTimeFromString("2022-02-21T23:29:55.837Z"),
        PageSize: ats.Int64(649832),
        RemoteID: ats.String("ab"),
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
    xAccountToken := "corrupti"
    id := "40394c26-071f-493f-9f06-42dac7af515c"
    expand := operations.JobInterviewStagesRetrieveExpandJob
    includeRemoteData := false
    operationSecurity := operations.JobInterviewStagesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.JobInterviewStages.Retrieve(ctx, operationSecurity, xAccountToken, id, expand, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobInterviewStage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `security`                                                                                                             | [operations.JobInterviewStagesRetrieveSecurity](../../models/operations/jobinterviewstagesretrievesecurity.md)         | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |
| `xAccountToken`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Token identifying the end user.                                                                                        |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `expand`                                                                                                               | [*operations.JobInterviewStagesRetrieveExpand](../../models/operations/jobinterviewstagesretrieveexpand.md)            | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
| `includeRemoteData`                                                                                                    | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                       |


### Response

**[*operations.JobInterviewStagesRetrieveResponse](../../models/operations/jobinterviewstagesretrieveresponse.md), error**

