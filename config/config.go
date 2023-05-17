package config

import (
	"BukaLelang/config/common"
	"log"
	"os"
	"sync"

	// "google.golang.org/genproto/googleapis/cloud/oslogin/common"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Port    string 
	Database struct {
		Driver   string 
		Host     string
		Name     string
		Address  string 
		Port     string 
		Username string 
		Password string 
	}
} 

var lock = &sync.Mutex{}
var appConfig *Configuration 

func GetConfiguration() *Configuration {
	lock.Lock() 
	defer lock.Unlock() 

	if appConfig == nil {
		appConfig = InitConfiguration()
	} 
	return appConfig
} 

func InitConfiguration() *Configuration {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ")
	}

	var defaultConfig Configuration 
	defaultConfig.Port = os.Getenv("AppPort")
	defaultConfig.Database.Host = os.Getenv("Host")
	defaultConfig.Database.Port = os.Getenv("Port")
	defaultConfig.Database.Username = os.Getenv("Username")
	defaultConfig.Database.Password = os.Getenv("Password")
	defaultConfig.Database.Name = os.Getenv("Name")
	common.JWTSecret = os.Getenv("JWTSecret")
	common.AWS_REGION = os.Getenv("AWS_REGION")
	common.ACCESS_KEY_ID = os.Getenv("ACCESS_KEY_ID") 
	common.ACCESS_SECRET_KEY = os.Getenv("AWS_ACCESS_SECRET_KEY")
	common.SERVER_KEY_MIDTRANS = os.Getenv("SERVER_KEY_MIDTRANS")

	return &defaultConfig 
}