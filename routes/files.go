package routes

import (
    "net/http"
    "strings"
    "fmt"
    "github.com/PROger4ever/GoRA/request"
)

func showHandler(wr http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    path := req.FormValue("path")
    request.PrintFile(wr, req, path)
}

func downloadHandler(wr http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    path := req.FormValue("path")
    filename := path[strings.LastIndex(path, "/")+1:]
    wr.Header().Set("Content-Type", "application/octet-stream")
    wr.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%v\"", filename))
    request.PrintFile(wr, req, path)
}
