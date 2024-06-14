package controller

import (
	"crud_api/api/handler"

	"github.com/labstack/echo/v4"
)

func Createroutes(e *echo.Echo) {
	e.POST("/students", handler.CreateStudent)
	e.GET("/students", handler.GetStudents)
	e.GET("/students/:id", handler.GetStudentByID)
	e.PUT("/students/:id", handler.UpdateStudentByID)
	e.DELETE("/students/:id", handler.DeleteStudentByID)
	e.GET("/students/:id/summary", handler.GenerateStudentSummary)
}
