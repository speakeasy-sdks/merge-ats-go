# Passthrough
(*Passthrough*)

### Available Operations

* [Create](#create) - Pull data from an endpoint not currently supported by Merge.

## Create

Pull data from an endpoint not currently supported by Merge.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    dataPassthroughRequest := shared.DataPassthroughRequest{
        BaseURLOverride: mergeatsgo.String("labore"),
        Data: mergeatsgo.String("{"company": "Lime", "model": "Gen 2.5"}"),
        Headers: map[string]interface{}{
            "quidem": "atque",
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
        NormalizeResponse: mergeatsgo.Bool(false),
        Path: "/scooters",
        RequestFormat: shared.DataPassthroughRequestRequestFormatJSON.ToPointer(),
    }
    xAccountToken := "laborum"

    ctx := context.Background()
    res, err := s.Passthrough.Create(ctx, dataPassthroughRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoteResponse != nil {
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

**[*operations.PassthroughCreateResponse](../../models/operations/passthroughcreateresponse.md), error**

