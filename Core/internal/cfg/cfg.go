package cfg

import (
	"core/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
	IP   string

	DbName     string
	DbPassword string
	DbUser     string
	DbPort     string
	DbHost     string
}

var Cfg Config

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func InitCfg() {
	log := logger.GetLogger()

	port, exst := os.LookupEnv("PORT")
	if !exst {
		log.Errorf("Error get PORT")
	}

	id, exst := os.LookupEnv("ID")
	if !exst {
		log.Errorf("Error get ID")
	}

	dbName, exst := os.LookupEnv("POSTGRES_DB")
	if !exst {
		log.Errorf("Error get DbName")
	}

	dbPassword, exst := os.LookupEnv("POSTGRES_PASSWORD")
	if !exst {
		log.Errorf("Error get DBPASSWORD")
	}

	dbUser, exst := os.LookupEnv("POSTGRES_USER")
	if !exst {
		log.Errorf("Error get DBUSER")
	}

	dbHost, exst := os.LookupEnv("DB_HOST")
	if !exst {
		log.Errorf("Error get DBUSER")
	}

	Cfg = Config{
		Port: port,
		IP:   id,

		DbName:     dbName,
		DbPassword: dbPassword,
		DbPort:     "5432",
		DbUser:     dbUser,
		DbHost:     dbHost,
	}
}
