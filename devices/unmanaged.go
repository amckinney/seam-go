// This file was auto-generated by Fern from our API Definition.

package devices

import (
	json "encoding/json"
	fmt "fmt"
	seamapigo "github.com/seamapi/go"
	core "github.com/seamapi/go/core"
	time "time"
)

type UnmanagedGetRequest struct {
	DeviceId *string `json:"device_id,omitempty" url:"device_id,omitempty"`
	Name     *string `json:"name,omitempty" url:"name,omitempty"`
}

type UnmanagedListRequest struct {
	// List all devices owned by this connected account
	ConnectedAccountId  *string                                                `json:"connected_account_id,omitempty" url:"connected_account_id,omitempty"`
	ConnectedAccountIds []string                                               `json:"connected_account_ids,omitempty" url:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                                                `json:"connect_webview_id,omitempty" url:"connect_webview_id,omitempty"`
	DeviceType          *seamapigo.DeviceType                                  `json:"device_type,omitempty" url:"device_type,omitempty"`
	DeviceTypes         []seamapigo.DeviceType                                 `json:"device_types,omitempty" url:"device_types,omitempty"`
	Manufacturer        *seamapigo.Manufacturer                                `json:"manufacturer,omitempty" url:"manufacturer,omitempty"`
	DeviceIds           []string                                               `json:"device_ids,omitempty" url:"device_ids,omitempty"`
	Limit               *float64                                               `json:"limit,omitempty" url:"limit,omitempty"`
	CreatedBefore       *time.Time                                             `json:"created_before,omitempty" url:"created_before,omitempty"`
	UserIdentifierKey   *string                                                `json:"user_identifier_key,omitempty" url:"user_identifier_key,omitempty"`
	CustomMetadataHas   map[string]*UnmanagedListRequestCustomMetadataHasValue `json:"custom_metadata_has,omitempty" url:"custom_metadata_has,omitempty"`
	IncludeIf           []UnmanagedListRequestIncludeIfItem                    `json:"include_if,omitempty" url:"include_if,omitempty"`
	ExcludeIf           []UnmanagedListRequestExcludeIfItem                    `json:"exclude_if,omitempty" url:"exclude_if,omitempty"`
}

func (u *UnmanagedListRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedListRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*u = UnmanagedListRequest(body)
	return nil
}

func (u *UnmanagedListRequest) MarshalJSON() ([]byte, error) {
	type embed UnmanagedListRequest
	var marshaler = struct {
		embed
		CreatedBefore *core.DateTime `json:"created_before,omitempty"`
	}{
		embed:         embed(*u),
		CreatedBefore: core.NewOptionalDateTime(u.CreatedBefore),
	}
	return json.Marshal(marshaler)
}

type UnmanagedGetResponse struct {
	Device *seamapigo.UnmanagedDevice `json:"device,omitempty" url:"device,omitempty"`
	Ok     bool                       `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (u *UnmanagedGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnmanagedGetResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnmanagedGetResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UnmanagedListRequestCustomMetadataHasValue struct {
	typeName string
	String   string
	Boolean  bool
}

func NewUnmanagedListRequestCustomMetadataHasValueFromString(value string) *UnmanagedListRequestCustomMetadataHasValue {
	return &UnmanagedListRequestCustomMetadataHasValue{typeName: "string", String: value}
}

func NewUnmanagedListRequestCustomMetadataHasValueFromBoolean(value bool) *UnmanagedListRequestCustomMetadataHasValue {
	return &UnmanagedListRequestCustomMetadataHasValue{typeName: "boolean", Boolean: value}
}

func (u *UnmanagedListRequestCustomMetadataHasValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		u.typeName = "string"
		u.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		u.typeName = "boolean"
		u.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UnmanagedListRequestCustomMetadataHasValue) MarshalJSON() ([]byte, error) {
	switch u.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "string":
		return json.Marshal(u.String)
	case "boolean":
		return json.Marshal(u.Boolean)
	}
}

type UnmanagedListRequestCustomMetadataHasValueVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
}

func (u *UnmanagedListRequestCustomMetadataHasValue) Accept(visitor UnmanagedListRequestCustomMetadataHasValueVisitor) error {
	switch u.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "string":
		return visitor.VisitString(u.String)
	case "boolean":
		return visitor.VisitBoolean(u.Boolean)
	}
}

