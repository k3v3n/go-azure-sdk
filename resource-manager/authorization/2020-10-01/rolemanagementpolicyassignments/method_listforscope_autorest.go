package rolemanagementpolicyassignments

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListForScopeOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RoleManagementPolicyAssignment

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListForScopeOperationResponse, error)
}

type ListForScopeCompleteResult struct {
	Items []RoleManagementPolicyAssignment
}

func (r ListForScopeOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListForScopeOperationResponse) LoadMore(ctx context.Context) (resp ListForScopeOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListForScope ...
func (c RoleManagementPolicyAssignmentsClient) ListForScope(ctx context.Context, id commonids.ScopeId) (resp ListForScopeOperationResponse, err error) {
	req, err := c.preparerForListForScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient", "ListForScope", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient", "ListForScope", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListForScope(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient", "ListForScope", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListForScope prepares the ListForScope request.
func (c RoleManagementPolicyAssignmentsClient) preparerForListForScope(ctx context.Context, id commonids.ScopeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Authorization/roleManagementPolicyAssignments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListForScopeWithNextLink prepares the ListForScope request with the given nextLink token.
func (c RoleManagementPolicyAssignmentsClient) preparerForListForScopeWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListForScope handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (c RoleManagementPolicyAssignmentsClient) responderForListForScope(resp *http.Response) (result ListForScopeOperationResponse, err error) {
	type page struct {
		Values   []RoleManagementPolicyAssignment `json:"value"`
		NextLink *string                          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListForScopeOperationResponse, err error) {
			req, err := c.preparerForListForScopeWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient", "ListForScope", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient", "ListForScope", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListForScope(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient", "ListForScope", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListForScopeComplete retrieves all of the results into a single object
func (c RoleManagementPolicyAssignmentsClient) ListForScopeComplete(ctx context.Context, id commonids.ScopeId) (ListForScopeCompleteResult, error) {
	return c.ListForScopeCompleteMatchingPredicate(ctx, id, RoleManagementPolicyAssignmentOperationPredicate{})
}

// ListForScopeCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RoleManagementPolicyAssignmentsClient) ListForScopeCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate RoleManagementPolicyAssignmentOperationPredicate) (resp ListForScopeCompleteResult, err error) {
	items := make([]RoleManagementPolicyAssignment, 0)

	page, err := c.ListForScope(ctx, id)
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

	out := ListForScopeCompleteResult{
		Items: items,
	}
	return out, nil
}
