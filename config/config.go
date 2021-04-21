package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configs struct {
	AppPort int
	AppName string

	DBConfigs DBConfigs
}

type DBConfigs struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBDriver   string
	DBName     string
	DBuri      string
}

var configs Configs

func ReadEnvInt(env string) int {
	envInt, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		panic("Config: unable to convert environment variable from string to int")
	}
	return envInt
}

func ReadEnvString(env string) string {
	envString := os.Getenv(env)
	if envString == "" {
		panic("Config: environment variable not set")
	}
	return envString
}

func InitConfigs() {

	err := godotenv.Load("/app/application.env")
	if err != nil {
		log.Fatalf("Config: error loading environment variables, %s", err.Error())
	} else {
		fmt.Println("Config: successfully loaded environment variables")
	}

	configs.AppPort = ReadEnvInt("APP_PORT")
	configs.AppName = ReadEnvString("APP_NAME")
	configs.DBConfigs.DBHost = ReadEnvString("DB_HOST")
	configs.DBConfigs.DBPort = ReadEnvInt("DB_PORT")
	configs.DBConfigs.DBUser = ReadEnvString("DB_USER")
	configs.DBConfigs.DBPassword = ReadEnvString("DB_PASSWORD")
	configs.DBConfigs.DBName = ReadEnvString("DB_NAME")
	configs.DBConfigs.DBDriver = ReadEnvString("DB_DRIVER")
	configs.DBConfigs.DBuri = ReadEnvString("DB_URI")

}

func AppPort() string {
	return strconv.Itoa(configs.AppPort)
}

func AppName() string {
	return configs.AppName
}

func GetDBConfig() DBConfigs {
	return configs.DBConfigs
}
