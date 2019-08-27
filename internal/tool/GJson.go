package utils

import "encoding/json"

func JsonEncode(v interface{}) (string, error) {

	if by, error := json.Marshal(v); error != nil {
		return "", error
	} else {
		return string(by), nil
	}
}

func JsonDecode(data string, v interface{}) error {
	if error := json.Unmarshal([]byte(data), v); error != nil {
		return error
	}
	return nil
}
