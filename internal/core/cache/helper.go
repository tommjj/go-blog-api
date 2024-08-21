package cache

import (
	"encoding/json"
	"fmt"
)

func unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func generateCacheKeyParams(params ...any) string {
	var str string

	last := len(params) - 1
	for i, param := range params {
		str += fmt.Sprintf("%v", param)

		if i != last {
			str += "-"
		}
	}

	return str
}
