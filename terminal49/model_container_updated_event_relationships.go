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
// ContainerUpdatedEventRelationships struct for ContainerUpdatedEventRelationships
type ContainerUpdatedEventRelationships struct {
	Container ContainerUpdatedEventRelationshipsContainer `json:"container"`
	Terminal ContainerUpdatedEventRelationshipsTerminal `json:"terminal"`
}
