# SelectiveSync

### Available Operations

* [List](#list) - Get a linked account's selective syncs.
* [RetrievePostMetadata](#retrievepostmetadata) - Get metadata for the conditions available to a linked account.
* [SelectiveSyncConfigurationsUpdate](#selectivesyncconfigurationsupdate) - Replace a linked account's selective syncs.

## List

Get a linked account's selective syncs.

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
    operationSecurity := operations.SelectiveSyncConfigurationsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SelectiveSync.List(ctx, operations.SelectiveSyncConfigurationsListRequest{
        XAccountToken: "facere",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkedAccountSelectiveSyncConfigurations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.SelectiveSyncConfigurationsListRequest](../../models/operations/selectivesyncconfigurationslistrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.SelectiveSyncConfigurationsListSecurity](../../models/operations/selectivesyncconfigurationslistsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |


### Response

**[*operations.SelectiveSyncConfigurationsListResponse](../../models/operations/selectivesyncconfigurationslistresponse.md), error**


## RetrievePostMetadata

Get metadata for the conditions available to a linked account.

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
    operationSecurity := operations.SelectiveSyncMetaListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SelectiveSync.RetrievePostMetadata(ctx, operations.SelectiveSyncMetaListRequest{
        XAccountToken: "libero",
        CommonModel: ats.String("architecto"),
        Cursor: ats.String("voluptatibus"),
        PageSize: ats.Int64(156383),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedConditionSchemaList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.SelectiveSyncMetaListRequest](../../models/operations/selectivesyncmetalistrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.SelectiveSyncMetaListSecurity](../../models/operations/selectivesyncmetalistsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.SelectiveSyncMetaListResponse](../../models/operations/selectivesyncmetalistresponse.md), error**


## SelectiveSyncConfigurationsUpdate

Replace a linked account's selective syncs.

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
    operationSecurity := operations.SelectiveSyncConfigurationsUpdateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SelectiveSync.SelectiveSyncConfigurationsUpdate(ctx, operations.SelectiveSyncConfigurationsUpdateRequest{
        LinkedAccountSelectiveSyncConfigurationListRequest: shared.LinkedAccountSelectiveSyncConfigurationListRequest{
            SyncConfigurations: []shared.LinkedAccountSelectiveSyncConfigurationRequest{
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "310661e9-6349-4e1c-b9e0-6e3a437000ae",
                            Operator: "ea",
                            Value: "quidem",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "6bc9b8f7-59ea-4c55-a974-1d311352965b",
                            Operator: "libero",
                            Value: "rem",
                        },
                    },
                },
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "72026114-35e1-439d-bc22-59b1abda8c07",
                            Operator: "ipsa",
                            Value: "voluptates",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "1084cb06-72d1-4ad8-b9ee-b9665b85efbd",
                            Operator: "alias",
                            Value: "quia",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "bae0be2d-7822-459e-bea4-b5197f92443d",
                            Operator: "officia",
                            Value: "dignissimos",
                        },
                    },
                },
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "e52b895c-537c-4645-8efb-0b34896c3ca5",
                            Operator: "est",
                            Value: "impedit",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "fbe2fd57-0757-4792-9177-deac646ecb57",
                            Operator: "dolorem",
                            Value: "modi",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "09e3eb1e-5a2b-412e-b07f-116db99545fc",
                            Operator: "sint",
                            Value: "enim",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "fa88970e-189d-4bb3-8fcb-33ea055b197c",
                            Operator: "possimus",
                            Value: "non",
                        },
                    },
                },
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "e2f52d82-d351-43bb-af48-b656bcdb35ff",
                            Operator: "consequuntur",
                            Value: "debitis",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "4b27537a-8cd9-4e73-99c1-77d525f77b11",
                            Operator: "incidunt",
                            Value: "accusamus",
                        },
                    },
                },
            },
        },
        XAccountToken: "saepe",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkedAccountSelectiveSyncConfigurations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.SelectiveSyncConfigurationsUpdateRequest](../../models/operations/selectivesyncconfigurationsupdaterequest.md)   | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `security`                                                                                                                   | [operations.SelectiveSyncConfigurationsUpdateSecurity](../../models/operations/selectivesyncconfigurationsupdatesecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |


### Response

**[*operations.SelectiveSyncConfigurationsUpdateResponse](../../models/operations/selectivesyncconfigurationsupdateresponse.md), error**

