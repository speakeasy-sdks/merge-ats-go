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
        Cursor: ats.String("officiis"),
        EndUserEmailAddress: ats.String("esse"),
        EndUserOrganizationName: ats.String("necessitatibus"),
        EndUserOriginID: ats.String("sed"),
        EndUserOriginIds: ats.String("veniam"),
        ID: ats.String("3b668451-c6c6-4e20-9e16-deab3fec9578"),
        Ids: ats.String("officia"),
        IncludeDuplicates: ats.Bool(false),
        IntegrationName: ats.String("voluptas"),
        IsTestAccount: ats.String("numquam"),
        PageSize: ats.Int64(364912),
        Status: ats.String("quos"),
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

