package base

import (
	"encoding/json"
	"fmt"
)

func EncodeBody(body any) ([]byte, error) {
	switch body := body.(type) {
	case string:
		return []byte(`"` + body + `"`), nil
	case bool:
		return []byte(fmt.Sprintf("%v", body)), nil
	case int64:
		return []byte(fmt.Sprintf("%v", body)), nil
	case float64:
		return []byte(fmt.Sprintf("%v", body)), nil
	case any:
		return json.Marshal(body)
	}

	return nil, nil
}
