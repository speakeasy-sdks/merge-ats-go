# Attachments
(*Attachments*)

### Available Operations

* [Create](#create) - Creates an `Attachment` object with the given values.
* [List](#list) - Returns a list of `Attachment` objects.
* [Retrieve](#retrieve) - Returns an `Attachment` object with the given `id`.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `Attachment` POSTs.

## Create

Creates an `Attachment` object with the given values.

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


    attachmentEndpointRequest := shared.AttachmentEndpointRequest{
        Model: shared.AttachmentRequest{
            AttachmentType: shared.AttachmentRequestAttachmentTypeResume.ToPointer(),
            Candidate: mergeatsgo.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
            FileName: mergeatsgo.String("Candidate Resume"),
            FileURL: mergeatsgo.String("http://alturl.com/p749b"),
            IntegrationParams: map[string]interface{}{
                "unique_integration_field": "unique_integration_field_value",
            },
            LinkedAccountParams: map[string]interface{}{
                "unique_linked_account_field": "unique_linked_account_field_value",
            },
        },
        RemoteUserID: "<value>",
    }

    var xAccountToken string = "<value>"

    var isDebugMode *bool = mergeatsgo.Bool(false)

    var runAsync *bool = mergeatsgo.Bool(false)

    ctx := context.Background()
    res, err := s.Attachments.Create(ctx, attachmentEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }
    if res.AttachmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `attachmentEndpointRequest`                                                              | [shared.AttachmentEndpointRequest](../../pkg/models/shared/attachmentendpointrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `xAccountToken`                                                                          | *string*                                                                                 | :heavy_check_mark:                                                                       | Token identifying the end user.                                                          |
| `isDebugMode`                                                                            | **bool*                                                                                  | :heavy_minus_sign:                                                                       | Whether to include debug fields (such as log file links) in the response.                |
| `runAsync`                                                                               | **bool*                                                                                  | :heavy_minus_sign:                                                                       | Whether or not third-party updates should be run asynchronously.                         |


### Response

**[*operations.AttachmentsCreateResponse](../../pkg/models/operations/attachmentscreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Returns a list of `Attachment` objects.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Attachments.List(ctx, operations.AttachmentsListRequest{
        XAccountToken: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedAttachmentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.AttachmentsListRequest](../../pkg/models/operations/attachmentslistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.AttachmentsListResponse](../../pkg/models/operations/attachmentslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Returns an `Attachment` object with the given `id`.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Attachments.Retrieve(ctx, operations.AttachmentsRetrieveRequest{
        XAccountToken: "<value>",
        ID: "5fea5659-1081-4ad2-8d60-4c8e92b241fa",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AttachmentsRetrieveRequest](../../pkg/models/operations/attachmentsretrieverequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.AttachmentsRetrieveResponse](../../pkg/models/operations/attachmentsretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RetrievePostMetadata

Returns metadata for `Attachment` POSTs.

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

    ctx := context.Background()
    res, err := s.Attachments.RetrievePostMetadata(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xAccountToken`                                       | *string*                                              | :heavy_check_mark:                                    | Token identifying the end user.                       |


### Response

**[*operations.AttachmentsMetaPostRetrieveResponse](../../pkg/models/operations/attachmentsmetapostretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
