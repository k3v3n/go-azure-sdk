package charges

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChargeSummary = LegacyChargeSummary{}

type LegacyChargeSummary struct {
	Properties LegacyChargeSummaryProperties `json:"properties"`

	// Fields inherited from ChargeSummary
	ETag *string `json:"eTag,omitempty"`
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = LegacyChargeSummary{}

func (s LegacyChargeSummary) MarshalJSON() ([]byte, error) {
	type wrapper LegacyChargeSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LegacyChargeSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LegacyChargeSummary: %+v", err)
	}
	decoded["kind"] = "legacy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LegacyChargeSummary: %+v", err)
	}

	return encoded, nil
}
