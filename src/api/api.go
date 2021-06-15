package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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
	Locale      *Locale `json:"localidad,omitempty"`
}

type Locale struct {
	Department string `json:"departamento"`
	Province   string `json:"provincia"`
	District   string `json:"distrito"`
}

var Pruebas []Prueba

var id_prueba int

//EndPoints
func GetPruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range Pruebas {
		id_prueba, _ = strconv.Atoi(params["id"])
		if item.Id == id_prueba {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	out_msg := fmt.Sprint("Prueba no encontrada con id ", id_prueba)
	json.NewEncoder(w).Encode(out_msg)
	// json.NewEncoder(w).Encode(&Prueba{})
}

func GetPruebasEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Pruebas)
}

func CreatePruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var Prueba Prueba
	_ = json.NewDecoder(req.Body).Decode(&Prueba)
	id_prueba, _ = strconv.Atoi(params["id"])
	Prueba.Id = id_prueba
	Pruebas = append(Pruebas, Prueba)
	out_msg := fmt.Sprint("Prueba creada con id ", id_prueba)
	json.NewEncoder(w).Encode(out_msg)
}

func DeletePruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id_prueba, _ = strconv.Atoi(params["id"])
	for index, item := range Pruebas {
		if item.Id == id_prueba {
			Pruebas = append(Pruebas[:index], Pruebas[index+1:]...)
			break
		}
	}
	out_msg := fmt.Sprint("Prueba eliminada con id ", id_prueba)
	json.NewEncoder(w).Encode(out_msg)
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
	// Pruebas = append(Pruebas, Prueba{Id: 1, Date: "15/04/2020", Type: "HISOPADO NASAL Y FARINGEO", Result: 1, Age: 27, Sex: 0, Institution: "ESSALUD", Locale: &Locale{Department: "Lima", Province: "Huaral", District: "Huaral"}})

	// adding example data
	filepath := "data/data2.csv"
	openfile, err := os.Open(filepath)
	checkError("Error in opening the file\n", err)
	filedata, err := csv.NewReader(openfile).ReadAll()
	checkError("Error in reading the file\n", err)

	for i, value := range filedata {
		Pruebas = append(Pruebas, Prueba{
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
	port := ":3000"
	router.HandleFunc("/pruebas", GetPruebasEndpoint).Methods("GET")
	router.HandleFunc("/pruebas/{id}", GetPruebaEndpoint).Methods("GET")
	router.HandleFunc("/pruebas/{id}", CreatePruebaEndpoint).Methods("POST")
	router.HandleFunc("/pruebas/{id}", DeletePruebaEndpoint).Methods("DELETE")

	fmt.Printf("\n Corriendo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
