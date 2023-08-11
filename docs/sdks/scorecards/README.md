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
        XAccountToken: "ex",
        ApplicationID: ats.String("nemo"),
        CreatedAfter: types.MustTimeFromString("2021-07-19T20:29:58.626Z"),
        CreatedBefore: types.MustTimeFromString("2021-08-24T00:23:56.434Z"),
        Cursor: ats.String("odio"),
        Expand: operations.ScorecardsListExpandApplicationInterview.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        InterviewID: ats.String("alias"),
        InterviewerID: ats.String("magni"),
        ModifiedAfter: types.MustTimeFromString("2022-12-08T11:32:42.651Z"),
        ModifiedBefore: types.MustTimeFromString("2022-09-26T12:26:41.334Z"),
        PageSize: ats.Int64(208253),
        RemoteFields: operations.ScorecardsListRemoteFieldsOverallRecommendation.ToPointer(),
        RemoteID: ats.String("exercitationem"),
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
        XAccountToken: "itaque",
        Expand: operations.ScorecardsRetrieveExpandApplication.ToPointer(),
        ID: "39dbc225-9b1a-4bda-8c07-0e1084cb0672",
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

