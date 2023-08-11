# Passthrough

### Available Operations

* [Create](#create) - Pull data from an endpoint not currently supported by Merge.
* [Create](#create) - Pull data from an endpoint not currently supported by Merge.
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
    operationSecurity := operations.PassthroughCreateMultipartSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.Create(ctx, operations.PassthroughCreateMultipartRequest{
        DataPassthroughRequest3: shared.DataPassthroughRequest3{
            BaseURLOverride: ats.String("quibusdam"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "odit": "voluptatibus",
                "vel": "magnam",
            },
            Method: "quibusdam",
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
            RequestFormat: ats.String("facere"),
        },
        XAccountToken: "libero",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PassthroughCreateMultipartRequest](../../models/operations/passthroughcreatemultipartrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.PassthroughCreateMultipartSecurity](../../models/operations/passthroughcreatemultipartsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.PassthroughCreateMultipartResponse](../../models/operations/passthroughcreatemultipartresponse.md), error**


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
    operationSecurity := operations.PassthroughCreateFormSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.Create(ctx, operations.PassthroughCreateFormRequest{
        DataPassthroughRequest3: shared.DataPassthroughRequest3{
            BaseURLOverride: ats.String("architecto"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "quia": "porro",
                "aliquam": "velit",
                "illo": "accusantium",
                "vel": "ea",
            },
            Method: "beatae",
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
            RequestFormat: ats.String("excepturi"),
        },
        XAccountToken: "eum",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PassthroughCreateFormRequest](../../models/operations/passthroughcreateformrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.PassthroughCreateFormSecurity](../../models/operations/passthroughcreateformsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.PassthroughCreateFormResponse](../../models/operations/passthroughcreateformresponse.md), error**


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
    operationSecurity := operations.PassthroughCreateJSONSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Passthrough.Create(ctx, operations.PassthroughCreateJSONRequest{
        DataPassthroughRequest: shared.DataPassthroughRequest{
            BaseURLOverride: ats.String("velit"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "perspiciatis": "earum",
                "dicta": "impedit",
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
        XAccountToken: "iste",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PassthroughCreateJSONRequest](../../models/operations/passthroughcreatejsonrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.PassthroughCreateJSONSecurity](../../models/operations/passthroughcreatejsonsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.PassthroughCreateJSONResponse](../../models/operations/passthroughcreatejsonresponse.md), error**

