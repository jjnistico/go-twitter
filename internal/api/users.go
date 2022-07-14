package api

import (
	"gotwitter/internal/endpoint"
	"gotwitter/internal/network"
	"gotwitter/internal/types"
)

type Users struct{}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
//   Parameters:
//     `ids` []string - array of user ids to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `tweet.fields` []string - array of tweet fields to include.
//     `user.fields`  []string - array of user fields to include.
//
func (*Users) Get(options types.GOTOptions) ([]types.UserData, []types.Error) {
	response, errors := network.Get[types.UsersResponse](endpoint.Users, options, []string{"ids"})
	if errors != nil {
		return []types.UserData{}, errors
	}
	return response.Data, response.Errors
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-user
//  Parameters:
//     `user_name` string - username to query **REQUIRED**
//     `expansions` []string - array of expansions (see link for available expansions)
//     `tweet.fields` []string - array of tweet fields to include.
//     `user.fields`  []string - array of user fields to include.
//
func (*Users) GetByUsername(user_name string, options types.GOTOptions) (types.UserData, []types.Error) {
	response, errors := network.Get[types.UserResponse](endpoint.UserByUsername(user_name), options, nil)
	if errors != nil {
		return types.UserData{}, errors
	}
	return response.Data, response.Errors
}
