# Tags

### Available Operations

* [List](#list) - Returns a list of `Tag` objects.

## List

Returns a list of `Tag` objects.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.TagsListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Tags.List(ctx, operations.TagsListRequest{
        XAccountToken: "modi",
        CreatedAfter: types.MustTimeFromString("2020-07-30T21:02:27.356Z"),
        CreatedBefore: types.MustTimeFromString("2022-11-02T07:13:29.133Z"),
        Cursor: ats.String("assumenda"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2020-01-18T07:20:38.221Z"),
        ModifiedBefore: types.MustTimeFromString("2020-08-26T04:20:10.413Z"),
        PageSize: ats.Int64(883819),
        RemoteID: ats.String("totam"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedTagList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.TagsListRequest](../../models/operations/tagslistrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.TagsListSecurity](../../models/operations/tagslistsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.TagsListResponse](../../models/operations/tagslistresponse.md), error**

