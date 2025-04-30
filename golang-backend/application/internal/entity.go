package entity

import "time"

type Appointment struct {
	AppID     string
	Title     string
	Diagnosis string
	Treatment string
	Status    string
	Date      time.Time
}

type Patient struct {
	Pid          string
	Name         string
	Email        string
	PhNo         string
	DOB          string
	Status       string
	Appointments []Appointment
}

type Staff struct {
	StaffId        string
	Name           string
	HashedPassword string
	Type           string
}
