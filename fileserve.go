package main

import (
  "net/http"
  "flag"
  "strconv"
  "fmt"
)

func main() {
  var (
    address = flag.String("listen", "", "address to bind")
    port = flag.Int("port", 8080, "port to listen")
    path = flag.String("path", ".", "path to serve")
    host = *address + ":" + strconv.Itoa(*port)
  )
  flag.Parse()
  fmt.Printf("Serving %s at %s\n", *path, host)
  http.ListenAndServe(host, http.FileServer(http.Dir(*path)))
}
