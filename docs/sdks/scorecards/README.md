# Scorecards

### Available Operations

* [List](#list) - Returns a list of `Scorecard` objects.
* [Retrieve](#retrieve) - Returns a `Scorecard` object with the given `id`.

## List

Returns a list of `Scorecard` objects.

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
    operationSecurity := operations.ScorecardsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Scorecards.List(ctx, operations.ScorecardsListRequest{
        XAccountToken: "maiores",
        ApplicationID: ats.String("enim"),
        CreatedAfter: types.MustTimeFromString("2021-04-14T02:13:07.391Z"),
        CreatedBefore: types.MustTimeFromString("2022-01-31T07:01:27.539Z"),
        Cursor: ats.String("nemo"),
        Expand: operations.ScorecardsListExpandApplicationInterviewer.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        InterviewID: ats.String("est"),
        InterviewerID: ats.String("quis"),
        ModifiedAfter: types.MustTimeFromString("2021-03-29T02:31:09.447Z"),
        ModifiedBefore: types.MustTimeFromString("2020-03-02T04:41:46.760Z"),
        PageSize: ats.Int64(900103),
        RemoteFields: operations.ScorecardsListRemoteFieldsOverallRecommendation.ToPointer(),
        RemoteID: ats.String("asperiores"),
        ShowEnumOrigins: operations.ScorecardsListShowEnumOriginsOverallRecommendation.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedScorecardList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ScorecardsListRequest](../../models/operations/scorecardslistrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ScorecardsListSecurity](../../models/operations/scorecardslistsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ScorecardsListResponse](../../models/operations/scorecardslistresponse.md), error**


## Retrieve

Returns a `Scorecard` object with the given `id`.

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
    operationSecurity := operations.ScorecardsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Scorecards.Retrieve(ctx, operations.ScorecardsRetrieveRequest{
        XAccountToken: "ex",
        Expand: operations.ScorecardsRetrieveExpandApplicationInterviewInterviewer.ToPointer(),
        ID: "ef1caa33-83c2-4beb-8773-73c8d72f64d1",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.ScorecardsRetrieveRemoteFieldsOverallRecommendation.ToPointer(),
        ShowEnumOrigins: operations.ScorecardsRetrieveShowEnumOriginsOverallRecommendation.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Scorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ScorecardsRetrieveRequest](../../models/operations/scorecardsretrieverequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ScorecardsRetrieveSecurity](../../models/operations/scorecardsretrievesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ScorecardsRetrieveResponse](../../models/operations/scorecardsretrieveresponse.md), error**

