package gotwit

import (
	"strings"
	"testing"
)

func TestPercentEncode(t *testing.T) {
	testStr := "test_param=hello world&another_param=test?%*#@ no, test"
	expected := "test_param%3Dhello%20world%26another_param%3Dtest%3F%25%2A%23%40%20no%2C%20test"

	encodedStr := percentEncode(testStr)

	if strings.Compare(encodedStr, expected) != 0 {
		t.Errorf("\nexpected: %s, got: %s\n", expected, encodedStr)
	}
}

func TestEmptyPercentEncode(t *testing.T) {
	testStr := ""
	expected := ""

	encodedStr := percentEncode(testStr)

	if strings.Compare(encodedStr, expected) != 0 {
		t.Errorf("\nexpected: %s, got: %s\n", expected, encodedStr)
	}
}
