// This file was auto-generated by Fern from our API Definition.

package acs

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
)

type CredentialPoolsListRequest struct {
	AcsSystemId string `json:"acs_system_id" url:"acs_system_id"`
}

type CredentialPoolsListResponse struct {
	AcsCredentialPools []*CredentialPoolsListResponseAcsCredentialPoolsItem `json:"acs_credential_pools,omitempty" url:"acs_credential_pools,omitempty"`
	Ok                 bool                                                 `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialPoolsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialPoolsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialPoolsListResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialPoolsListResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}
