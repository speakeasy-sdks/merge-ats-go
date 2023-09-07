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
        Category: operations.LinkedAccountsListCategoryAccounting.ToPointer(),
        Cursor: ats.String("voluptates"),
        EndUserEmailAddress: ats.String("libero"),
        EndUserOrganizationName: ats.String("vitae"),
        EndUserOriginID: ats.String("accusamus"),
        EndUserOriginIds: ats.String("similique"),
        ID: ats.String("426555ba-3c28-4744-ad53-b88f3a8d8f5c"),
        Ids: ats.String("sit"),
        IncludeDuplicates: ats.Bool(false),
        IntegrationName: ats.String("rerum"),
        IsTestAccount: ats.String("sed"),
        PageSize: ats.Int64(967966),
        Status: ats.String("explicabo"),
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

