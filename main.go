package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.ListenAndServe(":9090", nil)
}
