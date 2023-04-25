package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SwitchConfigurationRequest struct {

	// 语言。
	XLanguage *SwitchConfigurationRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationRequestBody `json:"body,omitempty"`
}

func (o SwitchConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationRequest struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationRequest", string(data)}, " ")
}

type SwitchConfigurationRequestXLanguage struct {
	value string
}

type SwitchConfigurationRequestXLanguageEnum struct {
	ZH_CN SwitchConfigurationRequestXLanguage
	EN_US SwitchConfigurationRequestXLanguage
}

func GetSwitchConfigurationRequestXLanguageEnum() SwitchConfigurationRequestXLanguageEnum {
	return SwitchConfigurationRequestXLanguageEnum{
		ZH_CN: SwitchConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
