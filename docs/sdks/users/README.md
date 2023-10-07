# Users
(*Users*)

### Available Operations

* [List](#list) - Returns a list of `RemoteUser` objects.
* [Retrieve](#retrieve) - Returns a `RemoteUser` object with the given `id`.

## List

Returns a list of `RemoteUser` objects.

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
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Users.List(ctx, operations.UsersListRequest{
        XAccountToken: "Northeast Metal Canada",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedRemoteUserList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UsersListRequest](../../models/operations/userslistrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.UsersListResponse](../../models/operations/userslistresponse.md), error**


## Retrieve

Returns a `RemoteUser` object with the given `id`.

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
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Users.Retrieve(ctx, operations.UsersRetrieveRequest{
        XAccountToken: "tracksuit Markets",
        ID: "1081ad20-d604-4c8e-92b2-41fa379087a1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoteUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UsersRetrieveRequest](../../models/operations/usersretrieverequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UsersRetrieveResponse](../../models/operations/usersretrieveresponse.md), error**

