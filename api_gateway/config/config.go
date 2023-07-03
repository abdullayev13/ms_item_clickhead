package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServicePort string

	AuthServiceHost string
	AuthServicePort string

	ArticleServiceHost string
	ArticleServicePort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("ErrEnvNodFound")
	}

	config := Config{}

	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":9000"))

	config.AuthServiceHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_HOST", "localhost"))
	config.AuthServicePort = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_PORT", ":9001"))

	config.ArticleServiceHost = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_HOST", "localhost"))
	config.ArticleServicePort = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_PORT", ":9002"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
