package azurebackupjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserFacingWarningDetail struct {
	ResourceName *string         `json:"resourceName,omitempty"`
	Warning      UserFacingError `json:"warning"`
}
