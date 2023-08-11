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
    attachmentEndpointRequest := shared.AttachmentEndpointRequest{
        Model: shared.AttachmentRequest{
            AttachmentType: shared.AttachmentRequestAttachmentTypeResume.ToPointer(),
            Candidate: ats.String("2872ba14-4084-492b-be96-e5eee6fc33ef"),
            FileName: ats.String("Candidate Resume"),
            FileURL: ats.String("http://alturl.com/p749b"),
            IntegrationParams: map[string]interface{}{
                "blanditiis": "deleniti",
                "sapiente": "amet",
                "deserunt": "nisi",
            },
            LinkedAccountParams: map[string]interface{}{
                "natus": "omnis",
                "molestiae": "perferendis",
            },
        },
        RemoteUserID: "nihil",
    }
    xAccountToken := "magnam"
    isDebugMode := false
    runAsync := false
    operationSecurity := operations.AttachmentsCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Attachments.Create(ctx, operationSecurity, attachmentEndpointRequest, xAccountToken, isDebugMode, runAsync)
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
| `security`                                                                                   | [operations.AttachmentsCreateSecurity](../../models/operations/attachmentscreatesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |
| `attachmentEndpointRequest`                                                                  | [shared.AttachmentEndpointRequest](../../models/shared/attachmentendpointrequest.md)         | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `xAccountToken`                                                                              | *string*                                                                                     | :heavy_check_mark:                                                                           | Token identifying the end user.                                                              |
| `isDebugMode`                                                                                | **bool*                                                                                      | :heavy_minus_sign:                                                                           | Whether to include debug fields (such as log file links) in the response.                    |
| `runAsync`                                                                                   | **bool*                                                                                      | :heavy_minus_sign:                                                                           | Whether or not third-party updates should be run asynchronously.                             |


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
        XAccountToken: "distinctio",
        CandidateID: ats.String("id"),
        CreatedAfter: types.MustTimeFromString("2022-09-17T02:55:11.783Z"),
        CreatedBefore: types.MustTimeFromString("2022-05-20T10:11:05.115Z"),
        Cursor: ats.String("nobis"),
        Expand: operations.AttachmentsListExpandCandidate.ToPointer(),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-02-14T08:45:31.579Z"),
        ModifiedBefore: types.MustTimeFromString("2022-11-24T10:55:00.183Z"),
        PageSize: ats.Int64(298282),
        RemoteFields: operations.AttachmentsListRemoteFieldsAttachmentType.ToPointer(),
        RemoteID: ats.String("et"),
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
        XAccountToken: "excepturi",
        Expand: operations.AttachmentsRetrieveExpandCandidate.ToPointer(),
        ID: "59890afa-563e-4251-afe4-c8b711e5b7fd",
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
    xAccountToken := "sed"
    operationSecurity := operations.AttachmentsMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Attachments.RetrievePostMetadata(ctx, operationSecurity, xAccountToken)
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
| `security`                                                                                                       | [operations.AttachmentsMetaPostRetrieveSecurity](../../models/operations/attachmentsmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `xAccountToken`                                                                                                  | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Token identifying the end user.                                                                                  |


### Response

**[*operations.AttachmentsMetaPostRetrieveResponse](../../models/operations/attachmentsmetapostretrieveresponse.md), error**

