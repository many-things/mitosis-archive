package config

import "time"

// TofNConfig contains connection configuration of T-of-N Daemon
type TofNConfig struct {
	Host        string        `mapstructure:"tofnd-host"`
	Port        string        `mapstructure:"tofnd-port"`
	DialTimeout time.Duration `mapstrcture:"tofnd-dial-timeout"`
}

type TendermintConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Denom  string `mapstructure:"denom"`
	Prefix string `mapstructure:"prefix"`
}

// SidecarConfig contains configuration for all Sidecar Program
type SidecarConfig struct {
	TofNConfig TofNConfig       `mapstructure:"tofn"`
	MitoConfig TendermintConfig `mapstructure:"mito"`
}

func DefaultTofNConfig() TofNConfig {
	return TofNConfig{
		Host:        "localhost",
		Port:        "50051",
		DialTimeout: 15 * time.Second,
	}
}

func DefaultMitoConfig() TendermintConfig {
	return TendermintConfig{
		Host:   "localhost",
		Port:   9090,
		Denom:  "mito",
		Prefix: "mito",
	}
}

func DefaultSidecarConfig() SidecarConfig {
	return SidecarConfig{
		TofNConfig: DefaultTofNConfig(),
		MitoConfig: DefaultMitoConfig(),
	}
}
