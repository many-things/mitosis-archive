package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type NodeConfig struct {
	Validator string `mapstructure:"validator" yaml:"validator"`
	Host      string `mapstructure:"host" yaml:"host"`
}

// TofNConfig contains connection configuration of T-of-N Daemon
type TofNConfig struct {
	Host        string        `mapstructure:"tofnd-host" yaml:"host"`
	Port        int           `mapstructure:"tofnd-port" yaml:"port"`
	DialTimeout time.Duration `mapstructure:"tofnd-dial-timeout" yaml:"dial-timeout"`
	Validator   string        `mapstructure:"validator" yaml:"validator"`
	Nodes       []NodeConfig  `mapstructure:"nodes" yaml:"nodes"`
}

type TmConfig struct {
	Host            string `mapstructure:"host" yaml:"host"`
	Port            int    `mapstructure:"port" yaml:"port"`
	Denom           string `mapstructure:"denom" yaml:"denom"`
	Prefix          string `mapstructure:"prefix" yaml:"prefix"`
	ValidatorPrefix string `mapstructure:"validator-prefix" yaml:"validator-prefix"`
	ChainID         string `mapstructure:"chain-id" yaml:"chain-id"`
	PrivKey         string `mapstructure:"privkey" yaml:"priv-key"`
}

// SidecarConfig contains configuration for all Sidecar Program
type SidecarConfig struct {
	Home       string     `mapstructure:"home" yaml:"home"`
	TofNConfig TofNConfig `mapstructure:"tofn" yaml:"tofn"`
	MitoConfig TmConfig   `mapstructure:"mito" yaml:"mitosis"`
}

func DefaultTofNConfig() TofNConfig {
	return TofNConfig{
		Host:        "localhost",
		Port:        50051,
		DialTimeout: 15 * time.Second,
		Validator:   "",
		Nodes:       []NodeConfig{},
	}
}

func DefaultMitoConfig() TmConfig {
	return TmConfig{
		Host:    "localhost",
		Port:    9090,
		Denom:   "mito",
		Prefix:  "mito",
		ChainID: "",
		PrivKey: "",
	}
}

func DefaultSidecarConfig() SidecarConfig {
	homeEnvDir, _ := os.LookupEnv("HOME")
	return SidecarConfig{
		Home:       homeEnvDir + "/.sidecar",
		TofNConfig: DefaultTofNConfig(),
		MitoConfig: DefaultMitoConfig(),
	}
}

func GetConfigFromFile(file string) (SidecarConfig, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return SidecarConfig{}, err
	}

	cfg := &SidecarConfig{}
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return SidecarConfig{}, err
	}

	return *cfg, nil
}
