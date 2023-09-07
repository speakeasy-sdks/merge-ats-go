# Departments

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.DepartmentsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Departments.List(ctx, operations.DepartmentsListRequest{
        XAccountToken: "vel",
        CreatedAfter: types.MustTimeFromString("2022-02-08T20:43:00.221Z"),
        CreatedBefore: types.MustTimeFromString("2022-08-07T14:04:49.880Z"),
        Cursor: ats.String("laudantium"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-10-11T02:58:57.556Z"),
        ModifiedBefore: types.MustTimeFromString("2022-09-15T20:38:47.955Z"),
        PageSize: ats.Int64(406120),
        RemoteID: ats.String("nulla"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedDepartmentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DepartmentsListRequest](../../models/operations/departmentslistrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DepartmentsListSecurity](../../models/operations/departmentslistsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "excepturi"
    id := "f5fce6c5-5614-46c3-a250-fb008c42e141"
    includeRemoteData := false
    operationSecurity := operations.DepartmentsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Departments.Retrieve(ctx, operationSecurity, xAccountToken, id, includeRemoteData)
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
| `security`                                                                                       | [operations.DepartmentsRetrieveSecurity](../../models/operations/departmentsretrievesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |
| `xAccountToken`                                                                                  | *string*                                                                                         | :heavy_check_mark:                                                                               | Token identifying the end user.                                                                  |
| `id`                                                                                             | *string*                                                                                         | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `includeRemoteData`                                                                              | **bool*                                                                                          | :heavy_minus_sign:                                                                               | Whether to include the original data Merge fetched from the third-party to produce these models. |


### Response

**[*operations.DepartmentsRetrieveResponse](../../models/operations/departmentsretrieveresponse.md), error**

