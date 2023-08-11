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
        AccountToken: ats.String("at"),
        Cursor: ats.String("impedit"),
        EndDate: ats.String("eos"),
        EndUserOrganizationName: ats.String("sapiente"),
        FirstIncidentTimeAfter: types.MustTimeFromString("2022-11-19T04:16:20.363Z"),
        FirstIncidentTimeBefore: types.MustTimeFromString("2022-11-22T22:38:36.599Z"),
        IncludeMuted: ats.String("cupiditate"),
        IntegrationName: ats.String("provident"),
        LastIncidentTimeAfter: types.MustTimeFromString("2020-10-06T18:56:18.922Z"),
        LastIncidentTimeBefore: types.MustTimeFromString("2020-06-16T06:36:28.349Z"),
        PageSize: ats.Int64(52508),
        StartDate: ats.String("earum"),
        Status: operations.IssuesListStatusResolved.ToPointer(),
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
        ID: "fe6c632c-a3ae-4d01-9799-6312fde04771",
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

