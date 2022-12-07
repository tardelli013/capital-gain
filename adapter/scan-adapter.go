package adapter

import (
	"bufio"
	"desafio-nu/core/domain"
	"encoding/json"
	"fmt"
	"os"
)

type ScanAdapter struct {
}

func NewScan() Scan {
	return &ScanAdapter{}
}

func (s ScanAdapter) ScanOperations() ([]*domain.Oper, error) {
	var operations []*domain.Oper

	fmt.Println("Enter a operations json: ")
	reader := bufio.NewReader(os.Stdin)
	return decode(reader, operations)
}

func decode(reader *bufio.Reader, oper []*domain.Oper) ([]*domain.Oper, error) {
	err := json.NewDecoder(reader).Decode(&oper)
	if err != nil {
		return nil, err
	}
	return oper, nil
}
