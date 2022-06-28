package endpoint

import (
	"reflect"
	"testing"
)

func TestGetEndpointOptions(t *testing.T) {
	endpoints := []string{Tweets, Users, "invalid"}

	expected_options := [][]string{
		{"ids*", "expansions", "media.fields", "place.fields", "poll.fields", "tweet.fields", "user.fields"},
		{"ids*", "expansions", "tweet.fields", "user.fields"},
		{},
	}

	for idx, endpoint := range endpoints {
		options := GetEndpointOptions(endpoint)

		if !reflect.DeepEqual(options, expected_options[idx]) {
			t.Errorf("\nexpected: %#v, got: %#v\n", expected_options[idx], options)
		}
	}
}
