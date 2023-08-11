# LinkedAccounts

### Available Operations

* [List](#list) - List linked accounts for your organization.

## List

List linked accounts for your organization.

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
    operationSecurity := operations.LinkedAccountsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.LinkedAccounts.List(ctx, operations.LinkedAccountsListRequest{
        Category: operations.LinkedAccountsListCategoryMktg.ToPointer(),
        Cursor: ats.String("libero"),
        EndUserEmailAddress: ats.String("architecto"),
        EndUserOrganizationName: ats.String("voluptatibus"),
        EndUserOriginID: ats.String("quia"),
        EndUserOriginIds: ats.String("porro"),
        ID: ats.String("4310661e-9634-49e1-8f9e-06e3a437000a"),
        Ids: ats.String("recusandae"),
        IncludeDuplicates: ats.Bool(false),
        IntegrationName: ats.String("ea"),
        IsTestAccount: ats.String("quidem"),
        PageSize: ats.Int64(377406),
        Status: ats.String("facilis"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedAccountDetailsAndActionsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.LinkedAccountsListRequest](../../models/operations/linkedaccountslistrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.LinkedAccountsListSecurity](../../models/operations/linkedaccountslistsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.LinkedAccountsListResponse](../../models/operations/linkedaccountslistresponse.md), error**

