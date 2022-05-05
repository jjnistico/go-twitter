# GO-TWITTER

### _This is a WIP_

Go-Twitter is a go rest server that can be used to programatically read, write, delete tweets in addition to other functionality like liking/retweeting

## Installation

Navigate to the root directory of this project and run

```
go run .
```

A client for the front-end is on the way and will be packaged with the server in a docker container for ease of use.

## Current Endpoints

`/callback`: used for the oauth protocol implementation. Twitter calls this endpoint with access token query params

`/authenticate`: implements the 3-step oauth protocol and stores access tokens in environment variables

`/tweets/by/id`: returns tweets for a comma-separated list of tweet ids, passed as query params `ids=`

`/users/by/username`: returns user data for a comma-separated list of usernames, passed as query params `usernames=`
