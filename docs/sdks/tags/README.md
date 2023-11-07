# Tags
(*.Tags*)

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
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := mergeatsgo.New(
        mergeatsgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Tags.List(ctx, operations.TagsListRequest{
        XAccountToken: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedTagList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.TagsListRequest](../../models/operations/tagslistrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.TagsListResponse](../../models/operations/tagslistresponse.md), error**

