package routes

//func publicHandler(wr http.ResponseWriter, req *http.Request) {
//    publicPath := req.URL.Path[len("/"):]
//    body, err := ioutil.ReadFile(publicPath)
//    if err != nil {
//        errText := fmt.Sprintf("500 Internal Server Error while reading page: %v", err)
//        http.Error(wr, errText, http.StatusInternalServerError)
//        return
//    }
//    wr.Header().Set("Content-Type", "text/html; charset=utf-8")
//    wr.Write(body)
//}
