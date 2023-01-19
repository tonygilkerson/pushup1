package student

type Student struct {
	Name string
	ID   int
}

type Students map[int]Student

func (sl Students) InitStudentList() {

	sl[111] = Student{"One",111}
	sl[222] = Student{"Two",222}
	sl[333] = Student{"Three",333}
	
}

