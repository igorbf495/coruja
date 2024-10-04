package main

import (
    "bytes"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
)

func makeRequest(method, url, body string) {
    var reqBody *bytes.Buffer
    if body != "" {
        reqBody = bytes.NewBuffer([]byte(body))
    } else {
        reqBody = nil
    }

    req, err := http.NewRequest(method, url, reqBody)
    if err != nil {
        fmt.Println("Erro ao criar requisição:", err)
        return
    }

    // Adicionando cabeçalhos (opcional)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Erro na requisição:", err)
        return
    }
    defer resp.Body.Close()

    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Erro ao ler o corpo da resposta:", err)
        return
    }

    fmt.Println("Status:", resp.Status)
    fmt.Println("Resposta:", string(responseBody))
}

func main() {
    method := flag.String("method", "GET", "Método da requisição (GET, POST, PUT, DELETE)")
    url := flag.String("url", "https://jsonplaceholder.typicode.com/posts/1", "URL para a requisição")
    body := flag.String("body", "", "Corpo da requisição (para POST e PUT)")

    flag.Parse()

    makeRequest(*method, *url, *body)
}
