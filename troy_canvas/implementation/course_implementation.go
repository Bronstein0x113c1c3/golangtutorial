package implementation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"troy_canvas/entity"
)

func (c *Course) GetAllAnnounements(token string) error {
	/*
		Step 1: Preparing the request with Authorization
		Don't forget to prevent error!
	*/

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://troy.instructure.com/api/v1/announcements?context_codes[]=course_%v", c.Info.ID), &bytes.Buffer{})
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	// var announcement entity.Course_Announement_Instructure
	if err != nil {
		return err
	}

	/*
		Step 2: Get the whole response
	*/

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var all_announcements []entity.Course_Announement_Instructure
	err = json.NewDecoder(resp.Body).Decode(&all_announcements)
	if err != nil {
		return err
	}
	// io.Copy(os.Stdout, resp.Body)

	/*
		Step 3: Import all announement to the courses
	*/
	for _, v := range all_announcements {

		c.Announcements = append(c.Announcements, &Announcement{
			Info: &v,
		})
	}
	return nil
}
func (c *Course) GetAllQuizzes(token string) error {

	/*
		Step 1: Preparing the client, authorizing
	*/
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://troy.instructure.com/api/v1/courses/%v/quizzes/", c.Info.ID), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	// var announcement entity.Course_Announement_Instructure
	if err != nil {
		return err
	}
	//Step 2: Processing the request, get the response, deserializing
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var all_quizzes []entity.Quiz_Instructure
	err = json.NewDecoder(resp.Body).Decode(&all_quizzes)
	//Step 3: Import all the quizzes information in there!
	if err != nil {
		return err
	}
	for _, v := range all_quizzes {
		c.Quizzes = append(c.Quizzes, &Quiz{
			Info: &v,
		})

	}
	// io.Copy(os.Stdout, resp.Body)
	return nil
}

// func (c *Course) GetInfo(token string) error {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", fmt.Sprintf("https://troy.instructure.com/api/v1/courses/%v", c.Info.ID), nil)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
// 	// var announcement entity.Course_Announement_Instructure
// 	if err != nil {
// 		return err
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	io.Copy(os.Stdout, resp.Body)

// 	return nil
// }
