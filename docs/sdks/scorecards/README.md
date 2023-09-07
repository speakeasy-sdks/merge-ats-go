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
        XAccountToken: "praesentium",
        ApplicationID: ats.String("mollitia"),
        CreatedAfter: types.MustTimeFromString("2022-12-21T09:05:01.168Z"),
        CreatedBefore: types.MustTimeFromString("2020-03-30T02:42:49.718Z"),
        Cursor: ats.String("quasi"),
        Expand: operations.ScorecardsListExpandApplicationInterviewer.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        InterviewID: ats.String("reprehenderit"),
        InterviewerID: ats.String("asperiores"),
        ModifiedAfter: types.MustTimeFromString("2022-03-27T08:02:14.031Z"),
        ModifiedBefore: types.MustTimeFromString("2021-05-22T03:09:11.884Z"),
        PageSize: ats.Int64(90885),
        RemoteFields: operations.ScorecardsListRemoteFieldsOverallRecommendation.ToPointer(),
        RemoteID: ats.String("esse"),
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
        XAccountToken: "amet",
        Expand: operations.ScorecardsRetrieveExpandInterviewInterviewer.ToPointer(),
        ID: "689eee95-26f8-4d98-ae88-1ead4f0e1012",
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

