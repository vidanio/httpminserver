package main

import (
	"fmt"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // recupera campos del form tanto GET como POST r.Method(string)
	for k, v := range r.Form {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Printf("IP: %s\t Referrer: %s\n", r.RemoteAddr, r.Referer())
	// END OF FUNC
	http.Redirect(w, r, r.Referer(), http.StatusFound) // vuelve al formulario que la llamó

}
