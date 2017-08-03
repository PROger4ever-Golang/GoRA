package request

import (
    "net/http"
    "os"
    "fmt"
    "io"
)

func ParseParam(req *http.Request, param string) []string {
    res := req.Form[param]
    if len(res) == 0 {
        res = req.MultipartForm.Value[param]
    }
    return res
}

func PrintFile(wr http.ResponseWriter, req *http.Request, path string) {
    var err error
    file, err := os.Open(path)

    if err != nil {
        errText := fmt.Sprintf("500 Internal Server Error while opening file: %v", err)
        http.Error(wr, errText, http.StatusInternalServerError)
        return
    }
    defer file.Close()

    _, err = io.Copy(wr, file)
    if err != nil {
        errText := fmt.Sprintf("500 Internal Server Error while reading file: %v", err)
        http.Error(wr, errText, http.StatusInternalServerError)
        return
    }
}
