package routes

import (
    "time"
    "os/exec"
    "bytes"
    "fmt"
    "net/http"
    "math"
    "strconv"
    "golang.org/x/net/context"
    "github.com/PROger4ever/GoRA/request"
    "strings"
)

func runHandler(wr http.ResponseWriter, req *http.Request) {
    var err error

    err = req.ParseForm()
    if err != nil {
        fmt.Fprintf(wr, "Error while ParseForm: %v", err)
        return
    }

    if strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") {
        err = req.ParseMultipartForm(defaultMaxMemory)
        if err != nil {
            fmt.Fprintf(wr, "Error while ParseMultipartForm: %v", err)
            return
        }
    }

    //log.Printf("%q", req.Form)
    //log.Printf("%q", req.MultipartForm)
    //log.Printf("%q", parseParam(req,"cmd"))
    //log.Printf("%q", parseParam(req,"params[]"))
    //log.Printf("%q", parseParam(req,"timeout"))
    //log.Printf("%q", parseParam(req,"nowait"))

    //TODO: check for parameters exist!
    //TODO: "if method == 'GET' {}"

    cmdLine := request.ParseParam(req, "cmd")[0]
    params := request.ParseParam(req, "params[]")
    timeout := request.ParseParam(req, "timeout")[0]
    nowait := request.ParseParam(req, "nowait")
    if len(cmdLine) == 0 {
        fmt.Fprint(wr, "No cmd parameter\n")
        return
    }
    fmt.Fprintf(wr, "The cmd is: %v\n", cmdLine)
    fmt.Fprintf(wr, "The params is: %q\n", params)
    fmt.Fprintf(wr, "The timeout is: %q\n", timeout)
    fmt.Fprintf(wr, "The nowait is: %q\n\n", nowait)

    t := math.MaxInt64
    if len(timeout) > 0 {
        t1, err := strconv.Atoi(timeout)
        t = t1

        if err != nil {
            fmt.Fprintf(wr, "Error while Atoi(timeout): %v", err)
            return
        }
    }

    if t < 0 {
        fmt.Fprint(wr, "Parameter 'timeout': incorrect value (must be correct time in ms)")
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Millisecond)
    defer cancel()

    cmd := exec.CommandContext(ctx, cmdLine, params...)

    var b bytes.Buffer
    cmd.Stdout = &b
    cmd.Stderr = &b
    err = cmd.Start()
    if err != nil {
        fmt.Fprintf(wr, "Error while starting the command: %v\n", err)
        return
    }
    fmt.Fprint(wr, "Started.\n")

    if len(nowait) == 0 || nowait[0] == "" {
        err := cmd.Wait()

        deadline, ok := ctx.Deadline()
        if !ok {
            fmt.Fprint(wr, "Error while getting deadline\n")
            return
        }
        if !deadline.After(time.Now()) {
            fmt.Fprintf(wr, "The command timed out. Killed. %v", err)
            return
        }

        if err != nil {
            fmt.Fprintf(wr, "Error while getting combined output: %v\n", err)
        }
        combinedString := string(b.Bytes()[:])
        fmt.Fprintf(wr, "Output:\n%v", combinedString)
        fmt.Fprintf(wr, "Exit code: %v\n", cmd.ProcessState.Sys())
    }
}
