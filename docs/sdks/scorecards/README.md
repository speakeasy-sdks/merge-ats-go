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
        XAccountToken: "repudiandae",
        ApplicationID: ats.String("nam"),
        CreatedAfter: types.MustTimeFromString("2022-07-08T17:52:09.255Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-22T04:21:35.753Z"),
        Cursor: ats.String("dignissimos"),
        Expand: operations.ScorecardsListExpandApplicationInterview.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        InterviewID: ats.String("quo"),
        InterviewerID: ats.String("deleniti"),
        ModifiedAfter: types.MustTimeFromString("2021-09-08T14:14:28.598Z"),
        ModifiedBefore: types.MustTimeFromString("2022-01-10T03:05:21.296Z"),
        PageSize: ats.Int64(426904),
        RemoteFields: operations.ScorecardsListRemoteFieldsOverallRecommendation.ToPointer(),
        RemoteID: ats.String("magnam"),
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
        XAccountToken: "quibusdam",
        Expand: operations.ScorecardsRetrieveExpandApplication.ToPointer(),
        ID: "db1f2c43-1066-41e9-a349-e1cf9e06e3a4",
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

