#!/bin/bash

# Função para exibir a ajuda
function show_help() {
    echo "Uso: ./coruja.sh -method=<método> -url=<url> -body=<corpo>"
    echo "Método: GET, POST, PUT, DELETE"
    echo "Corpo: Opcional, usado apenas para POST e PUT"
    exit 1
}

# Verificando se pelo menos um argumento foi passado
if [ "$#" -eq 0 ]; then
    show_help
fi

# Inicializando variáveis
METHOD=""
URL=""
BODY=""

# Analisando os argumentos
for arg in "$@"; do
    case $arg in
        -method=*)
            METHOD="${arg#*=}"
            shift
            ;;
        -url=*)
            URL="${arg#*=}"
            shift
            ;;
        -body=*)
            BODY="${arg#*=}"
            shift
            ;;
        *)
            echo "Argumento desconhecido: $arg"
            show_help
            ;;
    esac
done

# Verificando se o método e a URL foram fornecidos
if [ -z "$METHOD" ] || [ -z "$URL" ]; then
    echo "Método e URL são obrigatórios."
    show_help
fi

# Verificando se o método é válido
if [[ "$METHOD" != "GET" && "$METHOD" != "POST" && "$METHOD" != "PUT" && "$METHOD" != "DELETE" ]]; then
    echo "Método inválido. Use GET, POST, PUT ou DELETE."
    exit 1
fi

# Executando a requisição usando curl
if [ "$METHOD" == "GET" ]; then
    response=$(curl -s -w "%{http_code}" -o /dev/null "$URL")
else
    response=$(curl -s -w "%{http_code}" -o /dev/null -X "$METHOD" -H "Content-Type: application/json" -d "$BODY" "$URL")
fi

# Exibindo o resultado
echo "Resposta HTTP: $response"
