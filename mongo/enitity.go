package mongo

//User table for storing user data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Admin table
type Admin struct {
	ID int `json:"id"`
}

//Request table
type Request struct {
	ID     int `json:"ID"`
	UserID int `json:"userId"`
	
}

//Slot table
type Slot struct {
	ID int       `json:"id"`
	S  []Request `json:"s"`
}
