package storage

type mockLocalFileMgr struct {
	mockedKeyMap map[string]string
}

func (m mockLocalFileMgr) ExportKey(key, value string) error {
	return nil
}

func (m mockLocalFileMgr) ExportKeyMap(keys map[string]string) error {
	return nil
}

func (m mockLocalFileMgr) ImportKeyMap() (map[string]string, error) {
	return m.mockedKeyMap, nil
}
