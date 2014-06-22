package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func main() {
    for i := 1; i < len(exec.Args); i++ {
        getURL := string(exec.Args[i])
        resp, err := http.Get(getURL)
        if err != nil {
            fmt.Println(err)
            return
        }
        finalURL := resp.Request.URL.String()
        fmt.Println(finalURL)
    }
}
