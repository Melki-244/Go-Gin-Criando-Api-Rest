package models

import "gorm.io/gorm"

type Aluno struct {
  gorm.Model
  Nome string `json:"nome"`
  Cpf string `json:"cpf"`
  Rg string `json:"rg"`
}
