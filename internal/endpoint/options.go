package endpoint

// returns query parameters for endpoints. Parameters followed by '*' are required
func GetEndpointOptions(endpoint string) []string {
	switch endpoint {
	case "/api/tweets":
		return []string{"ids*", "expansions", "media.fields", "place.fields", "poll.fields", "tweet.fields", "user.fields"}
	case "/api/users":
		return []string{"ids*", "expansions", "tweet.fields", "user.fields"}
	case "/api/timeline_tweets":
		return []string{
			"id*",
			"end_time",
			"exclude",
			"expansions",
			"max_results",
			"media.fields",
			"pagination_token",
			"place.fields",
			"poll.fields",
			"since_id",
			"start_time",
			"tweet.fields",
			"until_id",
			"user.fields",
		}
	default:
		return []string{}
	}
}
