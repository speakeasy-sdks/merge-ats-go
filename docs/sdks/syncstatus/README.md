# SyncStatus

### Available Operations

* [List](#list) - Get syncing status. Possible values: `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`

## List

Get syncing status. Possible values: `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`

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
    operationSecurity := operations.SyncStatusListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SyncStatus.List(ctx, operations.SyncStatusListRequest{
        XAccountToken: "maiores",
        Cursor: ats.String("veritatis"),
        PageSize: ats.Int64(96450),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedSyncStatusList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.SyncStatusListRequest](../../models/operations/syncstatuslistrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.SyncStatusListSecurity](../../models/operations/syncstatuslistsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.SyncStatusListResponse](../../models/operations/syncstatuslistresponse.md), error**

