# GO-TWITTER

### _This is a WIP_

Go-Twitter is a go client that can be used to programatically read, write, delete tweets in addition to other functionality like liking/retweeting

## Installation

Coming Soon: After version 1.0, this will be packaged into a go module that can be installed with `go install`

## Credentials

In order to use many of the features of this application, you must provide api keys for your twitter account. The easiest way to do this is to navigate to `https://developer.twitter.com/` in your browser and follow these steps:

1) Click `developer portal` in the navigation bar at the top right.

2) Create a new app under `Projects & Apps` in the left hand menu. You can name the app whatever you would like (and is not already in use on twitter).

3) Make sure to copy the API key and Secret generated by twitter.

4) Create a file named `.env` in the root directory of this application following the example of `.env.sample`. Fill in the API key and Secret generated in the previous step in the `.env` file.

5) The API Key and Secret are used to identify your app. The quickest and easiest way to authenticate requests to the twitter api from within this client would be to generate an Access Token and Secret within the twitter developer portal. Under your application name, click `Keys and tokens` then click `Generate` next to `Access Token and Secret`. Copy these to your .env file under `OAUTH_TOKEN` and `OAUTH_SECRET`. You are now ready to make requests to the twitter api!