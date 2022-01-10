package main

import (
	"fmt"
	"log"
	"net/http"
	"tralc/api"
);

func main() {
	const SERVER_PORT string = "5035";
 	
	api.Route();

	fmt.Print("Server running at port ", SERVER_PORT, "\n");
	serverStatus := http.ListenAndServe(":" + SERVER_PORT, nil);
	if serverStatus != nil {
		log.Fatal(serverStatus)
	}
}

