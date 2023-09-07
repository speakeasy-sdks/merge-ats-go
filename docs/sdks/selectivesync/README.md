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
    xAccountToken := "enim"
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
    xAccountToken := "laboriosam"
    commonModel := "velit"
    cursor := "a"
    pageSize := 562783
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
                        ConditionSchemaID: "4e29e973-e922-4a57-a15b-e3e060807e2b",
                        Operator: "iure",
                        Value: "necessitatibus",
                    },
                },
            },
        },
    }
    xAccountToken := "ratione"
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

