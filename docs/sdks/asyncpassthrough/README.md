# AsyncPassthrough

### Available Operations

* [Create](#create) - Asynchronously pull data from an endpoint not currently supported by Merge.
* [Create](#create) - Asynchronously pull data from an endpoint not currently supported by Merge.
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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    operationSecurity := operations.AsyncPassthroughCreateMultipartSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Create(ctx, operations.AsyncPassthroughCreateMultipartRequest{
        DataPassthroughRequest3: shared.DataPassthroughRequest3{
            BaseURLOverride: ats.String("excepturi"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "facilis": "tempore",
                "labore": "delectus",
            },
            Method: "eum",
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
            RequestFormat: ats.String("eligendi"),
        },
        XAccountToken: "sint",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AsyncPassthroughReciept != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.AsyncPassthroughCreateMultipartRequest](../../models/operations/asyncpassthroughcreatemultipartrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.AsyncPassthroughCreateMultipartSecurity](../../models/operations/asyncpassthroughcreatemultipartsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.AsyncPassthroughCreateMultipartResponse](../../models/operations/asyncpassthroughcreatemultipartresponse.md), error**


## Create

Asynchronously pull data from an endpoint not currently supported by Merge.

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
    operationSecurity := operations.AsyncPassthroughCreateFormSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Create(ctx, operations.AsyncPassthroughCreateFormRequest{
        DataPassthroughRequest3: shared.DataPassthroughRequest3{
            BaseURLOverride: ats.String("aliquid"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "necessitatibus": "sint",
                "officia": "dolor",
                "debitis": "a",
            },
            Method: "dolorum",
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
            },
            NormalizeResponse: ats.Bool(false),
            Path: "/scooters",
            RequestFormat: ats.String("in"),
        },
        XAccountToken: "illum",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AsyncPassthroughReciept != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AsyncPassthroughCreateFormRequest](../../models/operations/asyncpassthroughcreateformrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.AsyncPassthroughCreateFormSecurity](../../models/operations/asyncpassthroughcreateformsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.AsyncPassthroughCreateFormResponse](../../models/operations/asyncpassthroughcreateformresponse.md), error**


## Create

Asynchronously pull data from an endpoint not currently supported by Merge.

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
    operationSecurity := operations.AsyncPassthroughCreateJSONSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Create(ctx, operations.AsyncPassthroughCreateJSONRequest{
        DataPassthroughRequest: shared.DataPassthroughRequest{
            BaseURLOverride: ats.String("maiores"),
            Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
            Headers: map[string]interface{}{
                "dicta": "magnam",
                "cumque": "facere",
                "ea": "aliquid",
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
        XAccountToken: "accusamus",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AsyncPassthroughReciept != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AsyncPassthroughCreateJSONRequest](../../models/operations/asyncpassthroughcreatejsonrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.AsyncPassthroughCreateJSONSecurity](../../models/operations/asyncpassthroughcreatejsonsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.AsyncPassthroughCreateJSONResponse](../../models/operations/asyncpassthroughcreatejsonresponse.md), error**


## Retrieve

Retrieves data from earlier async-passthrough POST request

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
    operationSecurity := operations.AsyncPassthroughRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Retrieve(ctx, operations.AsyncPassthroughRetrieveRequest{
        XAccountToken: "non",
        AsyncPassthroughReceiptID: "95efb9ba-88f3-4a66-9970-74ba4469b6e2",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AsyncPassthroughRetrieveRequest](../../models/operations/asyncpassthroughretrieverequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.AsyncPassthroughRetrieveSecurity](../../models/operations/asyncpassthroughretrievesecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.AsyncPassthroughRetrieveResponse](../../models/operations/asyncpassthroughretrieveresponse.md), error**

