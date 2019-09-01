package jhfmt

import (
	"encoding/json"
	"fmt"
)

func PrintBeautyStruct(value interface{}) {
	marshalStruct, _ := json.MarshalIndent(value, "", "\t")
	fmt.Println(string(marshalStruct))
}
