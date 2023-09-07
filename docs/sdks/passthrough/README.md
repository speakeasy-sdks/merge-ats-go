# Passthrough

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    dataPassthroughRequest := shared.DataPassthroughRequest{
        BaseURLOverride: ats.String("consequatur"),
        Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
        Headers: map[string]interface{}{
            "esse": "ipsam",
        },
        Method: shared.DataPassthroughRequestMethodPost,
        MultipartFormData: []shared.MultipartFormFieldRequest{
            shared.MultipartFormFieldRequest{
                ContentType: ats.String("application/pdf"),
                Data: "SW50ZWdyYXRlIGZhc3QKSW50ZWdyYXRlIG9uY2U=",
                Encoding: shared.MultipartFormFieldRequestEncodingBase64.ToPointer(),
                FileName: ats.String("resume.pdf"),
                Name: "resume",
            },
        },
        NormalizeResponse: ats.Bool(false),
        Path: "/scooters",
        RequestFormat: shared.DataPassthroughRequestRequestFormatJSON.ToPointer(),
    }
    xAccountToken := "sit"
    operationSecurity := operations.PassthroughCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.Create(ctx, operationSecurity, dataPassthroughRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `security`                                                                                   | [operations.PassthroughCreateSecurity](../../models/operations/passthroughcreatesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `dataPassthroughRequest`                                                                     | [shared.DataPassthroughRequest](../../models/shared/datapassthroughrequest.md)               | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `xAccountToken`                                                                              | *string*                                                                                     | :heavy_check_mark:                                                                           | Token identifying the end user.                                                              |


### Response

**[*operations.PassthroughCreateResponse](../../models/operations/passthroughcreateresponse.md), error**

