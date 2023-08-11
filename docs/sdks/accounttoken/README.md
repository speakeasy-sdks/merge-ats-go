# AccountToken

### Available Operations

* [Retrieve](#retrieve) - Returns the account token for the end user with the provided public token.

## Retrieve

Returns the account token for the end user with the provided public token.

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
    publicToken := "distinctio"
    operationSecurity := operations.AccountTokenRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AccountToken.Retrieve(ctx, operationSecurity, publicToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.AccountTokenRetrieveSecurity](../../models/operations/accounttokenretrievesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `publicToken`                                                                                      | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |


### Response

**[*operations.AccountTokenRetrieveResponse](../../models/operations/accounttokenretrieveresponse.md), error**
