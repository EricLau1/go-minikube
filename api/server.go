package api

import (
	"fmt"
	"go-minikube/api/routes"
	"log"
	"net/http"
)

func Run() {
	listen(9000)
}

func listen(port int) {
	fmt.Printf("\n\nListening on port %d...\n", port)
	fmt.Printf("\n\tAccess ==> localhost:%d/help\n\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.LoadCors(r)))
}
