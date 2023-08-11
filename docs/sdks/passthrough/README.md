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
    operationSecurity := operations.PassthroughCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.Create(ctx, operations.PassthroughCreateRequest{
        DataPassthroughRequest: shared.DataPassthroughRequest{
            BaseURLOverride: ats.String("quos"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "aspernatur": "ducimus",
                "nesciunt": "fuga",
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
                shared.MultipartFormFieldRequest{
                    ContentType: ats.String("application/pdf"),
                    Data: "SW50ZWdyYXRlIGZhc3QKSW50ZWdyYXRlIG9uY2U=",
                    Encoding: shared.MultipartFormFieldRequestEncodingBase64.ToPointer(),
                    FileName: ats.String("resume.pdf"),
                    Name: "resume",
                },
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
        },
        XAccountToken: "incidunt",
    }, operationSecurity)
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
| `request`                                                                                    | [operations.PassthroughCreateRequest](../../models/operations/passthroughcreaterequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.PassthroughCreateSecurity](../../models/operations/passthroughcreatesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.PassthroughCreateResponse](../../models/operations/passthroughcreateresponse.md), error**

