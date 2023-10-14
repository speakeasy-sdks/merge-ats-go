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
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    dataPassthroughRequest := shared.DataPassthroughRequest{
        Data: mergeatsgo.String("{\"company\": \"Lime\", \"model\": \"Gen 2.5\"}"),
        Headers: map[string]interface{}{
            "EXTRA-HEADER": "online",
        },
        Method: shared.DataPassthroughRequestMethodPost,
        MultipartFormData: []shared.MultipartFormFieldRequest{
            shared.MultipartFormFieldRequest{
                ContentType: mergeatsgo.String("application/pdf"),
                Data: "SW50ZWdyYXRlIGZhc3QKSW50ZWdyYXRlIG9uY2U=",
                Encoding: shared.MultipartFormFieldRequestEncodingBase64.ToPointer(),
                FileName: mergeatsgo.String("resume.pdf"),
                Name: "resume",
            },
        },
        Path: "/scooters",
        RequestFormat: shared.DataPassthroughRequestRequestFormatJSON.ToPointer(),
    }

    var xAccountToken string = "Configuration"

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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `dataPassthroughRequest`                                                       | [shared.DataPassthroughRequest](../../models/shared/datapassthroughrequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `xAccountToken`                                                                | *string*                                                                       | :heavy_check_mark:                                                             | Token identifying the end user.                                                |


### Response

**[*operations.AsyncPassthroughCreateResponse](../../models/operations/asyncpassthroughcreateresponse.md), error**


## Retrieve

Retrieves data from earlier async-passthrough POST request

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
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "till"

    var asyncPassthroughReceiptID string = "56591081-ad20-4d60-8c8e-92b241fa3790"

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

**[*operations.AsyncPassthroughRetrieveResponse](../../models/operations/asyncpassthroughretrieveresponse.md), error**

