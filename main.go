package main

import (
	"desafio-nu/adapter"
	"desafio-nu/core/ports"
	"desafio-nu/core/usecases"
	_ "desafio-nu/docs"
	"desafio-nu/helpers"
	"fmt"
)

// @title           Desafio Nu
// @version         1.0
// @description     Desafio de calculo de ganho de capital.
// @termsOfService  https://google.com
// @contact.name   Tardelli Moura
// @contact.url    https://google.com
// @contact.email  tardelli.m@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	useCase := usecases.NewOperationUseCase()
	//runWithTerminalScan(useCase)

	adapter.Router(useCase)

}

func runWithTerminalScan(useCase ports.OperationUseCase) {
	scan := adapter.NewScan()

	operationsFromScan, err := scan.ScanOperations()
	if err != nil {
		panic(err)
	}

	feeResults, err := useCase.CalcCapitalGain(operationsFromScan)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nResult: ")
	helpers.PrettyPrint(&feeResults)
}
