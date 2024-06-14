package service

import (
	"bytes"
	"crud_api/pkg/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var (
	Students = make(map[int]model.Student)
	nextID   = 1
	Mu       sync.Mutex
)

const ollamaAPIURL = "http://localhost:11434/api/generate"

func CreateStudent(student model.Student) error {
	Mu.Lock()
	defer Mu.Unlock()

	student.ID = nextID
	nextID++
	Students[student.ID] = student
	return nil
}

func GetAllStudents() ([]model.Student, error) {
	Mu.Lock()
	defer Mu.Unlock()

	studentList := make([]model.Student, 0, len(Students))
	for _, student := range Students {
		studentList = append(studentList, student)
	}

	return studentList, nil
}

func GetStudentByID(id int) (model.Student, error) {
	Mu.Lock()
	defer Mu.Unlock()

	student, exists := Students[id]
	if !exists {
		return model.Student{}, errors.New("student not found")
	}

	return student, nil
}

func UpdateStudentByID(id int, student model.Student) error {
	Mu.Lock()
	defer Mu.Unlock()

	_, exists := Students[id]
	if !exists {
		err := errors.New("student not found")
		if err != nil {
			return err
		}
	}

	student.ID = id
	Students[id] = student

	return nil
}

func DeleteStudentByID(id int) error {
	Mu.Lock()
	defer Mu.Unlock()

	_, exists := Students[id]
	if !exists {
		return errors.New("student not found")
	}

	delete(Students, id)
	return nil
}

func GetStudentSummaryFromOllama(student model.Student) (model.OllamaResponse, error) {
	prompt := fmt.Sprintf("Generate a summary for the student:\nName: %s\nAge: %d\nEmail: %s", student.Name, student.Age, student.Email)
	ollamaReq := model.OllamaRequest{
		Model:  "llama3", // Specify the model you are using
		Prompt: prompt,
		Stream: false,
	}
	var ollamaResp model.OllamaResponse

	reqBody, err := json.Marshal(ollamaReq)
	if err != nil {
		return ollamaResp, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := http.Post(ollamaAPIURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return ollamaResp, fmt.Errorf("failed to call Ollama API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return ollamaResp, fmt.Errorf("ollama API returned error: %s", string(bodyBytes))
	}

	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		log.Println(err.Error())
		log.Println((*resp).Body)
		return ollamaResp, fmt.Errorf("failed to decode Ollama response: %w", err)
	}

	return ollamaResp, nil
}
