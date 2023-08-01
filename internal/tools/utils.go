package tools

import (
	"encoding/json"
	"fmt"
)

func prettyJSON(i interface{}) string {
	res, err := json.MarshalIndent(i, " ", "\t")
	if err != nil {
		fmt.Errorf("marshalling is failed: %v", err.Error())
	}

	return string(res)
}
