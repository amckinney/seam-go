// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

// ThermostatsListRequest is an in-lined request used by the List endpoint.
type ThermostatsListRequest struct {
	ConnectedAccountId  *string       `json:"connected_account_id,omitempty"`
	ConnectedAccountIds *[]string     `json:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string       `json:"connect_webview_id,omitempty"`
	DeviceType          *DeviceType   `json:"device_type,omitempty"`
	DeviceTypes         *[]DeviceType `json:"device_types,omitempty"`
	Manufacturer        *Manufacturer `json:"manufacturer,omitempty"`
	DeviceIds           *[]string     `json:"device_ids,omitempty"`
	Limit               *float64      `json:"limit,omitempty"`
	CreatedBefore       *time.Time    `json:"created_before,omitempty"`
}
