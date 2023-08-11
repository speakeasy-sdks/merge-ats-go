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
        XAccountToken: "inventore",
        Code: ats.String("fuga"),
        CreatedAfter: types.MustTimeFromString("2020-01-27T09:38:54.834Z"),
        CreatedBefore: types.MustTimeFromString("2021-10-13T17:26:42.757Z"),
        Cursor: ats.String("delectus"),
        Expand: operations.JobsListExpandDepartmentsOfficesHiringManagers.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-05-25T00:50:01.296Z"),
        ModifiedBefore: types.MustTimeFromString("2022-03-01T17:27:03.796Z"),
        Offices: ats.String("quos"),
        PageSize: ats.Int64(415280),
        RemoteFields: operations.JobsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("itaque"),
        ShowEnumOrigins: operations.JobsListShowEnumOriginsStatus.ToPointer(),
        Status: operations.JobsListStatusDraft.ToPointer(),
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
        XAccountToken: "totam",
        Expand: operations.JobsRetrieveExpandRecruiters.ToPointer(),
        ID: "4be05601-3f59-4da7-97a5-9ecfef66ef1c",
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

