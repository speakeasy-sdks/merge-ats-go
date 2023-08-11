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
        XAccountToken: "veritatis",
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
        XAccountToken: "quasi",
        CommonModel: ats.String("laboriosam"),
        Cursor: ats.String("pariatur"),
        PageSize: ats.Int64(729448),
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
                            ConditionSchemaID: "545fc95f-a889-470e-989d-bb30fcb33ea0",
                            Operator: "quis",
                            Value: "veniam",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "b197cd44-e2f5-42d8-ad35-13bb6f48b656",
                            Operator: "libero",
                            Value: "minus",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "db35ff2e-4b27-4537-a8cd-9e7319c177d5",
                            Operator: "aspernatur",
                            Value: "enim",
                        },
                    },
                },
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "77b114ee-b52f-4f78-9fc3-7814d4c98e0c",
                            Operator: "aspernatur",
                            Value: "nam",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "b89eb75d-ad63-46c6-8050-3d8bb31180f7",
                            Operator: "amet",
                            Value: "provident",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "ae9e057e-b809-4e28-9033-1f3981d4c700",
                            Operator: "rerum",
                            Value: "ea",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "07f3c93c-73b9-4da3-b2ce-da7e23f22574",
                            Operator: "ab",
                            Value: "illo",
                        },
                    },
                },
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "af4b7544-e472-4e80-a857-a5b40463a7d5",
                            Operator: "molestiae",
                            Value: "veniam",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "f1400e76-4ad7-4334-ac1b-781b36a08088",
                            Operator: "repellendus",
                            Value: "veritatis",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "00efada2-00ef-4042-aeb2-164cf9ab8366",
                            Operator: "impedit",
                            Value: "ducimus",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "23ffda9e-06be-4e48-a5c1-fc0e115c80bf",
                            Operator: "a",
                            Value: "iste",
                        },
                    },
                },
            },
        },
        XAccountToken: "dicta",
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

