package config

type Kafka struct {
	Brokers              []string `mapstructure:"brokers"`
	Partition            int32    `mapstructure:"partition"`
	Partitioner          string   `mapstructure:"partitioner"` // "The partitioning scheme to use. Can be `hash`, `manual`, or `random`")
	SaslProducerUsername string   `mapstructure:"sasl_producer_username"`
	SaslProducerPassword string   `mapstructure:"sasl_producer_password"`
	SaslConsumerUsername string   `mapstructure:"sasl_consumer_username"`
	SaslConsumerPassword string   `mapstructure:"sasl_consumer_password"`
	KafkaEnabled         bool     `mapstructure:"enable"`
	UserProfileGroupID   string   `mapstructure:"user_profile_group_id"`
}
