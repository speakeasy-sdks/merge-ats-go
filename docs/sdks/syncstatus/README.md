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
    xAccountToken := "laborum"
    cursor := "distinctio"
    pageSize := 528940
    operationSecurity := operations.SyncStatusListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.SyncStatus.List(ctx, operationSecurity, xAccountToken, cursor, pageSize)
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
| `security`                                                                             | [operations.SyncStatusListSecurity](../../models/operations/syncstatuslistsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |
| `xAccountToken`                                                                        | *string*                                                                               | :heavy_check_mark:                                                                     | Token identifying the end user.                                                        |
| `cursor`                                                                               | **string*                                                                              | :heavy_minus_sign:                                                                     | The pagination cursor value.                                                           |
| `pageSize`                                                                             | **int64*                                                                               | :heavy_minus_sign:                                                                     | Number of results to return per page.                                                  |


### Response

**[*operations.SyncStatusListResponse](../../models/operations/syncstatuslistresponse.md), error**

