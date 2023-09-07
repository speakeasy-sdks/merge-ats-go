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
        XAccountToken: "reiciendis",
        CreatedAfter: types.MustTimeFromString("2022-05-03T16:31:59.395Z"),
        CreatedBefore: types.MustTimeFromString("2022-09-16T00:03:17.363Z"),
        Cursor: ats.String("dolorum"),
        Email: ats.String("Angus.Walsh@gmail.com"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2022-08-06T19:48:35.360Z"),
        ModifiedBefore: types.MustTimeFromString("2022-05-14T22:22:36.102Z"),
        PageSize: ats.Int64(238413),
        RemoteFields: operations.UsersListRemoteFieldsAccessRole.ToPointer(),
        RemoteID: ats.String("debitis"),
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
        XAccountToken: "laudantium",
        ID: "65e7956f-9251-4a5a-9da6-60ff57bfaad4",
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

