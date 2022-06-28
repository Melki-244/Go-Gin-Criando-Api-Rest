package controllers

import (
	"net/http"

	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/database"
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
func CriaNovoAluno(c *gin.Context)  {
  var aluno models.Aluno
  if err := c.ShouldBindJSON(&aluno); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  database.DB.Create(&aluno)
  c.JSON(http.StatusOK, aluno)
}
