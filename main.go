package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Ola mundo3</h1>")
}

func main() {
	godotenv.Load()

	http.HandleFunc("/probe", Index)

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")), nil)
}
