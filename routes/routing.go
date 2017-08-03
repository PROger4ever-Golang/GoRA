package routes

import (
    "regexp"
    "net/http"
)

const (
    defaultMaxMemory = 32 << 20 // 32 MB
)

type RequestHandler func(wr http.ResponseWriter, req *http.Request)
type Route struct {
    Pattern *regexp.Regexp
    Handler RequestHandler
}

var routes = []*Route{
    {
        regexp.MustCompile("^/run"),
        runHandler,
    }, {
        regexp.MustCompile("^/pages/run"),
        runPageHandler,
    }, {
        regexp.MustCompile("^/pages/upload"),
        uploadPageHandler,
    }, {
        regexp.MustCompile("^/show"),
        showHandler,
    }, {
        regexp.MustCompile("^/download"),
        downloadHandler,
    }, {
        regexp.MustCompile("^/upload"),
        uploadHandler,
    },
}

func findRoute(path string) *Route {
    for _, r := range routes {
        if r.Pattern.MatchString(path) {
            return r
        }
    }
    return nil
}

func Router(wr http.ResponseWriter, req *http.Request) {
    route := findRoute(req.URL.Path)
    if route == nil {
        http.NotFound(wr, req)
        return
    }
    route.Handler(wr, req)
}
