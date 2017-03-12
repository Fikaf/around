package main
import (
  "log"
  "fmt"
  "net/http"
  "os"
  "around/libs/area_tweets"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
  search, resp, err := area_tweets.LocalizedSearch()

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
