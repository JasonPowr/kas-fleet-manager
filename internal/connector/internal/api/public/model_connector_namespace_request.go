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

// ConnectorNamespaceRequest A connector namespace create request
type ConnectorNamespaceRequest struct {
	// Namespace name must match pattern `^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$`, or it may be empty to be auto-generated.
	Name        string                       `json:"name"`
	Annotations map[string]string            `json:"annotations,omitempty"`
	ClusterId   string                       `json:"cluster_id"`
	Kind        ConnectorNamespaceTenantKind `json:"kind"`
}
