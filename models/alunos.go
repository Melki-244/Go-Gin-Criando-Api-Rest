package models

import "gorm.io/gorm"

type Aluno struct {
  gorm.Model
  Nome string `json:"aluno"`
  Cpf string `json:"cpf"`
  Rg string `json:"rg"`
}

var Alunos []Aluno
