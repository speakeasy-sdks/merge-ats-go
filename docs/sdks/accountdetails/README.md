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
    operationSecurity := operations.AccountDetailsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AccountDetails.Retrieve(ctx, operations.AccountDetailsRetrieveRequest{
        XAccountToken: "provident",
    }, operationSecurity)
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
| `request`                                                                                              | [operations.AccountDetailsRetrieveRequest](../../models/operations/accountdetailsretrieverequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.AccountDetailsRetrieveSecurity](../../models/operations/accountdetailsretrievesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.AccountDetailsRetrieveResponse](../../models/operations/accountdetailsretrieveresponse.md), error**

