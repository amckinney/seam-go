// This file was auto-generated by Fern from our API Definition.

package thermostats

// ClimateSettingSchedulesGetRequest is an in-lined request used by the Get endpoint.
type ClimateSettingSchedulesGetRequest struct {
	ClimateSettingScheduleId *string `json:"climate_setting_schedule_id,omitempty"`
	DeviceId                 *string `json:"device_id,omitempty"`
}