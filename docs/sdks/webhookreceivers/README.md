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
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    webhookReceiverRequest := shared.WebhookReceiverRequest{
        Event: "bluetooth Extended",
        IsActive: false,
    }
    var xAccountToken string = "South"

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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `webhookReceiverRequest`                                                       | [shared.WebhookReceiverRequest](../../models/shared/webhookreceiverrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `xAccountToken`                                                                | *string*                                                                       | :heavy_check_mark:                                                             | Token identifying the end user.                                                |


### Response

**[*operations.WebhookReceiversCreateResponse](../../models/operations/webhookreceiverscreateresponse.md), error**


## List

Returns a list of `WebhookReceiver` objects.

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
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    var xAccountToken string = "Bicycle"

    ctx := context.Background()
    res, err := s.WebhookReceivers.List(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookReceivers != nil {
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

**[*operations.WebhookReceiversListResponse](../../models/operations/webhookreceiverslistresponse.md), error**

