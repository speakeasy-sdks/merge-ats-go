# LinkedAccounts
(*LinkedAccounts*)

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
    res, err := s.LinkedAccounts.List(ctx, operations.LinkedAccountsListRequest{
        Category: operations.LinkedAccountsListCategoryMktg.ToPointer(),
        Cursor: mergeatsgo.String("compress"),
        EndUserEmailAddress: mergeatsgo.String("Canada that orchid"),
        EndUserOrganizationName: mergeatsgo.String("West"),
        EndUserOriginID: mergeatsgo.String("boil primary synthesize"),
        EndUserOriginIds: mergeatsgo.String("hacking Paradigm"),
        ID: mergeatsgo.String("08055574-19e7-490e-a720-55dd402eb66e"),
        Ids: mergeatsgo.String("lavender Borders often"),
        IncludeDuplicates: mergeatsgo.Bool(false),
        IntegrationName: mergeatsgo.String("nor"),
        IsTestAccount: mergeatsgo.String("Namibia Executive card"),
        PageSize: mergeatsgo.Int64(841031),
        Status: mergeatsgo.String("Account"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedAccountDetailsAndActionsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.LinkedAccountsListRequest](../../models/operations/linkedaccountslistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.LinkedAccountsListResponse](../../models/operations/linkedaccountslistresponse.md), error**

