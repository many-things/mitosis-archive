package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// TODO: path~~ to libraries, convert path controller into interfaces (kvManager)

type LocalFileMgr interface {
	ImportKeyMap() (map[string]string, error)
	ExportKeyMap(keys map[string]string) error
	ExportKey(key, value string) error
}

type localFileMgr struct {
	basePath string
}

func NewLocalFileMgr(basePath string) LocalFileMgr {
	return localFileMgr{
		basePath: basePath,
	}
}

// pathExists check given absolute path exists.
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

// dirExists chechk given absolute path is dir
func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return info.IsDir(), nil
}

// readFile returns value of file
func readFile(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", nil
	}

	return string(f), nil
}

// writeFile write file with value on given path
func writeFile(path, value string) error {
	return os.WriteFile(path, []byte(value), 0644)
}

func (m localFileMgr) ExportKey(key, value string) error {
	targetPath := filepath.Join(m.basePath, key)
	return writeFile(targetPath, value)
}

// ImportKeyMap import keys from storage
func (m localFileMgr) ImportKeyMap() (map[string]string, error) {
	files, err := os.ReadDir(m.basePath)
	if err != nil {
		return nil, err
	}

	kvStore := make(map[string]string)
	for _, file := range files {
		value, err := readFile(filepath.Join(m.basePath, file.Name()))
		if err != nil {
			continue
		}
		kvStore[file.Name()] = value
	}

	return kvStore, nil
}

// ExportKeyMap export given map to target folder for each files
func (m localFileMgr) ExportKeyMap(keys map[string]string) error {
	if dir, err := dirExists(m.basePath); err != nil || !dir {
		return fmt.Errorf("target folder is not exists")
	}

	for key, value := range keys {
		targetPath := filepath.Join(m.basePath, key)

		if exist, err := pathExists(targetPath); err == nil || exist {
			continue
		}

		err := writeFile(targetPath, value)
		if err != nil {
			return err
		}
	}

	return nil
}
