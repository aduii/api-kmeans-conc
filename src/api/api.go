package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	km "github.com/aduii/api-kmeans-conc/src/kmeans"

	"github.com/gorilla/mux"
)

var id_prueba int

//EndPoints
func GetPruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	for _, item := range km.Pruebas2 {
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

func GetClusterEndpoint(w http.ResponseWriter, req *http.Request) {
	var Cluster []km.Prueba
	params := mux.Vars(req)
	for _, item := range km.Pruebas2 {
		w.Header().Set("Content-Type", "application/json")
		id_prueba, _ = strconv.Atoi(params["id"])
		if item.Cluster == id_prueba {
			Cluster = append(Cluster, item)
		}
	}
	if Cluster != nil {
		json.NewEncoder(w).Encode(Cluster)
		return
	}
	out_msg := fmt.Sprint("Prueba no encontrada con id ", id_prueba)
	json.NewEncoder(w).Encode(out_msg)
}

func GetPruebasEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(km.Pruebas2)
}

func CreatePruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var Prueba km.Prueba
	_ = json.NewDecoder(req.Body).Decode(&Prueba)
	id_prueba, _ = strconv.Atoi(params["id"])
	Prueba.Id = id_prueba
	km.Pruebas2 = append(km.Pruebas2, Prueba)
	out_msg := fmt.Sprint("Prueba creada con id ", id_prueba)
	json.NewEncoder(w).Encode(out_msg)
}

func DeletePruebaEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id_prueba, _ = strconv.Atoi(params["id"])
	for index, item := range km.Pruebas2 {
		if item.Id == id_prueba {
			km.Pruebas2 = append(km.Pruebas2[:index], km.Pruebas2[index+1:]...)
			break
		}
	}
	out_msg := fmt.Sprint("Prueba eliminada con id ", id_prueba)
	json.NewEncoder(w).Encode(out_msg)
}

func HandleFunc() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/pruebas", GetPruebasEndpoint).Methods("GET")
	apiRouter.HandleFunc("/pruebas/{id}", GetPruebaEndpoint).Methods("GET")
	apiRouter.HandleFunc("/clusters/{id}", GetClusterEndpoint).Methods("GET")
	apiRouter.HandleFunc("/pruebas/{id}", CreatePruebaEndpoint).Methods("POST")
	apiRouter.HandleFunc("/pruebas/{id}", DeletePruebaEndpoint).Methods("DELETE")

	fmt.Printf("\n Corriendo en http://localhost:%s", port)
	portd := ":" + port
	log.Fatal(http.ListenAndServe(portd, router))
}
