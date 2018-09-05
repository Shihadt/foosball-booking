package bootstrapper

import (
	"log"

	"github.com/spf13/viper"
	utils "github.com/user/work/foosball-booking/apputil"
)

var mode = "development"

//Startup intitialize project configuration.
func Startup() {

	//Initialise RSA keys
	utils.InitRSAKeys()
	// Initialize Logger objects with Log Level
	utils.SetLogLevel(utils.Level(AppConfig.LogLevel))

	
}

// AppConfig - application configuration will be stored in this module
var AppConfig configuration

type configuration struct {
	Server, DBHost, DBUser, DBPassword, DBName string
	LogLevel                                   int
}

func init() {
	// set filename , We are using app.toml as configuration file
	viper.SetConfigName("app")
	// path to app.toml
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config file not found:", err)
	}
	AppConfig = configuration{}
	AppConfig.DBHost = viper.GetString(mode + ".PostgresHost")
	AppConfig.DBUser = viper.GetString(mode + ".PostgresUser")
	AppConfig.DBPassword = viper.GetString(mode + ".PostgresPassword")
	AppConfig.DBName = viper.GetString(mode + ".Database")
	AppConfig.LogLevel = viper.GetInt(mode + ".LogLevel")
}
