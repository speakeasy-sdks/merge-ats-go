# RegenerateKey

### Available Operations

* [Create](#create) - Exchange remote keys.

## Create

Exchange remote keys.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.RegenerateKeyCreateSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.RegenerateKey.Create(ctx, shared.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoteKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [shared.RemoteKeyForRegenerationRequest](../../models/shared/remotekeyforregenerationrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.RegenerateKeyCreateSecurity](../../models/operations/regeneratekeycreatesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.RegenerateKeyCreateResponse](../../models/operations/regeneratekeycreateresponse.md), error**

