package controllers

import (
	"net/http"

	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/database"
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
  var alunos []models.Aluno
  database.DB.Find(&alunos)
  c.JSON(200, alunos)
}
func Saudacao(c *gin.Context)  {
  nome := c.Params.ByName("nome")
  c.JSON(200,gin.H{
    "Api Diz": "Eai " + nome + ", Tudo Beleza?",
  })
}
func CriaNovoAluno(c *gin.Context)  {
  var novoaluno models.Aluno
  if err := c.ShouldBindJSON(&novoaluno); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  database.DB.Create(&novoaluno)
  c.JSON(http.StatusOK, novoaluno)
}
