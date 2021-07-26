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
// WebhookNotificationAttributes struct for WebhookNotificationAttributes
type WebhookNotificationAttributes struct {
	Event string `json:"event"`
	// Whether the notification has been delivered to the webhook endpoint
	DeliveryStatus string `json:"delivery_status"`
	CreatedAt string `json:"created_at"`
}
