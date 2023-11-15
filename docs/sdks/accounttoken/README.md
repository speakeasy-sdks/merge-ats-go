# AccountToken
(*AccountToken*)

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
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New()


    var publicToken string = "string"

    ctx := context.Background()
    res, err := s.AccountToken.Retrieve(ctx, publicToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `publicToken`                                         | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.AccountTokenRetrieveResponse](../../pkg/models/operations/accounttokenretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
