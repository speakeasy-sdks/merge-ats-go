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
)

func main() {
    s := mergeatsgo.New()

    ctx := context.Background()
    res, err := s.JobInterviewStages.List(ctx, operations.JobInterviewStagesListRequest{
        XAccountToken: "string",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.JobInterviewStagesListRequest](../../pkg/models/operations/jobinterviewstageslistrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.JobInterviewStagesListResponse](../../pkg/models/operations/jobinterviewstageslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
    s := mergeatsgo.New()


    var xAccountToken string = "string"

    var id string = "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    var expand *operations.JobInterviewStagesRetrieveQueryParamExpand = operations.JobInterviewStagesRetrieveQueryParamExpandJob

    var includeRemoteData *bool = false

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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `xAccountToken`                                                                                                                        | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | Token identifying the end user.                                                                                                        |
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `expand`                                                                                                                               | [*operations.JobInterviewStagesRetrieveQueryParamExpand](../../../pkg/models/operations/jobinterviewstagesretrievequeryparamexpand.md) | :heavy_minus_sign:                                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.                 |
| `includeRemoteData`                                                                                                                    | **bool*                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                                       |


### Response

**[*operations.JobInterviewStagesRetrieveResponse](../../pkg/models/operations/jobinterviewstagesretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
