package helpers

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) {
	fmt.Printf(ToJsonString(data))
}

func ToJsonString(data interface{}) string {
	p, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(p)
}
