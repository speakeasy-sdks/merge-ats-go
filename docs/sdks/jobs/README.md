# Jobs

### Available Operations

* [List](#list) - Returns a list of `Job` objects.
* [Retrieve](#retrieve) - Returns a `Job` object with the given `id`.

## List

Returns a list of `Job` objects.

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
    operationSecurity := operations.JobsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Jobs.List(ctx, operations.JobsListRequest{
        XAccountToken: "ex",
        Code: ats.String("ut"),
        CreatedAfter: types.MustTimeFromString("2022-04-15T09:59:51.518Z"),
        CreatedBefore: types.MustTimeFromString("2022-06-10T00:34:58.414Z"),
        Cursor: ats.String("cum"),
        Expand: operations.JobsListExpandDepartmentsOfficesHiringManagersRecruiters.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-06-21T05:39:54.582Z"),
        ModifiedBefore: types.MustTimeFromString("2022-10-30T18:42:59.470Z"),
        Offices: ats.String("rerum"),
        PageSize: ats.Int64(665678),
        RemoteFields: operations.JobsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("culpa"),
        ShowEnumOrigins: operations.JobsListShowEnumOriginsStatus.ToPointer(),
        Status: operations.JobsListStatusArchived.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedJobList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.JobsListRequest](../../models/operations/jobslistrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.JobsListSecurity](../../models/operations/jobslistsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.JobsListResponse](../../models/operations/jobslistresponse.md), error**


## Retrieve

Returns a `Job` object with the given `id`.

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
    operationSecurity := operations.JobsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Jobs.Retrieve(ctx, operations.JobsRetrieveRequest{
        XAccountToken: "sapiente",
        Expand: operations.JobsRetrieveExpandOfficesRecruiters.ToPointer(),
        ID: "1ade008e-6f8c-45f3-90d8-cdb5a3418143",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.JobsRetrieveRemoteFieldsStatus.ToPointer(),
        ShowEnumOrigins: operations.JobsRetrieveShowEnumOriginsStatus.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Job != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.JobsRetrieveRequest](../../models/operations/jobsretrieverequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.JobsRetrieveSecurity](../../models/operations/jobsretrievesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.JobsRetrieveResponse](../../models/operations/jobsretrieveresponse.md), error**

