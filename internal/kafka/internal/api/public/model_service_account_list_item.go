/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

import (
	"time"
)

// ServiceAccountListItem struct for ServiceAccountListItem
type ServiceAccountListItem struct {
	// server generated unique id of the service account
	Id   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Href string `json:"href,omitempty"`
	// client id of the service account
	ClientId string `json:"client_id,omitempty"`
	// name of the service account
	Name string `json:"name,omitempty"`
	// owner of the service account
	// Deprecated
	DeprecatedOwner string `json:"owner,omitempty"`
	// service account created by the user
	CreatedBy string `json:"created_by,omitempty"`
	// service account creation timestamp
	CreatedAt time.Time `json:"created_at,omitempty"`
	// description of the service account
	Description string `json:"description,omitempty"`
}
