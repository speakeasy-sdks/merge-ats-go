# SelectiveSync
(*SelectiveSync*)

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
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "Bicycle"

    ctx := context.Background()
    res, err := s.SelectiveSync.List(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkedAccountSelectiveSyncConfigurations != nil {
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

**[*operations.SelectiveSyncConfigurationsListResponse](../../models/operations/selectivesyncconfigurationslistresponse.md), error**


## RetrievePostMetadata

Get metadata for the conditions available to a linked account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    var xAccountToken string = "Borders"

    var commonModel *string = "Home"

    var cursor *string = "Country"

    var pageSize *int64 = 925730

    ctx := context.Background()
    res, err := s.SelectiveSync.RetrievePostMetadata(ctx, xAccountToken, commonModel, cursor, pageSize)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedConditionSchemaList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xAccountToken`                                       | *string*                                              | :heavy_check_mark:                                    | Token identifying the end user.                       |
| `commonModel`                                         | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | The pagination cursor value.                          |
| `pageSize`                                            | **int64*                                              | :heavy_minus_sign:                                    | Number of results to return per page.                 |


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
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )


    linkedAccountSelectiveSyncConfigurationListRequest := shared.LinkedAccountSelectiveSyncConfigurationListRequest{
        SyncConfigurations: []shared.LinkedAccountSelectiveSyncConfigurationRequest{
            shared.LinkedAccountSelectiveSyncConfigurationRequest{
                LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "d0905bf4-aa77-4f20-8e77-54c352acfe54",
                        Operator: "JBOD",
                        Value: "Bicycle",
                    },
                },
            },
        },
    }

    var xAccountToken string = "guestbook"

    ctx := context.Background()
    res, err := s.SelectiveSync.Update(ctx, linkedAccountSelectiveSyncConfigurationListRequest, xAccountToken)
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
| `linkedAccountSelectiveSyncConfigurationListRequest`                                                                                   | [shared.LinkedAccountSelectiveSyncConfigurationListRequest](../../models/shared/linkedaccountselectivesyncconfigurationlistrequest.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `xAccountToken`                                                                                                                        | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | Token identifying the end user.                                                                                                        |


### Response

**[*operations.SelectiveSyncConfigurationsUpdateResponse](../../models/operations/selectivesyncconfigurationsupdateresponse.md), error**

