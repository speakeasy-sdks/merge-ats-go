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
            shared.CategoriesEnumAts,
            shared.CategoriesEnumAts,
        },
        CommonModels: []shared.CommonModelScopesBodyRequest{
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "impedit",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumWrite,
                },
                ModelID: "hris.Employee",
            },
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "nam",
                    "dolore",
                    "iusto",
                    "voluptate",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumRead,
                },
                ModelID: "hris.Employee",
            },
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "quo",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumWrite,
                    shared.EnabledActionsEnumRead,
                    shared.EnabledActionsEnumRead,
                },
                ModelID: "hris.Employee",
            },
        },
        EndUserEmailAddress: "voluptatibus",
        EndUserOrganizationName: "vel",
        EndUserOriginID: "magnam",
        Integration: ats.String("quibusdam"),
        LinkExpiryMins: ats.Int64(78969),
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

