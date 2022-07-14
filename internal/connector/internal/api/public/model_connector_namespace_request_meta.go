/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// ConnectorNamespaceRequestMeta struct for ConnectorNamespaceRequestMeta
type ConnectorNamespaceRequestMeta struct {
	// Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
	Name        string            `json:"name,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}
