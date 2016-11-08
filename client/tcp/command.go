package tcp

import (
	"fmt"
	"strings"
)

// ProcessCommand processes the command read from the user. It separates the payload
// from the message, and verifies if it are a valid JSON string.
func ProcessCommand(message string) bool {

	var payload string
	var verifyJSON bool
	if strings.HasPrefix(message, "save ") {

		payload = message[len("save "):len(message)]
		verifyJSON = true

	} else if strings.HasPrefix(message, "update ") {

		payload = message[len("update "):len(message)]
		verifyJSON = true

	} else if strings.HasPrefix(message, "delete ") && strings.HasPrefix(message, "lookup ") &&
		strings.HasPrefix(message, "create_db ") && strings.HasPrefix(message, "connect ") {

		verifyJSON = false

	}

	if verifyJSON && !IsJSON(payload) {
		fmt.Println("Não é JSON. Digite novamente")
		return false
	}

	return true
}
