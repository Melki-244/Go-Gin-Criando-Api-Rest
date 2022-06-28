package database

import (
	"log"

	"github.com/Melki-244/Go-Gin-Criando-Api-Rest/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
  DB *gorm.DB 
  err error
)

func ConectaComBancoDeDados()  {
  // Capturando Variavéis de Ambiente do Arquivo .env
  var envs map[string]string
  envs, err := godotenv.Read(".env")

  if err != nil {
    log.Panic("Erro Ao Carregar As Variavéis de Ambiente")
  }

  user := envs["POSTGRES_USER"]
  password := envs["POSTGRES_PASSWORD"]
  dbname := envs["POSTGRES_DB"]

  // Conetando Com O Banco De Dados
  stringconection := "host=localhost user="+user+" password="+password+" dbname="+dbname+" port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  DB, err = gorm.Open(postgres.Open(stringconection)) 
  if err != nil {
    log.Panic("Erro Ao Conectar Com O Banco De Dados")
  }

  // Migrando Para O Banco De Dados
  DB.AutoMigrate(&models.Aluno{})
}
