# Candidates
(*Candidates*)

### Available Operations

* [Create](#create) - Creates a `Candidate` object with the given values.
* [IgnoreCreate](#ignorecreate) - Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
* [List](#list) - Returns a list of `Candidate` objects.
* [Retrieve](#retrieve) - Returns a `Candidate` object with the given `id`.
* [RetrievePatchMetadata](#retrievepatchmetadata) - Returns metadata for `Candidate` PATCHs.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `Candidate` POSTs.
* [Update](#update) - Updates a `Candidate` object with the given `id`.

## Create

Creates a `Candidate` object with the given values.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    candidateEndpointRequest := shared.CandidateEndpointRequest{
        Model: shared.CandidateRequest{
            Applications: []string{
                "29eb9867-ce2a-403f-b8ce-f2844b89f078",
                "b4d08e5c-de00-4d64-a29f-66addac9af99",
                "4ff877d2-fb3e-4a5b-a7a5-168ddf2ffa56",
            },
            Attachments: []string{
                "bea08964-32b4-4a20-8bb4-2612ba09de1d",
            },
            CanEmail: mergeatsgo.Bool(true),
            Company: mergeatsgo.String("Columbia Dining App."),
            EmailAddresses: []shared.EmailAddressRequest{
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "unique_integration_field": "unique_integration_field_value",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "unique_linked_account_field": "unique_linked_account_field_value",
                    },
                    Value: mergeatsgo.String("hello@merge.dev"),
                },
            },
            FirstName: mergeatsgo.String("Gil"),
            IntegrationParams: map[string]interface{}{
                "unique_integration_field": "unique_integration_field_value",
            },
            IsPrivate: mergeatsgo.Bool(true),
            LastInteractionAt: types.MustNewTimeFromString("2021-10-17T00:00:00Z"),
            LastName: mergeatsgo.String("Feig"),
            LinkedAccountParams: map[string]interface{}{
                "unique_linked_account_field": "unique_linked_account_field_value",
            },
            Locations: []string{
                "San Francisco",
                "New York",
                "Miami",
            },
            PhoneNumbers: []shared.PhoneNumberRequest{
                shared.PhoneNumberRequest{
                    IntegrationParams: map[string]interface{}{
                        "unique_integration_field": "unique_integration_field_value",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "unique_linked_account_field": "unique_linked_account_field_value",
                    },
                    PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeMobile.ToPointer(),
                    Value: mergeatsgo.String("+1234567890"),
                },
            },
            RemoteTemplateID: mergeatsgo.String("92830948203"),
            Tags: []string{
                "High-Priority",
            },
            Title: mergeatsgo.String("Software Engineer"),
            Urls: []shared.URLRequest{
                shared.URLRequest{
                    IntegrationParams: map[string]interface{}{
                        "unique_integration_field": "unique_integration_field_value",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "unique_linked_account_field": "unique_linked_account_field_value",
                    },
                    URLType: shared.URLRequestURLTypeBlog.ToPointer(),
                    Value: mergeatsgo.String("http://alturl.com/p749b"),
                },
            },
        },
        RemoteUserID: "<value>",
    }

    var xAccountToken string = "<value>"

    var isDebugMode *bool = mergeatsgo.Bool(false)

    var runAsync *bool = mergeatsgo.Bool(false)

    ctx := context.Background()
    res, err := s.Candidates.Create(ctx, candidateEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }
    if res.CandidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `candidateEndpointRequest`                                                             | [shared.CandidateEndpointRequest](../../pkg/models/shared/candidateendpointrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `xAccountToken`                                                                        | *string*                                                                               | :heavy_check_mark:                                                                     | Token identifying the end user.                                                        |
| `isDebugMode`                                                                          | **bool*                                                                                | :heavy_minus_sign:                                                                     | Whether to include debug fields (such as log file links) in the response.              |
| `runAsync`                                                                             | **bool*                                                                                | :heavy_minus_sign:                                                                     | Whether or not third-party updates should be run asynchronously.                       |


### Response

**[*operations.CandidatesCreateResponse](../../pkg/models/operations/candidatescreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## IgnoreCreate

Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.

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


    ignoreCommonModelRequest := shared.IgnoreCommonModelRequest{
        Message: mergeatsgo.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
        Reason: shared.ReasonGeneralCustomerRequest,
    }

    var xAccountToken string = "<value>"

    var modelID string = "c7c17a75-bf57-4b4d-967c-fdddc7d08842"

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, ignoreCommonModelRequest, xAccountToken, modelID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `ignoreCommonModelRequest`                                                             | [shared.IgnoreCommonModelRequest](../../pkg/models/shared/ignorecommonmodelrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `xAccountToken`                                                                        | *string*                                                                               | :heavy_check_mark:                                                                     | Token identifying the end user.                                                        |
| `modelID`                                                                              | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |


### Response

**[*operations.CandidatesIgnoreCreateResponse](../../pkg/models/operations/candidatesignorecreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Returns a list of `Candidate` objects.

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
    res, err := s.Candidates.List(ctx, operations.CandidatesListRequest{
        XAccountToken: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaginatedCandidateList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CandidatesListRequest](../../pkg/models/operations/candidateslistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CandidatesListResponse](../../pkg/models/operations/candidateslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Retrieve

Returns a `Candidate` object with the given `id`.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var xAccountToken string = "<value>"

    var id string = "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    var expand *operations.CandidatesRetrieveQueryParamExpand = operations.CandidatesRetrieveQueryParamExpandApplications.ToPointer()

    var includeRemoteData *bool = mergeatsgo.Bool(false)

    ctx := context.Background()
    res, err := s.Candidates.Retrieve(ctx, xAccountToken, id, expand, includeRemoteData)
    if err != nil {
        log.Fatal(err)
    }
    if res.Candidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `xAccountToken`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | Token identifying the end user.                                                                                        |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `expand`                                                                                                               | [*operations.CandidatesRetrieveQueryParamExpand](../../pkg/models/operations/candidatesretrievequeryparamexpand.md)    | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
| `includeRemoteData`                                                                                                    | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                       |


### Response

**[*operations.CandidatesRetrieveResponse](../../pkg/models/operations/candidatesretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RetrievePatchMetadata

Returns metadata for `Candidate` PATCHs.

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

    var id string = "1e4cbcd8-c9a8-4659-9682-a54bbda71d58"

    ctx := context.Background()
    res, err := s.Candidates.RetrievePatchMetadata(ctx, xAccountToken, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.CandidatesMetaPatchRetrieveResponse](../../pkg/models/operations/candidatesmetapatchretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RetrievePostMetadata

Returns metadata for `Candidate` POSTs.

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
    res, err := s.Candidates.RetrievePostMetadata(ctx, xAccountToken)
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

**[*operations.CandidatesMetaPostRetrieveResponse](../../pkg/models/operations/candidatesmetapostretrieveresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Updates a `Candidate` object with the given `id`.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Candidates.Update(ctx, operations.PartialUpdateRequest{
        PatchedCandidateEndpointRequest: shared.PatchedCandidateEndpointRequest{
            Model: shared.PatchedCandidateRequest{
                Applications: []string{
                    "29eb9867-ce2a-403f-b8ce-f2844b89f078",
                    "b4d08e5c-de00-4d64-a29f-66addac9af99",
                    "4ff877d2-fb3e-4a5b-a7a5-168ddf2ffa56",
                },
                Attachments: []string{
                    "bea08964-32b4-4a20-8bb4-2612ba09de1d",
                },
                CanEmail: mergeatsgo.Bool(true),
                Company: mergeatsgo.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "unique_integration_field": "unique_integration_field_value",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "unique_linked_account_field": "unique_linked_account_field_value",
                        },
                        Value: mergeatsgo.String("hello@merge.dev"),
                    },
                },
                FirstName: mergeatsgo.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "unique_integration_field": "unique_integration_field_value",
                },
                IsPrivate: mergeatsgo.Bool(true),
                LastInteractionAt: types.MustNewTimeFromString("2021-10-17T00:00:00Z"),
                LastName: mergeatsgo.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "unique_linked_account_field": "unique_linked_account_field_value",
                },
                Locations: []string{
                    "San Francisco",
                    "New York",
                    "Miami",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "unique_integration_field": "unique_integration_field_value",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "unique_linked_account_field": "unique_linked_account_field_value",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeMobile.ToPointer(),
                        Value: mergeatsgo.String("+1234567890"),
                    },
                },
                RemoteTemplateID: mergeatsgo.String("92830948203"),
                Tags: []string{
                    "High-Priority",
                },
                Title: mergeatsgo.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "unique_integration_field": "unique_integration_field_value",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "unique_linked_account_field": "unique_linked_account_field_value",
                        },
                        URLType: shared.URLRequestURLTypeBlog.ToPointer(),
                        Value: mergeatsgo.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "<value>",
        },
        XAccountToken: "<value>",
        ID: "d0905bf4-aa77-4f20-8e77-54c352acfe54",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CandidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PartialUpdateRequest](../../pkg/models/operations/partialupdaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PartialUpdateResponse](../../pkg/models/operations/partialupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
