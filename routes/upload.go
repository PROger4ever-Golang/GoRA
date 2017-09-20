package routes

import (
    "fmt"
    "os"
    "io"
    "net/http"
    "github.com/PROger4ever/GoRA/request"
)

func uploadHandler(wr http.ResponseWriter, req *http.Request) {
    var err error

    err = req.ParseMultipartForm(defaultMaxMemory)
    if err != nil {
        fmt.Fprintf(wr, "Error while ParseMultipartForm: %v", err)
        return
    }

    filepathSlice := request.ParseParam(req, "filepath")
    if len(filepathSlice) == 0 || filepathSlice[0] == "" {
        fmt.Fprint(wr, "Error: local filepath isn't specified")
        return
    }
    filepath := filepathSlice[0]

    uploadFile, _ /*handler*/ , err := req.FormFile("uploadFile")
    if err != nil {
        fmt.Fprintf(wr, "Error while getting uploadFile from MultipartForm: %v", err)
        return
    }
    defer uploadFile.Close()

    //fmt.Fprintf(wr, "%v", handler.Header)
    localFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Fprintf(wr, "Error while opening local file: %v", err)
        return
    }
    defer localFile.Close()

    n, err := io.Copy(localFile, uploadFile)
    if err != nil {
        fmt.Fprintf(wr, "Error while copying to local file: %v %v", err, n)
        return
    }
    fmt.Fprintf(wr, "File successfully uploaded to %v", filepath)
}
