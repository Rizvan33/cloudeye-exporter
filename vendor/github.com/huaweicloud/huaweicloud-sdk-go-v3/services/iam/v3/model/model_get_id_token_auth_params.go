package model

import (
	"encoding/json"

	"strings"
)

// auth信息
type GetIdTokenAuthParams struct {
	IdToken *GetIdTokenIdTokenBody `json:"id_token"`

	Scope *GetIdTokenIdScopeBody `json:"scope,omitempty"`
}

func (o GetIdTokenAuthParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetIdTokenAuthParams struct{}"
	}

	return strings.Join([]string{"GetIdTokenAuthParams", string(data)}, " ")
}
