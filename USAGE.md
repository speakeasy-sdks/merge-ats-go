<!-- Start SDK Example Usage -->


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
        mergeatsgo.WithSecurity(shared.Security{
            TokenAuth: "",
        }),
    )
    var xAccountToken string = "till"

    ctx := context.Background()
    res, err := s.AccountDetails.Retrieve(ctx, xAccountToken)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountDetails != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->