package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowConfigurationDetailRequest struct {

	// 语言,默认：en-us。
	XLanguage *ShowConfigurationDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 参数模板ID
	ConfigId string `json:"config_id"`
}

func (o ShowConfigurationDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigurationDetailRequest", string(data)}, " ")
}

type ShowConfigurationDetailRequestXLanguage struct {
	value string
}

type ShowConfigurationDetailRequestXLanguageEnum struct {
	ZH_CN ShowConfigurationDetailRequestXLanguage
	EN_US ShowConfigurationDetailRequestXLanguage
}

func GetShowConfigurationDetailRequestXLanguageEnum() ShowConfigurationDetailRequestXLanguageEnum {
	return ShowConfigurationDetailRequestXLanguageEnum{
		ZH_CN: ShowConfigurationDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowConfigurationDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowConfigurationDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowConfigurationDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigurationDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
