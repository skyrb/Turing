package main

import "fmt"

// Configuration represents the configuration settings
type Configuration struct {
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost      string
	DatabasePort      int
	LogLevel          string
	APIKey            string
}

// ConfigurationBuilder is a builder for the Configuration object
type ConfigurationBuilder struct {
	config *Configuration
}

// NewConfigurationBuilder creates a new ConfigurationBuilder
func NewConfigurationBuilder() *ConfigurationBuilder {
	return &ConfigurationBuilder{config: &Configuration{}}
}

// WithDatabaseUsername sets the database username
func (b *ConfigurationBuilder) WithDatabaseUsername(username string) *ConfigurationBuilder {
	b.config.DatabaseUsername = username
	return b
}

// WithDatabasePassword sets the database password
func (b *ConfigurationBuilder) WithDatabasePassword(password string) *ConfigurationBuilder {
	b.config.DatabasePassword = password
	return b
}

// WithDatabaseHost sets the database host
func (b *ConfigurationBuilder) WithDatabaseHost(host string) *ConfigurationBuilder {
	b.config.DatabaseHost = host
	return b
}

// WithDatabasePort sets the database port
func (b *ConfigurationBuilder) WithDatabasePort(port int) *ConfigurationBuilder {
	b.config.DatabasePort = port
	return b
}

// WithLogLevel sets the log level
func (b *ConfigurationBuilder) WithLogLevel(logLevel string) *ConfigurationBuilder {
	b.config.LogLevel = logLevel
	return b
}

// WithAPIKey sets the API key
func (b *ConfigurationBuilder) WithAPIKey(apiKey string) *ConfigurationBuilder {
	b.config.APIKey = apiKey
	return b
}

// Build constructs and returns the Configuration object
func (b *ConfigurationBuilder) Build() *Configuration {
	return b.config
}

func main() {
	// Creating a configuration object using a builder
	config := NewConfigurationBuilder().
		WithDatabaseUsername("admin").
		WithDatabasePassword("secret").
		WithDatabaseHost("localhost").
		WithDatabasePort(5432).
		WithLogLevel("INFO").
		WithAPIKey("api-12345").
		Build()

	fmt.Printf("Configuration: %+v\n", config)
}