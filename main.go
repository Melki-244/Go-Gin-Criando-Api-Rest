package main

import (
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/database"
	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/routes"
)

func main() {
  database.ConectaComBancoDeDados()
  routes.HandleRequests() 
}
