package common

type College struct {
	database map[int]Student
}

type Student struct {
	ID        int
	firstname string
	lastname  string
}

func (s *Student) Add(payload Student, reply *Student) error {
	exists, ok := 
	
	return nil
}

func Get(payload Student, reply *Student) error {

	return nil
}

func  NewCollege() *College {
	//return &College{
	//	database: make(map[int]Student),
	//}
	//
	c := new(College)
	c.database = make(map[int]Student)
	return c
}