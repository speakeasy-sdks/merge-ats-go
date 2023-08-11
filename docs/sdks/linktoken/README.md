# LinkToken

### Available Operations

* [Create](#create) - Creates a link token to be used when linking a new end user.

## Create

Creates a link token to be used when linking a new end user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.LinkTokenCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.LinkToken.Create(ctx, shared.EndUserDetailsRequest{
        Categories: []shared.CategoriesEnum{
            shared.CategoriesEnumCrm,
            shared.CategoriesEnumAccounting,
        },
        CommonModels: []shared.CommonModelScopesBodyRequest{
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "eum",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumWrite,
                    shared.EnabledActionsEnumWrite,
                },
                ModelID: "hris.Employee",
            },
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "error",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumWrite,
                    shared.EnabledActionsEnumWrite,
                    shared.EnabledActionsEnumWrite,
                    shared.EnabledActionsEnumWrite,
                },
                ModelID: "hris.Employee",
            },
        },
        EndUserEmailAddress: "ad",
        EndUserOrganizationName: "deleniti",
        EndUserOriginID: "enim",
        Integration: ats.String("vitae"),
        LinkExpiryMins: ats.Int64(833316),
        ShouldCreateMagicLinkURL: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.EndUserDetailsRequest](../../models/shared/enduserdetailsrequest.md)             | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.LinkTokenCreateSecurity](../../models/operations/linktokencreatesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.LinkTokenCreateResponse](../../models/operations/linktokencreateresponse.md), error**

