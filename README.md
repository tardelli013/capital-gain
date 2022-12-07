
A aplicação foi desenvolvida com Golang 1.19

Para rodar a aplicação, dentro do diretório raiz rodar o comando:  go run main.go

go test ./... -coverprofile coverage.out && go tool cover -func coverage.out && go tool cover -html=coverage.out -o coverage.html
