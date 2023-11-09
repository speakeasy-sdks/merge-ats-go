# Eeoc

# The EEOC Object
### Description
The `EEOC` object is used to represent the Equal Employment Opportunity Commission information for a candidate (race, gender, veteran status, disability status).
### Usage Example
Fetch from the `LIST EEOCs` endpoint and filter by `candidate` to show all EEOC information for a candidate.


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Candidate`                                                                                                                                                                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | The candidate being represented.                                                                                                                                                                                                                                                                                                                                                                                                       | f963f34d-3d2f-4f77-b557-cf36bc7e6498                                                                                                                                                                                                                                                                                                                                                                                                   |
| `DisabilityStatus`                                                                                                                                                                                                                                                                                                                                                                                                                     | [*shared.DisabilityStatus](../../../pkg/models/shared/disabilitystatus.md)                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | The candidate's disability status.<br/><br/>* `YES_I_HAVE_A_DISABILITY_OR_PREVIOUSLY_HAD_A_DISABILITY` - YES_I_HAVE_A_DISABILITY_OR_PREVIOUSLY_HAD_A_DISABILITY<br/>* `NO_I_DONT_HAVE_A_DISABILITY` - NO_I_DONT_HAVE_A_DISABILITY<br/>* `I_DONT_WISH_TO_ANSWER` - I_DONT_WISH_TO_ANSWER                                                                                                                                                | I_DONT_WISH_TO_ANSWER                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `FieldMappings`                                                                                                                                                                                                                                                                                                                                                                                                                        | map[string]*interface{}*                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                    | [object Object]                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `Gender`                                                                                                                                                                                                                                                                                                                                                                                                                               | [*shared.Gender](../../../pkg/models/shared/gender.md)                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | The candidate's gender.<br/><br/>* `MALE` - MALE<br/>* `FEMALE` - FEMALE<br/>* `NON-BINARY` - NON-BINARY<br/>* `OTHER` - OTHER<br/>* `DECLINE_TO_SELF_IDENTIFY` - DECLINE_TO_SELF_IDENTIFY                                                                                                                                                                                                                                             | FEMALE                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `ID`                                                                                                                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                    | f7dd7b4f-237e-4772-8bd4-3246384c6c58                                                                                                                                                                                                                                                                                                                                                                                                   |
| `ModifiedAt`                                                                                                                                                                                                                                                                                                                                                                                                                           | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | This is the datetime that this object was last updated by Merge                                                                                                                                                                                                                                                                                                                                                                        | 2021-10-16T00:00:00Z                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `Race`                                                                                                                                                                                                                                                                                                                                                                                                                                 | [*shared.Race](../../../pkg/models/shared/race.md)                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | The candidate's race.<br/><br/>* `AMERICAN_INDIAN_OR_ALASKAN_NATIVE` - AMERICAN_INDIAN_OR_ALASKAN_NATIVE<br/>* `ASIAN` - ASIAN<br/>* `BLACK_OR_AFRICAN_AMERICAN` - BLACK_OR_AFRICAN_AMERICAN<br/>* `HISPANIC_OR_LATINO` - HISPANIC_OR_LATINO<br/>* `WHITE` - WHITE<br/>* `NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER` - NATIVE_HAWAIIAN_OR_OTHER_PACIFIC_ISLANDER<br/>* `TWO_OR_MORE_RACES` - TWO_OR_MORE_RACES<br/>* `DECLINE_TO_SELF_IDENTIFY` - DECLINE_TO_SELF_IDENTIFY | HISPANIC_OR_LATINO                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `RemoteData`                                                                                                                                                                                                                                                                                                                                                                                                                           | [][shared.RemoteData](../../../pkg/models/shared/remotedata.md)                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                    | [object Object]                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `RemoteID`                                                                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | The third-party API ID of the matching object.                                                                                                                                                                                                                                                                                                                                                                                         | 76                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `RemoteWasDeleted`                                                                                                                                                                                                                                                                                                                                                                                                                     | **bool*                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | Indicates whether or not this object has been deleted by third party webhooks.                                                                                                                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `SubmittedAt`                                                                                                                                                                                                                                                                                                                                                                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | When the information was submitted.                                                                                                                                                                                                                                                                                                                                                                                                    | 2021-10-15T00:00:00Z                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `VeteranStatus`                                                                                                                                                                                                                                                                                                                                                                                                                        | [*shared.VeteranStatus](../../../pkg/models/shared/veteranstatus.md)                                                                                                                                                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                     | The candidate's veteran status.<br/><br/>* `I_AM_NOT_A_PROTECTED_VETERAN` - I_AM_NOT_A_PROTECTED_VETERAN<br/>* `I_IDENTIFY_AS_ONE_OR_MORE_OF_THE_CLASSIFICATIONS_OF_A_PROTECTED_VETERAN` - I_IDENTIFY_AS_ONE_OR_MORE_OF_THE_CLASSIFICATIONS_OF_A_PROTECTED_VETERAN<br/>* `I_DONT_WISH_TO_ANSWER` - I_DONT_WISH_TO_ANSWER                                                                                                               | I_AM_NOT_A_PROTECTED_VETERAN                                                                                                                                                                                                                                                                                                                                                                                                           |