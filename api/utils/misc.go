package utils

import (
	"bytes"
	"encoding/json"
)

func JsonPrettyPrint(input []byte) []byte {
	var output bytes.Buffer
	err := json.Indent(&output, input, "", "\t")
	if err != nil {
		return []byte(input)
	}
	return output.Bytes()
}
