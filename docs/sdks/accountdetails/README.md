# AccountDetails

### Available Operations

* [Retrieve](#retrieve) - Get details for a linked account.

## Retrieve

Get details for a linked account.

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
    xAccountToken := "provident"
    operationSecurity := operations.AccountDetailsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AccountDetails.Retrieve(ctx, operationSecurity, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.AccountDetailsRetrieveSecurity](../../models/operations/accountdetailsretrievesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `xAccountToken`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | Token identifying the end user.                                                                        |


### Response

**[*operations.AccountDetailsRetrieveResponse](../../models/operations/accountdetailsretrieveresponse.md), error**

