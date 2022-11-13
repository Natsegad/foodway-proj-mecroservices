package cfg

type Config struct {
	Port string
	IP   string

	DbName     string
	DbPassword string
	DbUser     string
	DbPort     string
}

var Cfg Config

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

	//Cfg = Config{
	//	Port: port,
	//	IP:   id,
	//
	//	DbName:     dbName,
	//	DbPassword: dbPassword,
	//	DbPort:     dbPort,
	//	DbUser:     dbUser,
	//}

	Cfg = Config{
		Port: "8080",
		IP:   "localhost",

		DbName:     "foodway",
		DbPassword: "foodwaypass",
		DbPort:     "5432",
		DbUser:     "foodway",
	}
}
