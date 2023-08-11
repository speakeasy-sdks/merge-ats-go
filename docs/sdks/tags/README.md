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
        XAccountToken: "laboriosam",
        CreatedAfter: types.MustTimeFromString("2020-10-24T06:06:12.533Z"),
        CreatedBefore: types.MustTimeFromString("2021-11-04T21:45:16.507Z"),
        Cursor: ats.String("nemo"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-08-29T01:04:50.334Z"),
        ModifiedBefore: types.MustTimeFromString("2020-09-23T16:28:13.144Z"),
        PageSize: ats.Int64(574032),
        RemoteID: ats.String("enim"),
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

