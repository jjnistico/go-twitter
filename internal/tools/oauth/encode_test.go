package oauth

import (
	"strings"
	"testing"
)

func TestPercentEncode(t *testing.T) {
	test_str := "test_param=hello world&another_param=test?%*#@ no, test"
	expected := "test_param%3Dhello%20world%26another_param%3Dtest%3F%25%2A%23%40%20no%2C%20test"

	encoded_str := PercentEncode(test_str)

	if strings.Compare(encoded_str, expected) != 0 {
		t.Errorf("\nexpected: %s, got: %s\n", expected, encoded_str)
	}
}

func TestEmptyPercentEncode(t *testing.T) {
	test_str := ""
	expected := ""

	encoded_str := PercentEncode(test_str)

	if strings.Compare(encoded_str, expected) != 0 {
		t.Errorf("\nexpected: %s, got: %s\n", expected, encoded_str)
	}
}
