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
    attachmentEndpointRequest := shared.AttachmentEndpointRequest{
        Model: shared.AttachmentRequest{
            AttachmentType: shared.AttachmentRequestAttachmentTypeResume.ToPointer(),
            Candidate: mergeatsgo.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
            FileName: mergeatsgo.String("Candidate Resume"),
            FileURL: mergeatsgo.String("http://alturl.com/p749b"),
            IntegrationParams: map[string]interface{}{
                "odio": "bluetooth",
            },
            LinkedAccountParams: map[string]interface{}{
                "nulla": "Money",
            },
        },
        RemoteUserID: "Cambridgeshire grey technology",
    }
    xAccountToken := "East"
    isDebugMode := false
    runAsync := false

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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `attachmentEndpointRequest`                                                          | [shared.AttachmentEndpointRequest](../../models/shared/attachmentendpointrequest.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `xAccountToken`                                                                      | *string*                                                                             | :heavy_check_mark:                                                                   | Token identifying the end user.                                                      |
| `isDebugMode`                                                                        | **bool*                                                                              | :heavy_minus_sign:                                                                   | Whether to include debug fields (such as log file links) in the response.            |
| `runAsync`                                                                           | **bool*                                                                              | :heavy_minus_sign:                                                                   | Whether or not third-party updates should be run asynchronously.                     |


### Response

**[*operations.AttachmentsCreateResponse](../../models/operations/attachmentscreateresponse.md), error**


## List

Returns a list of `Attachment` objects.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Attachments.List(ctx, operations.AttachmentsListRequest{
        XAccountToken: "Northeast Metal Canada",
        CandidateID: mergeatsgo.String("Data Response West"),
        CreatedAfter: types.MustTimeFromString("2023-02-26T13:00:25.189Z"),
        CreatedBefore: types.MustTimeFromString("2023-12-20T19:28:33.339Z"),
        Cursor: mergeatsgo.String("primary"),
        Expand: operations.AttachmentsListExpandCandidate.ToPointer(),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-09-01T09:24:49.063Z"),
        ModifiedBefore: types.MustTimeFromString("2023-01-24T07:34:45.759Z"),
        PageSize: mergeatsgo.Int64(559247),
        RemoteFields: operations.AttachmentsListRemoteFieldsAttachmentType.ToPointer(),
        RemoteID: mergeatsgo.String("orchid synergies"),
        ShowEnumOrigins: operations.AttachmentsListShowEnumOriginsAttachmentType.ToPointer(),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.AttachmentsListRequest](../../models/operations/attachmentslistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.AttachmentsListResponse](../../models/operations/attachmentslistresponse.md), error**


## Retrieve

Returns an `Attachment` object with the given `id`.

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

    ctx := context.Background()
    res, err := s.Attachments.Retrieve(ctx, operations.AttachmentsRetrieveRequest{
        XAccountToken: "tracksuit Markets",
        Expand: operations.AttachmentsRetrieveExpandCandidate.ToPointer(),
        ID: "1081ad20-d604-4c8e-92b2-41fa379087a1",
        IncludeRemoteData: mergeatsgo.Bool(false),
        RemoteFields: operations.AttachmentsRetrieveRemoteFieldsAttachmentType.ToPointer(),
        ShowEnumOrigins: operations.AttachmentsRetrieveShowEnumOriginsAttachmentType.ToPointer(),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AttachmentsRetrieveRequest](../../models/operations/attachmentsretrieverequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.AttachmentsRetrieveResponse](../../models/operations/attachmentsretrieveresponse.md), error**


## RetrievePostMetadata

Returns metadata for `Attachment` POSTs.

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
    xAccountToken := "Borders"

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

**[*operations.AttachmentsMetaPostRetrieveResponse](../../models/operations/attachmentsmetapostretrieveresponse.md), error**

