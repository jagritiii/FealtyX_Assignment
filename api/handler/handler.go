package handler

import (
	"crud_api/api/service"
	"crud_api/pkg/enums"
	"crud_api/pkg/model"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func CreateStudent(c echo.Context) error {
	var reqBody model.Student
	err := json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.Faileddecode)
	}

	v := validator.New()
	err = v.StructExcept(&reqBody, "ID")
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.Validationerror)
	}
	err = service.CreateStudent(reqBody)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, enums.InsertFailed)
	}

	return c.JSON(http.StatusOK, enums.InsertSucceeded)
}

func GetStudents(c echo.Context) error {
	studentList, err := service.GetAllStudents()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, enums.FailedAllStudents)
	}
	return c.JSON(http.StatusOK, studentList)
}

func GetStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.InvalidStudentId)
	}

	student, err := service.GetStudentByID(id)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.NoStudentById)
	}
	return c.JSON(http.StatusOK, student)
}

func UpdateStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.InvalidStudentId)
	}

	var reqBody model.Student
	err = json.NewDecoder(c.Request().Body).Decode(&reqBody)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.Faileddecode)
	}

	v := validator.New()
	err = v.StructExcept(&reqBody, "ID")
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.Validationerror)
	}

	err = service.UpdateStudentByID(id, reqBody)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, enums.UpdateFailed)
	}

	return c.JSON(http.StatusOK, enums.UpdateSucceeded)
}

func DeleteStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, enums.InvalidStudentId)
	}
	err = service.DeleteStudentByID(id)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, enums.DeleteFailed)
	}
	return c.NoContent(http.StatusNoContent)
}

func GenerateStudentSummary(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid student ID")
	}

	student, ok := service.Students[id]

	if !ok {
		return c.JSON(http.StatusNotFound, "Student not found")
	}

	summary, err := service.GetStudentSummaryFromOllama(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, summary)
}
