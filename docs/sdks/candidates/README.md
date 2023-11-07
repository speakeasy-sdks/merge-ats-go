# Candidates
(*.Candidates*)

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
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
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
                        "unique_integration_field": "string",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "unique_linked_account_field": "string",
                    },
                    Value: mergeatsgo.String("merge_is_hiring@merge.dev"),
                },
            },
            FirstName: mergeatsgo.String("Gil"),
            IntegrationParams: map[string]interface{}{
                "unique_integration_field": "string",
            },
            IsPrivate: mergeatsgo.Bool(true),
            LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
            LastName: mergeatsgo.String("Feig"),
            LinkedAccountParams: map[string]interface{}{
                "unique_linked_account_field": "string",
            },
            Locations: []string{
                "San Francisco",
                "New York",
                "Miami",
            },
            PhoneNumbers: []shared.PhoneNumberRequest{
                shared.PhoneNumberRequest{
                    IntegrationParams: map[string]interface{}{
                        "unique_integration_field": "string",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "unique_linked_account_field": "string",
                    },
                    PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                    Value: mergeatsgo.String("+3198675309"),
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
                        "unique_integration_field": "string",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "unique_linked_account_field": "string",
                    },
                    URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                    Value: mergeatsgo.String("http://alturl.com/p749b"),
                },
            },
        },
        RemoteUserID: "string",
    }

    var xAccountToken string = "string"

    var isDebugMode *bool = false

    var runAsync *bool = false

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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `candidateEndpointRequest`                                                         | [shared.CandidateEndpointRequest](../../models/shared/candidateendpointrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `xAccountToken`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | Token identifying the end user.                                                    |
| `isDebugMode`                                                                      | **bool*                                                                            | :heavy_minus_sign:                                                                 | Whether to include debug fields (such as log file links) in the response.          |
| `runAsync`                                                                         | **bool*                                                                            | :heavy_minus_sign:                                                                 | Whether or not third-party updates should be run asynchronously.                   |


### Response

**[*operations.CandidatesCreateResponse](../../models/operations/candidatescreateresponse.md), error**


## IgnoreCreate

Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    ignoreCommonModelRequest := shared.IgnoreCommonModelRequest{
        Message: mergeatsgo.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
        Reason: shared.ReasonGeneralCustomerRequest,
    }

    var xAccountToken string = "string"

    var modelID string = "c7c17a75-bf57-4b4d-967c-fdddc7d08842"

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, ignoreCommonModelRequest, xAccountToken, modelID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `ignoreCommonModelRequest`                                                         | [shared.IgnoreCommonModelRequest](../../models/shared/ignorecommonmodelrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `xAccountToken`                                                                    | *string*                                                                           | :heavy_check_mark:                                                                 | Token identifying the end user.                                                    |
| `modelID`                                                                          | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |


### Response

**[*operations.CandidatesIgnoreCreateResponse](../../models/operations/candidatesignorecreateresponse.md), error**


## List

Returns a list of `Candidate` objects.

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
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Candidates.List(ctx, operations.CandidatesListRequest{
        XAccountToken: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CandidatesListRequest](../../models/operations/candidateslistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CandidatesListResponse](../../models/operations/candidateslistresponse.md), error**


## Retrieve

Returns a `Candidate` object with the given `id`.

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
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "string"

    var id string = "5fea5659-1081-4ad2-8d60-4c8e92b241fa"

    var expand *operations.CandidatesRetrieveQueryParamExpand = operations.CandidatesRetrieveQueryParamExpandApplications

    var includeRemoteData *bool = false

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
| `expand`                                                                                                               | [*operations.CandidatesRetrieveQueryParamExpand](../../models/operations/candidatesretrievequeryparamexpand.md)        | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
| `includeRemoteData`                                                                                                    | **bool*                                                                                                                | :heavy_minus_sign:                                                                                                     | Whether to include the original data Merge fetched from the third-party to produce these models.                       |


### Response

**[*operations.CandidatesRetrieveResponse](../../models/operations/candidatesretrieveresponse.md), error**


## RetrievePatchMetadata

Returns metadata for `Candidate` PATCHs.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "string"

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

**[*operations.CandidatesMetaPatchRetrieveResponse](../../models/operations/candidatesmetapatchretrieveresponse.md), error**


## RetrievePostMetadata

Returns metadata for `Candidate` POSTs.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "string"

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

**[*operations.CandidatesMetaPostRetrieveResponse](../../models/operations/candidatesmetapostretrieveresponse.md), error**


## Update

Updates a `Candidate` object with the given `id`.

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
        mergeatsgo.WithSecurity(""),
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
                            "unique_integration_field": "string",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "unique_linked_account_field": "string",
                        },
                        Value: mergeatsgo.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: mergeatsgo.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "unique_integration_field": "string",
                },
                IsPrivate: mergeatsgo.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: mergeatsgo.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "unique_linked_account_field": "string",
                },
                Locations: []string{
                    "San Francisco",
                    "New York",
                    "Miami",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "unique_integration_field": "string",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "unique_linked_account_field": "string",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: mergeatsgo.String("+3198675309"),
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
                            "unique_integration_field": "string",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "unique_linked_account_field": "string",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: mergeatsgo.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "string",
        },
        XAccountToken: "string",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PartialUpdateRequest](../../models/operations/partialupdaterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PartialUpdateResponse](../../models/operations/partialupdateresponse.md), error**

