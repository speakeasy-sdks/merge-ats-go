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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.IssuesListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Issues.List(ctx, operations.IssuesListRequest{
        AccountToken: ats.String("a"),
        Cursor: ats.String("error"),
        EndDate: ats.String("sint"),
        EndUserOrganizationName: ats.String("pariatur"),
        FirstIncidentTimeAfter: types.MustTimeFromString("2022-07-12T09:25:25.257Z"),
        FirstIncidentTimeBefore: types.MustTimeFromString("2020-01-10T06:57:07.831Z"),
        IncludeMuted: ats.String("facere"),
        IntegrationName: ats.String("veritatis"),
        LastIncidentTimeAfter: types.MustTimeFromString("2022-11-27T12:32:54.264Z"),
        LastIncidentTimeBefore: types.MustTimeFromString("2021-09-25T11:11:22.943Z"),
        PageSize: ats.Int64(398434),
        StartDate: ats.String("tenetur"),
        Status: operations.IssuesListStatusOngoing.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedIssueList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.IssuesListRequest](../../models/operations/issueslistrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.IssuesListSecurity](../../models/operations/issueslistsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    id := "e674bdb0-4f15-4756-882d-68ea19f1d170"
    operationSecurity := operations.IssuesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Issues.Retrieve(ctx, operationSecurity, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Issue != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.IssuesRetrieveSecurity](../../models/operations/issuesretrievesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |


### Response

**[*operations.IssuesRetrieveResponse](../../models/operations/issuesretrieveresponse.md), error**

