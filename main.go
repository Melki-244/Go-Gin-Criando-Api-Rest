package main

import (
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/models"
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/routes"
)

func main() {
  models.Alunos = []models.Aluno{
    {Nome: "Melki", Cpf: "000000000", Rg: "000000000"},
    {Nome: "Guilherme", Cpf: "000000000", Rg: "000000000"},
  }
  routes.HandleRequests() 
}
