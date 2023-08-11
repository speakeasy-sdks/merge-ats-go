# DeleteAccount

### Available Operations

* [DeleteAccountDelete](#deleteaccountdelete) - Delete a linked account.

## DeleteAccountDelete

Delete a linked account.

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
    operationSecurity := operations.DeleteAccountDeleteSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.DeleteAccount.DeleteAccountDelete(ctx, operations.DeleteAccountDeleteRequest{
        XAccountToken: "magnam",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteAccountDeleteRequest](../../models/operations/deleteaccountdeleterequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeleteAccountDeleteSecurity](../../models/operations/deleteaccountdeletesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeleteAccountDeleteResponse](../../models/operations/deleteaccountdeleteresponse.md), error**

