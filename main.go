package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/senodp/go-rest-api/database"
	"github.com/labstack/echo"
)

type CreateRequest struct{
	Title 			string `json:"title"`
	Description 	string `json:"description"`
}

func main(){
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil{
		panic(err)
	}

	e := echo.New()

	e.POST("/todo", func(ctx echo.Context)error{

		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		fmt.Println(request)

		return ctx.String(http.StatusOK, "OK")
	})

	e.Start(":8080")
}