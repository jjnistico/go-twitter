package auth

import (
	"regexp"
	"testing"
)

func TestPercentEncode(t *testing.T) {
	testStr := "test_param=hello world&another_param=test?%*#@ no, test"
	expected := "test_param%3Dhello%20world%26another_param%3Dtest%3F%25%2A%23%40%20no%2C%20test"

	encodedStr := percentEncode(testStr)

	compareResults(t, encodedStr, expected)
}

func TestEmptyPercentEncode(t *testing.T) {
	testStr := ""
	expected := ""

	encodedStr := percentEncode(testStr)

	compareResults(t, encodedStr, expected)
}

func TestGenerateNonce(t *testing.T) {
	nonceEmpty := generateNonce(0)

	if len(nonceEmpty) != 0 {
		t.Errorf("expected nonce of length 0, got length of %d", len(nonceEmpty))
	}

	nonce42 := generateNonce(42)

	if len(nonce42) != 42 {
		t.Errorf("expected nonce of length 42, got length of %d", len(nonce42))
	}

	r, _ := regexp.Compile("[a-zA-Z]")

	if !r.MatchString(nonce42) {
		t.Errorf("nonce %s is not a valid nonce", nonce42)
	}
}
