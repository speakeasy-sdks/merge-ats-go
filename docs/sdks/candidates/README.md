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
    candidateEndpointRequest := shared.CandidateEndpointRequest{
        Model: shared.CandidateRequest{
            Applications: []string{
                "2516fe4c-8b71-41e5-b7fd-2ed028921cdd",
            },
            Attachments: []string{
                "c692601f-b576-4b0d-9f0d-30c5fbb25870",
            },
            CanEmail: mergeatsgo.Bool(true),
            Company: mergeatsgo.String("Columbia Dining App."),
            EmailAddresses: []shared.EmailAddressRequest{
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "quis": "nesciunt",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "eos": "perferendis",
                    },
                    Value: mergeatsgo.String("merge_is_hiring@merge.dev"),
                },
            },
            FirstName: mergeatsgo.String("Gil"),
            IntegrationParams: map[string]interface{}{
                "dolores": "minus",
            },
            IsPrivate: mergeatsgo.Bool(true),
            LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
            LastName: mergeatsgo.String("Feig"),
            LinkedAccountParams: map[string]interface{}{
                "quam": "dolor",
            },
            Locations: []string{
                "vero",
            },
            PhoneNumbers: []shared.PhoneNumberRequest{
                shared.PhoneNumberRequest{
                    IntegrationParams: map[string]interface{}{
                        "nostrum": "hic",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "recusandae": "omnis",
                    },
                    PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                    Value: mergeatsgo.String("+3198675309"),
                },
            },
            RemoteTemplateID: mergeatsgo.String("92830948203"),
            Tags: []string{
                "facilis",
            },
            Title: mergeatsgo.String("Software Engineer"),
            Urls: []shared.URLRequest{
                shared.URLRequest{
                    IntegrationParams: map[string]interface{}{
                        "perspiciatis": "voluptatem",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "porro": "consequuntur",
                    },
                    URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                    Value: mergeatsgo.String("http://alturl.com/p749b"),
                },
            },
        },
        RemoteUserID: "blanditiis",
    }
    xAccountToken := "error"
    isDebugMode := false
    runAsync := false

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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    ignoreCommonModelRequest := shared.IgnoreCommonModelRequest{
        Message: mergeatsgo.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
        Reason: shared.IgnoreCommonModelRequestReasonGeneralCustomerRequest,
    }
    xAccountToken := "eaque"
    modelID := "9b3fe49a-8d9c-4bf4-8633-323f9b77f3a4"

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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidates.List(ctx, operations.CandidatesListRequest{
        XAccountToken: "veritatis",
        CreatedAfter: types.MustTimeFromString("2022-12-11T09:46:30.457Z"),
        CreatedBefore: types.MustTimeFromString("2022-07-06T22:32:29.592Z"),
        Cursor: mergeatsgo.String("quaerat"),
        EmailAddresses: mergeatsgo.String("accusamus"),
        Expand: operations.CandidatesListExpandAttachments.ToPointer(),
        FirstName: mergeatsgo.String("Weston"),
        IncludeDeletedData: mergeatsgo.Bool(false),
        IncludeRemoteData: mergeatsgo.Bool(false),
        LastName: mergeatsgo.String("Hodkiewicz"),
        ModifiedAfter: types.MustTimeFromString("2022-08-22T21:20:36.034Z"),
        ModifiedBefore: types.MustTimeFromString("2022-12-13T23:37:42.244Z"),
        PageSize: mergeatsgo.Int64(854614),
        RemoteID: mergeatsgo.String("ab"),
        Tags: mergeatsgo.String("soluta"),
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
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    xAccountToken := "dolorum"
    id := "77a89ebf-737a-4e42-83ce-5e6a95d8a0d4"
    expand := operations.CandidatesRetrieveExpandApplications
    includeRemoteData := false

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
| `expand`                                                                                                               | [*operations.CandidatesRetrieveExpand](../../models/operations/candidatesretrieveexpand.md)                            | :heavy_minus_sign:                                                                                                     | Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces. |
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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    xAccountToken := "vel"
    id := "ce2af7a7-3cf3-4be4-93f8-70b326b5a734"

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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    xAccountToken := "qui"

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
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Candidates.Update(ctx, operations.PartialUpdateRequest{
        PatchedCandidateEndpointRequest: shared.PatchedCandidateEndpointRequest{
            Model: shared.PatchedCandidateRequest{
                Applications: []string{
                    "9cdb1a84-22bb-4679-9232-2715bf0cbb1e",
                },
                Attachments: []string{
                    "31b8b90f-3443-4a11-88e0-adcf4b921879",
                },
                CanEmail: mergeatsgo.Bool(true),
                Company: mergeatsgo.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "voluptatibus": "quisquam",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "vero": "omnis",
                        },
                        Value: mergeatsgo.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: mergeatsgo.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "quis": "ipsum",
                },
                IsPrivate: mergeatsgo.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: mergeatsgo.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "delectus": "voluptate",
                },
                Locations: []string{
                    "consectetur",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "vero": "tenetur",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "dignissimos": "hic",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: mergeatsgo.String("+3198675309"),
                    },
                },
                RemoteTemplateID: mergeatsgo.String("92830948203"),
                Tags: []string{
                    "distinctio",
                },
                Title: mergeatsgo.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "quod": "odio",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "similique": "facilis",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: mergeatsgo.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "vero",
        },
        XAccountToken: "ducimus",
        ID: "4dd39c0f-5d2c-4ff7-870a-45626d436813",
        IsDebugMode: mergeatsgo.Bool(false),
        RunAsync: mergeatsgo.Bool(false),
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

