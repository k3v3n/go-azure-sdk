package targetcomputesizes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByReplicationProtectedItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TargetComputeSize
}

type ListByReplicationProtectedItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TargetComputeSize
}

// ListByReplicationProtectedItems ...
func (c TargetComputeSizesClient) ListByReplicationProtectedItems(ctx context.Context, id ReplicationProtectedItemId) (result ListByReplicationProtectedItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/targetComputeSizes", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]TargetComputeSize `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByReplicationProtectedItemsComplete retrieves all the results into a single object
func (c TargetComputeSizesClient) ListByReplicationProtectedItemsComplete(ctx context.Context, id ReplicationProtectedItemId) (ListByReplicationProtectedItemsCompleteResult, error) {
	return c.ListByReplicationProtectedItemsCompleteMatchingPredicate(ctx, id, TargetComputeSizeOperationPredicate{})
}

// ListByReplicationProtectedItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TargetComputeSizesClient) ListByReplicationProtectedItemsCompleteMatchingPredicate(ctx context.Context, id ReplicationProtectedItemId, predicate TargetComputeSizeOperationPredicate) (result ListByReplicationProtectedItemsCompleteResult, err error) {
	items := make([]TargetComputeSize, 0)

	resp, err := c.ListByReplicationProtectedItems(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByReplicationProtectedItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
