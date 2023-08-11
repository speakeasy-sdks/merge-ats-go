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
                "028921cd-dc69-4260-9fb5-76b0d5f0d30c",
                "5fbb2587-0532-402c-b3d5-fe9b90c28909",
                "b3fe49a8-d9cb-4f48-a333-23f9b77f3a41",
                "00674ebf-6928-40d1-ba77-a89ebf737ae4",
            },
            Attachments: []string{
                "03ce5e6a-95d8-4a0d-846c-e2af7a73cf3b",
            },
            CanEmail: ats.Bool(true),
            Company: ats.String("Columbia Dining App."),
            EmailAddresses: []shared.EmailAddressRequest{
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "enim": "dolorem",
                        "sapiente": "totam",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "sit": "expedita",
                        "neque": "sed",
                    },
                    Value: ats.String("merge_is_hiring@merge.dev"),
                },
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "libero": "voluptas",
                        "deserunt": "quam",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "incidunt": "qui",
                    },
                    Value: ats.String("merge_is_hiring@merge.dev"),
                },
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "maxime": "pariatur",
                        "soluta": "dicta",
                        "laborum": "totam",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "aspernatur": "dolores",
                        "distinctio": "facilis",
                    },
                    Value: ats.String("merge_is_hiring@merge.dev"),
                },
                shared.EmailAddressRequest{
                    EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                    IntegrationParams: map[string]interface{}{
                        "quam": "molestias",
                        "temporibus": "qui",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "fugit": "magni",
                    },
                    Value: ats.String("merge_is_hiring@merge.dev"),
                },
            },
            FirstName: ats.String("Gil"),
            IntegrationParams: map[string]interface{}{
                "sunt": "ullam",
                "nam": "hic",
            },
            IsPrivate: ats.Bool(true),
            LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
            LastName: ats.String("Feig"),
            LinkedAccountParams: map[string]interface{}{
                "cumque": "soluta",
            },
            Locations: []string{
                "et",
                "saepe",
                "ipsum",
            },
            PhoneNumbers: []shared.PhoneNumberRequest{
                shared.PhoneNumberRequest{
                    IntegrationParams: map[string]interface{}{
                        "quos": "tempore",
                        "cupiditate": "aperiam",
                        "delectus": "dolorem",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "labore": "adipisci",
                        "dolorum": "architecto",
                    },
                    PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                    Value: ats.String("+3198675309"),
                },
            },
            RemoteTemplateID: ats.String("92830948203"),
            Tags: []string{
                "aut",
            },
            Title: ats.String("Software Engineer"),
            Urls: []shared.URLRequest{
                shared.URLRequest{
                    IntegrationParams: map[string]interface{}{
                        "consequatur": "est",
                        "repellendus": "porro",
                        "doloribus": "ut",
                        "facilis": "cupiditate",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "quae": "laudantium",
                    },
                    URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                    Value: ats.String("http://alturl.com/p749b"),
                },
                shared.URLRequest{
                    IntegrationParams: map[string]interface{}{
                        "occaecati": "voluptatibus",
                        "quisquam": "vero",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "quis": "ipsum",
                        "delectus": "voluptate",
                        "consectetur": "vero",
                    },
                    URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                    Value: ats.String("http://alturl.com/p749b"),
                },
                shared.URLRequest{
                    IntegrationParams: map[string]interface{}{
                        "dignissimos": "hic",
                        "distinctio": "quod",
                        "odio": "similique",
                        "facilis": "vero",
                    },
                    LinkedAccountParams: map[string]interface{}{
                        "dolore": "quibusdam",
                        "illum": "sequi",
                    },
                    URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                    Value: ats.String("http://alturl.com/p749b"),
                },
            },
        },
        RemoteUserID: "natus",
    }
    xAccountToken := "impedit"
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
    xAccountToken := "aut"
    modelID := "f5d2cff7-c70a-4456-a6d4-36813f16d9f5"
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
        XAccountToken: "sapiente",
        CreatedAfter: types.MustTimeFromString("2020-04-13T07:43:17.350Z"),
        CreatedBefore: types.MustTimeFromString("2022-03-24T11:20:42.976Z"),
        Cursor: ats.String("corporis"),
        EmailAddresses: ats.String("veniam"),
        Expand: operations.CandidatesListExpandApplicationsAttachments.ToPointer(),
        FirstName: ats.String("Armando"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        LastName: ats.String("Gutkowski"),
        ModifiedAfter: types.MustTimeFromString("2022-03-24T01:04:28.722Z"),
        ModifiedBefore: types.MustTimeFromString("2022-01-27T22:22:28.881Z"),
        PageSize: ats.Int64(132487),
        RemoteID: ats.String("minima"),
        Tags: ats.String("eaque"),
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
    xAccountToken := "a"
    id := "b008c42e-141a-4ac3-a6c8-dd6b14429074"
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
    xAccountToken := "eius"
    id := "778a7bd4-66d2-48c1-8ab3-cdca4251904e"
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
    xAccountToken := "ipsam"
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
                    "3c7e0bc7-178e-4479-af2a-70c688282aa4",
                },
                Attachments: []string{
                    "2562f222-e981-47ee-97cb-e61e6b7b95bc",
                    "0ab3c20c-4f37-489f-9871-f99dd2efd121",
                    "aa6f1e67-4bdb-404f-9575-6082d68ea19f",
                },
                CanEmail: ats.Bool(true),
                Company: ats.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "et": "voluptate",
                            "ipsa": "minima",
                            "veritatis": "consectetur",
                            "adipisci": "iste",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "accusantium": "rem",
                            "aut": "laudantium",
                            "eum": "mollitia",
                            "ab": "corrupti",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: ats.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "voluptatem": "dolor",
                    "occaecati": "numquam",
                },
                IsPrivate: ats.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: ats.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "explicabo": "voluptas",
                    "aut": "dignissimos",
                    "dicta": "maiores",
                    "natus": "velit",
                },
                Locations: []string{
                    "voluptas",
                    "asperiores",
                    "aperiam",
                    "ea",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "repellendus": "officia",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "dignissimos": "officia",
                            "asperiores": "nemo",
                            "quae": "quaerat",
                            "porro": "quod",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "ab": "adipisci",
                            "fuga": "id",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "velit": "culpa",
                            "est": "recusandae",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                },
                RemoteTemplateID: ats.String("92830948203"),
                Tags: []string{
                    "fugiat",
                    "vel",
                    "ducimus",
                },
                Title: ats.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "labore": "possimus",
                            "facilis": "cum",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "in": "corporis",
                            "reiciendis": "assumenda",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "recusandae": "aliquid",
                            "aperiam": "cum",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "in": "exercitationem",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "facere": "numquam",
                            "doloribus": "suscipit",
                            "reiciendis": "quidem",
                            "saepe": "necessitatibus",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "sunt": "asperiores",
                            "adipisci": "non",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "amet",
        },
        XAccountToken: "beatae",
        ID: "7fe35b60-eb1e-4a42-a555-ba3c28744ed5",
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

