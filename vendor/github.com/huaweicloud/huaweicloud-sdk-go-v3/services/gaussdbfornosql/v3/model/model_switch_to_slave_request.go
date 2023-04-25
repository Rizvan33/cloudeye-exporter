package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchToSlaveRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o SwitchToSlaveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchToSlaveRequest struct{}"
	}

	return strings.Join([]string{"SwitchToSlaveRequest", string(data)}, " ")
}
