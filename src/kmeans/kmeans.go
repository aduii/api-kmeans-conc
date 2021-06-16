package kmeans

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Prueba struct {
	Id          int     `json:"id"`
	Date        string  `json:"fecha"`
	Type        string  `json:"tipomuestra"`
	Result      int     `json:"resultado"`
	Age         int     `json:"edad"`
	Sex         int     `json:"sexo"`
	Institution string  `json:"institucion"`
	Locale      *Locale `json:"localidad,omitempty"`
	Cluster     int     `json:"cluster,omitempty"`
}

type Locale struct {
	Department string `json:"departamento"`
	Province   string `json:"provincia"`
	District   string `json:"distrito"`
}

var Pruebas []Prueba
var Pruebas2 []Prueba

var Centroides [][]float64
var DistIndCen [][]float64
var DistMinIndCen []float64
var IndClusterAsignado []int
var firstTime bool = true

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func Calc_Centroides() {
	min := 0
	if firstTime {
		Centroides = [][]float64{
			{float64(Pruebas[randInt(min, len(Pruebas))].Sex), float64(Pruebas[randInt(min, len(Pruebas))].Age)},
			{float64(Pruebas[randInt(min, len(Pruebas))].Sex), float64(Pruebas[randInt(min, len(Pruebas))].Age)},
		}
		firstTime = false
	}
}

func Calc_Dist_Euclidiana() {
	for i := 0; i < len(Pruebas); i++ {
		ind_v1 := float64(Pruebas[i].Sex)
		ind_v2 := float64(Pruebas[i].Age)
		dist_ind_c1 := math.Pow(ind_v1-Centroides[0][0], 2) + math.Pow(ind_v2-Centroides[0][1], 2)
		dist_ind_c2 := math.Pow(ind_v1-Centroides[1][0], 2) + math.Pow(ind_v2-Centroides[1][1], 2)

		a := []float64{dist_ind_c1, dist_ind_c2}
		DistIndCen = append(DistIndCen, a)
	}
}

func Calc_Dist_Min() {
	cluster1 := 1
	cluster2 := 2
	for i, _ := range DistIndCen {
		if DistIndCen[i][0] < DistIndCen[i][1] {
			a := DistIndCen[i][0]
			DistMinIndCen = append(DistMinIndCen, a)
			IndClusterAsignado = append(IndClusterAsignado, cluster1)
		} else {
			b := DistIndCen[i][1]
			DistMinIndCen = append(DistMinIndCen, b)
			IndClusterAsignado = append(IndClusterAsignado, cluster2)
		}
	}
}

func Calc_PromxClus() {
	var cclus1, cclus2 int = 1, 1
	var sclus1_v1, sclus1_v2, sclus2_v1, sclus2_v2 float64 = 0, 0, 0, 0
	var prom_clus1_v1, prom_clus1_v2, prom_clus2_v1, prom_clus2_v2 int = 0, 0, 0, 0
	for i := 0; i < len(Pruebas); i++ {
		if IndClusterAsignado[i] == 1 {
			sclus1_v1 += float64(Pruebas[i].Sex)
			sclus1_v2 += float64(Pruebas[i].Age)
			cclus1++
		} else {
			sclus2_v1 += float64(Pruebas[i].Sex)
			sclus2_v2 += float64(Pruebas[i].Age)
			cclus2++
		}
	}
	prom_clus1_v1 += int(sclus1_v1) / cclus1
	prom_clus1_v2 += int(sclus1_v2) / cclus1
	prom_clus2_v1 += int(sclus2_v1) / cclus2
	prom_clus2_v2 += int(sclus2_v2) / cclus2
	Centroides = [][]float64{
		{float64(prom_clus1_v1), float64(prom_clus1_v2)},
		{float64(prom_clus2_v1), float64(prom_clus2_v2)},
	}
}

func Iteraciones() {
	iteraciones := 0
	for iteraciones < 20 {
		DistIndCen = [][]float64{}
		DistMinIndCen = []float64{}
		IndClusterAsignado = []int{}
		Calc_Centroides()
		Calc_Dist_Euclidiana()
		Calc_Dist_Min()
		Calc_PromxClus()
		iteraciones++
	}
}

func stoInt(v string) int {
	vint, _ := strconv.Atoi(v)
	return vint
}

func ReadCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Add() {
	url := "https://raw.githubusercontent.com/aduii/api-kmeans-conc/master/src/data/data.csv"
	filedata, err := ReadCSVFromUrl(url)
	if err != nil {
		panic(err)
	}
	// adding example data

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

func AddFinal() {
	for i, _ := range Pruebas {
		Pruebas2 = append(Pruebas2, Prueba{
			Id:          i + 1,
			Date:        Pruebas[i].Date,
			Type:        Pruebas[i].Type,
			Result:      Pruebas[i].Result,
			Age:         Pruebas[i].Age,
			Sex:         Pruebas[i].Sex,
			Institution: Pruebas[i].Institution,
			Locale: &Locale{
				Department: Pruebas[i].Locale.Department,
				Province:   Pruebas[i].Locale.Province,
				District:   Pruebas[i].Locale.District,
			},
			Cluster: IndClusterAsignado[i]})
	}
}

func Kmeans() {
	Add()
	Iteraciones()
	AddFinal()
	fmt.Println("Cluster Asignado: ")
	fmt.Println(IndClusterAsignado)
}
