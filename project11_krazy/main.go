package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"

	// "sync"

	_ "strconv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
)

type Student_Entity struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}
type List_Student struct {
	Student_List []Student_Entity `json:"students"`
}

var student_map = make(map[int]bool)

func Crawling(linktocrawl string, pagesize int, page_status int /*entity_chan chan []Student_Entity, wg *sync.WaitGroup*/, entity_chan chan []Student_Entity, worker_id chan int, check map[int]bool) {
	// Start a query instance (HTML page) to crawl (webpage?, size of page?, which page?)
	doc, err := goquery.NewDocument(linktocrawl + "&pageSize=" + strconv.Itoa(pagesize) + "&SinhvienLmh_page=" + strconv.Itoa(page_status))
	if err != nil {
		fmt.Println("error countered")
	}
	//
	student_entity := make([]Student_Entity, 0, pagesize)
	doc.Find("tbody").Find("tr").Each(func(i int, s1 *goquery.Selection) {
		var student Student_Entity
		count := 0
		s1.Find("td").Each(func(j int, s *goquery.Selection) {
			attr, _ := s.Attr("style")
			if attr == "width: 40px" {
				student.Id, _ = strconv.Atoi(s.Text())
				count += 1
			} else if attr == "width: 100px" && count == 1 {
				student.Name = s.Text()
				count = 0
			} else if attr == "width: 20px" {
				if random := rand.Float64() * 2; random >= 1 {
					student.Checked = true
				} else {
					student.Checked = false
				}
			}
		})
		// fmt.Println(student)
		if _, exist := check[student.Id]; exist == false {
			student_entity = append(student_entity, student)
			check[student.Id] = true
		}
	})
	entity_chan <- student_entity
	worker_id <- page_status
	// close(entity_chan)
}

func main() {
	// wg.Add(1)
	list_worker_done := make([]int, 0, 8)
	var student_entity []Student_Entity = make([]Student_Entity, 0, 100000)
	student_chan := make(chan []Student_Entity)
	id_chan := make(chan int)
	// go func() {
	// 	Crawling("http://112.137.129.87/qldt/?SinhvienLmh%5Bterm_id%5D=035&ajax=sinhvien-lmh-grid", 5000, 3, student_chan)
	// }()
	for i := 1; i <= 8; i++ {
		go Crawling("http://112.137.129.87/qldt/?SinhvienLmh%5Bterm_id%5D=035&ajax=sinhvien-lmh-grid", 5000, i, student_chan, id_chan, student_map)
	}
	for {
		student_list := <-student_chan
		log.Printf("this list has %v students \n", len(student_list))
		student_entity = append(student_entity, student_list...)
		log.Printf("the total students imported: %v students \n", len(student_entity))
		worker_done_id := <-id_chan
		log.Printf("worker %v done", worker_done_id)
		list_worker_done = append(list_worker_done, worker_done_id)
		if len(list_worker_done) == 8 {
			break
		}
	}
	log.Printf("total students after crawling: %v", len(student_entity))
	students_data := &List_Student{
		Student_List: student_entity,
	}
	// _ = students_data
	byte_data, err2 := json.MarshalIndent(*students_data, "", "    ")
	if err2 != nil {
		panic(err2)
	}
	err := ioutil.WriteFile("./students.json", byte_data, 0644)
	if err != nil {
		log.Fatalf("error in writing file")
	}
}
