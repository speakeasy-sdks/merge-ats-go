# WebhookReceivers
(*WebhookReceivers*)

### Available Operations

* [Create](#create) - Creates a `WebhookReceiver` object with the given values.
* [List](#list) - Returns a list of `WebhookReceiver` objects.

## Create

Creates a `WebhookReceiver` object with the given values.

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


    webhookReceiverRequest := shared.WebhookReceiverRequest{
        Event: "<value>",
        IsActive: false,
    }

    var xAccountToken string = "<value>"

    ctx := context.Background()
    res, err := s.WebhookReceivers.Create(ctx, webhookReceiverRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookReceiver != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `webhookReceiverRequest`                                                           | [shared.WebhookReceiverRequest](../../pkg/models/shared/webhookreceiverrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `xAccountToken`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | Token identifying the end user.                                                    |


### Response

**[*operations.WebhookReceiversCreateResponse](../../pkg/models/operations/webhookreceiverscreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Returns a list of `WebhookReceiver` objects.

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
    res, err := s.WebhookReceivers.List(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
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

**[*operations.WebhookReceiversListResponse](../../pkg/models/operations/webhookreceiverslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
