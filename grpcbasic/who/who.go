package who

import "fmt"

type Key struct {
	Name  string
	State string
}
type Country struct {
	K             Key
	PositiveCases int
}

type WHO struct {
	database map[Key]Country
}

func NewWHO() *WHO {
	return &WHO{
		database: make(map[Key]Country),
	}
}

func (w *WHO) Add(payload Country, reply *Country) error {
	if _, ok := w.database[payload.K]; ok {
		return fmt.Errorf("already in the list '%v'", payload.K)
	}
	w.database[payload.K] = payload
	*reply = payload
	return nil
}

func (w *WHO) Get(payload Country, reply *Country) error {
	results, ok := w.database[payload.K]
	if !ok {
		return fmt.Errorf("unable to fine in the database '%v'", payload.K)
	}
	*reply = results
	return nil
}
