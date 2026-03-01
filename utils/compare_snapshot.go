package utils

func CompareSnapshot(oldSnapshot, newSnapshot map[string]string) (added, modified, deleted []string) {
	//added and modified files
	for file, hash := range newSnapshot {
		if oldHash, exists := oldSnapshot[file]; !exists {
			added = append(added, file)
		} else if oldHash != hash {
			modified = append(modified, file)
		}
	}

	//deleted files
	for file := range oldSnapshot {
		if _, exists := newSnapshot[file]; !exists {
			deleted = append(deleted, file)
		}
	}

	return added, modified, deleted
}
