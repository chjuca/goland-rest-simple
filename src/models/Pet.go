package models

type Pet struct {
	ID   int    `json:ID`
	Name string `json:Name`
	Age  int    `json:Age`
}
