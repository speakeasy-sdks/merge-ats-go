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
        XAccountToken: "odit",
        ApplicationID: ats.String("repellat"),
        CreatedAfter: types.MustTimeFromString("2021-11-29T15:18:11.704Z"),
        CreatedBefore: types.MustTimeFromString("2022-12-15T11:44:44.229Z"),
        Cursor: ats.String("odio"),
        Expand: operations.ScorecardsListExpandApplicationInterviewInterviewer.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        InterviewID: ats.String("in"),
        InterviewerID: ats.String("ducimus"),
        ModifiedAfter: types.MustTimeFromString("2022-08-28T07:07:51.380Z"),
        ModifiedBefore: types.MustTimeFromString("2022-10-30T21:28:00.704Z"),
        PageSize: ats.Int64(498180),
        RemoteFields: operations.ScorecardsListRemoteFieldsOverallRecommendation.ToPointer(),
        RemoteID: ats.String("voluptate"),
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
        XAccountToken: "pariatur",
        Expand: operations.ScorecardsRetrieveExpandInterviewer.ToPointer(),
        ID: "ac646ecb-5734-409e-beb1-e5a2b12eb07f",
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

