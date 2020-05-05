package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dcaponi/echo/animal"
	"github.com/dcaponi/echo/database"
	"github.com/dcaponi/echo/handler"
	"github.com/dcaponi/echo/mineral"
	"github.com/dcaponi/echo/vegetable"
)

func main() {
	db := database.New(database.PSQLConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Db:       os.Getenv("POSTGRES_DB"),
	})
	defer db.Close()

	animalRepo, _ := animal.New(&db)
	aController := animal.NewController(&animalRepo)
	aHandler := handler.NewHandler(&aController)
	http.Handle("/animal", http.HandlerFunc(aHandler.Handle))

	mineralRepo, _ := mineral.New(&db)
	mController := mineral.NewController(&mineralRepo)
	mHandler := handler.NewHandler(&mController)
	http.Handle("/mineral", http.HandlerFunc(mHandler.Handle))

	vegetableRepo, _ := vegetable.New(&db)
	vController := vegetable.NewController(&vegetableRepo)
	vHandler := handler.NewHandler(&vController)
	http.Handle("/vegetable", http.HandlerFunc(vHandler.Handle))

	fmt.Println(http.ListenAndServe(":8080", nil))
}
