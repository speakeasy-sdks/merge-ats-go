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
        Category: operations.LinkedAccountsListCategoryCrm.ToPointer(),
        Cursor: ats.String("quo"),
        EndUserEmailAddress: ats.String("ex"),
        EndUserOrganizationName: ats.String("ut"),
        EndUserOriginID: ats.String("ad"),
        EndUserOriginIds: ats.String("expedita"),
        ID: ats.String("08b61891-baa0-4fe1-ade0-08e6f8c5f350"),
        Ids: ats.String("illum"),
        IncludeDuplicates: ats.Bool(false),
        IntegrationName: ats.String("totam"),
        IsTestAccount: ats.String("impedit"),
        PageSize: ats.Int64(842777),
        Status: ats.String("nam"),
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

