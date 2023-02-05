package ollipackage

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hallo Olli\n")
}

func headers(responseWriter http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		for _, header := range headers {
			fmt.Fprintf(responseWriter, "%v: %v\n", name, header)
		}
	}
}

func StartHttp() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
}
