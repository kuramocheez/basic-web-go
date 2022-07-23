package main

import (
	"github.com/kuramocheez/belajar-go/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", controller.GetData)
	r.HandleFunc("/notes/tambah-data", controller.PostData)
	r.HandleFunc("/notes/ubah-data/{id}", controller.UpdateData)
	r.HandleFunc("/notes/hapus-data/{id}", controller.DeleteData)
	http.ListenAndServe(":8000", r)
}


