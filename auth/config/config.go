package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	AuthGRPCPort string

	PostgresMaxConnections int32
	TokenSecretKey         string
	AccessTokenExpiring    time.Duration
	RefreshTokenExpiring   time.Duration
}

// Load ...
func Load() Config {
	if err := godotenv.Load("/article_go_user_service.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "postgres"))

	config.AuthGRPCPort = cast.ToString(getOrReturnDefaultValue("AUTH_GRPC_PORT", ":9001"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.TokenSecretKey = cast.ToString(getOrReturnDefaultValue("TOKEN_SECRET_KEY", "qWerTy"))
	config.AccessTokenExpiring = cast.ToDuration(getOrReturnDefaultValue("ACCESS_TOKEN_EXPIRING", time.Hour*24*7))
	config.RefreshTokenExpiring = cast.ToDuration(getOrReturnDefaultValue("REFRESH_TOKEN_EXPIRING", time.Hour*24*28))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
