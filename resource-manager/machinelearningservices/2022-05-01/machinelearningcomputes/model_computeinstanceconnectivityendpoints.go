package machinelearningcomputes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeInstanceConnectivityEndpoints struct {
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
}