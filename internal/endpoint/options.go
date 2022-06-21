package endpoint

// returns query parameters for endpoints. Parameters followed by '*' are required
func GetEndpointOptions(endpoint string) []string {
	switch endpoint {
	case Tweets:
		return []string{"id*", "expansions", "media.fields", "place.fields", "poll.fields", "tweet.fields", "user.fields"}
	case Users:
		return []string{"id*", "expansions", "tweet.fields", "user.fields"}
	default:
		return []string{}
	}
}