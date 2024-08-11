package cache

import "encoding/json"

func unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
