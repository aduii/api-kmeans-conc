package kmeans

import (
	api "api-kmeans-conc/src/api"
	"fmt"
	"math"
	"math/rand"
	"time"
)

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
			{float64(api.Pruebas[randInt(min, len(api.Pruebas))].Sex), float64(api.Pruebas[randInt(min, len(api.Pruebas))].Age)},
			{float64(api.Pruebas[randInt(min, len(api.Pruebas))].Sex), float64(api.Pruebas[randInt(min, len(api.Pruebas))].Age)},
		}
		firstTime = false
	}
}

func Calc_Dist_Euclidiana() {
	for i := 0; i < len(api.Pruebas); i++ {
		ind_v1 := float64(api.Pruebas[i].Sex)
		ind_v2 := float64(api.Pruebas[i].Age)
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
	var cclus1, cclus2 int = 0, 0
	var sclus1_v1, sclus1_v2, sclus2_v1, sclus2_v2 float64 = 0, 0, 0, 0
	var prom_clus1_v1, prom_clus1_v2, prom_clus2_v1, prom_clus2_v2 int = 0, 0, 0, 0
	for i := 0; i < len(api.Pruebas); i++ {
		if IndClusterAsignado[i] == 1 {
			sclus1_v1 += float64(api.Pruebas[i].Sex)
			sclus1_v2 += float64(api.Pruebas[i].Age)
			cclus1++
		} else {
			sclus2_v1 += float64(api.Pruebas[i].Sex)
			sclus2_v2 += float64(api.Pruebas[i].Age)
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

func Kmeans() {
	Iteraciones()
	fmt.Println("Cluster Asignado: ")
	fmt.Println(IndClusterAsignado)
}
