package psql

import "fmt"

type MapperI interface {
	PostgresExec() string
}

type Mapper struct{}

func (m Mapper) PostgresExec() string {
	return fmt.Sprintf(InsertExec, "users", "first_name", "caleb")
}
