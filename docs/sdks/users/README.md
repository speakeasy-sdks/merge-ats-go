# Users

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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/types"
)

func main() {
    s := ats.New()
    operationSecurity := operations.UsersListSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Users.List(ctx, operations.UsersListRequest{
        XAccountToken: "a",
        CreatedAfter: types.MustTimeFromString("2022-06-02T10:52:44.468Z"),
        CreatedBefore: types.MustTimeFromString("2022-10-14T13:04:49.369Z"),
        Cursor: ats.String("at"),
        Email: ats.String("Percy2@gmail.com"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-03-06T12:29:45.421Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-06T08:20:32.050Z"),
        PageSize: ats.Int64(998026),
        RemoteFields: operations.UsersListRemoteFieldsAccessRole.ToPointer(),
        RemoteID: ats.String("velit"),
        ShowEnumOrigins: operations.UsersListShowEnumOriginsAccessRole.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedRemoteUserList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UsersListRequest](../../models/operations/userslistrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.UsersListSecurity](../../models/operations/userslistsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


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
	"github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
)

func main() {
    s := ats.New()
    operationSecurity := operations.UsersRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.Users.Retrieve(ctx, operations.UsersRetrieveRequest{
        XAccountToken: "porro",
        ID: "93c73b9d-a3f2-4ced-a7e2-3f2257411faf",
        IncludeRemoteData: ats.Bool(false),
        RemoteFields: operations.UsersRetrieveRemoteFieldsAccessRole.ToPointer(),
        ShowEnumOrigins: operations.UsersRetrieveShowEnumOriginsAccessRole.ToPointer(),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoteUser != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UsersRetrieveRequest](../../models/operations/usersretrieverequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.UsersRetrieveSecurity](../../models/operations/usersretrievesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.UsersRetrieveResponse](../../models/operations/usersretrieveresponse.md), error**

