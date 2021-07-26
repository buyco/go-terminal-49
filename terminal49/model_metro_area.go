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
// MetroArea struct for MetroArea
type MetroArea struct {
	Id string `json:"id"`
	Attributes MetroAreaAttributes `json:"attributes,omitempty"`
	Type string `json:"type"`
}
