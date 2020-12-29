package tutils

import (
	"fmt"
	"io/ioutil"
	"os"
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
