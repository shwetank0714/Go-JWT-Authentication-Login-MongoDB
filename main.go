package main

import (
	"fmt"
	"log"
	"net/http"

	routers "github.com/shwetank0714/jwtmodfile/routes"
)

func main() {

	router := routers.UserRoutes()

	fmt.Println("Listening on Port 8004 --- ")
	log.Fatal(http.ListenAndServe(":8004", router))

}
