package dinowebportal

import (
	"fmt"
	"net/http"
)

func RunWebPortal(addr string) error {
	http.HandleFunc("/", roothandler)
	return http.ListenAndServe(addr, nil)
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal %s", r.RemoteAddr)
}
