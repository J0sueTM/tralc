package api

import (
	"fmt"
	"net/http"
);

func checkRoute(r *http.Request, routePath string, routeMethod string) string {
	if (r.URL.Path != routePath) {
		return "404 not found";
	}	else if (r.Method != routeMethod) {
		return "Method not supported";
	}

	return "";
}

func Home(rw http.ResponseWriter, r *http.Request) {
	routeStatus := checkRoute(r, "/", "GET");

	if routeStatus != "" {
		http.Error(rw, routeStatus, http.StatusNotFound);
	}

	fmt.Fprintf(rw, "");
}

func Route() {
	fileServer := http.FileServer(http.Dir("../../public"));
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/": Home,
		"/dummy": Home,
	};

	for routePath, routeFunc := range routes {
		if routePath == "/" {
			http.Handle(routePath, fileServer);

			continue;
		}

		http.HandleFunc(routePath, routeFunc);
	}
}
