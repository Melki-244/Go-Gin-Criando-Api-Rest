package models

type Aluno struct {
  Nome string `json:"aluno"`
  Cpf string `json:"cpf"`
  Rg string `json:"rg"`
}

var Alunos []Aluno
