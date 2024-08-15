package routes

import (
	"net/http"

	"github.com/Go-Dev-and-Chill/chat-app-server/controllers"
	"github.com/gin-gonic/gin"
)

func HandleFunc() {
	r := gin.Default()

	r.GET("/", controllers.HelloWorld)
	r.NoRoute(controllers.NotFound)
	http.ListenAndServe("localhost:8080", r)
}
