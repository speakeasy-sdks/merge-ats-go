<!-- Start SDK Example Usage -->


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
    operationSecurity := operations.AccountDetailsRetrieveSecurity{
            TokenAuth: "",
        }

    ctx := context.Background()
    res, err := s.AccountDetails.Retrieve(ctx, operations.AccountDetailsRetrieveRequest{
        XAccountToken: "corrupti",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountDetails != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->