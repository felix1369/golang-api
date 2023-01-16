package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// EnvType is variable in .env file
type EnvType struct {
	DbConn  string
	Timeout int
}

// Env is global var for EnvType
var Env = EnvType{}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	Env.DbConn = getEnv("DB_CONN", "default conn")
	Env.Timeout = getEnvAsInt("TIMEOUT", 60)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}

func getEnvAsBool(name string, defaultValue bool) bool {
	valueStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valueStr); err == nil {
		return val
	}

	return defaultValue
}

func getEnvAsSlice(name string, defaultValue []string, separator string) []string {
	valueStr := getEnv(name, "")

	if valueStr == "" {
		return defaultValue
	}

	val := strings.Split(valueStr, separator)

	return val
}
