package config

func GetDbUrl() string { return "postgres://gossip:gossip@127.0.0.1/gossip?sslmode=disable" }

func GetAmqpUrl() string { return "amqp://localhost/" }
