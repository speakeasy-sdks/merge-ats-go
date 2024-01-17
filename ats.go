// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package mergeatsgo

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production
	"https://api.merge.dev/api/ats/v1",
	// Sandbox
	"https://api-sandbox.merge.dev/api/ats/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Ats - Merge ATS API: The unified API for building rich integrations with multiple Applicant Tracking System platforms.
type Ats struct {
	AccountDetails     *AccountDetails
	AccountToken       *AccountToken
	Activities         *Activities
	Applications       *Applications
	AsyncPassthrough   *AsyncPassthrough
	Attachments        *Attachments
	AvailableActions   *AvailableActions
	Candidates         *Candidates
	DeleteAccount      *DeleteAccount
	Departments        *Departments
	Eeocs              *Eeocs
	GenerateKey        *GenerateKey
	Interviews         *Interviews
	Issues             *Issues
	JobInterviewStages *JobInterviewStages
	Jobs               *Jobs
	LinkToken          *LinkToken
	LinkedAccounts     *LinkedAccounts
	Offers             *Offers
	Offices            *Offices
	Passthrough        *Passthrough
	RegenerateKey      *RegenerateKey
	RejectReasons      *RejectReasons
	Scorecards         *Scorecards
	SelectiveSync      *SelectiveSync
	SyncStatus         *SyncStatus
	ForceResync        *ForceResync
	Tags               *Tags
	Users              *Users
	WebhookReceivers   *WebhookReceivers

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Ats)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Ats) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Ats) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Ats) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Ats) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(tokenAuth string) SDKOption {
	return func(sdk *Ats) {
		security := shared.Security{TokenAuth: &tokenAuth}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Ats) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Ats {
	sdk := &Ats{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0",
			SDKVersion:        "0.9.2",
			GenVersion:        "2.234.3",
			UserAgent:         "speakeasy-sdk/go 0.9.2 2.234.3 1.0 github.com/speakeasy-sdks/merge-ats-go",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.AccountDetails = newAccountDetails(sdk.sdkConfiguration)

	sdk.AccountToken = newAccountToken(sdk.sdkConfiguration)

	sdk.Activities = newActivities(sdk.sdkConfiguration)

	sdk.Applications = newApplications(sdk.sdkConfiguration)

	sdk.AsyncPassthrough = newAsyncPassthrough(sdk.sdkConfiguration)

	sdk.Attachments = newAttachments(sdk.sdkConfiguration)

	sdk.AvailableActions = newAvailableActions(sdk.sdkConfiguration)

	sdk.Candidates = newCandidates(sdk.sdkConfiguration)

	sdk.DeleteAccount = newDeleteAccount(sdk.sdkConfiguration)

	sdk.Departments = newDepartments(sdk.sdkConfiguration)

	sdk.Eeocs = newEeocs(sdk.sdkConfiguration)

	sdk.GenerateKey = newGenerateKey(sdk.sdkConfiguration)

	sdk.Interviews = newInterviews(sdk.sdkConfiguration)

	sdk.Issues = newIssues(sdk.sdkConfiguration)

	sdk.JobInterviewStages = newJobInterviewStages(sdk.sdkConfiguration)

	sdk.Jobs = newJobs(sdk.sdkConfiguration)

	sdk.LinkToken = newLinkToken(sdk.sdkConfiguration)

	sdk.LinkedAccounts = newLinkedAccounts(sdk.sdkConfiguration)

	sdk.Offers = newOffers(sdk.sdkConfiguration)

	sdk.Offices = newOffices(sdk.sdkConfiguration)

	sdk.Passthrough = newPassthrough(sdk.sdkConfiguration)

	sdk.RegenerateKey = newRegenerateKey(sdk.sdkConfiguration)

	sdk.RejectReasons = newRejectReasons(sdk.sdkConfiguration)

	sdk.Scorecards = newScorecards(sdk.sdkConfiguration)

	sdk.SelectiveSync = newSelectiveSync(sdk.sdkConfiguration)

	sdk.SyncStatus = newSyncStatus(sdk.sdkConfiguration)

	sdk.ForceResync = newForceResync(sdk.sdkConfiguration)

	sdk.Tags = newTags(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.WebhookReceivers = newWebhookReceivers(sdk.sdkConfiguration)

	return sdk
}
