package redis

import (
	json "github.com/json-iterator/go"
)

// GetJSON - gets a string-serialized struct in Redis at the given key
func GetJSON(key string, object interface{}) error {
	value, err := Get(key)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(value), object)
	if err != nil {
		return err
	}
	return nil
}
