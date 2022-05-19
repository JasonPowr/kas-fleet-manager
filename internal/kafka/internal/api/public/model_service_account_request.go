/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// ServiceAccountRequest Schema for the request to create a service account
type ServiceAccountRequest struct {
	// The name of the service account
	Name string `json:"name"`
	// A description for the service account
	Description string `json:"description,omitempty"`
}
