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
        XAccountToken: "consequatur",
        CreatedAfter: types.MustTimeFromString("2022-08-17T23:19:52.079Z"),
        CreatedBefore: types.MustTimeFromString("2022-06-21T08:29:37.931Z"),
        Cursor: ats.String("quas"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-12-01T12:56:16.650Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-01T12:15:51.307Z"),
        PageSize: ats.Int64(405942),
        RemoteID: ats.String("sed"),
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
    operationSecurity := operations.DepartmentsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Departments.Retrieve(ctx, operations.DepartmentsRetrieveRequest{
        XAccountToken: "sit",
        ID: "65e904f3-b119-44b8-abf6-03a79f9dfe0a",
        IncludeRemoteData: ats.Bool(false),
    }, operationSecurity)
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
| `request`                                                                                        | [operations.DepartmentsRetrieveRequest](../../models/operations/departmentsretrieverequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DepartmentsRetrieveSecurity](../../models/operations/departmentsretrievesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DepartmentsRetrieveResponse](../../models/operations/departmentsretrieveresponse.md), error**

