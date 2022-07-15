package utils

import (
	"gotwitter/internal/utils/test"
	"testing"
)

func TestSortByMapKey(t *testing.T) {
	testMaps := [][]map[string]string{
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

	expectedMaps := [][]map[string]string{
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

	for i := 0; i < len(testMaps); i++ {
		SortByMapKey(testMaps[i])

		pass := test.CompareObjectMaps(testMaps[i], expectedMaps[i])

		if !pass {
			t.Errorf("\nexpected: %#v,\n actual: %#v\n", expectedMaps[i], testMaps[i])
		}
	}

}
