package entity

import (
	"fmt"
	"os"
	"troy_canvas_optimized/entity"

	"github.com/joho/godotenv"
)

var Student *entity.Student

func init() {
	_ = godotenv.Load("././env/.env")
	token := os.Getenv("TOKEN")
	fmt.Println(token)
	Student = entity.Init(token)
}
