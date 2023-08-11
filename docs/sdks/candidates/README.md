# Candidates

### Available Operations

* [CandidatesPartialUpdate](#candidatespartialupdate) - Updates a `Candidate` object with the given `id`.
* [Create](#create) - Creates a `Candidate` object with the given values.
* [IgnoreCreate](#ignorecreate) - Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
* [IgnoreCreate](#ignorecreate) - Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
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
                    "0d5f0d30-c5fb-4b25-8705-3202c73d5fe9",
                    "b90c2890-9b3f-4e49-a8d9-cbf48633323f",
                    "9b77f3a4-1006-474e-bf69-280d1ba77a89",
                },
                Attachments: []string{
                    "bf737ae4-203c-4e5e-aa95-d8a0d446ce2a",
                    "f7a73cf3-be45-43f8-b0b3-26b5a73429cd",
                    "b1a8422b-b679-4d23-a271-5bf0cbb1e31b",
                    "8b90f344-3a11-408e-8adc-f4b921879fce",
                },
                CanEmail: ats.Bool(true),
                Company: ats.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "ipsum": "delectus",
                            "voluptate": "consectetur",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "tenetur": "dignissimos",
                            "hic": "distinctio",
                            "quod": "odio",
                            "similique": "facilis",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "ducimus": "dolore",
                            "quibusdam": "illum",
                            "sequi": "natus",
                            "impedit": "aut",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "exercitationem": "nulla",
                            "fugit": "porro",
                            "maiores": "doloribus",
                            "iusto": "eligendi",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "alias": "officia",
                            "tempora": "ipsam",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "aspernatur": "vel",
                            "possimus": "magnam",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: ats.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "ex": "laudantium",
                },
                IsPrivate: ats.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: ats.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "dolor": "maiores",
                },
                Locations: []string{
                    "ex",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "voluptatibus": "nostrum",
                            "sapiente": "quisquam",
                            "saepe": "ea",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "corporis": "veniam",
                            "aliquid": "inventore",
                            "magnam": "ea",
                            "quo": "consectetur",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "aspernatur": "minima",
                            "eaque": "a",
                            "libero": "aut",
                            "aut": "deleniti",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "aliquam": "fugit",
                            "accusamus": "inventore",
                            "non": "et",
                            "dolorum": "laborum",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "velit": "eum",
                            "autem": "nobis",
                            "quas": "assumenda",
                            "nulla": "voluptas",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "quasi": "tempora",
                            "numquam": "explicabo",
                            "provident": "ipsa",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "magnam": "odio",
                            "eius": "esse",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "rem": "fuga",
                            "reprehenderit": "quidem",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                },
                RemoteTemplateID: ats.String("92830948203"),
                Tags: []string{
                    "ut",
                    "eum",
                    "suscipit",
                    "assumenda",
                },
                Title: ats.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "quisquam": "veritatis",
                            "ipsa": "id",
                            "quidem": "neque",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "illum": "quo",
                            "fuga": "eius",
                            "eos": "voluptas",
                            "ab": "cupiditate",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "consequatur",
        },
        XAccountToken: "tempora",
        ID: "e523c7e0-bc71-478e-8796-f2a70c688282",
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
                    "a482562f-222e-4981-bee1-7cbe61e6b7b9",
                    "5bc0ab3c-20c4-4f37-89fd-871f99dd2efd",
                    "121aa6f1-e674-4bdb-84f1-5756082d68ea",
                },
                Attachments: []string{
                    "9f1d1705-1339-4d08-886a-1840394c2607",
                },
                CanEmail: ats.Bool(true),
                Company: ats.String("Columbia Dining App."),
                EmailAddresses: []shared.EmailAddressRequest{
                    shared.EmailAddressRequest{
                        EmailAddressType: shared.EmailAddressRequestEmailAddressTypePersonal.ToPointer(),
                        IntegrationParams: map[string]interface{}{
                            "natus": "velit",
                            "voluptatibus": "voluptas",
                            "asperiores": "aperiam",
                            "ea": "quaerat",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "repellendus": "officia",
                        },
                        Value: ats.String("merge_is_hiring@merge.dev"),
                    },
                },
                FirstName: ats.String("Gil"),
                IntegrationParams: map[string]interface{}{
                    "dignissimos": "officia",
                    "asperiores": "nemo",
                    "quae": "quaerat",
                    "porro": "quod",
                },
                IsPrivate: ats.Bool(true),
                LastInteractionAt: types.MustTimeFromString("2021-10-17T00:00:00Z"),
                LastName: ats.String("Feig"),
                LinkedAccountParams: map[string]interface{}{
                    "ab": "adipisci",
                    "fuga": "id",
                },
                Locations: []string{
                    "velit",
                    "culpa",
                },
                PhoneNumbers: []shared.PhoneNumberRequest{
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "totam": "fugiat",
                            "vel": "ducimus",
                            "quos": "vel",
                            "labore": "possimus",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "cum": "commodi",
                            "in": "corporis",
                            "reiciendis": "assumenda",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
                        IntegrationParams: map[string]interface{}{
                            "recusandae": "aliquid",
                            "aperiam": "cum",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "in": "exercitationem",
                        },
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                    shared.PhoneNumberRequest{
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
                        PhoneNumberType: shared.PhoneNumberRequestPhoneNumberTypeHome.ToPointer(),
                        Value: ats.String("+3198675309"),
                    },
                },
                RemoteTemplateID: ats.String("92830948203"),
                Tags: []string{
                    "beatae",
                },
                Title: ats.String("Software Engineer"),
                Urls: []shared.URLRequest{
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "debitis": "consectetur",
                            "corporis": "harum",
                            "laboriosam": "ipsa",
                            "voluptates": "libero",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "accusamus": "similique",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                    shared.URLRequest{
                        IntegrationParams: map[string]interface{}{
                            "aspernatur": "voluptas",
                            "voluptas": "voluptas",
                        },
                        LinkedAccountParams: map[string]interface{}{
                            "nobis": "dolorum",
                            "adipisci": "minus",
                        },
                        URLType: shared.URLRequestURLTypePersonal.ToPointer(),
                        Value: ats.String("http://alturl.com/p749b"),
                    },
                },
            },
            RemoteUserID: "dolores",
        },
        XAccountToken: "blanditiis",
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
    operationSecurity := operations.CandidatesIgnoreCreateMultipartSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, operations.CandidatesIgnoreCreateMultipartRequest{
        IgnoreCommonModelRequest2: shared.IgnoreCommonModelRequest2{
            Message: ats.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
            Reason: "in",
        },
        XAccountToken: "dolore",
        ModelID: "4ed53b88-f3a8-4d8f-9c0b-2f2fb7b194a2",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CandidatesIgnoreCreateMultipartRequest](../../models/operations/candidatesignorecreatemultipartrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.CandidatesIgnoreCreateMultipartSecurity](../../models/operations/candidatesignorecreatemultipartsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.CandidatesIgnoreCreateMultipartResponse](../../models/operations/candidatesignorecreatemultipartresponse.md), error**


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
    operationSecurity := operations.CandidatesIgnoreCreateFormSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, operations.CandidatesIgnoreCreateFormRequest{
        IgnoreCommonModelRequest2: shared.IgnoreCommonModelRequest2{
            Message: ats.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
            Reason: "in",
        },
        XAccountToken: "commodi",
        ModelID: "b26916fe-1f08-4f42-94e3-698f447f603e",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CandidatesIgnoreCreateFormRequest](../../models/operations/candidatesignorecreateformrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.CandidatesIgnoreCreateFormSecurity](../../models/operations/candidatesignorecreateformsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.CandidatesIgnoreCreateFormResponse](../../models/operations/candidatesignorecreateformresponse.md), error**


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
    operationSecurity := operations.CandidatesIgnoreCreateJSONSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Candidates.IgnoreCreate(ctx, operations.CandidatesIgnoreCreateJSONRequest{
        IgnoreCommonModelRequest: shared.IgnoreCommonModelRequest{
            Message: ats.String("deletion request by user id 51903790-7dfe-4053-8d63-5a10cc4ffd39"),
            Reason: shared.IgnoreCommonModelRequestReasonGeneralCustomerRequest,
        },
        XAccountToken: "praesentium",
        ModelID: "b445e80c-a55e-4fd2-8e45-7e1858b6a89f",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.CandidatesIgnoreCreateJSONRequest](../../models/operations/candidatesignorecreatejsonrequest.md)   | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `security`                                                                                                     | [operations.CandidatesIgnoreCreateJSONSecurity](../../models/operations/candidatesignorecreatejsonsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |


### Response

**[*operations.CandidatesIgnoreCreateJSONResponse](../../models/operations/candidatesignorecreatejsonresponse.md), error**


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
        XAccountToken: "expedita",
        CreatedAfter: types.MustTimeFromString("2022-05-21T14:36:19.485Z"),
        CreatedBefore: types.MustTimeFromString("2022-04-26T13:26:55.921Z"),
        Cursor: ats.String("officia"),
        EmailAddresses: ats.String("dolorum"),
        Expand: operations.CandidatesListExpandApplicationsAttachments.ToPointer(),
        FirstName: ats.String("Shane"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        LastName: ats.String("Goldner"),
        ModifiedAfter: types.MustTimeFromString("2022-09-14T18:19:59.469Z"),
        ModifiedBefore: types.MustTimeFromString("2022-02-22T10:47:09.142Z"),
        PageSize: ats.Int64(30235),
        RemoteID: ats.String("culpa"),
        Tags: ats.String("expedita"),
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
        XAccountToken: "magnam",
        Expand: operations.CandidatesRetrieveExpandApplications.ToPointer(),
        ID: "75088e51-8620-465e-904f-3b1194b8abf6",
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
        XAccountToken: "alias",
        ID: "3a79f9df-e0ab-47da-8a50-ce187f86bc17",
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
        XAccountToken: "amet",
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

