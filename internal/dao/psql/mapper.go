package psql

import (
	"fmt"
	"github.com/calebtracey/mind-your-business-api/external"
	log "github.com/sirupsen/logrus"
	"strings"
)

type MapperI interface {
	PostgresExec(request *external.ApiRequest) string
}

type Mapper struct{}

func (m Mapper) PostgresExec(request *external.ApiRequest) string {
	columns, values := ParseStructToSlices(request.Request.User)
	query := fmt.Sprintf(InsertExec, "users", strings.Join(columns, ", "), strings.Join(values, ", "))
	log.Infof("PostgresExec: query:\n%s\n", query)
	return query
}
