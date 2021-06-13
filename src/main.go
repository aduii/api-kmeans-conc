package main

import (
	"github.com/gorilla/mux"
)

type Prueba struct {
	Id          int
	Date        string
	Type        string
	Result      int
	Age         int
	Sex         int
	Institution string
	Locale      *Locale
}

type Locale struct {
	Department string
	Province   string
	District   string
}

var pruebas []Prueba

//EndPoints

func main() {
	router := mux.NewRouter()
}
