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
        XAccountToken: "rem",
        CreatedAfter: types.MustTimeFromString("2022-09-05T23:50:51.335Z"),
        CreatedBefore: types.MustTimeFromString("2022-12-28T14:32:37.413Z"),
        Cursor: ats.String("corporis"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-01-22T10:13:48.972Z"),
        ModifiedBefore: types.MustTimeFromString("2022-03-30T21:08:39.591Z"),
        PageSize: ats.Int64(5189),
        RemoteID: ats.String("maiores"),
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

