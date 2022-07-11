package utils

import "sort"

func SortByMapKey(map_to_sort []map[string]string) {
	sort.Slice(map_to_sort, func(i, j int) bool {
		var a_key, b_key string
		for key := range map_to_sort[i] {
			a_key = key
		}
		for key := range map_to_sort[j] {
			b_key = key
		}
		return a_key < b_key
	})
}
