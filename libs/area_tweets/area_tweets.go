package area_tweets

import (
  "net/http"
  "os"
  "github.com/dghubble/go-twitter/twitter"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/clientcredentials"
)

func twitterClient() *twitter.Client {
  config := &clientcredentials.Config{
      ClientID: os.Getenv("TWITTER_CONSUMER_KEY"),
      ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
      TokenURL: "https://api.twitter.com/oauth2/token",
    }
  httpClient := config.Client(oauth2.NoContext)
  client := twitter.NewClient(httpClient)
  return client
}

func localizedSearch() (*twitter.Search, *http.Response, error) {
  config := &clientcredentials.Config{
      ClientID: os.Getenv("TWITTER_CONSUMER_KEY"),
      ClientSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
      TokenURL: "https://api.twitter.com/oauth2/token",
    }
  httpClient := config.Client(oauth2.NoContext)
  client := twitter.NewClient(httpClient)

  search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
      Query: "",
      Geocode: "43.6425662,-79.3958115,1km",
      ResultType: "recent",
      Count: 100,
  })
  return search, resp, err
}
