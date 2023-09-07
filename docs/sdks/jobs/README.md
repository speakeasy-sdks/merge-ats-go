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
        XAccountToken: "quod",
        Code: ats.String("labore"),
        CreatedAfter: types.MustTimeFromString("2022-10-04T21:10:21.697Z"),
        CreatedBefore: types.MustTimeFromString("2021-09-04T08:55:11.388Z"),
        Cursor: ats.String("suscipit"),
        Expand: operations.JobsListExpandDepartmentsOffices.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-09-01T22:08:53.618Z"),
        ModifiedBefore: types.MustTimeFromString("2021-06-13T13:07:55.703Z"),
        Offices: ats.String("fugiat"),
        PageSize: ats.Int64(424089),
        RemoteFields: operations.JobsListRemoteFieldsStatus.ToPointer(),
        RemoteID: ats.String("ducimus"),
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
        XAccountToken: "vel",
        Expand: operations.JobsRetrieveExpandDepartmentsOfficesHiringManagers.ToPointer(),
        ID: "dbb675fd-5e60-4b37-9ed4-f6fbee41f333",
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

