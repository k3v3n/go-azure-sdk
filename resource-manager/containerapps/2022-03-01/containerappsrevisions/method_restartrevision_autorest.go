package containerappsrevisions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type RestartRevisionOperationResponse struct {
	HttpResponse *http.Response
}

// RestartRevision ...
func (c ContainerAppsRevisionsClient) RestartRevision(ctx context.Context, id RevisionId) (result RestartRevisionOperationResponse, err error) {
	req, err := c.preparerForRestartRevision(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsrevisions.ContainerAppsRevisionsClient", "RestartRevision", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsrevisions.ContainerAppsRevisionsClient", "RestartRevision", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestartRevision(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerappsrevisions.ContainerAppsRevisionsClient", "RestartRevision", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestartRevision prepares the RestartRevision request.
func (c ContainerAppsRevisionsClient) preparerForRestartRevision(ctx context.Context, id RevisionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestartRevision handles the response to the RestartRevision request. The method always
// closes the http.Response Body.
func (c ContainerAppsRevisionsClient) responderForRestartRevision(resp *http.Response) (result RestartRevisionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}