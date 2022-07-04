package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mbaezahuenupil/go-mongodb-test/src/server"
	"github.com/mbaezahuenupil/go-mongodb-test/src/util"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading environments %s", err.Error()))
	}
	util.ConfigureLog()
	router := mux.NewRouter()
	server.ConfigureServer(router)

	timeWrite, err := strconv.Atoi(os.Getenv("HTTP_WRITE_TIMEOUT"))
	if err != nil {
		panic(fmt.Sprintf("time write problems to cast %v", err.Error()))
	}
	timeRead, err := strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	if err != nil {
		panic(fmt.Sprintf("time read problems to cast %v", err.Error()))
	}
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT")),
		WriteTimeout: time.Duration(timeWrite) * time.Second,
		ReadTimeout:  time.Duration(timeRead) * time.Second,
	}
	log.Printf("Server Running. Version: %v on Port: %v\n", os.Getenv("VERSION"), os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
