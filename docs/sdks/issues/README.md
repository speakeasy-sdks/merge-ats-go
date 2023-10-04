# Issues
(*Issues*)

### Available Operations

* [List](#list) - Gets issues.
* [Retrieve](#retrieve) - Get a specific issue.

## List

Gets issues.

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
    res, err := s.Issues.List(ctx, operations.IssuesListRequest{
        AccountToken: mergeatsgo.String("Northeast Metal Canada"),
        Cursor: mergeatsgo.String("Data Response West"),
        EndDate: mergeatsgo.String("boil primary synthesize"),
        EndUserOrganizationName: mergeatsgo.String("hacking Paradigm"),
        FirstIncidentTimeAfter: types.MustTimeFromString("2021-01-09T04:15:41.822Z"),
        FirstIncidentTimeBefore: types.MustTimeFromString("2022-07-07T22:30:40.342Z"),
        IncludeMuted: mergeatsgo.String("explicit"),
        IntegrationName: mergeatsgo.String("virtual"),
        LastIncidentTimeAfter: types.MustTimeFromString("2021-03-22T04:25:28.253Z"),
        LastIncidentTimeBefore: types.MustTimeFromString("2022-09-19T00:01:59.827Z"),
        PageSize: mergeatsgo.Int64(931165),
        StartDate: mergeatsgo.String("accusantium defensive"),
        Status: operations.IssuesListStatusOngoing.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedIssueList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.IssuesListRequest](../../models/operations/issueslistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.IssuesListResponse](../../models/operations/issueslistresponse.md), error**


## Retrieve

Get a specific issue.

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
    id := "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    ctx := context.Background()
    res, err := s.Issues.Retrieve(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Issue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.IssuesRetrieveResponse](../../models/operations/issuesretrieveresponse.md), error**

