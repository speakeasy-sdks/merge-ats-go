# Scorecards
(*Scorecards*)

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
    res, err := s.Scorecards.List(ctx, operations.ScorecardsListRequest{
        XAccountToken: "Northeast Metal Canada",
        ApplicationID: mergeatsgo.String("Data Response West"),
        CreatedAfter: types.MustTimeFromString("2023-02-26T13:00:25.189Z"),
        CreatedBefore: types.MustTimeFromString("2023-12-20T19:28:33.339Z"),
        Cursor: mergeatsgo.String("primary"),
        Expand: operations.ScorecardsListExpandApplicationInterviewer.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        InterviewID: mergeatsgo.String("generate orchid synergies"),
        InterviewerID: mergeatsgo.String("Guyana empowering"),
        ModifiedAfter: types.MustTimeFromString("2022-06-16T20:57:25.031Z"),
        ModifiedBefore: types.MustTimeFromString("2021-12-06T18:55:39.498Z"),
        PageSize: mergeatsgo.Int64(73227),
        RemoteFields: mergeatsgo.String("gah accusantium"),
        RemoteID: mergeatsgo.String("tan green Smart"),
        ShowEnumOrigins: mergeatsgo.String("North"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedScorecardList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ScorecardsListRequest](../../models/operations/scorecardslistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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

    ctx := context.Background()
    res, err := s.Scorecards.Retrieve(ctx, operations.ScorecardsRetrieveRequest{
        XAccountToken: "tracksuit Markets",
        Expand: operations.ScorecardsRetrieveExpandApplication.ToPointer(),
        ID: "081ad20d-604c-48e9-ab24-1fa379087a15",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: mergeatsgo.String("Mobility"),
        ShowEnumOrigins: mergeatsgo.String("recklessly program"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Scorecard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ScorecardsRetrieveRequest](../../models/operations/scorecardsretrieverequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ScorecardsRetrieveResponse](../../models/operations/scorecardsretrieveresponse.md), error**

