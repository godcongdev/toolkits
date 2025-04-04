/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

type ParamConfig interface {
	String() string
}

// Config represents the configuration for hash algorithms
type Config struct {
	SaltLength  int    `env:"HASH_SALTLENGTH"`
	ParamConfig string `env:"HASH_PARAM_CONFIG"`
}

// ConfigOption is a function that modifies a Config
type ConfigOption func(*Config)

// WithSaltLength sets the salt length
func WithSaltLength(length int) ConfigOption {
	return func(cfg *Config) {
		cfg.SaltLength = length
	}
}

func WithParams(paramConfig ParamConfig) ConfigOption {
	return func(cfg *Config) {
		if paramConfig == nil {
			return
		}
		cfg.ParamConfig = paramConfig.String()
	}
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		SaltLength: 16, // Default salt length
	}
}
