package targettypes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetTypesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewTargetTypesClientWithBaseURI(endpoint string) TargetTypesClient {
	return TargetTypesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}