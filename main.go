package main

import (
    "flag"
    "fmt"
    "net/http"
    "io/ioutil"
)

func getRequest(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erro na requisição:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Erro ao ler o corpo da resposta:", err)
        return
    }
    fmt.Println("Status:", resp.Status)
    fmt.Println("Resposta:", string(body))
}

func main() {
    url := flag.String("url", "https://jsonplaceholder.typicode.com/posts/1", "URL para requisição GET")
    flag.Parse()
    
    getRequest(*url)
}
