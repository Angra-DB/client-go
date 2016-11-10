package utils

import "encoding/json"

// IsJSON is a method that verifies if an entry is a valid JSON string
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
