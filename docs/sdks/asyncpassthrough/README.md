# AsyncPassthrough

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    dataPassthroughRequest := shared.DataPassthroughRequest{
        BaseURLOverride: ats.String("excepturi"),
        Data: ats.String("{"company": "Lime", "model": "Gen 2.5"}"),
        Headers: map[string]interface{}{
            "facilis": "tempore",
            "labore": "delectus",
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
        },
        NormalizeResponse: ats.Bool(false),
        Path: "/scooters",
        RequestFormat: shared.DataPassthroughRequestRequestFormatJSON.ToPointer(),
    }
    xAccountToken := "non"
    operationSecurity := operations.AsyncPassthroughCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Create(ctx, operationSecurity, dataPassthroughRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AsyncPassthroughReciept != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.AsyncPassthroughCreateSecurity](../../models/operations/asyncpassthroughcreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `dataPassthroughRequest`                                                                               | [shared.DataPassthroughRequest](../../models/shared/datapassthroughrequest.md)                         | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `xAccountToken`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | Token identifying the end user.                                                                        |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "eligendi"
    asyncPassthroughReceiptID := "969e9a3e-fa77-4dfb-94cd-66ae395efb9b"
    operationSecurity := operations.AsyncPassthroughRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AsyncPassthrough.Retrieve(ctx, operationSecurity, xAccountToken, asyncPassthroughReceiptID)
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
| `security`                                                                                                 | [operations.AsyncPassthroughRetrieveSecurity](../../models/operations/asyncpassthroughretrievesecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `xAccountToken`                                                                                            | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Token identifying the end user.                                                                            |
| `asyncPassthroughReceiptID`                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |


### Response

**[*operations.AsyncPassthroughRetrieveResponse](../../models/operations/asyncpassthroughretrieveresponse.md), error**

