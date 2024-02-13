# SyncStatus
(*SyncStatus*)

### Available Operations

* [List](#list) - Get syncing status. Possible values: `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`

## List

Get syncing status. Possible values: `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`

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


    var xAccountToken string = "string"

    var cursor *string = mergeatsgo.String("string")

    var pageSize *int64 = mergeatsgo.Int64(768578)

    ctx := context.Background()
    res, err := s.SyncStatus.List(ctx, xAccountToken, cursor, pageSize)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedSyncStatusList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `xAccountToken`                                       | *string*                                              | :heavy_check_mark:                                    | Token identifying the end user.                       |
| `cursor`                                              | **string*                                             | :heavy_minus_sign:                                    | The pagination cursor value.                          |
| `pageSize`                                            | **int64*                                              | :heavy_minus_sign:                                    | Number of results to return per page.                 |


### Response

**[*operations.SyncStatusListResponse](../../pkg/models/operations/syncstatuslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
