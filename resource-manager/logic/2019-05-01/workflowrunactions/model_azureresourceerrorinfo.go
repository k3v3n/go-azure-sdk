package workflowrunactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceErrorInfo struct {
	Code    string                    `json:"code"`
	Details *[]AzureResourceErrorInfo `json:"details,omitempty"`
	Message string                    `json:"message"`
}
