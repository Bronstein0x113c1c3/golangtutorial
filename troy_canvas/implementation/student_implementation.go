package implementation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"troy_canvas/entity"
)

func (s *StudentInfo) GetAllCourses(token string) error {
	/*
		Step 1: Preparing for request
	*/
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://troy.instructure.com/api/v1/courses", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	// var announcement entity.Course_Announement_Instructure
	if err != nil {
		return err
	}
	/*
		Get the whole request
	*/
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var courses []entity.Course_Info_Instructure
	json.NewDecoder(resp.Body).Decode(&courses)
	for _, v := range courses {
		s.Courses = append(s.Courses, &Course{
			Info: &v,
		})
	}
	return nil
}