type UnmanagedListRequestExcludeIfItem string

const (
	UnmanagedListRequestExcludeIfItemCanRemotelyUnlock            UnmanagedListRequestExcludeIfItem = "can_remotely_unlock"
	UnmanagedListRequestExcludeIfItemCanRemotelyLock              UnmanagedListRequestExcludeIfItem = "can_remotely_lock"
	UnmanagedListRequestExcludeIfItemCanProgramOfflineAccessCodes UnmanagedListRequestExcludeIfItem = "can_program_offline_access_codes"
	UnmanagedListRequestExcludeIfItemCanProgramOnlineAccessCodes  UnmanagedListRequestExcludeIfItem = "can_program_online_access_codes"
	UnmanagedListRequestExcludeIfItemCanSimulateRemoval           UnmanagedListRequestExcludeIfItem = "can_simulate_removal"
)

func NewUnmanagedListRequestExcludeIfItemFromString(s string) (UnmanagedListRequestExcludeIfItem, error) {
	switch s {
	case "can_remotely_unlock":
		return UnmanagedListRequestExcludeIfItemCanRemotelyUnlock, nil
	case "can_remotely_lock":
		return UnmanagedListRequestExcludeIfItemCanRemotelyLock, nil
	case "can_program_offline_access_codes":
		return UnmanagedListRequestExcludeIfItemCanProgramOfflineAccessCodes, nil
	case "can_program_online_access_codes":
		return UnmanagedListRequestExcludeIfItemCanProgramOnlineAccessCodes, nil
	case "can_simulate_removal":
		return UnmanagedListRequestExcludeIfItemCanSimulateRemoval, nil
	}
	var t UnmanagedListRequestExcludeIfItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UnmanagedListRequestExcludeIfItem) Ptr() *UnmanagedListRequestExcludeIfItem {
	return &u
}

type UnmanagedListRequestIncludeIfItem string

const (
	UnmanagedListRequestIncludeIfItemCanRemotelyUnlock            UnmanagedListRequestIncludeIfItem = "can_remotely_unlock"
	UnmanagedListRequestIncludeIfItemCanRemotelyLock              UnmanagedListRequestIncludeIfItem = "can_remotely_lock"
	UnmanagedListRequestIncludeIfItemCanProgramOfflineAccessCodes UnmanagedListRequestIncludeIfItem = "can_program_offline_access_codes"
	UnmanagedListRequestIncludeIfItemCanProgramOnlineAccessCodes  UnmanagedListRequestIncludeIfItem = "can_program_online_access_codes"
	UnmanagedListRequestIncludeIfItemCanSimulateRemoval           UnmanagedListRequestIncludeIfItem = "can_simulate_removal"
)

func NewUnmanagedListRequestIncludeIfItemFromString(s string) (UnmanagedListRequestIncludeIfItem, error) {
	switch s {
	case "can_remotely_unlock":
		return UnmanagedListRequestIncludeIfItemCanRemotelyUnlock, nil
	case "can_remotely_lock":
		return UnmanagedListRequestIncludeIfItemCanRemotelyLock, nil
	case "can_program_offline_access_codes":
		return UnmanagedListRequestIncludeIfItemCanProgramOfflineAccessCodes, nil
	case "can_program_online_access_codes":
		return UnmanagedListRequestIncludeIfItemCanProgramOnlineAccessCodes, nil
	case "can_simulate_removal":
		return UnmanagedListRequestIncludeIfItemCanSimulateRemoval, nil
	}
	var t UnmanagedListRequestIncludeIfItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UnmanagedListRequestIncludeIfItem) Ptr() *UnmanagedListRequestIncludeIfItem {
	return &u
}

type UnmanagedListResponse struct {
	Devices []*seamapigo.UnmanagedDevice `json:"devices,omitempty" url:"devices,omitempty"`
	Ok      bool                         `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (u *UnmanagedListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnmanagedListResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnmanagedListResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UnmanagedUpdateResponse struct {
	Ok bool `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (u *UnmanagedUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UnmanagedUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UnmanagedUpdateResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UnmanagedUpdateResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UnmanagedUpdateRequest struct {
	DeviceId  string `json:"device_id" url:"device_id"`
	IsManaged bool   `json:"is_managed" url:"is_managed"`
}
