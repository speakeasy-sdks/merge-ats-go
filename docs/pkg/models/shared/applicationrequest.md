# ApplicationRequest

# The Application Object
### Description
The Application Object is used to represent a candidate's journey through a particular Job's recruiting process. If a Candidate applies for multiple Jobs, there will be a separate Application for each Job if the third-party integration allows it.

### Usage Example
Fetch from the `LIST Applications` endpoint and filter by `ID` to show all applications.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `AppliedAt`                                | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | When the application was submitted.        | 2021-10-15T00:00:00Z                       |
| `Candidate`                                | **string*                                  | :heavy_minus_sign:                         | The candidate applying.                    | 2872ba14-4084-492b-be96-e5eee6fc33ef       |
| `CreditedTo`                               | **string*                                  | :heavy_minus_sign:                         | The user credited for this application.    | 58166795-8d68-4b30-9bfb-bfd402479484       |
| `CurrentStage`                             | **string*                                  | :heavy_minus_sign:                         | The application's current stage.           | d578dfdc-7b0a-4ab6-a2b0-4b40f20eb9ea       |
| `IntegrationParams`                        | map[string]*interface{}*                   | :heavy_minus_sign:                         | N/A                                        | [object Object]                            |
| `Job`                                      | **string*                                  | :heavy_minus_sign:                         | The job being applied for.                 | 52bf9b5e-0beb-4f6f-8a72-cd4dca7ca633       |
| `LinkedAccountParams`                      | map[string]*interface{}*                   | :heavy_minus_sign:                         | N/A                                        | [object Object]                            |
| `RejectReason`                             | **string*                                  | :heavy_minus_sign:                         | The application's reason for rejection.    | 59b25f2b-da02-40f5-9656-9fa0db555784       |
| `RejectedAt`                               | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | When the application was rejected.         | 2021-11-15T00:00:00Z                       |
| `RemoteTemplateID`                         | **string*                                  | :heavy_minus_sign:                         | N/A                                        | 92830948203                                |
| `Source`                                   | **string*                                  | :heavy_minus_sign:                         | The application's source.                  | Campus recruiting event                    |