# AvailableActions
(*AvailableActions*)

### Available Operations

* [Retrieve](#retrieve) - Returns a list of models and actions available for an account.

## Retrieve

Returns a list of models and actions available for an account.

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
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var xAccountToken string = "<value>"

    ctx := context.Background()
    res, err := s.AvailableActions.Retrieve(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.AvailableActions != nil {
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

**[*operations.AvailableActionsRetrieveResponse](../../pkg/models/operations/availableactionsretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
