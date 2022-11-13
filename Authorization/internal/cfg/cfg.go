package cfg

import (
	"foodway/internal/domain"
)

var Cfg domain.Config

func LoadEnv() {

}

func InitCfg() {
	//log := logger.GetLogger()
	//
	//port, exst := os.LookupEnv("PORT")
	//if !exst {
	//	log.Errorf("Error get PORT")
	//}
	//
	//id, exst := os.LookupEnv("ID")
	//if !exst {
	//	log.Errorf("Error get ID")
	//}
	//
	//dbName, exst := os.LookupEnv("DBNAME")
	//if !exst {
	//	log.Errorf("Error get DbName")
	//}
	//
	//dbPassword, exst := os.LookupEnv("DBPASSWORD")
	//if !exst {
	//	log.Errorf("Error get DBPASSWORD")
	//}
	//
	//dbPort, exst := os.LookupEnv("DBPORT")
	//if !exst {
	//	log.Errorf("Error get DBPORT")
	//}
	//
	//dbUser, exst := os.LookupEnv("DBUSER")
	//if !exst {
	//	log.Errorf("Error get DBUSER")
	//}

	Cfg = domain.Config{
		Port: "8080",
		IP:   "localhost",

		DbName:     "test-db",
		DbPassword: "13134777",
		DbPort:     "5436",
		DbUser:     "postgres",
	}
}
