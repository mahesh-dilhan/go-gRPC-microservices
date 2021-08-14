package who

import "fmt"

type KEY struct {
	name  string
	state string
}
type Country struct {
	key           KEY
	positiveCases int
}

type WHO struct {
	database map[KEY]Country
}

func NewWHO() *WHO {
	return &WHO{
		database: make(map[KEY]Country),
	}
}

func (w *WHO) Add(payload Country, reply *Country) error {
	if _, ok := w.database[payload.key]; ok {
		return fmt.Errorf("already in the list '%v'", payload.key)
	}
	w.database[payload.key] = payload
	*reply = payload
	return nil
}

func (w *WHO) Get(payload Country, reply *Country) error {
	results, ok := w.database[payload.key]
	if !ok {
		return fmt.Errorf("unable to fine in the database '%v'", payload.key)
	}
	*reply = results
	return nil
}
