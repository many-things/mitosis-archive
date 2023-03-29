package config

import "time"

// TofNConfig contains connection configuration of T-of-N Daemon
type TofNConfig struct {
	Host        string        `mapstructure:"tofnd-host"`
	Port        int           `mapstructure:"tofnd-port"`
	DialTimeout time.Duration `mapstrcture:"tofnd-dial-timeout"`
}

type TmConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Denom   string `mapstructure:"denom"`
	Prefix  string `mapstructure:"prefix"`
	ChainID string `mapstructure:"chain-id"`
	PrivKey string `mapstructure:"privkey"`
}

// SidecarConfig contains configuration for all Sidecar Program
type SidecarConfig struct {
	TofNConfig TofNConfig `mapstructure:"tofn"`
	MitoConfig TmConfig   `mapstructure:"mitosis"`
}

func DefaultTofNConfig() TofNConfig {
	return TofNConfig{
		Host:        "localhost",
		Port:        50051,
		DialTimeout: 15 * time.Second,
	}
}

func DefaultMitoConfig() TmConfig {
	return TmConfig{
		Host:    "localhost",
		Port:    9090,
		Denom:   "mitosis",
		Prefix:  "mitosis",
		ChainID: "",
		PrivKey: "",
	}
}

func DefaultSidecarConfig() SidecarConfig {
	return SidecarConfig{
		TofNConfig: DefaultTofNConfig(),
		MitoConfig: DefaultMitoConfig(),
	}
}
