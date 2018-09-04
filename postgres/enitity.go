package postgres

import "github.com/user/work/foosball-booking/date"



//Employee table for storing Employee data
type Employee struct {
	EmployeeID   int    `json:"employeeID"`
	EmployeeName string `json:"employeeName"`
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

