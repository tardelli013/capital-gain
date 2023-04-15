
# Desafio Ganho de Capital

[![Go](https://github.com/tardelli013/capital-gain/actions/workflows/go.yml/badge.svg?branch=develop)](https://github.com/tardelli013/capital-gain/actions/workflows/go.yml)

A aplicação foi desenvolvida utilizando apenas Golang 1.20 e seus recursos, nenhuma biblioteca externa foi adicionada para facilitar o build e a execução.

A modelagem da aplicação foi baseada no modelo hexagonal (_ports and adapters_) e as regras de negócio estão centralizadas no pacote `/core/usecases`

Todos os casos te teste da documentação foram implementados no arquivo `/core/usecases/capital-gain-test.go` e o pacote usecases possui 100% de cobertura de testes.

Para rodar a aplicação é necessário ter o ambiente de desenvolvimento com Golang > 1.19, aqui o [link](https://go.dev/doc/install) para instalação.

## Estrutura de diretórios e arquivos
```shell
  ├── adapter
      └── scan.go
      └── scan-adapter.go
  ├── core
      └── domain
          └── fee-result.go
          └── operation.go
      └── ports
          └── capital-gain-ports.go
      └── usecases
          └── calc
              └── buy.go
              └── sell.go
          └── mocks
              └── operations.go
          └── capital-gain.go
          └── capital-gain-test.go
  ├── helpers
      └── utils.go
  ├── go.mod
      └── go.sum
  ├── main.go
  ├── README.md
```

### Para rodar a aplicação, dentro do diretório raiz rodar o comando:
```go
go mod tidy
go run main.go
```
Ao rodar a aplicação toda a iteração poderá ser feita via console ou via API:

Documentação da API com Swagger disponível na rota: http://localhost:8080/docs/index.html

Abaixo um exemplo de iteração via console:
````go
> go run main.go

> Enter a operations json:
> // aqui abaixo devemos colar o json com as operações e finalizar com a tecla "Enter"
> [{"operation":"buy", "unit-cost":10.00, "quantity": 100},{"operation":"sell", "unit-cost":15.00, "quantity": 50},{"operation":"sell", "unit-cost":15.00, "quantity": 50}]
> // abaixo receberemos o resultado da operação
> Result:
> [{"tax":0},{"tax":0},{"tax":0}]

> Enter a operations json:
````

### Para rodar os testes e gerar um relatório de cobertura podemos usar o seguinte comando:
```go
go test ./... -coverprofile coverage.out && go tool cover -func coverage.out && go tool cover -html=coverage.out -o coverage.html
```
Os artefatos de cobertura serão gerados na raiz com os nomes `coverage.html` e `coverage.out`
