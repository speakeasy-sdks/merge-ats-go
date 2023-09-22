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
        XAccountToken: "minima",
        ApplicationID: mergeatsgo.String("aspernatur"),
        CreatedAfter: types.MustTimeFromString("2022-01-08T02:07:04.894Z"),
        CreatedBefore: types.MustTimeFromString("2021-04-04T14:18:49.875Z"),
        Cursor: mergeatsgo.String("error"),
        Expand: operations.ScorecardsListExpandApplicationInterviewer.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        InterviewID: mergeatsgo.String("suscipit"),
        InterviewerID: mergeatsgo.String("repudiandae"),
        ModifiedAfter: types.MustTimeFromString("2021-12-01T19:00:30.130Z"),
        ModifiedBefore: types.MustTimeFromString("2022-01-28T23:50:19.904Z"),
        PageSize: mergeatsgo.Int64(680697),
        RemoteFields: mergeatsgo.String("repellendus"),
        RemoteID: mergeatsgo.String("labore"),
        ShowEnumOrigins: mergeatsgo.String("reiciendis"),
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
        XAccountToken: "doloremque",
        Expand: operations.ScorecardsRetrieveExpandInterviewer.ToPointer(),
        ID: "1012563f-94e2-49e9-b3e9-22a57a15be3e",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: mergeatsgo.String("doloremque"),
        ShowEnumOrigins: mergeatsgo.String("iure"),
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

