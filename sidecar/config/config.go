package config

import "time"

// TofNConfig contains connection configuration of T-of-N Daemon
type TofNConfig struct {
	Host        string        `mapstructure:"tofnd-host"`
	Port        string        `mapstructure:"tofnd-port"`
	DialTimeout time.Duration `mapstrcture:"tofnd-dial-timeout"`
}

// SidecarConfig contains configuration for all Sidecar Program
type SidecarConfig struct {
	TofNConfig TofNConfig `mapstructure:"tofn"`
}

func DefaultTofNConfig() TofNConfig {
	return TofNConfig{
		Host:        "localhost",
		Port:        "50051",
		DialTimeout: 15 * time.Second,
	}
}

func DefaultSidecarConfig() SidecarConfig {
	return SidecarConfig{
		TofNConfig: DefaultTofNConfig(),
	}
}
