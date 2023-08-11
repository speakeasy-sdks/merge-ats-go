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
    operationSecurity := operations.WebhookReceiversCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.WebhookReceivers.Create(ctx, operations.WebhookReceiversCreateRequest{
        WebhookReceiverRequest: shared.WebhookReceiverRequest{
            Event: "nisi",
            IsActive: false,
            Key: ats.String("voluptatibus"),
        },
        XAccountToken: "quaerat",
    }, operationSecurity)
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
| `request`                                                                                              | [operations.WebhookReceiversCreateRequest](../../models/operations/webhookreceiverscreaterequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.WebhookReceiversCreateSecurity](../../models/operations/webhookreceiverscreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
    operationSecurity := operations.WebhookReceiversListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.WebhookReceivers.List(ctx, operations.WebhookReceiversListRequest{
        XAccountToken: "blanditiis",
    }, operationSecurity)
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
| `request`                                                                                          | [operations.WebhookReceiversListRequest](../../models/operations/webhookreceiverslistrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.WebhookReceiversListSecurity](../../models/operations/webhookreceiverslistsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.WebhookReceiversListResponse](../../models/operations/webhookreceiverslistresponse.md), error**

