package main

import (
	"net/http"
	//"fmt"
	"encoding/json"
	"github.com/senodp/go-rest-api/database"
	"github.com/labstack/echo"
)

type CreateRequest struct{
	Title 			string `json:"title"`
	Description 	string `json:"description"`
}

type TodoResponse struct{
	Id 				int 	`json:"id"`
	Title 			string 	`json:"title"`
	Description 	string 	`json:"description"`
	Done 			bool 	`json:"bool"`
}

func main(){
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil{
		panic(err)
	}

	e := echo.New()

	e.GET("/todo", func(ctx echo.Context)error{
		rows, err := db.Query("SELECT * FROM todolist")

		if err != nil{
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []TodoResponse

		for rows.Next(){
			var id int 
			var title string
			var description string
			var done int

			err = rows.Scan(&id, &title, &description, &done)

			if err != nil{
			return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var todo TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1{
				todo.Done = true
			}

			res = append(res, todo)
		}

		return ctx.JSON(http.StatusOK, res)
	})

	e.POST("/todo", func(ctx echo.Context)error{

		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todolist (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)
		if err != nil{
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		// fmt.Println(request)

		return ctx.String(http.StatusOK, "OK")
	})

	e.DELETE("/todo/:id", func(ctx echo.Context)error{

		//ambil data id
		id := ctx.Param("id")

		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"DELETE FROM todolist WHERE id = ?",
			id,
		)
		if err != nil{
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		// fmt.Println(request)

		return ctx.String(http.StatusOK, "OK")
	})

	e.Start(":8080")
}