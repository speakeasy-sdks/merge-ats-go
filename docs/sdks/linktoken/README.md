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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.LinkToken.Create(ctx, shared.EndUserDetailsRequest{
        Categories: []shared.CategoriesEnum{
            shared.CategoriesEnumTicketing,
        },
        CommonModels: []shared.CommonModelScopesBodyRequest{
            shared.CommonModelScopesBodyRequest{
                DisabledFields: []string{
                    "first_name",
                },
                EnabledActions: []shared.EnabledActionsEnum{
                    shared.EnabledActionsEnumRead,
                    shared.EnabledActionsEnumWrite,
                },
                ModelID: "hris.Employee",
            },
        },
        EndUserEmailAddress: "string",
        EndUserOrganizationName: "string",
        EndUserOriginID: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.EndUserDetailsRequest](../../pkg/models/shared/enduserdetailsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.LinkTokenCreateResponse](../../pkg/models/operations/linktokencreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
