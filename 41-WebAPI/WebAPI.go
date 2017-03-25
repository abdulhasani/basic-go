package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**WEB API adalah sebuah web yang menerima request dari client dan menghasilkan response
biasa JSON/XML
*/

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	student{"E001", "Hasani", 24},
	student{"E002", "Bridge", 25},
	student{"E003", "Abdul Kadir Hasani", 26},
}

// untuk handle route
func users(w http.ResponseWriter, r *http.Request) {
	//digunakan untuk menentukan response
	w.Header().Set("Content-Type", "application/json")
	//deteksi jenis request
	if r.Method == "POST" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

//siapkan juga handler untuk rute /user. perbedaan rute ini dengan rute /users diatas adalah
// /users menghasilkan semua sample data yang ada (array)
// /user menghasilkan satu buah data saja, diambil dari data sample berdasarkan IDnya
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Tpye", "application/json")
	if r.Method == "POST" {
		var id = r.FormValue("id") //mengambil data data payload yang dikirim dari client pada hal ini merupakan data ID
		var resut []byte
		var err error
		//mencari sesuai yang direquest
		for _, each := range data {
			if each.ID == id {
				resut, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(resut)
				return
			}
		}
		http.Error(w, "User not found", http.StatusBadRequest)
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
