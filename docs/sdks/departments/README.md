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
        XAccountToken: "cum",
        CreatedAfter: types.MustTimeFromString("2021-11-21T14:04:14.152Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-25T08:06:51.452Z"),
        Cursor: ats.String("culpa"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2021-04-07T21:05:22.480Z"),
        ModifiedBefore: types.MustTimeFromString("2021-02-13T15:31:34.151Z"),
        PageSize: ats.Int64(348783),
        RemoteID: ats.String("nobis"),
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
    xAccountToken := "sit"
    id := "b2f2fb7b-194a-4276-b269-16fe1f08f429"
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

