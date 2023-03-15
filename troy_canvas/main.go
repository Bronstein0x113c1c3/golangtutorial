package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"troy_canvas/entity"
	"troy_canvas/implementation"

	"github.com/joho/godotenv"
)

func main() {
	student := &implementation.StudentInfo{}
	Loading(student)

	/*
		Step 3: Some commands
		1. Student:
			-Get All Courses (Name, ID, Index)
		2. Courses:
			-Get course info
			-Get Announements
		3. Announements
		4. Quizzes
	*/

}

func Loading(s *implementation.StudentInfo) error {
	/*
		Step 1: Init the credential
	*/

	godotenv.Load(".env")
	token := os.Getenv("TOKEN")
	log.Println("Load the env file done!")
	/*
		Step 2: Wait for loading information
			2.1: Load student information
			2.2: Load courses
			2.3.1: Load quizzes
			2.3.2: Load Announcements
	*/
	//2.1: Load the student info
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://troy.instructure.com/api/v1/users/self", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	// var announcement entity.Course_Announement_Instructure
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// io.Copy(os.Stdout, resp.Body)
	var student_info entity.Student_Info_Instructure
	err = json.NewDecoder(resp.Body).Decode(&student_info)
	if err != nil {
		return err
	}
	s.Info = &student_info
	log.Println("Load the student information done!")
	//2.2: Load courses
	err = s.GetAllCourses(token)
	if err != nil {
		return err
	}
	log.Println("Load the courses done!")
	//2.3.1: Load quizzes, announcement
	for _, c := range s.Courses {
		c.GetAllAnnounements(token)
		c.GetAllQuizzes(token)
	}
	//Done!
	return nil
}
func Command() {

}

// for testing purposes
func GetUser(token string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://troy.instructure.com/api/v1/users/self", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	// var announcement entity.Course_Announement_Instructure
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// io.Copy(os.Stdout, resp.Body)
	var student_info entity.Student_Info_Instructure
	json.NewDecoder(resp.Body).Decode(&student_info)
	fmt.Printf("%+v\n", student_info)
	return nil
}
func GetCourses(token string) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://troy.instructure.com/api/v1/courses", &bytes.Buffer{})
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	// var announcement entity.Course_Announement_Instructure
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var courses []entity.Course_Info_Instructure
	json.NewDecoder(resp.Body).Decode(&courses)
	fmt.Printf("%+v \n", courses)
	fmt.Printf("%v", len(courses))

	return nil
}
