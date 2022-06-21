package intelligencepacks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type EnableOperationResponse struct {
	HttpResponse *http.Response
}

// Enable ...
func (c IntelligencePacksClient) Enable(ctx context.Context, id IntelligencePackId) (result EnableOperationResponse, err error) {
	req, err := c.preparerForEnable(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intelligencepacks.IntelligencePacksClient", "Enable", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "intelligencepacks.IntelligencePacksClient", "Enable", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEnable(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intelligencepacks.IntelligencePacksClient", "Enable", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEnable prepares the Enable request.
func (c IntelligencePacksClient) preparerForEnable(ctx context.Context, id IntelligencePackId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/enable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEnable handles the response to the Enable request. The method always
// closes the http.Response Body.
func (c IntelligencePacksClient) responderForEnable(resp *http.Response) (result EnableOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}