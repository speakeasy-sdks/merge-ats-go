// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package mergeatsgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/merge-ats-go/internal/hooks"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/operations"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/models/shared"
	"github.com/speakeasy-sdks/merge-ats-go/pkg/utils"
	"io"
	"net/http"
	"net/url"
)

type AsyncPassthrough struct {
	sdkConfiguration sdkConfiguration
}

func newAsyncPassthrough(sdkConfig sdkConfiguration) *AsyncPassthrough {
	return &AsyncPassthrough{
		sdkConfiguration: sdkConfig,
	}
}

// Create - Asynchronously pull data from an endpoint not currently supported by Merge.
func (s *AsyncPassthrough) Create(ctx context.Context, dataPassthroughRequest shared.DataPassthroughRequest, xAccountToken string) (*operations.AsyncPassthroughCreateResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "async_passthrough_create",
		SecuritySource: s.sdkConfiguration.Security,
	}

	request := operations.AsyncPassthroughCreateRequest{
		DataPassthroughRequest: dataPassthroughRequest,
		XAccountToken:          xAccountToken,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/async-passthrough")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "DataPassthroughRequest", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.AsyncPassthroughCreateResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.AsyncPassthroughReciept
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.AsyncPassthroughReciept = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// Retrieve - Retrieves data from earlier async-passthrough POST request
func (s *AsyncPassthrough) Retrieve(ctx context.Context, xAccountToken string, asyncPassthroughReceiptID string) (*operations.AsyncPassthroughRetrieveResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "async_passthrough_retrieve",
		SecuritySource: s.sdkConfiguration.Security,
	}

	request := operations.AsyncPassthroughRetrieveRequest{
		XAccountToken:             xAccountToken,
		AsyncPassthroughReceiptID: asyncPassthroughReceiptID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/async-passthrough/{async_passthrough_receipt_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.AsyncPassthroughRetrieveResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RemoteResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RemoteResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
