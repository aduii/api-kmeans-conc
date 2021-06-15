package main

import (
	api "api-kmeans-conc/src/api"
	km "api-kmeans-conc/src/kmeans"
)

func main() {
	api.Add()
	km.Kmeans()
	api.HandleFunc()
}
