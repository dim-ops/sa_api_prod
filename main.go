package main

import (
	"context"
	"go-sa/config"
	"go-sa/database"
	"go-sa/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	entreprise := &database.DataEntreprise{
		Col: collection,
		Ctx: ctx,
	}

	r := mux.NewRouter()

	r.HandleFunc("/datas", handlers.SearchDatas(entreprise)).Methods("GET")
	r.HandleFunc("/datas/{id}", handlers.GetData(entreprise)).Methods("GET")
	r.HandleFunc("/datas", handlers.InsertData(entreprise)).Methods("POST")
	r.HandleFunc("/datas/{id}", handlers.UpdateData(entreprise)).Methods("PATCH")
	r.HandleFunc("/datas/{id}", handlers.DeleteData(entreprise)).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
