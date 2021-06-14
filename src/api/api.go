package api

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Prueba struct {
	Id          int     `json:"id"`
	Date        string  `json:"fecha"`
	Type        string  `json:"tipo-muestra"`
	Result      int     `json:"resultado"`
	Age         int     `json:"edad"`
	Sex         int     `json:"sexo"`
	Institution string  `json:"institucion"`
	Locale      *Locale `json:"localidad"`
}

type Locale struct {
	Department string `json:"departamento"`
	Province   string `json:"provincia"`
	District   string `json:"distrito"`
}

var pruebas []Prueba

var id_prueba int

//EndPoints
func GetPruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range pruebas {
		id_prueba, _ = strconv.Atoi(params["id"])
		if item.Id == id_prueba {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Prueba{})
}

func GetpruebasEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(pruebas)
}

func CreatePruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var Prueba Prueba
	_ = json.NewDecoder(req.Body).Decode(&Prueba)
	id_prueba, _ = strconv.Atoi(params["id"])
	Prueba.Id = id_prueba
	pruebas = append(pruebas, Prueba)
	json.NewEncoder(w).Encode(pruebas)

}

func DeletePruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id_prueba, _ = strconv.Atoi(params["id"])
	for index, item := range pruebas {
		if item.Id == id_prueba {
			pruebas = append(pruebas[:index], pruebas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(pruebas)
}

func stoInt(v string) int {
	vint, _ := strconv.Atoi(v)
	return vint
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func Add() {
	// pruebas = append(pruebas, Prueba{Id: 1, Date: "15/04/2020", Type: "Molecular", Result: 1, Age: 27, Sex: 0, Institution: "ESSALUD", Locale: &Locale{Department: "Lima", Province: "Huaral", District: "Huaral"}})
	// pruebas = append(pruebas, Prueba{Id: 2, Date: "15/04/2020", Type: "Molecular", Result: 1, Age: 27, Sex: 0, Institution: "ESSALUD", Locale: &Locale{Department: "Lima", Province: "Huaral", District: "Huaral"}})

	// adding example data
	filepath := "data/data2.csv"
	openfile, err := os.Open(filepath)
	checkError("Error in opening the file\n", err)
	filedata, err := csv.NewReader(openfile).ReadAll()
	checkError("Error in reading the file\n", err)

	for i, value := range filedata {
		pruebas = append(pruebas, Prueba{
			Id:          i + 1,
			Date:        value[1],
			Type:        value[2],
			Result:      stoInt(value[3]),
			Age:         stoInt(value[4]),
			Sex:         stoInt(value[5]),
			Institution: value[6],
			Locale: &Locale{
				Department: value[7],
				Province:   value[8],
				District:   value[9],
			}})
	}

}

func HandleFunc() {
	//EndPoints
	router := mux.NewRouter()
	router.HandleFunc("/pruebas", GetpruebasEndpoint).Methods("GET")
	router.HandleFunc("/pruebas/{id}", GetPruebaEndpoint).Methods("GET")
	router.HandleFunc("/pruebas/{id}", CreatePruebaEndpoint).Methods("POST")
	router.HandleFunc("/pruebas/{id}", DeletePruebaEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
