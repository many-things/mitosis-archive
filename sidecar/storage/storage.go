package storage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/many-things/mitosis/sidecar/config"
)

type Storage interface {
	SaveKey(keyID string, value []byte) error
	GetKey(keyID string) ([]byte, error)
	GetValidator() types.ValAddress

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

func convertBase64ToByte(value string) ([]byte, error) {
	value = strings.Trim(value, " ")
	return base64.StdEncoding.DecodeString(value)
}

func convertByteToBase64(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

func newStorage(cfg *config.SidecarConfig) Storage {
	mgr := NewLocalFileMgr(cfg.Home + "/storage")
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
func (s keyStorage) SaveKey(keyID string, value []byte) error {
	conv := convertByteToBase64(value)
	s.Keys[keyID] = conv

	err := s.fileMgr.ExportKey(keyID, conv)
	if err != nil {
		return err
	}

	return nil
}

// GetKey returns target key in storage
func (s keyStorage) GetKey(keyID string) ([]byte, error) {
	if val, ok := s.Keys[keyID]; ok {
		fmt.Printf("Called GetKey: %s\n", val)
		return convertBase64ToByte(val)
	}

	return nil, fmt.Errorf("cannot found key: %s", keyID)
}

// IsTarget returns given address is matches with storage Validator
func (s keyStorage) IsTarget(validator types.ValAddress) bool {
	return bytes.Equal(validator, s.ValidatorAddress)
}

// GetValidator returns validator info
func (s keyStorage) GetValidator() types.ValAddress {
	return s.ValidatorAddress
}
