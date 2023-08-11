# WebhookReceivers

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    webhookReceiverRequest := shared.WebhookReceiverRequest{
        Event: "non",
        IsActive: false,
        Key: ats.String("distinctio"),
    }
    xAccountToken := "in"
    operationSecurity := operations.WebhookReceiversCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.WebhookReceivers.Create(ctx, operationSecurity, webhookReceiverRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookReceiver != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.WebhookReceiversCreateSecurity](../../models/operations/webhookreceiverscreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `webhookReceiverRequest`                                                                               | [shared.WebhookReceiverRequest](../../models/shared/webhookreceiverrequest.md)                         | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `xAccountToken`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | Token identifying the end user.                                                                        |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "exercitationem"
    operationSecurity := operations.WebhookReceiversListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.WebhookReceivers.List(ctx, operationSecurity, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhookReceivers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `security`                                                                                         | [operations.WebhookReceiversListSecurity](../../models/operations/webhookreceiverslistsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |
| `xAccountToken`                                                                                    | *string*                                                                                           | :heavy_check_mark:                                                                                 | Token identifying the end user.                                                                    |


### Response

**[*operations.WebhookReceiversListResponse](../../models/operations/webhookreceiverslistresponse.md), error**

