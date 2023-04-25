package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectQuotasResponse struct {
	Quotas         *ProjectQuotasResult `json:"quotas,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowProjectQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectQuotasResponse", string(data)}, " ")
}
