package controllers

import (
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
  c.JSON(200, models.Alunos)
}
func Saudacao(c *gin.Context)  {
  nome := c.Params.ByName("nome")
  c.JSON(200,gin.H{
    "Api Diz": "Eai " + nome + ", Tudo Beleza?",
  })
}
