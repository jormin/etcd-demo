package helper

import (
	"encoding/json"
	"fmt"
)

// PrintJSON 打印Json
func PrintJSON(v interface{}) {
	b, _ := json.Marshal(v)
	fmt.Printf("%s\n", b)
}
