package psql

import (
	"fmt"
	"github.com/calebtracey/mind-your-business-api/external"
	"strings"
)

type MapperI interface {
	NewUserExec(request *external.ApiRequest) string
}

type Mapper struct{}

func (m Mapper) NewUserExec(request *external.ApiRequest) string {
	columns, values := ParseStructToSlices(request.Request.User)
	return fmt.Sprintf(InsertExec, "users", strings.Join(columns, ", "), strings.Join(values, ", "))
}
