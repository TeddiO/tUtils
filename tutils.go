package tutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// tUtils: A bunch of boilerplate that I end up using in different projects

func GetEnvVar(varName string, defaultValue string) string {
	stringValue, isSet := os.LookupEnv(varName)

	if !isSet {
		return defaultValue
	}

	return stringValue
}

func GetDockerSecret(secretName string) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("/run/secrets/%s", secretName))
	if err != nil {
		return ""
	}

	return string(data)
}

// Tobool is a helper function designed to take any string input and spit
// out the Go equivalent. Anything that isn't "true" (upper or lower)
// will return false. Useful for things like input from env vars.
func Tobool(value string) bool {
	if strings.ToLower(value) == "true" {
		return true
	}

	return false
}
