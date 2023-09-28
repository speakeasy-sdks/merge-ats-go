# JobInterviewStages
(*JobInterviewStages*)

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
    res, err := s.JobInterviewStages.List(ctx, operations.JobInterviewStagesListRequest{
        XAccountToken: "voluptatem",
        CreatedAfter: types.MustTimeFromString("2022-06-03T05:52:14.954Z"),
        CreatedBefore: types.MustTimeFromString("2022-03-25T13:15:35.424Z"),
        Cursor: mergeatsgo.String("explicabo"),
        Expand: mergeatsgo.String("voluptas"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        JobID: mergeatsgo.String("aut"),
        ModifiedAfter: types.MustTimeFromString("2022-11-19T20:21:26.456Z"),
        ModifiedBefore: types.MustTimeFromString("2021-02-22T18:20:18.826Z"),
        PageSize: mergeatsgo.Int64(244651),
        RemoteID: mergeatsgo.String("voluptatibus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedJobInterviewStageList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.JobInterviewStagesListRequest](../../models/operations/jobinterviewstageslistrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


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
    xAccountToken := "voluptas"
    id := "f0642dac-7af5-415c-8413-aa63aae8d678"
    expand := "vel"
    includeRemoteData := false

    ctx := context.Background()
    res, err := s.JobInterviewStages.Retrieve(ctx, xAccountToken, id, expand, includeRemoteData)
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
| `xAccountToken`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Token identifying the end user.                                                                                        |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `expand`                                                                                                               | **string*                                                                                                              | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
| `includeRemoteData`                                                                                                    | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                       |


### Response

**[*operations.JobInterviewStagesRetrieveResponse](../../models/operations/jobinterviewstagesretrieveresponse.md), error**

