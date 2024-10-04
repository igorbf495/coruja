package main

import (
    "bytes"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
)

func makeRequest(method, url string, body string) {
    var reqBody *bytes.Buffer

    // Inicializando o corpo da requisição apenas se for necessário
    if body != "" {
        reqBody = bytes.NewBuffer([]byte(body))
    }

    // Debug: Imprimindo os parâmetros antes de fazer a requisição
    fmt.Printf("Método: %s, URL: %s, Corpo: %v\n", method, url, body)

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

    // Verificação de método válido
    validMethods := map[string]bool{"GET": true, "POST": true, "PUT": true, "DELETE": true}
    if _, ok := validMethods[*method]; !ok {
        fmt.Println("Método inválido. Use GET, POST, PUT ou DELETE.")
        return
    }

    makeRequest(*method, *url, *body)
}
