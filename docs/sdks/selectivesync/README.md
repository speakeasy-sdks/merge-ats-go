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
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var xAccountToken string = "<value>"

    ctx := context.Background()
    res, err := s.SelectiveSync.List(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
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

**[*operations.SelectiveSyncConfigurationsListResponse](../../pkg/models/operations/selectivesyncconfigurationslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## RetrievePostMetadata

Get metadata for the conditions available to a linked account.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    var xAccountToken string = "<value>"

    var commonModel *string = mergeatsgo.String("<value>")

    var cursor *string = mergeatsgo.String("<value>")

    var pageSize *int64 = mergeatsgo.Int64(70130)

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

**[*operations.SelectiveSyncMetaListResponse](../../pkg/models/operations/selectivesyncmetalistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Replace a linked account's selective syncs.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"context"
	"log"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
    )


    linkedAccountSelectiveSyncConfigurationListRequest := shared.LinkedAccountSelectiveSyncConfigurationListRequest{
        SyncConfigurations: []shared.LinkedAccountSelectiveSyncConfigurationRequest{
            shared.LinkedAccountSelectiveSyncConfigurationRequest{
                LinkedAccountConditions: []shared.LinkedAccountConditionRequest{
                    shared.LinkedAccountConditionRequest{
                        ConditionSchemaID: "d0905bf4-aa77-4f20-8e77-54c352acfe54",
                        Operator: "<value>",
                        Value: "<value>",
                    },
                },
            },
        },
    }

    var xAccountToken string = "<value>"

    ctx := context.Background()
    res, err := s.SelectiveSync.Update(ctx, linkedAccountSelectiveSyncConfigurationListRequest, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `linkedAccountSelectiveSyncConfigurationListRequest`                                                                                       | [shared.LinkedAccountSelectiveSyncConfigurationListRequest](../../pkg/models/shared/linkedaccountselectivesyncconfigurationlistrequest.md) | :heavy_check_mark:                                                                                                                         | N/A                                                                                                                                        |
| `xAccountToken`                                                                                                                            | *string*                                                                                                                                   | :heavy_check_mark:                                                                                                                         | Token identifying the end user.                                                                                                            |


### Response

**[*operations.SelectiveSyncConfigurationsUpdateResponse](../../pkg/models/operations/selectivesyncconfigurationsupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
