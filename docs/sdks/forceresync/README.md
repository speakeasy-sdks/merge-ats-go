# ForceResync

### Available Operations

* [Create](#create) - Force re-sync of all models. This is available for all organizations via the dashboard. Force re-sync is also available programmatically via API for monthly, quarterly, and highest sync frequency customers on the Core, Professional, or Enterprise plans. Doing so will consume a sync credit for the relevant linked account.

## Create

Force re-sync of all models. This is available for all organizations via the dashboard. Force re-sync is also available programmatically via API for monthly, quarterly, and highest sync frequency customers on the Core, Professional, or Enterprise plans. Doing so will consume a sync credit for the relevant linked account.

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
    xAccountToken := "cupiditate"
    operationSecurity := operations.SyncStatusResyncCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.ForceResync.Create(ctx, operationSecurity, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.SyncStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.SyncStatusResyncCreateSecurity](../../models/operations/syncstatusresynccreatesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |
| `xAccountToken`                                                                                        | *string*                                                                                               | :heavy_check_mark:                                                                                     | Token identifying the end user.                                                                        |


### Response

**[*operations.SyncStatusResyncCreateResponse](../../models/operations/syncstatusresynccreateresponse.md), error**

