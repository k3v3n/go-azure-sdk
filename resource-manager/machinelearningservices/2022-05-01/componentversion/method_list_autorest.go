package componentversion

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ComponentVersionResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListOperationResponse, error)
}

type ListCompleteResult struct {
	Items []ComponentVersionResource
}

func (r ListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListOperationResponse) LoadMore(ctx context.Context) (resp ListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListOperationOptions struct {
	ListViewType *ListViewType
	OrderBy      *string
	Skip         *string
	Top          *int64
}

func DefaultListOperationOptions() ListOperationOptions {
	return ListOperationOptions{}
}

func (o ListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.OrderBy != nil {
		out["$orderBy"] = *o.OrderBy
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// List ...
func (c ComponentVersionClient) List(ctx context.Context, id ComponentId, options ListOperationOptions) (resp ListOperationResponse, err error) {
	req, err := c.preparerForList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "List", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "List", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "List", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListComplete retrieves all of the results into a single object
func (c ComponentVersionClient) ListComplete(ctx context.Context, id ComponentId, options ListOperationOptions) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, options, ComponentVersionResourceOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ComponentVersionClient) ListCompleteMatchingPredicate(ctx context.Context, id ComponentId, options ListOperationOptions, predicate ComponentVersionResourceOperationPredicate) (resp ListCompleteResult, err error) {
	items := make([]ComponentVersionResource, 0)

	page, err := c.List(ctx, id, options)
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

	out := ListCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForList prepares the List request.
func (c ComponentVersionClient) preparerForList(ctx context.Context, id ComponentId, options ListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWithNextLink prepares the List request with the given nextLink token.
func (c ComponentVersionClient) preparerForListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForList handles the response to the List request. The method always
// closes the http.Response Body.
func (c ComponentVersionClient) responderForList(resp *http.Response) (result ListOperationResponse, err error) {
	type page struct {
		Values   []ComponentVersionResource `json:"value"`
		NextLink *string                    `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListOperationResponse, err error) {
			req, err := c.preparerForListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "List", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "List", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "List", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}