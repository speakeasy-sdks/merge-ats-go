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
        XAccountToken: "earum",
        CreatedAfter: types.MustTimeFromString("2021-08-28T09:50:26.086Z"),
        CreatedBefore: types.MustTimeFromString("2022-03-14T23:12:21.252Z"),
        Cursor: ats.String("voluptatibus"),
        Expand: operations.JobInterviewStagesListExpandJob.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        JobID: ats.String("molestias"),
        ModifiedAfter: types.MustTimeFromString("2020-02-18T03:48:05.478Z"),
        ModifiedBefore: types.MustTimeFromString("2022-08-29T17:35:23.458Z"),
        PageSize: ats.Int64(698249),
        RemoteID: ats.String("tempora"),
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
    xAccountToken := "quis"
    id := "12c10326-48dc-42f6-9519-9ebfd0e9fe6c"
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

