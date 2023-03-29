package storage

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/assert"
)

func getMockedLocalFileMgr(validAddr types.ValAddress, keys map[string]string) Storage {
	return keyStorage{
		ValidatorAddress: validAddr,
		Keys:             keys,
		fileMgr:          mockLocalFileMgr{keys},
	}
}

func Test_SaveKey(t *testing.T) {
	keys := make(map[string]string)
	mgr := getMockedLocalFileMgr([]byte(""), keys)

	err := mgr.SaveKey("test_key", "test_value")
	assert.NilError(t, err)

	val := keys["test_key"]
	assert.Equal(t, val, "test_value")
}

func Test_GetKey(t *testing.T) {
	existKey := "exist_key"

	keys := make(map[string]string)
	keys[existKey] = "existing_key"
	mgr := getMockedLocalFileMgr([]byte(""), keys)

	result, err := mgr.GetKey(existKey)
	assert.NilError(t, err)
	assert.Equal(t, result, "existing_key")

	notExistKey := "not_exist_key"
	notExist, err := mgr.GetKey(notExistKey)
	assert.Equal(t, notExist, "")
	assert.Error(t, err, fmt.Sprintf("cannot found key: %s", notExistKey))
}
