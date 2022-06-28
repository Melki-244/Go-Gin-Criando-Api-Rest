package routes

import (
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests()  {
  r := gin.Default()
  r.GET("/alunos", controllers.ExibeTodosOsAlunos)
  r.Run()
}
