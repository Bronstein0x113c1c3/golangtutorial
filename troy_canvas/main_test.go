package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetUser(t *testing.T) {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")
	GetUser(token)
}
func TestGetCourses(t *testing.T) {
	godotenv.Load(".env")
	token := os.Getenv("TOKEN")
	GetCourses(token)
}
