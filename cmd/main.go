package main

import (
	"gocv-clean/cmd/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

// Frameworks más usados:
// Revel
// Beego
// Martini
// Gin
// Buffalo

// Librerías más usadas:
// Testify (testing)
// Logrus (log)
// pkg/errors (errors)
// cobra
// godog

// Librerías utilizadas para APIs
// Gorilla mux
// Negroni
// google/jsonapi
// grpc
// echo

// Web-scraping frameworks:
// Pholcus (chino)
// go_spider
// ants-go
// colly
// Dataflow kit

// Seguir:
// @golang
// @golang_news
// @golangch
// @golangweekly
// @francesc
// @bketelsen
// @goinggodotnet
// @matryer
// @golangbcn
// @FriendsOfGo
