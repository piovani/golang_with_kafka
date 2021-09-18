package main

import (
	"piovani/golang_with_kafka/route"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	route.InitRoute()
}
