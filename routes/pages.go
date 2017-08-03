package routes

import (
    "net/http"
    "fmt"
    "../package_pages"
)

func runPageHandler(wr http.ResponseWriter, req *http.Request) {
    wr.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(wr, package_pages.RunPage)
}

func uploadPageHandler(wr http.ResponseWriter, req *http.Request) {
    wr.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(wr, package_pages.UploadPage)
}
