package main

import (
	api "github.com/aduii/api-kmeans-conc/src/api"
	km "github.com/aduii/api-kmeans-conc/src/kmeans"
)

func main() {
	km.Kmeans()
	api.HandleFunc()
}
