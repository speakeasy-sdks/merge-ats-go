# Issues

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
        AccountToken: mergeatsgo.String("tenetur"),
        Cursor: mergeatsgo.String("quae"),
        EndDate: mergeatsgo.String("earum"),
        EndUserOrganizationName: mergeatsgo.String("vel"),
        FirstIncidentTimeAfter: types.MustTimeFromString("2022-09-28T13:55:38.564Z"),
        FirstIncidentTimeBefore: types.MustTimeFromString("2021-04-21T04:50:55.832Z"),
        IncludeMuted: mergeatsgo.String("soluta"),
        IntegrationName: mergeatsgo.String("accusantium"),
        LastIncidentTimeAfter: types.MustTimeFromString("2022-01-15T23:18:40.855Z"),
        LastIncidentTimeBefore: types.MustTimeFromString("2022-08-24T06:58:07.315Z"),
        PageSize: mergeatsgo.Int64(443879),
        StartDate: mergeatsgo.String("ullam"),
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
    id := "082d68ea-19f1-4d17-8513-39d08086a184"

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

