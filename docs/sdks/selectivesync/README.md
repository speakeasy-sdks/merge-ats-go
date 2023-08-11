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
        XAccountToken: "assumenda",
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
        XAccountToken: "beatae",
        CommonModel: ats.String("est"),
        Cursor: ats.String("facere"),
        PageSize: ats.Int64(545918),
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
                            ConditionSchemaID: "eeb9665b-85ef-4bd0-abae-0be2d782259e",
                            Operator: "adipisci",
                            Value: "recusandae",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "a4b5197f-9244-43da-bce5-2b895c537c64",
                            Operator: "corporis",
                            Value: "magnam",
                        },
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "efb0b348-96c3-4ca5-acfb-e2fd57075779",
                            Operator: "dolores",
                            Value: "error",
                        },
                    },
                },
                shared.LinkedAccountSelectiveSyncConfigurationRequest{
                    LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                        shared.LinkedAccountConditionRequest{
                            ConditionSchemaID: "77deac64-6ecb-4573-809e-3eb1e5a2b12e",
                            Operator: "nobis",
                            Value: "ipsa",
                        },
                    },
                },
            },
        },
        XAccountToken: "ducimus",
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

