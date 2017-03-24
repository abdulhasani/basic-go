package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Abdul Kadir Hasani",
			"Message": "have a nice day",
		}
		//parsing template
		var t, err = template.ParseFiles("Templating.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		//hasil parsing template diisi variable data yang berupa map
		t.Execute(w, data)
	})
	fmt.Println("starting web server http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
