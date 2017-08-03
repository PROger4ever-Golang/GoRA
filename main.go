package main

import (
    "net/http"
    "./routes"
    "flag"
    "strconv"
    "fmt"
)

var host = flag.String("h", "127.0.0.1", "host/ip to bind")
var port = flag.Int("p", 53837, "port to bind")

func main() {
    flag.Parse()
    bind := *host + ":" + strconv.Itoa(*port)

    http.HandleFunc("/", routes.Router)
    fmt.Printf("Listening on %v...", bind)
    http.ListenAndServe(bind, nil)
}