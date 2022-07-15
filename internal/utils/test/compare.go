package test

func CompareObjectMaps(mapA []map[string]string, mapB []map[string]string) bool {
	if len(mapA) != len(mapB) {
		return false
	}

	for i := 0; i < len(mapA); i++ {
		var keyA, valA, keyB, valB string

		for k, v := range mapA[i] {
			keyA = k
			valA = v
		}

		for k, v := range mapB[i] {
			keyB = k
			valB = v
		}

		if keyA != keyB || valA != valB {
			return false
		}
	}

	return true
}
