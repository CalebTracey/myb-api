package psql

import (
	"fmt"
	"github.com/calebtracey/mind-your-business-api/external"
)

type MapperI interface {
	PostgresExec(request *external.ApiRequest) string
}

type Mapper struct{}

func (m Mapper) PostgresExec(request *external.ApiRequest) string {
	columns, values := ParseStructToSlices(request.Request.User)
	return fmt.Sprintf(InsertExec, "users", columns, values)
}
