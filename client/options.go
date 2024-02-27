// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/seamapi/go/core"
	option "github.com/seamapi/go/option"
	http "net/http"
)

// WithBaseURL sets the base URL, overriding the default
// environment, if any.
func WithBaseURL(baseURL string) *core.BaseURLOption {
	return option.WithBaseURL(baseURL)
}

// WithHTTPClient uses the given HTTPClient to issue the request.
func WithHTTPClient(httpClient core.HTTPClient) *core.HTTPClientOption {
	return option.WithHTTPClient(httpClient)
}

// WithHTTPHeader adds the given http.Header to the request.
func WithHTTPHeader(httpHeader http.Header) *core.HTTPHeaderOption {
	return option.WithHTTPHeader(httpHeader)
}

// WithMaxAttempts configures the maximum number of retry attempts.
func WithMaxAttempts(attempts uint) *core.MaxAttemptsOption {
	return option.WithMaxAttempts(attempts)
}

// WithApiKey sets the 'Authorization: Bearer <apiKey>' request header.
func WithApiKey(apiKey string) *core.ApiKeyOption {
	return option.WithApiKey(apiKey)
}

// WithSeamWorkspace sets the seamWorkspace request header.
func WithSeamWorkspace(seamWorkspace string) *core.SeamWorkspaceOption {
	return option.WithSeamWorkspace(seamWorkspace)
}

// WithSeamClientSessionToken sets the seamClientSessionToken request header.
func WithSeamClientSessionToken(seamClientSessionToken string) *core.SeamClientSessionTokenOption {
	return option.WithSeamClientSessionToken(seamClientSessionToken)
}

// WithClientSessionToken sets the clientSessionToken request header.
func WithClientSessionToken(clientSessionToken string) *core.ClientSessionTokenOption {
	return option.WithClientSessionToken(clientSessionToken)
}
