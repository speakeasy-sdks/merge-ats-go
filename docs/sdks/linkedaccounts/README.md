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
        Category: operations.LinkedAccountsListCategoryTicketing.ToPointer(),
        Cursor: mergeatsgo.String("temporibus"),
        EndUserEmailAddress: mergeatsgo.String("ullam"),
        EndUserOrganizationName: mergeatsgo.String("adipisci"),
        EndUserOriginID: mergeatsgo.String("cum"),
        EndUserOriginIds: mergeatsgo.String("blanditiis"),
        ID: mergeatsgo.String("8f3a8d8f-5c0b-42f2-bb7b-194a276b2691"),
        Ids: mergeatsgo.String("suscipit"),
        IncludeDuplicates: mergeatsgo.Bool(false),
        IntegrationName: mergeatsgo.String("sapiente"),
        IsTestAccount: mergeatsgo.String("debitis"),
        PageSize: mergeatsgo.Int64(72434),
        Status: mergeatsgo.String("reiciendis"),
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

