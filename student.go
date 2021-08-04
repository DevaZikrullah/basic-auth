package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "01", Name: "suhas", Grade: 1})
	students = append(students, &Student{Id: "02", Name: "joko", Grade: 2})
	students = append(students, &Student{Id: "03", Name: "aljo", Grade: 3})
}
