package main

import (
	"net/http"

	"github.com/senodp/go-rest-api/database"
	"github.com/labstack/echo"
)

func main(){
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil{
		panic(err)
	}

	e := echo.New()

	e.Start(":8080")
}