package jhfmt

import (
	"encoding/json"
	"fmt"
)

// PrintBeautyStruct print struct with beauty format
func PrintBeautyStruct(value interface{}) {
	marshalStruct, _ := json.MarshalIndent(value, "", "\t")
	fmt.Println(string(marshalStruct))
}
