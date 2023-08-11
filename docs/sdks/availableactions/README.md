# AvailableActions

### Available Operations

* [Retrieve](#retrieve) - Returns a list of models and actions available for an account.

## Retrieve

Returns a list of models and actions available for an account.

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
    xAccountToken := "saepe"
    operationSecurity := operations.AvailableActionsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AvailableActions.Retrieve(ctx, operationSecurity, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AvailableActions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `security`                                                                                                 | [operations.AvailableActionsRetrieveSecurity](../../models/operations/availableactionsretrievesecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `xAccountToken`                                                                                            | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Token identifying the end user.                                                                            |


### Response

**[*operations.AvailableActionsRetrieveResponse](../../models/operations/availableactionsretrieveresponse.md), error**

