package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// TODO: path~~ to libraries, convert path controller into interfaces (kvManager)

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

func exportKeyToPath(path, key, value string) error {
	targetPath := filepath.Join(path, key)
	return writeFile(targetPath, value)
}

// importKeyFromPath import keys from storage
func importKeysFromPath(path string) (map[string]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	kvStore := make(map[string]string)
	for _, file := range files {
		value, err := readFile(filepath.Join(path, file.Name()))
		if err != nil {
			continue
		}
		kvStore[file.Name()] = value
	}

	return kvStore, nil
}

// exportKeyToPath export given map to target folder for each files
func exportKeysToPath(keys map[string]string, path string) error {
	if dir, err := dirExists(path); err != nil || !dir {
		return fmt.Errorf("target folder is not exists")
	}

	for key, value := range keys {
		targetPath := filepath.Join(path, key)

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
