package main
import (
  "fmt"
  "regexp"
  "net/http"
  "io/ioutil"
)
func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    parsedURL := regexp.MustCompile(`^\/(.+:\/)?\/*(.*)`).FindStringSubmatch(r.URL.Path)
    if (parsedURL[1] != "") {
      parsedURL[2] = parsedURL[1] + "/" + parsedURL[2]
    }
    resp, err := http.Get(parsedURL[2])
    if err != nil {
      fmt.Fprint(w, err)
    } else {
      defer resp.Body.Close()
      responseData, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Fprint(w, err)
      } else {
        fmt.Fprint(w, string(responseData))
      }
    }
  })
  http.ListenAndServe(":8000", nil)
}
