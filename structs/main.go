package main

import (
	"fmt"
	"server/data"
)

func main() {

	// var max data.Instructor
	max := data.Instructor{Id: 1, LastName: "freeman"}
	max.FirstName = "max"
	fmt.Println(max.Print())

	kyle := data.Instructor{Id: 2, FirstName: "kyle", LastName: "simpson"}
	fmt.Println(kyle.Print())

	goCourse := data.Course{Id: 2, Name: "Go fundamentals",
		Instructor: max}
	// fmt.Printf("%v", goCourse)

	// swiftWs := data.Workshop{
	// Course: data.Course{
	// Name:       "Swift for ios",
	// Instructor: max,
	// },
	// Date: time.Now(),
	// }
	swiftWs := data.NewWorkshop("Swift with iOS", max)
	// fmt.Println(swiftWs)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWs.Course

	for _, course := range courses {
		fmt.Println(course)
	}

}
