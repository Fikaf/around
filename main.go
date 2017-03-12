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

func get_tweets(w http.ResponseWriter, r *http.Request) {
  lat := r.URL.Query().Get("lat")
  long := r.URL.Query().Get("long")
  search, resp, err := area_tweets.LocalizedSearch(lat, long)

  fmt.Fprintln(w, err)
  fmt.Fprintln(w, search)
  fmt.Fprintln(w, resp)
}

func home_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello")
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/scan", get_tweets)
  http.HandleFunc("/", home_page)
  // http.Handle("/", http.FileServer(http.Dir("./static")))
  // http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
