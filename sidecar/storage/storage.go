package storage

import (
	"fmt"
	"github.com/many-things/mitosis/sidecar/config"
	"sync"
)

type Storage interface {
	SaveKey(keyId string, value string) error
	GetKey(keyId string) (string, error)

	IsTargetEvent(validator string, keyId string) bool
}

type keyStorage struct {
	ValidatorAddress string
	Keys             map[string]string
	fileMgr          LocalFileMgr
}

var (
	storage    Storage
	initConfig sync.Once
)

func newConfig(cfg *config.SidecarConfig) Storage {
	mgr := NewLocalFileMgr("")
	savedKey, err := mgr.ImportKeyMap()
	if err != nil {
		panic(err) // if this method is not executed, must not be runnable for sidecar logics
	}

	// TODO: add adoptable variables
	return keyStorage{
		ValidatorAddress: "",
		Keys:             savedKey,
		fileMgr:          mgr,
	}
}

// GetStorage returns global singleton storage variable
func GetStorage(cfg *config.SidecarConfig) Storage {
	initConfig.Do(func() {
		storage = newConfig(cfg)
	})

	return storage
}

// SaveKey save key, value on map and storage
func (s keyStorage) SaveKey(keyId, value string) error {
	s.Keys[keyId] = value

	err := s.fileMgr.ExportKey(keyId, value)
	if err != nil {
		return err
	}

	return nil
}

// GetKey returns target key in storage
func (s keyStorage) GetKey(keyId string) (string, error) {
	if val, ok := s.Keys[keyId]; ok {
		return val, nil
	}

	return "", fmt.Errorf("cannot found key: %s", keyId)
}

// IsTargetEvent returns TF variable is given event info is valid for the validator {
func (s keyStorage) IsTargetEvent(validator, keyId string) bool {
	if validator != s.ValidatorAddress {
		return false
	}

	if _, ok := s.Keys[keyId]; ok {
		return true
	}

	return false
}
