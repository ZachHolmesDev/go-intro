package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// method could be anywhere in package like in another
// file although that would be a little silly

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", i.LastName, i.FirstName, i.Score)
}
