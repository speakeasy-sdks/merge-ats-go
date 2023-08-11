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
        AccountToken: ats.String("impedit"),
        Cursor: ats.String("quibusdam"),
        EndDate: ats.String("nam"),
        EndUserOrganizationName: ats.String("ipsam"),
        FirstIncidentTimeAfter: types.MustTimeFromString("2022-07-22T07:25:32.550Z"),
        FirstIncidentTimeBefore: types.MustTimeFromString("2022-12-02T14:32:21.843Z"),
        IncludeMuted: ats.String("deleniti"),
        IntegrationName: ats.String("veritatis"),
        LastIncidentTimeAfter: types.MustTimeFromString("2022-10-12T04:34:00.871Z"),
        LastIncidentTimeBefore: types.MustTimeFromString("2022-11-25T11:52:35.630Z"),
        PageSize: ats.Int64(24944),
        StartDate: ats.String("modi"),
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
    operationSecurity := operations.IssuesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Issues.Retrieve(ctx, operations.IssuesRetrieveRequest{
        ID: "1813d520-8ece-47e2-93b6-68451c6c6e20",
    }, operationSecurity)
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
| `request`                                                                              | [operations.IssuesRetrieveRequest](../../models/operations/issuesretrieverequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.IssuesRetrieveSecurity](../../models/operations/issuesretrievesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.IssuesRetrieveResponse](../../models/operations/issuesretrieveresponse.md), error**

