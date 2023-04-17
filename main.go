package main

import (
	"desafio-nu/adapter"
	"desafio-nu/core/ports"
	"desafio-nu/core/usecases"
	_ "desafio-nu/docs"
	"desafio-nu/helpers"
	"fmt"
	"time"
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

	// up gin server in a new Goroutine
	go runWithGinHttp(useCase)

	// this delay is for gin server up
	time.Sleep(2 * time.Second)
	runWithTerminalScan(useCase)
}

func runWithGinHttp(useCase ports.OperationUseCase) {
	g := adapter.SetupRouter(useCase)
	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
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

	fmt.Println("\n[SCAN] Result: ")
	helpers.PrettyPrint(&feeResults)
	fmt.Print("\n")

	runWithTerminalScan(useCase)
}
