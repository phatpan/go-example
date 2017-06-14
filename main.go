package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type Books struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type Response struct {
	Result []Books
}

func GetBook(w rest.ResponseWriter, req *rest.Request) {
	result := Response{
		Result: []Books{
			Books{
				Name:        "Angular js",
				Price:       899,
				Description: "angular js",
			},
			Books{
				Name:        "robot framework",
				Price:       699,
				Description: "automate test using robot framework",
			},
		},
	}
	w.WriteJson(&result)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/books/", GetBook),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
