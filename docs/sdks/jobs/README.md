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
    res, err := s.Jobs.List(ctx, operations.JobsListRequest{
        XAccountToken: "labore",
        Code: mergeatsgo.String("possimus"),
        CreatedAfter: types.MustTimeFromString("2021-07-11T02:16:12.828Z"),
        CreatedBefore: types.MustTimeFromString("2022-07-21T19:01:11.341Z"),
        Cursor: mergeatsgo.String("corporis"),
        Expand: operations.JobsListExpandRecruiters.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-11-29T08:07:19.623Z"),
        ModifiedBefore: types.MustTimeFromString("2021-10-22T16:49:10.196Z"),
        Offices: mergeatsgo.String("aperiam"),
        PageSize: mergeatsgo.Int64(738683),
        RemoteFields: mergeatsgo.String("consectetur"),
        RemoteID: mergeatsgo.String("in"),
        ShowEnumOrigins: mergeatsgo.String("exercitationem"),
        Status: operations.JobsListStatusPending.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedJobList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.JobsListRequest](../../models/operations/jobslistrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


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
    res, err := s.Jobs.Retrieve(ctx, operations.JobsRetrieveRequest{
        XAccountToken: "facere",
        Expand: operations.JobsRetrieveExpandDepartmentsOffices.ToPointer(),
        ID: "f6fbee41-f333-417f-a35b-60eb1ea42655",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: mergeatsgo.String("minima"),
        ShowEnumOrigins: mergeatsgo.String("nobis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Job != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.JobsRetrieveRequest](../../models/operations/jobsretrieverequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.JobsRetrieveResponse](../../models/operations/jobsretrieveresponse.md), error**

