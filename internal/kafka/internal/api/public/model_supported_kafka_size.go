/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// SupportedKafkaSize Supported Kafka Size
type SupportedKafkaSize struct {
	// Unique identifier of this Kafka instance size.
	Id string `json:"id,omitempty"`
	// Display name of this Kafka instance size.
	DisplayName             string                           `json:"display_name,omitempty"`
	IngressThroughputPerSec SupportedKafkaSizeBytesValueItem `json:"ingress_throughput_per_sec,omitempty"`
	EgressThroughputPerSec  SupportedKafkaSizeBytesValueItem `json:"egress_throughput_per_sec,omitempty"`
	// Maximum amount of total connections available to this Kafka instance size.
	TotalMaxConnections  int32                            `json:"total_max_connections,omitempty"`
	MaxDataRetentionSize SupportedKafkaSizeBytesValueItem `json:"max_data_retention_size,omitempty"`
	// Maximum amount of total partitions available to this Kafka instance size.
	MaxPartitions int32 `json:"max_partitions,omitempty"`
	// Maximum data retention period available to this Kafka instance size.
	MaxDataRetentionPeriod string `json:"max_data_retention_period,omitempty"`
	// Maximium connection attempts per second available to this Kafka instance size.
	MaxConnectionAttemptsPerSec int32                            `json:"max_connection_attempts_per_sec,omitempty"`
	MaxMessageSize              SupportedKafkaSizeBytesValueItem `json:"max_message_size,omitempty"`
	// Minimum number of in-sync replicas.
	MinInSyncReplicas int32 `json:"min_in_sync_replicas,omitempty"`
	// Replication factor available to this Kafka instance size.
	ReplicationFactor int32 `json:"replication_factor,omitempty"`
	// List of Availability Zone modes that this Kafka instance size supports. The possible values are \"single\", \"multi\".
	SupportedAzModes []string `json:"supported_az_modes,omitempty"`
	// The limit lifespan of the kafka instance in seconds. If not specified then the instance never expires.
	LifespanSeconds *int32 `json:"lifespan_seconds,omitempty"`
	// Quota consumed by this Kafka instance size.
	QuotaConsumed int32 `json:"quota_consumed,omitempty"`
	// Quota type used by this Kafka instance size.
	QuotaType string `json:"quota_type,omitempty"`
	// Data plane cluster capacity consumed by this Kafka instance size.
	CapacityConsumed int32 `json:"capacity_consumed,omitempty"`
	// Maturity level of the size. Can be \"stable\" or \"preview\".
	MaturityStatus string `json:"maturity_status,omitempty"`
}
