package utils


import(
	"log"
	"os"
	"github.com/spf13/viper"
)

func LoadEnv(){
	viper.SetConfigFile(".env")
	if err := viper.ReadConfig(); err != nil{
		log.Fatal("Error reading .env file: %v", err)
	}

	viper.AutomaticEnv()
}

func GetEnv(key string) string{
	return viper.GetString(key)
}