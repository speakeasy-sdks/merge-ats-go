# github.com/speakeasy-sdks/merge-ats-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/merge-ats-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
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
* [CreateChangeState](docs/sdks/applications/README.md#createchangestate) - Updates the `current_stage` field of an `Application` object
* [List](docs/sdks/applications/README.md#list) - Returns a list of `Application` objects.
* [Retrieve](docs/sdks/applications/README.md#retrieve) - Returns an `Application` object with the given `id`.
* [RetrievePostMetadata](docs/sdks/applications/README.md#retrievepostmetadata) - Returns metadata for `Application` POSTs.

### [AsyncPassthrough](docs/sdks/asyncpassthrough/README.md)

* [Create](docs/sdks/asyncpassthrough/README.md#create) - Asynchronously pull data from an endpoint not currently supported by Merge.
* [Create](docs/sdks/asyncpassthrough/README.md#create) - Asynchronously pull data from an endpoint not currently supported by Merge.
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

* [CandidatesPartialUpdate](docs/sdks/candidates/README.md#candidatespartialupdate) - Updates a `Candidate` object with the given `id`.
* [Create](docs/sdks/candidates/README.md#create) - Creates a `Candidate` object with the given values.
* [IgnoreCreate](docs/sdks/candidates/README.md#ignorecreate) - Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
* [List](docs/sdks/candidates/README.md#list) - Returns a list of `Candidate` objects.
* [Retrieve](docs/sdks/candidates/README.md#retrieve) - Returns a `Candidate` object with the given `id`.
* [RetrievePatchMetadata](docs/sdks/candidates/README.md#retrievepatchmetadata) - Returns metadata for `Candidate` PATCHs.
* [RetrievePostMetadata](docs/sdks/candidates/README.md#retrievepostmetadata) - Returns metadata for `Candidate` POSTs.

### [DeleteAccount](docs/sdks/deleteaccount/README.md)

* [DeleteAccountDelete](docs/sdks/deleteaccount/README.md#deleteaccountdelete) - Delete a linked account.

### [Departments](docs/sdks/departments/README.md)

* [List](docs/sdks/departments/README.md#list) - Returns a list of `Department` objects.
* [Retrieve](docs/sdks/departments/README.md#retrieve) - Returns a `Department` object with the given `id`.

### [Eeocs](docs/sdks/eeocs/README.md)

* [List](docs/sdks/eeocs/README.md#list) - Returns a list of `EEOC` objects.
* [Retrieve](docs/sdks/eeocs/README.md#retrieve) - Returns an `EEOC` object with the given `id`.

### [ForceResync](docs/sdks/forceresync/README.md)

* [Create](docs/sdks/forceresync/README.md#create) - Force re-sync of all models. This is available for all organizations via the dashboard. Force re-sync is also available programmatically via API for monthly, quarterly, and highest sync frequency customers on the Core, Professional, or Enterprise plans. Doing so will consume a sync credit for the relevant linked account.

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
* [Create](docs/sdks/passthrough/README.md#create) - Pull data from an endpoint not currently supported by Merge.
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
* [SelectiveSyncConfigurationsUpdate](docs/sdks/selectivesync/README.md#selectivesyncconfigurationsupdate) - Replace a linked account's selective syncs.

### [SyncStatus](docs/sdks/syncstatus/README.md)

* [List](docs/sdks/syncstatus/README.md#list) - Get syncing status. Possible values: `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`

### [Tags](docs/sdks/tags/README.md)

* [List](docs/sdks/tags/README.md#list) - Returns a list of `Tag` objects.

### [Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - Returns a list of `RemoteUser` objects.
* [Retrieve](docs/sdks/users/README.md#retrieve) - Returns a `RemoteUser` object with the given `id`.

### [WebhookReceivers](docs/sdks/webhookreceivers/README.md)

* [Create](docs/sdks/webhookreceivers/README.md#create) - Creates a `WebhookReceiver` object with the given values.
* [List](docs/sdks/webhookreceivers/README.md#list) - Returns a list of `WebhookReceiver` objects.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
