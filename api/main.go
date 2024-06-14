package main

import (
	"crud_api/api/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.Createroutes(e)
	e.Logger.Fatal(e.Start(":7000"))
}
