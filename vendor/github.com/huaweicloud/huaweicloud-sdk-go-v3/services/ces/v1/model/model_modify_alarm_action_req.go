package model

import (
	"encoding/json"

	"strings"
)

//
type ModifyAlarmActionReq struct {
	// 告警是否启用。true：启动。false：停止

	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o ModifyAlarmActionReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyAlarmActionReq struct{}"
	}

	return strings.Join([]string{"ModifyAlarmActionReq", string(data)}, " ")
}
