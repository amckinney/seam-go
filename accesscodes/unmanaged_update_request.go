// This file was auto-generated by Fern from our API Definition.

package accesscodes

// UnmanagedUpdateRequest is an in-lined request used by the Update endpoint.
type UnmanagedUpdateRequest struct {
	AccessCodeId string `json:"access_code_id"`
	IsManaged    bool   `json:"is_managed"`
	Force        *bool  `json:"force,omitempty"`
}