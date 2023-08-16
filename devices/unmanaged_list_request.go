// This file was auto-generated by Fern from our API Definition.

package devices

import (
	seamapigo "github.com/seamapi/go"
	time "time"
)

// UnmanagedListRequest is an in-lined request used by the List endpoint.
type UnmanagedListRequest struct {
	ConnectedAccountId  *string                                          `json:"connected_account_id,omitempty"`
	ConnectedAccountIds *[]string                                        `json:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                                          `json:"connect_webview_id,omitempty"`
	DeviceType          *seamapigo.UnmanagedListRequestDeviceType        `json:"device_type,omitempty"`
	DeviceTypes         *[]seamapigo.UnmanagedListRequestDeviceTypesItem `json:"device_types,omitempty"`
	Manufacturer        *seamapigo.UnmanagedListRequestManufacturer      `json:"manufacturer,omitempty"`
	DeviceIds           *[]string                                        `json:"device_ids,omitempty"`
	Limit               *float64                                         `json:"limit,omitempty"`
	CreatedBefore       *time.Time                                       `json:"created_before,omitempty"`
}
