package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    for i := 1; i < len(os.Args); i++ {
        getURL := string(os.Args[i])
        resp, err := http.Get(getURL)
        if err != nil {
            fmt.Println(err)
            return
        }
        finalURL := resp.Request.URL.String()
        fmt.Printf(`{"url":"%s"}` + "\n", finalURL)
    }
}
