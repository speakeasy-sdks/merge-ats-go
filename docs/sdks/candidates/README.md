# Candidates

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    candidateEndpointRequest := shared.CandidateEndpointRequest{
        Model: shared.CandidateRequest{
            Applications: []string{
                "90afa563-e251-46fe-8c8b-711e5b7fd2ed",
            },
            Attachments: []string{
                "028921cd-dc69-4260-9fb5-76b0d5f0d30c",
            },
            CanEmail: ats.Bool(true),
            Company: ats.String("Columbia Dining App."),
            EmailAddresses: []shared.EmailAddressRequest{
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "corporis": "hic",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "libero": "nobis",
                    },
                    Value: ats.String("merge_is_hiring@merge.dev"),
                },
            },
            FirstName: ats.String("Gil"),
            IntegrationParams: map[string]interface{}{
                "dolores": "quis",
            },
            IsPrivate: ats.Bool(true),
            LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
            LastName: ats.String("Feig"),
            LinkedAccountParams: map[string]interface{}{
                "totam": "dignissimos",
            },
            Locations: []string{
                "eaque",
            },
            PhoneNumbers: []shared.PhoneNumberRequest{
                shared.PhoneNumberRequest{
                    IntegrationParams: map[string]interface{}{
                        "quis": "nesciunt",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "eos": "perferendis",
                    },
                    PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                    Value: ats.String("+3198675309"),
                },
            },
            RemoteTemplateID: ats.String("92830948203"),
            Tags: []string{
                "dolores",
            },
            Title: ats.String("Software Engineer"),
            Urls: []shared.URLRequest{
                shared.URLRequest{
                    IntegrationParams: map[string]interface{}{
                        "minus": "quam",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "dolor": "vero",
                    },
                    URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                    Value: ats.String("http://alturl.com/p749b"),
                },
            },
        },
        RemoteUserID: "nostrum",
    }
    xAccountToken := "hic"
    isDebugMode := false
    runAsync := false
    operationSecurity := operations.CandidatesCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.Create(ctx, operationSecurity, candidateEndpointRequest, xAccountToken, isDebugMode, runAsync)
    if err != nil {
        log.Fatal(err)
    }

    if res.CandidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `security`                                                                                 | [operations.CandidatesCreateSecurity](../../models/operations/candidatescreatesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `candidateEndpointRequest`                                                                 | [shared.CandidateEndpointRequest](../../models/shared/candidateendpointrequest.md)         | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `xAccountToken`                                                                            | *string*                                                                                   | :heavy_check_mark:                                                                         | Token identifying the end user.                                                            |
| `isDebugMode`                                                                              | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Whether to include debug fields (such as log file links) in the response.                  |
| `runAsync`                                                                                 | **bool*                                                                                    | :heavy_minus_sign:                                                                         | Whether or not third-party updates should be run asynchronously.                           |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := ats.New()
    ignoreCommonModelRequest := shared.IgnoreCommonModelRequest{
        Message: ats.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
        Reason: shared.IgnoreCommonModelRequestReasonGeneralCustomerRequest,
    }
    xAccountToken := "recusandae"
    modelID := "9b90c289-09b3-4fe4-9a8d-9cbf48633323"
    operationSecurity := operations.CandidatesIgnoreCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, operationSecurity, ignoreCommonModelRequest, xAccountToken, modelID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.CandidatesIgnoreCreateSecurity](../../models/operations/candidatesignorecreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `ignoreCommonModelRequest`                                                                             | [shared.IgnoreCommonModelRequest](../../models/shared/ignorecommonmodelrequest.md)                     | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `xAccountToken`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | Token identifying the end user.                                                                        |
| `modelID`                                                                                              | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.CandidatesListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.List(ctx, operations.CandidatesListRequest{
        XAccountToken: "hic",
        CreatedAfter: types.MustTimeFromString("2021-07-10T03:04:11.898Z"),
        CreatedBefore: types.MustTimeFromString("2022-07-05T23:34:50.715Z"),
        Cursor: ats.String("reiciendis"),
        EmailAddresses: ats.String("amet"),
        Expand: operations.CandidatesListExpandAttachments.ToPointer(),
        FirstName: ats.String("Deonte"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        LastName: ats.String("Bogan"),
        ModifiedAfter: types.MustTimeFromString("2022-12-11T09:46:30.457Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-06T22:32:29.592Z"),
        PageSize: ats.Int64(311796),
        RemoteID: ats.String("accusamus"),
        Tags: ats.String("quidem"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedCandidateList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CandidatesListRequest](../../models/operations/candidateslistrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CandidatesListSecurity](../../models/operations/candidateslistsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "voluptatibus"
    id := "69280d1b-a77a-489e-bf73-7ae4203ce5e6"
    expand := operations.CandidatesRetrieveExpandApplicationsAttachments
    includeRemoteData := false
    operationSecurity := operations.CandidatesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.Retrieve(ctx, operationSecurity, xAccountToken, id, expand, includeRemoteData)
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
| `security`                                                                                                             | [operations.CandidatesRetrieveSecurity](../../models/operations/candidatesretrievesecurity.md)                         | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |
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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "provident"
    id := "5d8a0d44-6ce2-4af7-a73c-f3be453f870b"
    operationSecurity := operations.CandidatesMetaPatchRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.RetrievePatchMetadata(ctx, operationSecurity, xAccountToken, id)
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
| `security`                                                                                                       | [operations.CandidatesMetaPatchRetrieveSecurity](../../models/operations/candidatesmetapatchretrievesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `xAccountToken`                                                                                                  | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Token identifying the end user.                                                                                  |
| `id`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    xAccountToken := "neque"
    operationSecurity := operations.CandidatesMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.RetrievePostMetadata(ctx, operationSecurity, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.MetaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `security`                                                                                                     | [operations.CandidatesMetaPostRetrieveSecurity](../../models/operations/candidatesmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `xAccountToken`                                                                                                | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Token identifying the end user.                                                                                |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.PartialUpdateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.Update(ctx, operations.PartialUpdateRequest{
        PatchedCandidateEndpointRequest: shared.PatchedCandidateEndpointRequest{
            Model: shared.PatchedCandidateRequest{
                Applications: []string{
                    "26b5a734-29cd-4b1a-8422-bb679d232271",
                },
                Attachments: []string{
                    "5bf0cbb1-e31b-48b9-8f34-43a1108e0adc",
                },
                CanEmail: ats.Bool(true),
                Company: ats.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "doloribus": "ut",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "facilis": "cupiditate",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: ats.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "qui": "quae",
                },
                IsPrivate: ats.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: ats.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "laudantium": "odio",
                },
                Locations: []string{
                    "occaecati",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "voluptatibus": "quisquam",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "vero": "omnis",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                },
                RemoteTemplateID: ats.String("92830948203"),
                Tags: []string{
                    "quis",
                },
                Title: ats.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "ipsum": "delectus",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "voluptate": "consectetur",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "vero",
        },
        XAccountToken: "tenetur",
        ID: "7fbc7abd-74dd-439c-8f5d-2cff7c70a456",
        IsDebugMode: ats.Bool(false),
        RunAsync: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CandidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PartialUpdateRequest](../../models/operations/partialupdaterequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.PartialUpdateSecurity](../../models/operations/partialupdatesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.PartialUpdateResponse](../../models/operations/partialupdateresponse.md), error**

