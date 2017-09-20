package main

import (
    "net/http"
    "flag"
    "strconv"
    "fmt"
    "log"
    "github.com/sevlyar/go-daemon"
    "./routes"
    "os"
)

var host = flag.String("h", "", "host/ip to bind")
var port = flag.Int("p", 53837, "port to bind")

//func main() {
//    flag.Parse()
//    bind := *host + ":" + strconv.Itoa(*port)
//
//    http.HandleFunc("/", routes.Router)
//    fmt.Printf("Listening on %v...\n", bind)
//
//    //syscall.Umask(0777)
//    //_, err := syscall.Setsid()
//    //if err != nil {
//    //    log.Fatalf("Error occured while Setsid: %v\n", err)
//    //}
//    //
//    //os.Stdin.Close()
//    //os.Stdout.Close()
//    //os.Stderr.Close()
//
//    err := http.ListenAndServe(bind, nil)
//    if err != nil {
//        //TODO: Output to log file
//        log.Fatalf("Error occured while ListenAndServe: %v\n", err)
//    }
//}

func main() {
    cntxt := &daemon.Context{
        PidFileName: "pid",
        PidFilePerm: 0644,
        LogFileName: "log",
        LogFilePerm: 0640,
        WorkDir:     "./",
        Umask:       027,
        Args:        os.Args,
    }

    d, err := cntxt.Reborn()
    if err != nil {
        log.Fatal("Unable to run: ", err)
    }
    if d != nil {
        return
    }
    defer cntxt.Release()

    log.Print("- - - - - - - - - - - - - - -")
    log.Print("daemon started")



    flag.Parse()
    bind := *host + ":" + strconv.Itoa(*port)

    http.HandleFunc("/", routes.Router)
    fmt.Printf("Listening on %v...\n", bind)

    //syscall.Umask(0777)
    //_, err := syscall.Setsid()
    //if err != nil {
    //    log.Fatalf("Error occured while Setsid: %v\n", err)
    //}
    //
    //os.Stdin.Close()
    //os.Stdout.Close()
    //os.Stderr.Close()

    err = http.ListenAndServe(bind, nil)
    if err != nil {
        //TODO: Output to log file
        log.Fatalf("Error occured while ListenAndServe: %v\n", err)
    }
}