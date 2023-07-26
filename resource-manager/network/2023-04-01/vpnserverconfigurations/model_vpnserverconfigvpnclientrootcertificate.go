package vpnserverconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnServerConfigVpnClientRootCertificate struct {
	Name           *string `json:"name,omitempty"`
	PublicCertData *string `json:"publicCertData,omitempty"`
}
