package common

import "fmt"

type College struct {
	database map[int]Student
}

type Student struct {
	ID        int
	firstname string
	lastname  string
}

func (c *College) Add(payload Student, reply *Student) error {
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("Already exits ID '%d' ", payload.ID)
	}

	return nil
}

func (c *College) Get(payload Student, reply *Student) error {

	return nil
}

func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}

	//c := new(College)
	//c.database = make(map[int]Student)
	//return c
}
