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
        XAccountToken: "suscipit",
        Code: ats.String("dolorem"),
        CreatedAfter: types.MustTimeFromString("2022-03-27T18:38:54.968Z"),
        CreatedBefore: types.MustTimeFromString("2022-08-15T23:37:19.342Z"),
        Cursor: ats.String("animi"),
        Expand: operations.JobsListExpandOfficesRecruiters.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-12-22T05:17:09.936Z"),
        ModifiedBefore: types.MustTimeFromString("2022-11-29T01:33:31.768Z"),
        Offices: ats.String("ducimus"),
        PageSize: ats.Int64(619183),
        RemoteFields: operations.JobsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("occaecati"),
        ShowEnumOrigins: operations.JobsListShowEnumOriginsStatus.ToPointer(),
        Status: operations.JobsListStatusClosed.ToPointer(),
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
        XAccountToken: "adipisci",
        Expand: operations.JobsRetrieveExpandDepartmentsHiringManagers.ToPointer(),
        ID: "2fde0477-1778-4ff6-9d01-7476360a15db",
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

