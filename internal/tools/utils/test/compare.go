package test

func CompareObjectMaps(map_a []map[string]string, map_b []map[string]string) bool {
	if len(map_a) != len(map_b) {
		return false
	}

	for i := 0; i < len(map_a); i++ {
		var key_a, val_a, key_b, val_b string

		for k, v := range map_a[i] {
			key_a = k
			val_a = v
		}

		for k, v := range map_b[i] {
			key_b = k
			val_b = v
		}

		if key_a != key_b || val_a != val_b {
			return false
		}
	}

	return true
}
