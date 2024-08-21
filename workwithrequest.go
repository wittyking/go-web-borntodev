package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Course struct {
	ID         int     `json: "id"`
	Name       string  `json: "name"`
	Price      float64 `json: "price`
	Instructor string  `json: "instructor`
}

var courseList []Course

func init() {
	courseJSON := `[
	{
		"id":101,
		"name": "Python",
		"price": 2590,
		"instructor": "BorntoDev"
	},
	{
		"id":102,
		"name": "JavaScript",
		"price": 2000,
		"instructor": "BorntoDev"
	},
	{
		"id":103,
		"name": "SQL",
		"price": 3000,
		"instructor": "BorntoDev"
	}
]`
	err := json.Unmarshal([]byte(courseJSON), &courseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, course := range courseList {
		if highestID < course.ID {
			highestID = course.ID
		}
	}
	return highestID + 1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(courseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		bodybytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("bodybytes conv :", string(bodybytes))
		err = json.Unmarshal(bodybytes, &newCourse)
		if err != nil {
			fmt.Println("error :", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newCourse.ID = getNextID()
		fmt.Println("unmarshal :", newCourse)
		courseList = append(courseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)
}
