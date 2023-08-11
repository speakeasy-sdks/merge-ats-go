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
        XAccountToken: "hic",
        CreatedAfter: types.MustTimeFromString("2021-11-18T08:45:19.384Z"),
        CreatedBefore: types.MustTimeFromString("2021-11-14T19:02:31.942Z"),
        Cursor: ats.String("odio"),
        Email: ats.String("Stephanie_Braun@gmail.com"),
        IncludeDeletedData: ats.Bool(false),
        IncludeRemoteData: ats.Bool(false),
        ModifiedAfter: types.MustTimeFromString("2020-10-21T20:39:57.323Z"),
        ModifiedBefore: types.MustTimeFromString("2022-07-06T20:33:05.731Z"),
        PageSize: ats.Int64(38557),
        RemoteFields: operations.UsersListRemoteFieldsAccessRole.ToPointer(),
        RemoteID: ats.String("delectus"),
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
        XAccountToken: "impedit",
        ID: "b33ea055-b197-4cd4-8e2f-52d82d3513bb",
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

