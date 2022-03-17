package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowOpenIdConnectConfigRequest struct {
	// 身份提供商ID

	IdpId string `json:"idp_id"`
}

func (o ShowOpenIdConnectConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOpenIdConnectConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowOpenIdConnectConfigRequest", string(data)}, " ")
}
