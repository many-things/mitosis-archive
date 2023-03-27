package storage

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func getMockedLocalFileMgr(validAddr string, keys map[string]string) Storage {
	return keyStorage{
		ValidatorAddress: validAddr,
		Keys:             keys,
		fileMgr:          mockLocalFileMgr{keys},
	}
}

func Test_SaveKey(t *testing.T) {
	keys := make(map[string]string)
	mgr := getMockedLocalFileMgr("", keys)

	err := mgr.SaveKey("test_key", "test_value")
	assert.NilError(t, err)

	val, _ := keys["test_key"]
	assert.Equal(t, val, "test_value")
}

func Test_GetKey(t *testing.T) {
	existKey := "exist_key"

	keys := make(map[string]string)
	keys[existKey] = "existing_key"
	mgr := getMockedLocalFileMgr("", keys)

	result, err := mgr.GetKey(existKey)
	assert.NilError(t, err)
	assert.Equal(t, result, "existing_key")

	notExistKey := "not_exist_key"
	notExist, err := mgr.GetKey(notExistKey)
	assert.Equal(t, notExist, "")
	assert.Error(t, err, fmt.Sprintf("cannot found key: %s", notExistKey))
}

func Test_IsTargetEvent(t *testing.T) {
	keys := make(map[string]string)
	existKey := "exist_key"
	validatorAddr := "validator1"

	keys[existKey] = "exist_value"
	mgr := getMockedLocalFileMgr(validatorAddr, keys)

	notMatchValid := mgr.IsTargetEvent("validator2", existKey)
	assert.Equal(t, notMatchValid, false)

	notMatchKey := mgr.IsTargetEvent(validatorAddr, "not_exist_key")
	assert.Equal(t, notMatchKey, false)

	matched := mgr.IsTargetEvent(validatorAddr, existKey)
	assert.Equal(t, matched, true)
}
