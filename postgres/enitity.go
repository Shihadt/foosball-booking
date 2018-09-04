package postgres

import "github.com/user/work/foosball-booking/date"

//User table for storing user data
type User struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
}

//Admin table
type Admin struct {
	AdminID   int    `json:"adminID"`
	AdminName string `json:"adminName"`
}

//Request table
type Request struct {
	RequestID   int  `json:"requestID"`
	UserID      int  `json:"userId"`
	TimeSlot    int  `json:"timeSlot"`
	IsAllocated bool `json:"isAllocated"`
}

//Slot table
type Slot struct {
	SlotID  int        `json:"slotID"`
	Request [4]Request `json:"request"`
}

//Period table
type Period struct {
	PeriodID int       `json:"periodID"`
	Date     date.Date `json:"date"`
	Slot     []Slot    `json:"slot"`
}
