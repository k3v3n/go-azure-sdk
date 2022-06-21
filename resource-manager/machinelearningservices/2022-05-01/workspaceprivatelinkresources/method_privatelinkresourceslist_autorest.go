package workspaceprivatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type PrivateLinkResourcesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResourceListResult
}

// PrivateLinkResourcesList ...
func (c WorkspacePrivateLinkResourcesClient) PrivateLinkResourcesList(ctx context.Context, id WorkspaceId) (result PrivateLinkResourcesListOperationResponse, err error) {
	req, err := c.preparerForPrivateLinkResourcesList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient", "PrivateLinkResourcesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient", "PrivateLinkResourcesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateLinkResourcesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient", "PrivateLinkResourcesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateLinkResourcesList prepares the PrivateLinkResourcesList request.
func (c WorkspacePrivateLinkResourcesClient) preparerForPrivateLinkResourcesList(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateLinkResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPrivateLinkResourcesList handles the response to the PrivateLinkResourcesList request. The method always
// closes the http.Response Body.
func (c WorkspacePrivateLinkResourcesClient) responderForPrivateLinkResourcesList(resp *http.Response) (result PrivateLinkResourcesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}