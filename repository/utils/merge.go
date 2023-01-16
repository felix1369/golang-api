package utils

import (
	"encoding/json"
)

// Merge two structs, where a's values take precendence over b's values (a's values will be kept over b's if each field has a value)
func Merge(a, b interface{}) (interface{}, error) {
	jb, err := json.Marshal(b)
	if err != nil {
		return a, err
	}
	err = json.Unmarshal(jb, &a)
	if err != nil {
		return a, err
	}
	return a, nil
}
