package tutils

import (
	"database/sql"
	"fmt"
	"log"
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
	data, err := os.ReadFile(fmt.Sprintf("/run/secrets/%s", secretName))
	if err != nil {
		return ""
	}

	return string(data)
}

// Tobool is a helper function designed to take any string input and spit
// out the Go equivalent. Anything that isn't "true" (upper or lower)
// will return false. Useful for things like input from env vars.
func Tobool(value string) bool {
	return strings.ToLower(value) == "true"
}

// CreateMySQLDatabaseInstance is a helper function that creates a new database connection
// to a MySQL database.
func CreateMySQLDatabaseInstance(dbHost string, dbUsername string, dbPassword string, dbSchema string, dbConnectionType string, dbOptionalParams string) *sql.DB {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?%s", dbUsername, dbPassword, dbConnectionType, dbHost, dbSchema, dbOptionalParams))

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
