package storage

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/sidecar/config"
)

type Storage interface {
	SaveKey(keyID string, value string) error
	GetKey(keyID string) (string, error)
	GetValidator() types.ValAddress

	IsTargetEvent(validator types.ValAddress, keyID string) bool
	IsTarget(validator types.ValAddress) bool
}

type keyStorage struct {
	ValidatorAddress []byte
	Keys             map[string]string
	fileMgr          LocalFileMgr
}

var (
	storage    Storage
	initConfig sync.Once
)

func newStorage(_ *config.SidecarConfig) Storage {
	mgr := NewLocalFileMgr("")
	savedKey, err := mgr.ImportKeyMap()
	if err != nil {
		panic(err) // if this method is not executed, must not be runnable for sidecar logics
	}

	// TODO: add adoptable variables
	return keyStorage{
		ValidatorAddress: []byte(""),
		Keys:             savedKey,
		fileMgr:          mgr,
	}
}

// GetStorage returns global singleton storage variable
func GetStorage(cfg *config.SidecarConfig) Storage {
	initConfig.Do(func() {
		storage = newStorage(cfg)
	})

	return storage
}

// SaveKey save key, value on map and storage
func (s keyStorage) SaveKey(keyID, value string) error {
	s.Keys[keyID] = value

	err := s.fileMgr.ExportKey(keyID, value)
	if err != nil {
		return err
	}

	return nil
}

// GetKey returns target key in storage
func (s keyStorage) GetKey(keyID string) (string, error) {
	if val, ok := s.Keys[keyID]; ok {
		return val, nil
	}

	return "", fmt.Errorf("cannot found key: %s", keyID)
}

// IsTargetEvent returns TF variable is given event info is valid for the validator {
func (s keyStorage) IsTargetEvent(validator types.ValAddress, keyID string) bool {
	if bytes.Equal(validator, s.ValidatorAddress) {
		return false
	}

	if _, ok := s.Keys[keyID]; ok {
		return true
	}

	return false
}

// IsTarget returns given address is matches with storage Validator
func (s keyStorage) IsTarget(validator types.ValAddress) bool {
	return bytes.Equal(validator, s.ValidatorAddress)
}

// GetValidator returns validator info
func (s keyStorage) GetValidator() types.ValAddress {
	return s.ValidatorAddress
}
