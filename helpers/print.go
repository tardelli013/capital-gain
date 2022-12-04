package helpers

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n%s\n", p)
}
