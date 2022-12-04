package main

import (
	"desafio-nu/adapter"
	"desafio-nu/core/usecases"
	"desafio-nu/helpers"
)

func main() {
	scan := adapter.NewScan()
	useCase := usecases.NewOperationUseCase()

	operationsFromScan, err := scan.ScanOperations()
	if err != nil {
		panic(err)
	}

	feeResults, err := useCase.CalcCapitalGain(operationsFromScan)
	if err != nil {
		panic(err)
	}

	helpers.PrettyPrint(&feeResults)
}
