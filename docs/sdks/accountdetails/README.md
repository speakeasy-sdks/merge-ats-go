# AccountDetails
(*AccountDetails*)

### Available Operations

* [Retrieve](#retrieve) - Get details for a linked account.

## Retrieve

Get details for a linked account.

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


    var xAccountToken string = "string"

    ctx := context.Background()
    res, err := s.AccountDetails.Retrieve(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xAccountToken`                                       | *string*                                              | :heavy_check_mark:                                    | Token identifying the end user.                       |


### Response

**[*operations.AccountDetailsRetrieveResponse](../../pkg/models/operations/accountdetailsretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
