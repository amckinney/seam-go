// This file was auto-generated by Fern from our API Definition.

package api

type AugustDeviceMetadata struct {
	LockId             string  `json:"lock_id"`
	LockName           string  `json:"lock_name"`
	HouseName          string  `json:"house_name"`
	HouseId            *string `json:"house_id,omitempty"`
	HasKeypad          bool    `json:"has_keypad"`
	Model              *string `json:"model,omitempty"`
	KeypadBatteryLevel *string `json:"keypad_battery_level,omitempty"`
}
