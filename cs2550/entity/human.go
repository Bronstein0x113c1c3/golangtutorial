package entity

import "time"

type Human struct {
	ID         int
	Surname    string
	Name       string
	Department string
	Postion    string
	DoB        time.Time
	Hometown   string
	Addr       string
	Email      string
	PhoneNum   string
	DayStarted time.Time
	Record     map[time.Time][2]time.Time
}
