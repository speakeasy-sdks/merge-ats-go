# AsyncPassthrough
(*AsyncPassthrough*)

### Available Operations

* [Create](#create) - Asynchronously pull data from an endpoint not currently supported by Merge.
* [Retrieve](#retrieve) - Retrieves data from earlier async-passthrough POST request

## Create

Asynchronously pull data from an endpoint not currently supported by Merge.

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


    dataPassthroughRequest := shared.DataPassthroughRequest{
        Data: mergeatsgo.String("{\"company\": \"Lime\", \"model\": \"Gen 2.5\"}"),
        Headers: map[string]interface{}{
            "EXTRA-HEADER": "value",
        },
        Method: shared.MethodPost,
        Path: "/scooters",
        RequestFormat: shared.RequestFormatJSON.ToPointer(),
    }

    var xAccountToken string = "<value>"

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Create(ctx, dataPassthroughRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.AsyncPassthroughReciept != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `dataPassthroughRequest`                                                           | [shared.DataPassthroughRequest](../../pkg/models/shared/datapassthroughrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `xAccountToken`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | Token identifying the end user.                                                    |


### Response

**[*operations.AsyncPassthroughCreateResponse](../../pkg/models/operations/asyncpassthroughcreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Retrieves data from earlier async-passthrough POST request

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

    var asyncPassthroughReceiptID string = "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Retrieve(ctx, xAccountToken, asyncPassthroughReceiptID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xAccountToken`                                       | *string*                                              | :heavy_check_mark:                                    | Token identifying the end user.                       |
| `asyncPassthroughReceiptID`                           | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.AsyncPassthroughRetrieveResponse](../../pkg/models/operations/asyncpassthroughretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
