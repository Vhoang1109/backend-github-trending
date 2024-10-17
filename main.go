package main

import (
	"backend-github-trending/db"
	"backend-github-trending/handler"
	"backend-github-trending/repository/repo_impl"
	"backend-github-trending/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "hoang1109",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
