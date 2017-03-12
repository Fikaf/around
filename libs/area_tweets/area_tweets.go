package area_tweets
// "os"
// "fmt"
// "strconv"
import (
  "net/http"
  "strings"
  "github.com/dghubble/go-twitter/twitter"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/clientcredentials"
  "os"
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

func LocalizedSearch(lat string, long string) (*twitter.Search, *http.Response, error) {
  // fmt.Printf("%f %f", lat , long )
  // lat_str := strconv.FormatFloat(lat, 'G', -1, 64)
  // long_str := strconv.FormatFloat(long, 'G', -1, 64)
  geocode := []string{lat, long, "1km"}
  // fmt.Printf("%s", strings.Join(geocode, ","))
  client := twitterClient()
  search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
      Query: "",
      Geocode: strings.Join(geocode, ","),
      ResultType: "recent",
      Count: 100,
  })
  return search, resp, err
}
