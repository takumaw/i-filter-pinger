package main

import (
  "flag"
  "net/http"
  "net/url"
  "os"

  "github.com/PuerkitoBio/goquery"
)

func main() {
  // Parse arguments
  flag.Parse()
  args := flag.Args()

  if (len(args) == 0) {
    os.Exit(1)
  }

  targetUrl := args[0]

  // Load url
  doc, err := goquery.NewDocument(targetUrl)

  if err != nil {
    os.Exit(1)
  }

  // Build form
  values := url.Values{}
  doc.Find("form[name=frm_info] input[type=hidden]").Each(func(_ int, s *goquery.Selection) {
    name, _ := s.Attr("name")
    value, _ := s.Attr("value")
    values.Add(name, value)
  })

  // Ping web site
  resp, err := http.Get("https://webrep.msg/envelope.cgi" + "?" + values.Encode())
  defer resp.Body.Close() // FIXME defer wont work with os.Exit(code).

  if err != nil {
    os.Exit(1)
  } else if resp.StatusCode != 200 {
    os.Exit(1)
  } else {
    os.Exit(0)
  }
}
