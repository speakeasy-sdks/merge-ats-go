# Attachments

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    operationSecurity := operations.AttachmentsCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Attachments.Create(ctx, operations.AttachmentsCreateRequest{
        AttachmentEndpointRequest: shared.AttachmentEndpointRequest{
            Model: shared.AttachmentRequest{
                AttachmentType: shared.AttachmentRequestAttachmentTypeResume.ToPointer(),
                Candidate: ats.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
                FileName: ats.String("Candidate Resume"),
                FileURL: ats.String("http://alturl.com/p749b"),
                IntegrationParams: map[string]interface{}{
                    "magnam": "et",
                },
                LinkedAccountParams: map[string]interface{}{
                    "ullam": "provident",
                    "quos": "sint",
                    "accusantium": "mollitia",
                },
            },
            RemoteUserID: "reiciendis",
        },
        XAccountToken: "mollitia",
        IsDebugMode: ats.Bool(false),
        RunAsync: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AttachmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AttachmentsCreateRequest](../../models/operations/attachmentscreaterequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.AttachmentsCreateSecurity](../../models/operations/attachmentscreatesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.AttachmentsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Attachments.List(ctx, operations.AttachmentsListRequest{
        XAccountToken: "ad",
        CandidateID: ats.String("eum"),
        CreatedAfter: types.MustTimeFromString("2022-02-07T18:15:06.372Z"),
        CreatedBefore: types.MustTimeFromString("2022-08-19T20:09:28.183Z"),
        Cursor: ats.String("quasi"),
        Expand: operations.AttachmentsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-01-06T19:47:24.047Z"),
        ModifiedBefore: types.MustTimeFromString("2022-03-21T22:14:24.691Z"),
        PageSize: ats.Int64(806194),
        RemoteFields: operations.AttachmentsListRemoteFieldsAttachmentType.ToPointer(),
        RemoteID: ats.String("deleniti"),
        ShowEnumOrigins: operations.AttachmentsListShowEnumOriginsAttachmentType.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedAttachmentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.AttachmentsListRequest](../../models/operations/attachmentslistrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.AttachmentsListSecurity](../../models/operations/attachmentslistsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.AttachmentsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Attachments.Retrieve(ctx, operations.AttachmentsRetrieveRequest{
        XAccountToken: "facilis",
        Expand: operations.AttachmentsRetrieveExpandCandidate.ToPointer(),
        ID: "711e5b7f-d2ed-4028-921c-ddc692601fb5",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.AttachmentsRetrieveRemoteFieldsAttachmentType.ToPointer(),
        ShowEnumOrigins: operations.AttachmentsRetrieveShowEnumOriginsAttachmentType.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AttachmentsRetrieveRequest](../../models/operations/attachmentsretrieverequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.AttachmentsRetrieveSecurity](../../models/operations/attachmentsretrievesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.AttachmentsMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Attachments.RetrievePostMetadata(ctx, operations.AttachmentsMetaPostRetrieveRequest{
        XAccountToken: "voluptate",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AttachmentsMetaPostRetrieveRequest](../../models/operations/attachmentsmetapostretrieverequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.AttachmentsMetaPostRetrieveSecurity](../../models/operations/attachmentsmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.AttachmentsMetaPostRetrieveResponse](../../models/operations/attachmentsmetapostretrieveresponse.md), error**

