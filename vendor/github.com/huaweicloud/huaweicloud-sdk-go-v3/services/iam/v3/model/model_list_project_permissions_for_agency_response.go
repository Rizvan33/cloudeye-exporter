package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProjectPermissionsForAgencyResponse struct {
	// 权限信息列表。

	Roles          *[]RoleResult `json:"roles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListProjectPermissionsForAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectPermissionsForAgencyResponse struct{}"
	}

	return strings.Join([]string{"ListProjectPermissionsForAgencyResponse", string(data)}, " ")
}
