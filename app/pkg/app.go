package build

import (
	"github.com/tonygilkerson/pushup1/app/pkg/student"
)

var pushupGreeting = "Welcome to Pushup!"

var StudentList student.Students

func init() {
	StudentList = student.Students(make(map[int]student.Student))
	StudentList.InitStudentList()
}






