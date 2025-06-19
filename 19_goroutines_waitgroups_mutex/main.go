package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex 

func main() {
    websiteList := []string{
        "https://youtube.com",
        "https://go.dev",
        "https://google.dev",
        "https://github.com",
    
    }

    for _, web:= range websiteList{
        go getStatusCode(web)
        wg.Add(1)
    }

    wg.Wait()
    fmt.Println(signals)
}
    
func greeter(s string) {
    //for i := 0; i < 5; i++ {
    //	time.Sleep(3 * time.Microsecond)
    //	fmt.Println(s)
    //}
}

func getStatusCode(endpoint string) {
    defer wg.Done()

    res, err := http.Get(endpoint)

    if err != nil {

        fmt.Println("OOPS problem at endpoint")

    } else {

        mut.Lock()
        signals = append(signals,endpoint)
        mut.Unlock()
        fmt.Printf("%d is the status code for %s \n", res.StatusCode, endpoint)
    }
}