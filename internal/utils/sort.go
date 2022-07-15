package utils

import "sort"

func SortByMapKey(mapToSort []map[string]string) {
	sort.Slice(mapToSort, func(i, j int) bool {
		var aKey, bKey string
		for key := range mapToSort[i] {
			aKey = key
		}
		for key := range mapToSort[j] {
			bKey = key
		}
		return aKey < bKey
	})
}
