package utils

func BuildCurrentSnapshot(root string) (map[string]string, error) {
	snapshot := make(map[string]string)

	files, err := ScanFiles(root)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		hash, err := Hashing(file)
		if err != nil {
			return nil, err
		}
		snapshot[file] = hash
	}

	return snapshot, nil
}
