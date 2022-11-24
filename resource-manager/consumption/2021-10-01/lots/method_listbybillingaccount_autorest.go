package lots

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]LotSummary

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByBillingAccountOperationResponse, error)
}

type ListByBillingAccountCompleteResult struct {
	Items []LotSummary
}

func (r ListByBillingAccountOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByBillingAccountOperationResponse) LoadMore(ctx context.Context) (resp ListByBillingAccountOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByBillingAccountOperationOptions struct {
	Filter *string
}

func DefaultListByBillingAccountOperationOptions() ListByBillingAccountOperationOptions {
	return ListByBillingAccountOperationOptions{}
}

func (o ListByBillingAccountOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByBillingAccountOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByBillingAccount ...
func (c LotsClient) ListByBillingAccount(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (resp ListByBillingAccountOperationResponse, err error) {
	req, err := c.preparerForListByBillingAccount(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByBillingAccount", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByBillingAccount(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByBillingAccount", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByBillingAccount prepares the ListByBillingAccount request.
func (c LotsClient) preparerForListByBillingAccount(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/lots", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByBillingAccountWithNextLink prepares the ListByBillingAccount request with the given nextLink token.
func (c LotsClient) preparerForListByBillingAccountWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByBillingAccount handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (c LotsClient) responderForListByBillingAccount(resp *http.Response) (result ListByBillingAccountOperationResponse, err error) {
	type page struct {
		Values   []LotSummary `json:"value"`
		NextLink *string      `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByBillingAccountOperationResponse, err error) {
			req, err := c.preparerForListByBillingAccountWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByBillingAccount", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByBillingAccount", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByBillingAccount(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "lots.LotsClient", "ListByBillingAccount", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByBillingAccountComplete retrieves all of the results into a single object
func (c LotsClient) ListByBillingAccountComplete(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (ListByBillingAccountCompleteResult, error) {
	return c.ListByBillingAccountCompleteMatchingPredicate(ctx, id, options, LotSummaryOperationPredicate{})
}

// ListByBillingAccountCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c LotsClient) ListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions, predicate LotSummaryOperationPredicate) (resp ListByBillingAccountCompleteResult, err error) {
	items := make([]LotSummary, 0)

	page, err := c.ListByBillingAccount(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByBillingAccountCompleteResult{
		Items: items,
	}
	return out, nil
}