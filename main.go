package main

import (
	"desafio-nu/adapter"
	"desafio-nu/core/usecases"
	"encoding/json"
	"fmt"
)

func main() {
	scan := adapter.NewScan()
	useCase := usecases.NewOperationUseCase()

	operationsFromScan, err := scan.ScanOperations()
	if err != nil {
		panic(err)
	}

	feeResults, err := useCase.CalcOperations(operationsFromScan)
	if err != nil {
		panic(err)
	}

	PrettyPrint(&feeResults)
}

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
