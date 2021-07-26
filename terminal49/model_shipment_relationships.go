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
// ShipmentRelationships struct for ShipmentRelationships
type ShipmentRelationships struct {
	Destination ShipmentRelationshipsDestination `json:"destination,omitempty"`
	PortOfLading ShipmentRelationshipsPortOfLading `json:"port_of_lading,omitempty"`
	Containers ShipmentRelationshipsContainers `json:"containers,omitempty"`
	PortOfDischarge ShipmentRelationshipsPortOfLading `json:"port_of_discharge,omitempty"`
	Customer ShipmentRelationshipsCustomer `json:"customer,omitempty"`
	PodTerminal ShipmentRelationshipsPodTerminal `json:"pod_terminal,omitempty"`
}
