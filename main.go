package main

import "github.com/gin-gonic/gin"

func ExibeTodosOsAlunos(c *gin.Context) {
  c.JSON(200, gin.H{
    "id":"1",
    "nome":"Melki",
  })
}
func main() {
  r := gin.Default()
  r.GET("/alunos", ExibeTodosOsAlunos)
  r.Run()
}
