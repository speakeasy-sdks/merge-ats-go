<div align="center">
    <img src="https://github.com/speakeasy-sdks/merge-ats-go/assets/6267663/5aa3b6d0-154c-4dbc-b402-2346f7d09354" width="350px">
    <h1>ATS Go SDK</h1>
   <p>Merge is a single API to add hundreds of integrations to your app.</p>
   <a href="https://docs.merge.dev/ats/overview/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/merge-ats-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/merge-ats-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/merge-ats-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/merge-ats-go?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/merge-ats-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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
		mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	var xAccountToken string = "<value>"

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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [AccountDetails](docs/sdks/accountdetails/README.md)

* [Retrieve](docs/sdks/accountdetails/README.md#retrieve) - Get details for a linked account.

### [AccountToken](docs/sdks/accounttoken/README.md)

* [Retrieve](docs/sdks/accounttoken/README.md#retrieve) - Returns the account token for the end user with the provided public token.

### [Activities](docs/sdks/activities/README.md)

* [Create](docs/sdks/activities/README.md#create) - Creates an `Activity` object with the given values.
* [List](docs/sdks/activities/README.md#list) - Returns a list of `Activity` objects.
* [Retrieve](docs/sdks/activities/README.md#retrieve) - Returns an `Activity` object with the given `id`.
* [RetrievePostMetadata](docs/sdks/activities/README.md#retrievepostmetadata) - Returns metadata for `Activity` POSTs.

### [Applications](docs/sdks/applications/README.md)

* [Create](docs/sdks/applications/README.md#create) - Creates an `Application` object with the given values.
* [List](docs/sdks/applications/README.md#list) - Returns a list of `Application` objects.
* [Retrieve](docs/sdks/applications/README.md#retrieve) - Returns an `Application` object with the given `id`.
* [RetrievePostMetadata](docs/sdks/applications/README.md#retrievepostmetadata) - Returns metadata for `Application` POSTs.
* [UpdateChangeState](docs/sdks/applications/README.md#updatechangestate) - Updates the `current_stage` field of an `Application` object

### [AsyncPassthrough](docs/sdks/asyncpassthrough/README.md)

* [Create](docs/sdks/asyncpassthrough/README.md#create) - Asynchronously pull data from an endpoint not currently supported by Merge.
* [Retrieve](docs/sdks/asyncpassthrough/README.md#retrieve) - Retrieves data from earlier async-passthrough POST request

### [Attachments](docs/sdks/attachments/README.md)

* [Create](docs/sdks/attachments/README.md#create) - Creates an `Attachment` object with the given values.
* [List](docs/sdks/attachments/README.md#list) - Returns a list of `Attachment` objects.
* [Retrieve](docs/sdks/attachments/README.md#retrieve) - Returns an `Attachment` object with the given `id`.
* [RetrievePostMetadata](docs/sdks/attachments/README.md#retrievepostmetadata) - Returns metadata for `Attachment` POSTs.

### [AvailableActions](docs/sdks/availableactions/README.md)

* [Retrieve](docs/sdks/availableactions/README.md#retrieve) - Returns a list of models and actions available for an account.

### [Candidates](docs/sdks/candidates/README.md)

* [Create](docs/sdks/candidates/README.md#create) - Creates a `Candidate` object with the given values.
* [IgnoreCreate](docs/sdks/candidates/README.md#ignorecreate) - Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
* [List](docs/sdks/candidates/README.md#list) - Returns a list of `Candidate` objects.
* [Retrieve](docs/sdks/candidates/README.md#retrieve) - Returns a `Candidate` object with the given `id`.
* [RetrievePatchMetadata](docs/sdks/candidates/README.md#retrievepatchmetadata) - Returns metadata for `Candidate` PATCHs.
* [RetrievePostMetadata](docs/sdks/candidates/README.md#retrievepostmetadata) - Returns metadata for `Candidate` POSTs.
* [Update](docs/sdks/candidates/README.md#update) - Updates a `Candidate` object with the given `id`.

### [DeleteAccount](docs/sdks/deleteaccount/README.md)

* [DeleteAccountDelete](docs/sdks/deleteaccount/README.md#deleteaccountdelete) - Delete a linked account.

### [Departments](docs/sdks/departments/README.md)

* [List](docs/sdks/departments/README.md#list) - Returns a list of `Department` objects.
* [Retrieve](docs/sdks/departments/README.md#retrieve) - Returns a `Department` object with the given `id`.

### [Eeocs](docs/sdks/eeocs/README.md)

* [List](docs/sdks/eeocs/README.md#list) - Returns a list of `EEOC` objects.
* [Retrieve](docs/sdks/eeocs/README.md#retrieve) - Returns an `EEOC` object with the given `id`.

### [GenerateKey](docs/sdks/generatekey/README.md)

* [Create](docs/sdks/generatekey/README.md#create) - Create a remote key.

### [Interviews](docs/sdks/interviews/README.md)

* [Create](docs/sdks/interviews/README.md#create) - Creates a `ScheduledInterview` object with the given values.
* [List](docs/sdks/interviews/README.md#list) - Returns a list of `ScheduledInterview` objects.
* [Retrieve](docs/sdks/interviews/README.md#retrieve) - Returns a `ScheduledInterview` object with the given `id`.
* [RetrievePostMetadata](docs/sdks/interviews/README.md#retrievepostmetadata) - Returns metadata for `ScheduledInterview` POSTs.

### [Issues](docs/sdks/issues/README.md)

* [List](docs/sdks/issues/README.md#list) - Gets issues.
* [Retrieve](docs/sdks/issues/README.md#retrieve) - Get a specific issue.

### [JobInterviewStages](docs/sdks/jobinterviewstages/README.md)

* [List](docs/sdks/jobinterviewstages/README.md#list) - Returns a list of `JobInterviewStage` objects.
* [Retrieve](docs/sdks/jobinterviewstages/README.md#retrieve) - Returns a `JobInterviewStage` object with the given `id`.

### [Jobs](docs/sdks/jobs/README.md)

* [List](docs/sdks/jobs/README.md#list) - Returns a list of `Job` objects.
* [Retrieve](docs/sdks/jobs/README.md#retrieve) - Returns a `Job` object with the given `id`.

### [LinkToken](docs/sdks/linktoken/README.md)

* [Create](docs/sdks/linktoken/README.md#create) - Creates a link token to be used when linking a new end user.

### [LinkedAccounts](docs/sdks/linkedaccounts/README.md)

* [List](docs/sdks/linkedaccounts/README.md#list) - List linked accounts for your organization.

### [Offers](docs/sdks/offers/README.md)

* [List](docs/sdks/offers/README.md#list) - Returns a list of `Offer` objects.
* [Retrieve](docs/sdks/offers/README.md#retrieve) - Returns an `Offer` object with the given `id`.

### [Offices](docs/sdks/offices/README.md)

* [List](docs/sdks/offices/README.md#list) - Returns a list of `Office` objects.
* [Retrieve](docs/sdks/offices/README.md#retrieve) - Returns an `Office` object with the given `id`.

### [Passthrough](docs/sdks/passthrough/README.md)

* [Create](docs/sdks/passthrough/README.md#create) - Pull data from an endpoint not currently supported by Merge.

### [RegenerateKey](docs/sdks/regeneratekey/README.md)

* [Create](docs/sdks/regeneratekey/README.md#create) - Exchange remote keys.

### [RejectReasons](docs/sdks/rejectreasons/README.md)

* [List](docs/sdks/rejectreasons/README.md#list) - Returns a list of `RejectReason` objects.
* [Retrieve](docs/sdks/rejectreasons/README.md#retrieve) - Returns a `RejectReason` object with the given `id`.

### [Scorecards](docs/sdks/scorecards/README.md)

* [List](docs/sdks/scorecards/README.md#list) - Returns a list of `Scorecard` objects.
* [Retrieve](docs/sdks/scorecards/README.md#retrieve) - Returns a `Scorecard` object with the given `id`.

### [SelectiveSync](docs/sdks/selectivesync/README.md)

* [List](docs/sdks/selectivesync/README.md#list) - Get a linked account's selective syncs.
* [RetrievePostMetadata](docs/sdks/selectivesync/README.md#retrievepostmetadata) - Get metadata for the conditions available to a linked account.
* [Update](docs/sdks/selectivesync/README.md#update) - Replace a linked account's selective syncs.

### [SyncStatus](docs/sdks/syncstatus/README.md)

* [List](docs/sdks/syncstatus/README.md#list) - Get syncing status. Possible values: `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`

### [ForceResync](docs/sdks/forceresync/README.md)

* [Create](docs/sdks/forceresync/README.md#create) - Force re-sync of all models. This is available for all organizations via the dashboard. Force re-sync is also available programmatically via API for monthly, quarterly, and highest sync frequency customers on the Core, Professional, or Enterprise plans. Doing so will consume a sync credit for the relevant linked account.

### [Tags](docs/sdks/tags/README.md)

* [List](docs/sdks/tags/README.md#list) - Returns a list of `Tag` objects.

### [Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - Returns a list of `RemoteUser` objects.
* [Retrieve](docs/sdks/users/README.md#retrieve) - Returns a `RemoteUser` object with the given `id`.

### [WebhookReceivers](docs/sdks/webhookreceivers/README.md)

* [Create](docs/sdks/webhookreceivers/README.md#create) - Creates a `WebhookReceiver` object with the given values.
* [List](docs/sdks/webhookreceivers/README.md#list) - Returns a list of `WebhookReceiver` objects.
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"log"
)

func main() {
	s := mergeatsgo.New(
		mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	var xAccountToken string = "<value>"

	ctx := context.Background()
	res, err := s.AccountDetails.Retrieve(ctx, xAccountToken)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.merge.dev/api/ats/v1` | None |
| 1 | `https://api-sandbox.merge.dev/api/ats/v1` | None |

#### Example

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
		mergeatsgo.WithServerIndex(1),
		mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	var xAccountToken string = "<value>"

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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		mergeatsgo.WithServerURL("https://api.merge.dev/api/ats/v1"),
		mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	var xAccountToken string = "<value>"

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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `TokenAuth` | apiKey      | API key     |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	mergeatsgo "github.com/speakeasy-sdks/merge-ats-go"
	"log"
)

func main() {
	s := mergeatsgo.New(
		mergeatsgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	var xAccountToken string = "<value>"

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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
