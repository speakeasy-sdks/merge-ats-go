<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"log"
)

func main() {
	s := mergeatsgo.New(
		mergeatsgo.WithSecurity(""),
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