package routes

import (
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests()  {
  r := gin.Default()
  r.GET("/alunos", controllers.ExibeTodosOsAlunos)
  r.POST("/alunos", controllers.CriaNovoAluno)
  r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
  r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
  r.GET("/:nome", controllers.Saudacao)
  r.Run()
}
