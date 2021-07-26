/*
 * Terminal49 API Reference
 *
 * The Terminal 49 API offers a convenient way to programmatically track your shipments from origin to destination.
 *
 * API version: 0.2.0
 * Contact: support@terminal49.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package terminal49
// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	Data []TrackingRequest `json:"data,omitempty"`
	Links Links `json:"links,omitempty"`
	Meta Meta `json:"meta,omitempty"`
	Included []interface{} `json:"included,omitempty"`
}
