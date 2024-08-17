package routes

import (
	"net/http"

	"github.com/Go-Dev-and-Chill/chat-app-server/controllers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// The HandleFunc function reads a configuration file, sets up a server using Gin framework, and
// listens on a specified server address and port.
func HandleFunc() {
	//readding the config file
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/", controllers.HelloWorld)
	r.NoRoute(controllers.NotFound)

	http.ListenAndServe(viper.GetString("SERVER_ADRESS") + ":" + viper.GetString("PORT"), r)

}
