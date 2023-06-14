package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Validate_Error string

type ServerConfig struct {
	Port            int
	AllowOrigins    []string
	LogLevel        string
	DefaultLanguage string
	Languages       []string
}

// MySqlConfig - MySqlConfig configuration
type MySqlConfig struct {
	URL      string
	Database string
}

// Config structure.
type Config struct {
	Server ServerConfig
	MySQL  MySqlConfig
}

// AppConfig - Appconfig object,.
var AppConfig = &Config{
	Server: ServerConfig{
		Port:            3000,
		AllowOrigins:    []string{"*"},
		LogLevel:        "info",
		DefaultLanguage: "en",
		Languages:       []string{"en"},
	},
	MySQL: MySqlConfig{
		URL:      "",
		Database: "",
	},
}

// LoadEnv - function load Enviroment variable from .env file.
func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		Server: ServerConfig{
			Port:            getEnvAsInt("API_PORT", 3000),
			AllowOrigins:    strings.Split(getEnv("ALLOW_ORIGIN", "*"), ","),
			LogLevel:        getEnv("LOG_LEVEL", "info"),
			DefaultLanguage: getEnv("DEFAULT_LANGUAGE", "en"),
			Languages:       strings.Split(getEnv("LANGUAGES", "en"), ","),
		},
		MySQL: MySqlConfig{
			URL:      getEnv("URL", ""),
			Database: getEnv("DATABASE", ""),
		},
	}
	AppConfig = config
	return config
}

// Simple helper function to read an environment or return a default value.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value.
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
