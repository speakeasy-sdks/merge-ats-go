# Candidates

### Available Operations

* [CandidatesPartialUpdate](#candidatespartialupdate) - Updates a `Candidate` object with the given `id`.
* [Create](#create) - Creates a `Candidate` object with the given values.
* [IgnoreCreate](#ignorecreate) - Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
* [List](#list) - Returns a list of `Candidate` objects.
* [Retrieve](#retrieve) - Returns a `Candidate` object with the given `id`.
* [RetrievePatchMetadata](#retrievepatchmetadata) - Returns metadata for `Candidate` PATCHs.
* [RetrievePostMetadata](#retrievepostmetadata) - Returns metadata for `Candidate` POSTs.

## CandidatesPartialUpdate

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
    operationSecurity := operations.CandidatesPartialUpdateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.CandidatesPartialUpdate(ctx, operations.CandidatesPartialUpdateRequest{
        PatchedCandidateEndpointRequest: shared.PatchedCandidateEndpointRequest{
            Model: shared.PatchedCandidateRequest{
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
        },
        XAccountToken: "impedit",
        ID: "0f5d2cff-7c70-4a45-a26d-436813f16d9f",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CandidatesPartialUpdateRequest](../../models/operations/candidatespartialupdaterequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.CandidatesPartialUpdateSecurity](../../models/operations/candidatespartialupdatesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.CandidatesPartialUpdateResponse](../../models/operations/candidatespartialupdateresponse.md), error**


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
    operationSecurity := operations.CandidatesCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.Create(ctx, operations.CandidatesCreateRequest{
        CandidateEndpointRequest: shared.CandidateEndpointRequest{
            Model: shared.CandidateRequest{
                Applications: []string{
                    "fce6c556-146c-43e2-90fb-008c42e141aa",
                    "c366c8dd-6b14-4429-8747-4778a7bd466d",
                },
                Attachments: []string{
                    "8c10ab3c-dca4-4251-904e-523c7e0bc717",
                },
                CanEmail: ats.Bool(true),
                Company: ats.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "aliquam": "odio",
                            "occaecati": "commodi",
                            "sapiente": "dolores",
                            "deserunt": "molestiae",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "porro": "eum",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "praesentium": "consequuntur",
                            "deleniti": "fugit",
                            "fuga": "mollitia",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "atque": "explicabo",
                            "minima": "nisi",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "sapiente": "consequuntur",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "explicabo": "saepe",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: ats.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "atque": "et",
                    "esse": "eveniet",
                    "accusamus": "veritatis",
                },
                IsPrivate: ats.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: ats.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "quod": "nam",
                    "vero": "aliquid",
                },
                Locations: []string{
                    "saepe",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "molestiae": "rerum",
                            "occaecati": "minima",
                            "distinctio": "eligendi",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "culpa": "tempore",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "cumque": "consequuntur",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "minus": "quaerat",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                },
                RemoteTemplateID: ats.String("92830948203"),
                Tags: []string{
                    "consectetur",
                    "esse",
                    "blanditiis",
                    "provident",
                },
                Title: ats.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "quas": "esse",
                            "quasi": "a",
                            "error": "sint",
                            "pariatur": "possimus",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "eveniet": "asperiores",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "veritatis": "consequuntur",
                            "quasi": "similique",
                            "culpa": "aliquid",
                            "tenetur": "quae",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "vel": "in",
                            "eius": "libero",
                            "illum": "soluta",
                            "accusantium": "aliquam",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "dicta": "ullam",
                            "reprehenderit": "ullam",
                            "nisi": "aut",
                            "voluptatum": "qui",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "ex": "deleniti",
                            "itaque": "dolorum",
                            "architecto": "omnis",
                            "tenetur": "quasi",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                    shared.URLRequest{
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
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "non",
        },
        XAccountToken: "voluptatem",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CandidatesCreateRequest](../../models/operations/candidatescreaterequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CandidatesCreateSecurity](../../models/operations/candidatescreatesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


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
    operationSecurity := operations.CandidatesIgnoreCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, operations.CandidatesIgnoreCreateRequest{
        IgnoreCommonModelRequest: shared.IgnoreCommonModelRequest{
            Message: ats.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
            Reason: shared.IgnoreCommonModelRequestReasonGeneralCustomerRequest,
        },
        XAccountToken: "dolor",
        ModelID: "94c26071-f93f-45f0-a42d-ac7af515cc41",
    }, operationSecurity)
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
| `request`                                                                                              | [operations.CandidatesIgnoreCreateRequest](../../models/operations/candidatesignorecreaterequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.CandidatesIgnoreCreateSecurity](../../models/operations/candidatesignorecreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


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
        XAccountToken: "adipisci",
        CreatedAfter: types.MustTimeFromString("2021-09-04T08:55:11.388Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-03T04:29:10.698Z"),
        Cursor: ats.String("culpa"),
        EmailAddresses: ats.String("est"),
        Expand: operations.CandidatesListExpandAttachments.ToPointer(),
        FirstName: ats.String("Joseph"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        LastName: ats.String("Stehr"),
        ModifiedAfter: types.MustTimeFromString("2022-07-03T08:20:26.765Z"),
        ModifiedBefore: types.MustTimeFromString("2022-02-22T16:20:29.225Z"),
        PageSize: ats.Int64(287051),
        RemoteID: ats.String("possimus"),
        Tags: ats.String("facilis"),
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
    operationSecurity := operations.CandidatesRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.Retrieve(ctx, operations.CandidatesRetrieveRequest{
        XAccountToken: "cum",
        Expand: operations.CandidatesRetrieveExpandApplicationsAttachments.ToPointer(),
        ID: "75fd5e60-b375-4ed4-b6fb-ee41f33317fe",
        IncludeRemoteData: ats.Bool(false),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Candidate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CandidatesRetrieveRequest](../../models/operations/candidatesretrieverequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CandidatesRetrieveSecurity](../../models/operations/candidatesretrievesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


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
    operationSecurity := operations.CandidatesMetaPatchRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.RetrievePatchMetadata(ctx, operations.CandidatesMetaPatchRetrieveRequest{
        XAccountToken: "consectetur",
        ID: "5b60eb1e-a426-4555-ba3c-28744ed53b88",
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
| `request`                                                                                                        | [operations.CandidatesMetaPatchRetrieveRequest](../../models/operations/candidatesmetapatchretrieverequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.CandidatesMetaPatchRetrieveSecurity](../../models/operations/candidatesmetapatchretrievesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


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
    operationSecurity := operations.CandidatesMetaPostRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.RetrievePostMetadata(ctx, operations.CandidatesMetaPostRetrieveRequest{
        XAccountToken: "hic",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CandidatesMetaPostRetrieveRequest](../../models/operations/candidatesmetapostretrieverequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.CandidatesMetaPostRetrieveSecurity](../../models/operations/candidatesmetapostretrievesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.CandidatesMetaPostRetrieveResponse](../../models/operations/candidatesmetapostretrieveresponse.md), error**

