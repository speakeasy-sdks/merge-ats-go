# Departments
(*Departments*)

### Available Operations

* [List](#list) - Returns a list of `Department` objects.
* [Retrieve](#retrieve) - Returns a `Department` object with the given `id`.

## List

Returns a list of `Department` objects.

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
    res, err := s.Departments.List(ctx, operations.DepartmentsListRequest{
        XAccountToken: "Northeast Metal Canada",
        CreatedAfter: types.MustTimeFromString("2023-10-02T13:41:25.267Z"),
        CreatedBefore: types.MustTimeFromString("2022-11-27T21:47:17.392Z"),
        Cursor: mergeatsgo.String("Response West male"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-04-05T07:45:35.369Z"),
        ModifiedBefore: types.MustTimeFromString("2022-06-14T06:59:01.434Z"),
        PageSize: mergeatsgo.Int64(68504),
        RemoteID: mergeatsgo.String("Designer hacking"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedDepartmentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DepartmentsListRequest](../../models/operations/departmentslistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DepartmentsListResponse](../../models/operations/departmentslistresponse.md), error**


## Retrieve

Returns a `Department` object with the given `id`.

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
    xAccountToken := "till"
    id := "56591081-ad20-4d60-8c8e-92b241fa3790"
    includeRemoteData := false

    ctx := context.Background()
    res, err := s.Departments.Retrieve(ctx, xAccountToken, id, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }

    if res.Department != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `xAccountToken`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | Token identifying the end user.                                                                  |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `includeRemoteData`                                                                              | **bool*                                                                                          | :heavy_minus_sign:                                                                               | Whether to include the original data Merge fetched from the third-party to produce these models. |


### Response

**[*operations.DepartmentsRetrieveResponse](../../models/operations/departmentsretrieveresponse.md), error**

