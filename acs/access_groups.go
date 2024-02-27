// This file was auto-generated by Fern from our API Definition.

package acs

import (
	json "encoding/json"
	fmt "fmt"
	seamapigo "github.com/seamapi/go"
	core "github.com/seamapi/go/core"
)

type AccessGroupsAddUserRequest struct {
	AcsAccessGroupId string `json:"acs_access_group_id" url:"acs_access_group_id"`
	AcsUserId        string `json:"acs_user_id" url:"acs_user_id"`
}

type AccessGroupsGetRequest struct {
	AcsAccessGroupId string `json:"acs_access_group_id" url:"acs_access_group_id"`
}

type AccessGroupsListRequest struct {
	AcsSystemId *string `json:"acs_system_id,omitempty" url:"acs_system_id,omitempty"`
	AcsUserId   *string `json:"acs_user_id,omitempty" url:"acs_user_id,omitempty"`
}

type AccessGroupsListUsersRequest struct {
	AcsAccessGroupId string `json:"acs_access_group_id" url:"acs_access_group_id"`
}

type AccessGroupsRemoveUserRequest struct {
	AcsAccessGroupId string `json:"acs_access_group_id" url:"acs_access_group_id"`
	AcsUserId        string `json:"acs_user_id" url:"acs_user_id"`
}

type AccessGroupsAddUserResponse struct {
	Ok bool `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (a *AccessGroupsAddUserResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccessGroupsAddUserResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccessGroupsAddUserResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccessGroupsAddUserResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AccessGroupsGetResponse struct {
	AcsAccessGroup *seamapigo.AcsAccessGroup `json:"acs_access_group,omitempty" url:"acs_access_group,omitempty"`
	Ok             bool                      `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (a *AccessGroupsGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccessGroupsGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccessGroupsGetResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccessGroupsGetResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AccessGroupsListResponse struct {
	AcsAccessGroups []*seamapigo.AcsAccessGroup `json:"acs_access_groups,omitempty" url:"acs_access_groups,omitempty"`
	Ok              bool                        `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (a *AccessGroupsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccessGroupsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccessGroupsListResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccessGroupsListResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AccessGroupsListUsersResponse struct {
	AcsUsers []*seamapigo.AcsUser `json:"acs_users,omitempty" url:"acs_users,omitempty"`
	Ok       bool                 `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (a *AccessGroupsListUsersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccessGroupsListUsersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccessGroupsListUsersResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccessGroupsListUsersResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AccessGroupsRemoveUserResponse struct {
	Ok bool `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (a *AccessGroupsRemoveUserResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AccessGroupsRemoveUserResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AccessGroupsRemoveUserResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *AccessGroupsRemoveUserResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}
