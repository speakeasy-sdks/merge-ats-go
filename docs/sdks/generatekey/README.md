# GenerateKey
(*GenerateKey*)

### Available Operations

* [Create](#create) - Create a remote key.

## Create

Create a remote key.

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

    ctx := context.Background()
    res, err := s.GenerateKey.Create(ctx, shared.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoteKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.GenerateRemoteKeyRequest](../../models/shared/generateremotekeyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GenerateKeyCreateResponse](../../models/operations/generatekeycreateresponse.md), error**

