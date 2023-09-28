# LinkToken
(*LinkToken*)

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
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.LinkToken.Create(ctx, shared.EndUserDetailsRequest{
        Categories: []shared.CategoriesEnum{
            shared.CategoriesEnumCrm,
        },
        CommonModels: []shared.CommonModelScopesBodyRequest{
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "adipisci",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumWrite,
                },
                ModelID: "hris.Employee",
            },
        },
        EndUserEmailAddress: "dolores",
        EndUserOrganizationName: "blanditiis",
        EndUserOriginID: "in",
        Integration: mergeatsgo.String("dolore"),
        LinkExpiryMins: mergeatsgo.Int64(304468),
        ShouldCreateMagicLinkURL: mergeatsgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.EndUserDetailsRequest](../../models/shared/enduserdetailsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.LinkTokenCreateResponse](../../models/operations/linktokencreateresponse.md), error**

