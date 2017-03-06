package main
import (
  "log"
  "fmt"
  "net/http"
  "os"
  "github.com/dghubble/go-twitter/twitter"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/clientcredentials"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
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
  fmt.Fprintln(w, err)
  fmt.Fprintln(w, search)
  fmt.Fprintln(w, resp)
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
