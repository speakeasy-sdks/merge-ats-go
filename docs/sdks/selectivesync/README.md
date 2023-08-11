# SelectiveSync

### Available Operations

* [List](#list) - Get a linked account's selective syncs.
* [RetrievePostMetadata](#retrievepostmetadata) - Get metadata for the conditions available to a linked account.
* [Update](#update) - Replace a linked account's selective syncs.

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
    xAccountToken := "dolor"
    operationSecurity := operations.SelectiveSyncConfigurationsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SelectiveSync.List(ctx, operationSecurity, xAccountToken)
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
| `security`                                                                                                               | [operations.SelectiveSyncConfigurationsListSecurity](../../models/operations/selectivesyncconfigurationslistsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |
| `xAccountToken`                                                                                                          | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | Token identifying the end user.                                                                                          |


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
    xAccountToken := "iusto"
    commonModel := "sit"
    cursor := "doloremque"
    pageSize := 7468
    operationSecurity := operations.SelectiveSyncMetaListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SelectiveSync.RetrievePostMetadata(ctx, operationSecurity, xAccountToken, commonModel, cursor, pageSize)
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
| `security`                                                                                           | [operations.SelectiveSyncMetaListSecurity](../../models/operations/selectivesyncmetalistsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |
| `xAccountToken`                                                                                      | *string*                                                                                             | :heavy_check_mark:                                                                                   | Token identifying the end user.                                                                      |
| `commonModel`                                                                                        | **string*                                                                                            | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `cursor`                                                                                             | **string*                                                                                            | :heavy_minus_sign:                                                                                   | The pagination cursor value.                                                                         |
| `pageSize`                                                                                           | **int64*                                                                                             | :heavy_minus_sign:                                                                                   | Number of results to return per page.                                                                |


### Response

**[*operations.SelectiveSyncMetaListResponse](../../models/operations/selectivesyncmetalistresponse.md), error**


## Update

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
    linkedAccountSelectiveSyncConfigurationListRequest := shared.LinkedAccountSelectiveSyncConfigurationListRequest{
        SyncConfigurations: []shared.LinkedAccountSelectiveSyncConfigurationRequest{
            shared.LinkedAccountSelectiveSyncConfigurationRequest{
                LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "6b6bc9b8-f759-4eac-95a9-741d31135296",
                        Operator: "nemo",
                        Value: "soluta",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "b8a72026-1143-45e1-b9db-c2259b1abda8",
                        Operator: "placeat",
                        Value: "sit",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "70e1084c-b067-42d1-ad87-9eeb9665b85e",
                        Operator: "voluptatibus",
                        Value: "cum",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "d02bae0b-e2d7-4822-99e3-ea4b5197f924",
                        Operator: "numquam",
                        Value: "nesciunt",
                    },
                },
            },
            shared.LinkedAccountSelectiveSyncConfigurationRequest{
                LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "a7ce52b8-95c5-437c-a454-efb0b34896c3",
                        Operator: "minus",
                        Value: "fuga",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "5acfbe2f-d570-4757-b929-177deac646ec",
                        Operator: "quidem",
                        Value: "exercitationem",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "73409e3e-b1e5-4a2b-92eb-07f116db9954",
                        Operator: "nostrum",
                        Value: "doloribus",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "c95fa889-70e1-489d-bb30-fcb33ea055b1",
                        Operator: "cupiditate",
                        Value: "molestiae",
                    },
                },
            },
            shared.LinkedAccountSelectiveSyncConfigurationRequest{
                LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "d44e2f52-d82d-4351-bbb6-f48b656bcdb3",
                        Operator: "ad",
                        Value: "voluptatibus",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "f2e4b275-37a8-4cd9-a731-9c177d525f77",
                        Operator: "libero",
                        Value: "illo",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "14eeb52f-f785-4fc3-b814-d4c98e0c2bb8",
                        Operator: "provident",
                        Value: "repudiandae",
                    },
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "b75dad63-6c60-4050-bd8b-b31180f739ae",
                        Operator: "provident",
                        Value: "repudiandae",
                    },
                },
            },
        },
    }
    xAccountToken := "consequatur"
    operationSecurity := operations.SelectiveSyncConfigurationsUpdateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SelectiveSync.Update(ctx, operationSecurity, linkedAccountSelectiveSyncConfigurationListRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkedAccountSelectiveSyncConfigurations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `security`                                                                                                                             | [operations.SelectiveSyncConfigurationsUpdateSecurity](../../models/operations/selectivesyncconfigurationsupdatesecurity.md)           | :heavy_check_mark:                                                                                                                     | The security requirements to use for the request.                                                                                      |
| `linkedAccountSelectiveSyncConfigurationListRequest`                                                                                   | [shared.LinkedAccountSelectiveSyncConfigurationListRequest](../../models/shared/linkedaccountselectivesyncconfigurationlistrequest.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `xAccountToken`                                                                                                                        | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | Token identifying the end user.                                                                                                        |


### Response

**[*operations.SelectiveSyncConfigurationsUpdateResponse](../../models/operations/selectivesyncconfigurationsupdateresponse.md), error**

