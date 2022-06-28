package utils

import (
	"gotwitter/internal/tools/test"
	"testing"
)

func TestSortByMapKey(t *testing.T) {
	test_maps := [][]map[string]string{
		{
			{"e_key": "e_val"},
			{"d_key": "d_val"},
			{"c_key": "c_val"},
			{"b_key": "b_val"},
			{"a_key": "a_val"},
		},
		{
			{"a_key": "a_val"},
		},
		{},
	}

	expected_maps := [][]map[string]string{
		{
			{"a_key": "a_val"},
			{"b_key": "b_val"},
			{"c_key": "c_val"},
			{"d_key": "d_val"},
			{"e_key": "e_val"},
		},
		{
			{"a_key": "a_val"},
		},
		{},
	}

	for i := 0; i < len(test_maps); i++ {
		SortByMapKey(test_maps[i])

		pass := test.CompareObjectMaps(test_maps[i], expected_maps[i])

		if !pass {
			t.Errorf("\nexpected: %#v,\n actual: %#v\n", expected_maps[i], test_maps[i])
		}
	}

}
